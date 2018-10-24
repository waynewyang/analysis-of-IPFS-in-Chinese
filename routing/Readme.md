# 路由层分析

peerid = cid
## 目录
- [目的](#目的)
- [对上层协议接口](#对上层协议接口)
- [DHT原理](/routing/dht.md)
- [源码中修改K桶 a并发与k值提高效率](/routing/修改并发属.md)

## 目的

- **peer routing **
> 节点路由，寻找网络中的对等节点

- **content routing**
> 寻址网络中的内容，基于IPLD的内容寻址

- **DHT操作方法**
> 通过PutValue以及GetValue方法，提供分布式的key-value键值对存储、获取

> 例子：内容DHT表、IPNS记录

> 扩展：其他的是否可行，比如分布式鉴权，将鉴权信息去中心化

## 对上层协议接口
- 接口目录 libp2p/go-libp2p-routing
- IpfsRouting接口
```
// IpfsRouting is the combination of different routing types that ipfs
// uses. It can be satisfied by a single item (such as a DHT) or multiple
// different pieces that are more optimized to each task.
type IpfsRouting interface {
	ContentRouting
	PeerRouting
	ValueStore

	// Bootstrap allows callers to hint to the routing system to get into a
	// Boostrapped state
	Bootstrap(context.Context) error

	// TODO expose io.Closer or plain-old Close error
}
```
- PeerRouting
```
// PeerRouting is a way to find information about certain peers.
// This can be implemented by a simple lookup table, a tracking server,
// or even a DHT.
type PeerRouting interface {
	// Find specific Peer
	// FindPeer searches for a peer with given ID, returns a pstore.PeerInfo
	// with relevant addresses.
	FindPeer(context.Context, peer.ID) (pstore.PeerInfo, error)
}
```
- ContentRouting
```
// ContentRouting is a value provider layer of indirection. It is used to find
// information about who has what content.
type ContentRouting interface {
	// Provide adds the given cid to the content routing system. If 'true' is
	// passed, it also announces it, otherwise it is just kept in the local
	// accounting of which objects are being provided.
	Provide(context.Context, cid.Cid, bool) error

	// Search for peers who are able to provide a given key
	FindProvidersAsync(context.Context, cid.Cid, int) <-chan pstore.PeerInfo
}
```
- ValueStore
```
// ValueStore is a basic Put/Get interface.
type ValueStore interface {

	// PutValue adds value corresponding to given Key.
	PutValue(context.Context, string, []byte, ...ropts.Option) error

	// GetValue searches for the value corresponding to given Key.
	GetValue(context.Context, string, ...ropts.Option) ([]byte, error)

	// SearchValue searches for better and better values from this value
	// store corresponding to the given Key. By default implementations must
	// stop the search after a good value is found. A 'good' value is a value
	// that would be returned from GetValue.
	//
	// Useful when you want a result *now* but still want to hear about
	// better/newer results.
	//
	// Implementations of this methods won't return ErrNotFound. When a value
	// couldn't be found, the channel will get closed without passing any results
	SearchValue(context.Context, string, ...ropts.Option) (<-chan []byte, error)
}
```
- [回到目录](#目录)

##  内部细节（待分析）
放到libp2p环节详细分析

- 底层路由方法
  - DHT
  - mDNS
  - snr
  - dns
- IPRS
- 其他