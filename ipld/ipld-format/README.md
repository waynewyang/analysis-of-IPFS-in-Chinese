## ipld-format

### 说明

format主要规范定义了IPLD的接口抽象层

对DAG的批量管理，缓冲大小限制8M（8左移20位），节点数目限制128个

对于不同格式的IPLD解析器，需要分别去实现。

### 源码分析

- format.go
    1 Node是所有IPLD对象必须实现的接口
    2 Node表征了IPLD对戏那个包含的数据以及其links
    3 节点与link之间的互相转换

```
package format

import (
	"context"
	"fmt"

	blocks "gx/ipfs/QmWAzSEoqZ6xU6pu8yL8e5WaMb7wtbfbhhN4p1DknUPtr3/go-block-format"

	cid "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"
)

//解析:需要进一步看实现确认？
type Resolver interface {
	// Resolve resolves a path through this node, stopping at any link boundary
	// and returning the object found as well as the remaining path to traverse
	Resolve(path []string) (interface{}, []string, error)

    // 查看某个节点的叶子节点?
	// Tree lists all paths within the object under 'path', and up to the given depth.
	// To list the entire object (similar to `find .`) pass "" and -1
	Tree(path string, depth int) []string
}

// Node is the base interface all IPLD nodes must implement.
//
// Nodes are **Immutable** and all methods defined on the interface are
// **Thread Safe**.
// 所有的IPLD节点均需要实现Node接口
type Node interface {
	blocks.Block //block包含了cid与实际的数据
	Resolver     //解析:需要进一步看实现确认？

	// ResolveLink is a helper function that calls resolve and asserts the
	// output is a link
	ResolveLink(path []string) (*Link, []string, error) //解析一个路径所包含的Link

	//深拷贝一个IPLD节点
	// Copy returns a deep copy of this node
	Copy() Node

	// 返回一个节点对象所有的Links
	// Links is a helper function that returns all links within this object
	Links() []*Link

	//返回节点状态
	// TODO: not sure if stat deserves to stay
	Stat() (*NodeStat, error)

	//返回对象的大小字节数
	// Size returns the size in bytes of the serialized object
	Size() (uint64, error)
}

// Link结构体，
// Link represents an IPFS Merkle DAG Link between Nodes.
type Link struct {
	// utf string name. should be unique per object
	Name string // utf8

	// cumulative size of target object
	Size uint64

	// multihash of the target object
	Cid *cid.Cid
}

// NodeStat is a statistics object for a Node. Mostly sizes.
type NodeStat struct {
	Hash string
	////有多少个links
	NumLinks int // number of links in link table
	////非link数据的大小
	BlockSize int // size of the raw, encoded data
	//links部分的大小
	LinksSize int // size of the links segment
	//数据大小
	DataSize int // size of the data segment
	//总大小
	CumulativeSize int // cumulative size of object and its references
}

func (ns NodeStat) String() string {
	f := "NodeStat{NumLinks: %d, BlockSize: %d, LinksSize: %d, DataSize: %d, CumulativeSize: %d}"
	return fmt.Sprintf(f, ns.NumLinks, ns.BlockSize, ns.LinksSize, ns.DataSize, ns.CumulativeSize)
}

//将某个节点转换为一个link对象,成为其他节点的一部分
// MakeLink creates a link to the given node
func MakeLink(n Node) (*Link, error) {
	s, err := n.Size()
	if err != nil {
		return nil, err
	}

	return &Link{
		Size: s,
		Cid:  n.Cid(),
	}, nil
}

//通过link获取其Node节点对象
// GetNode returns the MDAG Node that this link points to
func (l *Link) GetNode(ctx context.Context, serv NodeGetter) (Node, error) {
	return serv.Get(ctx, l.Cid)
}

```
- coding.go 

目的是对block进行解码为node，node是IPLD的最小单元(实现了IPLD的所有接口)
BlockDecoder  规范了注册解码函数，以及具体解码的接口，由safeBlockDecoder实现
```
package format

import (
	"fmt"
	"sync"

	blocks "gx/ipfs/QmWAzSEoqZ6xU6pu8yL8e5WaMb7wtbfbhhN4p1DknUPtr3/go-block-format"
)

// DecodeBlockFunc functions decode blocks into nodes.
//将block转换为Node
type DecodeBlockFunc func(block blocks.Block) (Node, error)

type BlockDecoder interface {
	Register(codec uint64, decoder DecodeBlockFunc)
	Decode(blocks.Block) (Node, error)
}

//实现了BlockDecoder接口
type safeBlockDecoder struct {
	// Can be replaced with an RCU if necessary.
	lock     sync.RWMutex
	decoders map[uint64]DecodeBlockFunc
}

//注册解码器,对不同的codec注册不同的block解码器
// Register registers decoder for all blocks with the passed codec.
//
// This will silently replace any existing registered block decoders.
func (d *safeBlockDecoder) Register(codec uint64, decoder DecodeBlockFunc) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.decoders[codec] = decoder
}

func (d *safeBlockDecoder) Decode(block blocks.Block) (Node, error) {
	// Short-circuit by cast if we already have a Node.
	if node, ok := block.(Node); ok {
		return node, nil
	}

	ty := block.Cid().Type()

	d.lock.RLock()
	decoder, ok := d.decoders[ty]
	d.lock.RUnlock()

	if ok {
		return decoder(block)
	} else {
		// TODO: get the *long* name for this format
		return nil, fmt.Errorf("unrecognized object type: %d", ty)
	}
}

var DefaultBlockDecoder BlockDecoder = &safeBlockDecoder{decoders: make(map[uint64]DecodeBlockFunc)}

// Decode decodes the given block using the default BlockDecoder.
func Decode(block blocks.Block) (Node, error) {
	return DefaultBlockDecoder.Decode(block)
}

// Register registers block decoders with the default BlockDecoder.
func Register(codec uint64, decoder DecodeBlockFunc) {
	DefaultBlockDecoder.Register(codec, decoder)
}
```


- merkledag.go
    1 NodeGetter接口，通过cid获取Node
    2 LinkGetter接口，通过cid获取Links
    3 DAGService 接口，对DAG进行增删操作

```
package format

import (
	"context"
	"fmt"

	cid "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"
)

var ErrNotFound = fmt.Errorf("merkledag: not found")

// Either a node or an error.
type NodeOption struct {
	Node Node
	Err  error
}

//从cid获取Node信息
// The basic Node resolution service.
type NodeGetter interface {
	// Get retrieves nodes by CID. Depending on the NodeGetter
	// implementation, this may involve fetching the Node from a remote
	// machine; consider setting a deadline in the context.
	Get(context.Context, *cid.Cid) (Node, error)

	// GetMany returns a channel of NodeOptions given a set of CIDs.
	GetMany(context.Context, []*cid.Cid) <-chan *NodeOption
}

//从cid获取Links信息
// NodeGetters can optionally implement this interface to make finding linked
// objects faster.
type LinkGetter interface {
	NodeGetter

	// TODO(ipfs/go-ipld-format#9): This should return []*cid.Cid

	// GetLinks returns the children of the node refered to by the given
	// CID.
	GetLinks(ctx context.Context, nd *cid.Cid) ([]*Link, error)
}

//DAG管理服务，增删
// DAGService is an IPFS Merkle DAG service.
type DAGService interface {
	NodeGetter

	// Add adds a node to this DAG.
	Add(context.Context, Node) error

	// Remove removes a node from this DAG.
	//
	// Remove returns no error if the requested node is not present in this DAG.
	Remove(context.Context, *cid.Cid) error

	// AddMany adds many nodes to this DAG.
	//
	// Consider using NewBatch instead of calling this directly if you need
	// to add an unbounded number of nodes to avoid buffering too much.
	AddMany(context.Context, []Node) error

	// RemoveMany removes many nodes from this DAG.
	//
	// It returns success even if the nodes were not present in the DAG.
	RemoveMany(context.Context, []*cid.Cid) error
}
```
- batch.go
    批量对dag进行操作 

```
package format

import (
	"context"
	"errors"
	"runtime"
)

// ParallelBatchCommits is the number of batch commits that can be in-flight before blocking.
// TODO(ipfs/go-ipfs#4299): Experiment with multiple datastores, storage
// devices, and CPUs to find the right value/formula.
var ParallelBatchCommits = runtime.NumCPU() * 2

// ErrNotCommited is returned when closing a batch that hasn't been successfully
// committed.
var ErrNotCommited = errors.New("error: batch not commited")

// ErrClosed is returned when operating on a batch that has already been closed.
var ErrClosed = errors.New("error: batch closed")

//wayne:批量增加，限制最大大小为8M（8<<20）,最大增加节点数目：128个
// NewBatch returns a node buffer (Batch) that buffers nodes internally and
// commits them to the underlying DAGService in batches. Use this if you intend
// to add or remove a lot of nodes all at once.
//
// If the passed context is canceled, any in-progress commits are aborted.
func NewBatch(ctx context.Context, ds DAGService) *Batch {
	ctx, cancel := context.WithCancel(ctx)
	return &Batch{
		ds:            ds,
		ctx:           ctx,
		cancel:        cancel,
		commitResults: make(chan error, ParallelBatchCommits),
		MaxSize:       8 << 20,

		// By default, only batch up to 128 nodes at a time.
		// The current implementation of flatfs opens this many file
		// descriptors at the same time for the optimized batch write.
		MaxNodes: 128,
	}
}

// Batch is a buffer for batching adds to a dag.
// 批量增删的结构体
type Batch struct {
	ds DAGService

	ctx    context.Context
	cancel func()

	activeCommits int
	err           error
	commitResults chan error

	nodes []Node
	size  int

	MaxSize  int
	MaxNodes int
}

func (t *Batch) processResults() {
	for t.activeCommits > 0 {
		select {
		case err := <-t.commitResults:
			t.activeCommits--
			if err != nil {
				t.setError(err)
				return
			}
		default:
			return
		}
	}
}

func (t *Batch) asyncCommit() {
	numBlocks := len(t.nodes)
	if numBlocks == 0 {
		return
	}
	if t.activeCommits >= ParallelBatchCommits {
		select {
		case err := <-t.commitResults:
			t.activeCommits--

			if err != nil {
				t.setError(err)
				return
			}
		case <-t.ctx.Done():
			t.setError(t.ctx.Err())
			return
		}
	}
	go func(ctx context.Context, b []Node, result chan error, ds DAGService) {
		select {
		case result <- ds.AddMany(ctx, b):
		case <-ctx.Done():
		}
	}(t.ctx, t.nodes, t.commitResults, t.ds)

	t.activeCommits++
	t.nodes = make([]Node, 0, numBlocks)
	t.size = 0

	return
}

//wayne:增加一个Node到批量缓冲中
// Add adds a node to the batch and commits the batch if necessary.
func (t *Batch) Add(nd Node) error {
	if t.err != nil {
		return t.err
	}
	// Not strictly necessary but allows us to catch errors early.
	t.processResults()

	if t.err != nil {
		return t.err
	}

	t.nodes = append(t.nodes, nd)
	t.size += len(nd.RawData())

	if t.size > t.MaxSize || len(t.nodes) > t.MaxNodes {
		t.asyncCommit()
	}
	return t.err
}

// wayne:提交,批量节点
// Commit commits batched nodes.
func (t *Batch) Commit() error {
	if t.err != nil {
		return t.err
	}

	t.asyncCommit()

loop:
	for t.activeCommits > 0 {
		select {
		case err := <-t.commitResults:
			t.activeCommits--
			if err != nil {
				t.setError(err)
				break loop
			}
		case <-t.ctx.Done():
			t.setError(t.ctx.Err())
			break loop
		}
	}

	return t.err
}

func (t *Batch) setError(err error) {
	t.err = err

	t.cancel()

	// Drain as much as we can without blocking.
loop:
	for {
		select {
		case <-t.commitResults:
		default:
			break loop
		}
	}

	// Be nice and cleanup. These can take a *lot* of memory.
	t.commitResults = nil
	t.ds = nil
	t.ctx = nil
	t.nodes = nil
	t.size = 0
	t.activeCommits = 0
}
```
- promise.go
    1 线程安全，阻塞性获取node
    2 Send 赋值
    3 Get 获取node

```
package format

import (
	"context"
)

// NodePromise provides a promise like interface for a dag Node
// the first call to Get will block until the Node is received
// from its internal channels, subsequent calls will return the
// cached node.
//
// Thread Safety: This is multiple-consumer/single-producer safe.
func NewNodePromise(ctx context.Context) *NodePromise {
	return &NodePromise{
		done: make(chan struct{}),
		ctx:  ctx,
	}
}

type NodePromise struct {
	value Node
	err   error
	done  chan struct{}

	ctx context.Context
}

// Call this function to fail a promise.
//
// Once a promise has been failed or fulfilled, further attempts to fail it will
// be silently dropped.
func (np *NodePromise) Fail(err error) {
	if np.err != nil || np.value != nil {
		// Already filled.
		return
	}
	np.err = err
	close(np.done)
}

// Fulfill this promise.
//
// Once a promise has been fulfilled or failed, calling this function will
// panic.
func (np *NodePromise) Send(nd Node) {
	// if promise has a value, don't fail it
	if np.err != nil || np.value != nil {
		panic("already filled")
	}
	np.value = nd
	close(np.done)
}

// Get the value of this promise.
//
// This function is safe to call concurrently from any number of goroutines.
func (np *NodePromise) Get(ctx context.Context) (Node, error) {
	select {
	case <-np.done:
		return np.value, np.err
	case <-np.ctx.Done():
		return nil, np.ctx.Err()
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
```
- daghelper.go
    dag的对象获取操作函数

```
package format

import (
	"context"

	cid "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"
)

// GetLinks returns the CIDs of the children of the given node. Prefer this
// method over looking up the node itself and calling `Links()` on it as this
// method may be able to use a link cache.
func GetLinks(ctx context.Context, ng NodeGetter, c *cid.Cid) ([]*Link, error) {
	if c.Type() == cid.Raw {
		return nil, nil
	}
	if gl, ok := ng.(LinkGetter); ok {
		return gl.GetLinks(ctx, c)
	}
	node, err := ng.Get(ctx, c)
	if err != nil {
		return nil, err
	}
	return node.Links(), nil
}

// GetDAG will fill out all of the links of the given Node.
// It returns an array of NodePromise with the linked nodes all in the proper
// order.
func GetDAG(ctx context.Context, ds NodeGetter, root Node) []*NodePromise {
	var cids []*cid.Cid
	for _, lnk := range root.Links() {
		cids = append(cids, lnk.Cid)
	}

	return GetNodes(ctx, ds, cids)
}

// GetNodes returns an array of 'FutureNode' promises, with each corresponding
// to the key with the same index as the passed in keys
func GetNodes(ctx context.Context, ds NodeGetter, keys []*cid.Cid) []*NodePromise {

	// Early out if no work to do
	if len(keys) == 0 {
		return nil
	}

	promises := make([]*NodePromise, len(keys))
	for i := range keys {
		promises[i] = NewNodePromise(ctx)
	}

	dedupedKeys := dedupeKeys(keys)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		nodechan := ds.GetMany(ctx, dedupedKeys)

		for count := 0; count < len(keys); {
			select {
			case opt, ok := <-nodechan:
				if !ok {
					for _, p := range promises {
						p.Fail(ErrNotFound)
					}
					return
				}

				if opt.Err != nil {
					for _, p := range promises {
						p.Fail(opt.Err)
					}
					return
				}

				nd := opt.Node
				c := nd.Cid()
				for i, lnk_c := range keys {
					if c.Equals(lnk_c) {
						count++
						promises[i].Send(nd)
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return promises
}

//去重
// Remove duplicates from a list of keys
func dedupeKeys(cids []*cid.Cid) []*cid.Cid {
	set := cid.NewSet()
	for _, c := range cids {
		set.Add(c)
	}
	return set.Keys()
}
```
