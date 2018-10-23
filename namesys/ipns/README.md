# IPNS 原理分析

## 底层go-ipns包

### 创建IPNS对象
> func Create(sk ic.PrivKey, val []byte, seq uint64, eol time.Time) (*pb.IpnsEntry, error) 
> 用私钥创建IPNS对象，映射具体的ipfsPath
> sk 私钥， val:真实的ipfs路径, seq 序号，eol 有效期

### IPNS对象与内容哈希的绑定
> func EmbedPublicKey(pk ic.PubKey, entry *pb.IpnsEntry) error
>  建立内容hash与ipns的映射,第一个参数是真实的内容哈希地址，第二个是IPNS对象
- 提供ipns对象创建方法

### IPNS对象的查找
> func selectRecord(recs []*pb.IpnsEntry, vals [][]byte) (int, error) 
> 有超过1条以上ipns记录，则返回最新的记录

> 目前的实现是没有历史记录的，历史记录在ipfs网络之中，但是没有没ipns映射

## namesys包
- 接口
```
go-ipfs/namesys/interface.go
// Namesys represents a cohesive name publishing and resolving system.
//
// Publishing a name is the process of establishing a mapping, a key-value
// pair, according to naming rules and databases.
//
// Resolving a name is the process of looking up the value associated with the
// key (name).
type NameSystem interface {
	Resolver
	Publisher
}

// Resolver is an object capable of resolving names.
type Resolver interface {

	// Resolve performs a recursive lookup, returning the dereferenced
	// path.  For example, if ipfs.io has a DNS TXT record pointing to
	//   /ipns/QmatmE9msSfkKxoffpHwNLNKgwZG8eT9Bud6YoPab52vpy
	// and there is a DHT IPNS entry for
	//   QmatmE9msSfkKxoffpHwNLNKgwZG8eT9Bud6YoPab52vpy
	//   -> /ipfs/Qmcqtw8FfrVSBaRmbWwHxt3AuySBhJLcvmFYi3Lbc4xnwj
	// then
	//   Resolve(ctx, "/ipns/ipfs.io")
	// will resolve both names, returning
	//   /ipfs/Qmcqtw8FfrVSBaRmbWwHxt3AuySBhJLcvmFYi3Lbc4xnwj
	//
	// There is a default depth-limit to avoid infinite recursion.  Most
	// users will be fine with this default limit, but if you need to
	// adjust the limit you can specify it as an option.
	Resolve(ctx context.Context, name string, options ...opts.ResolveOpt) (value path.Path, err error)
}

// Publisher is an object capable of publishing particular names.
type Publisher interface {

	// Publish establishes a name-value mapping.
	// TODO make this not PrivKey specific.
	Publish(ctx context.Context, name ci.PrivKey, value path.Path) error

	// TODO: to be replaced by a more generic 'PublishWithValidity' type
	// call once the records spec is implemented
	PublishWithEOL(ctx context.Context, name ci.PrivKey, value path.Path, eol time.Time) error
}

```

- publish
> Publish动作->更新本地的ipns记录->
> 调用route层(go-libp2p-routing)的PutValue方法发布ipns记录

> IpnsPublisher实现了Publish接口
func (p *IpnsPublisher) Publish(ctx context.Context, k ci.PrivKey, value path.Path) error

> func PutRecordToRouting(ctx context.Context, r routing.ValueStore, k ci.PubKey, entry *pb.IpnsEntry) error
> ipns.EmbedPublicKey(k, entry)

>func PublishEntry(ctx context.Context, r routing.ValueStore, ipnskey string, rec *pb.IpnsEntry)
>r.PutValue(timectx, ipnskey, data)

- Resolver
> Resolver动作->查询本地记录->缓存OK，结束
> 不OK，调用路由层GetValue方法

## 代码分析

- ipns包ipns.go

```
package ipns

import (
	"bytes"
	"fmt"
	"time"

	pb "gx/ipfs/QmNqBhXpBKa5jcjoUZHfxDgAFxtqK3rDA5jtW811GBvVob/go-ipns/pb"

	u "gx/ipfs/QmPdKqUcHGFdeSpvjVoaTRPPstGif9GBZb5Q56RVw9o69A/go-ipfs-util"
	ic "gx/ipfs/QmPvyPwuCgJ7pDmrKDxRtsScJgBaM5h4EpRL2qQJsmXf4n/go-libp2p-crypto"
	peer "gx/ipfs/QmQsErDt8Qgw1XrsXf2BpEzDgGWtB1YLsTAARBup5b6B9W/go-libp2p-peer"
)

// Create creates a new IPNS entry and signs it with the given private key.
//
// This function does not embed the public key. If you want to do that, use
// `EmbedPublicKey`.
// 用私钥创建IPNS对象，映射具体的ipfsPath
// sk 私钥， val:真实的ipfs路径,seq 序号，eol 有效期
func Create(sk ic.PrivKey, val []byte, seq uint64, eol time.Time) (*pb.IpnsEntry, error) {
	entry := new(pb.IpnsEntry)

	entry.Value = val
	typ := pb.IpnsEntry_EOL
	entry.ValidityType = &typ
	entry.Sequence = &seq
	entry.Validity = []byte(u.FormatRFC3339(eol))

	sig, err := sk.Sign(ipnsEntryDataForSig(entry))
	if err != nil {
		return nil, err
	}
	entry.Signature = sig

	return entry, nil
}

// 校验IPNS对象，签名、过期时间等
// Validates validates the given IPNS entry against the given public key.
func Validate(pk ic.PubKey, entry *pb.IpnsEntry) error {
	// Check the ipns record signature with the public key
	// 用公钥验证ipns对象合法性,验证签名
	if ok, err := pk.Verify(ipnsEntryDataForSig(entry), entry.GetSignature()); err != nil || !ok {
		return ErrSignature
	}
	//获取过期时间失败
	eol, err := GetEOL(entry)
	if err != nil {
		return err
	}
	//过期
	if time.Now().After(eol) {
		return ErrExpiredRecord
	}
	return nil
}

// GetEOL returns the EOL of this IPNS entry
//
// This function returns ErrUnrecognizedValidity if the validity type of the
// record isn't EOL. Otherwise, it returns an error if it can't parse the EOL.
func GetEOL(entry *pb.IpnsEntry) (time.Time, error) {
	if entry.GetValidityType() != pb.IpnsEntry_EOL {
		return time.Time{}, ErrUnrecognizedValidity
	}
	return u.ParseRFC3339(string(entry.GetValidity()))
}

// 建立内容与ipns的映射,第一个参数是真实的地址，第二个是IPNS对象
// EmbedPublicKey embeds the given public key in the given ipns entry. While not
// strictly required, some nodes (e.g., DHT servers) may reject IPNS entries
// that don't embed their public keys as they may not be able to validate them
// efficiently.
func EmbedPublicKey(pk ic.PubKey, entry *pb.IpnsEntry) error {
	// Try extracting the public key from the ID. If we can, *don't* embed
	// it.
	id, err := peer.IDFromPublicKey(pk)
	if err != nil {
		return err
	}
	extracted, err := id.ExtractPublicKey()
	if err != nil {
		return err
	}
	if extracted != nil {
		return nil
	}

	// We failed to extract the public key from the peer ID, embed it in the
	// record.
	pkBytes, err := pk.Bytes()
	if err != nil {
		return err
	}
	entry.PubKey = pkBytes
	return nil
}

// ExtractPublicKey extracts a public key matching `pid` from the IPNS record,
// if possible.
//
// This function returns (nil, nil) when no public key can be extracted and
// nothing is malformed.
func ExtractPublicKey(pid peer.ID, entry *pb.IpnsEntry) (ic.PubKey, error) {
	//存在ipns,从ipns对象中提取到公钥，然后通过公钥转换为peerid，判断是否与给定的pid相等
	if entry.PubKey != nil {
		pk, err := ic.UnmarshalPublicKey(entry.PubKey)
		if err != nil {
			return nil, fmt.Errorf("unmarshaling pubkey in record: %s", err)
		}

		expPid, err := peer.IDFromPublicKey(pk)
		if err != nil {
			return nil, fmt.Errorf("could not regenerate peerID from pubkey: %s", err)
		}

		if pid != expPid {
			return nil, ErrPublicKeyMismatch
		}
		return pk, nil
	}

	//不存在ipns
	return pid.ExtractPublicKey()
}

// Compare compares two IPNS entries. It returns:
//
// * -1 if a is older than b
// * 0 if a and b cannot be ordered (this doesn't mean that they are equal)
// * +1 if a is newer than b
//
// It returns an error when either a or b are malformed.
//
// NOTE: It *does not* validate the records, the caller is responsible for calling
// `Validate` first.
//
// NOTE: If a and b cannot be ordered by this function, you can determine their
// order by comparing their serialized byte representations (using
// `bytes.Compare`). You must do this if you are implementing a libp2p record
// validator (or you can just use the one provided for you by this package).
func Compare(a, b *pb.IpnsEntry) (int, error) {
	as := a.GetSequence()
	bs := b.GetSequence()

	if as > bs {
		return 1, nil
	} else if as < bs {
		return -1, nil
	}

	at, err := u.ParseRFC3339(string(a.GetValidity()))
	if err != nil {
		return 0, err
	}

	bt, err := u.ParseRFC3339(string(b.GetValidity()))
	if err != nil {
		return 0, err
	}

	if at.After(bt) {
		return 1, nil
	} else if bt.After(at) {
		return -1, nil
	}

	return 0, nil
}

func ipnsEntryDataForSig(e *pb.IpnsEntry) []byte {
	return bytes.Join([][]byte{
		e.Value,
		e.Validity,
		[]byte(fmt.Sprint(e.GetValidityType())),
	},
		[]byte{})
}
```
- ipns包record.go
```
package ipns

import (
	"bytes"
	"errors"

	pb "gx/ipfs/QmNqBhXpBKa5jcjoUZHfxDgAFxtqK3rDA5jtW811GBvVob/go-ipns/pb"

	ic "gx/ipfs/QmPvyPwuCgJ7pDmrKDxRtsScJgBaM5h4EpRL2qQJsmXf4n/go-libp2p-crypto"
	peer "gx/ipfs/QmQsErDt8Qgw1XrsXf2BpEzDgGWtB1YLsTAARBup5b6B9W/go-libp2p-peer"
	logging "gx/ipfs/QmRREK2CAZ5Re2Bd9zZFG6FeYDppUWt5cMgsoUEp3ktgSr/go-log"
	record "gx/ipfs/QmdHb9aBELnQKTVhvvA3hsQbRgUAwsWUzBP2vZ6Y5FBYvE/go-libp2p-record"
	proto "gx/ipfs/QmdxUuburamoF6zF9qjeQC4WYcWGbWuRmdLacMEsW8ioD8/gogo-protobuf/proto"
	pstore "gx/ipfs/QmeKD8YT7887Xu6Z86iZmpYNxrLogJexqxEugSmaf14k64/go-libp2p-peerstore"
)

var log = logging.Logger("ipns")

var _ record.Validator = Validator{}

// 返回ipns路径
// RecordKey returns the libp2p record key for a given peer ID.
func RecordKey(pid peer.ID) string {
	return "/ipns/" + string(pid)
}

// Validator is an IPNS record validator that satisfies the libp2p record
// validator interface.
type Validator struct {
	// KeyBook, if non-nil, will be used to lookup keys for validating IPNS
	// records.
	KeyBook pstore.KeyBook
}

// 验证一个IPNS记录的合法性
// Validate validates an IPNS record.
func (v Validator) Validate(key string, value []byte) error {
	ns, pidString, err := record.SplitKey(key)
	if err != nil || ns != "ipns" {
		return ErrInvalidPath
	}

	// Parse the value into an IpnsEntry
	entry := new(pb.IpnsEntry)
	err = proto.Unmarshal(value, entry)
	if err != nil {
		return ErrBadRecord
	}

	// Get the public key defined by the ipns path
	pid, err := peer.IDFromString(pidString)
	if err != nil {
		log.Debugf("failed to parse ipns record key %s into peer ID", pidString)
		return ErrKeyFormat
	}

	pubk, err := v.getPublicKey(pid, entry)
	if err != nil {
		return err
	}

	return Validate(pubk, entry)
}

// 提取公钥
func (v Validator) getPublicKey(pid peer.ID, entry *pb.IpnsEntry) (ic.PubKey, error) {
	pk, err := ExtractPublicKey(pid, entry)
	if err != nil {
		return nil, err
	}
	if pk != nil {
		return pk, nil
	}

	if v.KeyBook == nil {
		log.Debugf("public key with hash %s not found in IPNS record and no peer store provided", pid)
		return nil, ErrPublicKeyNotFound
	}

	pubk := v.KeyBook.PubKey(pid)
	if pubk == nil {
		log.Debugf("public key with hash %s not found in peer store", pid)
		return nil, ErrPublicKeyNotFound
	}
	return pubk, nil
}

// 查找选出最新的序号、过期时间的ipns记录
// Select selects the best record by checking which has the highest sequence
// number and latest EOL.
//
// This function returns an error if any of the records fail to parse. Validate
// your records first!
func (v Validator) Select(k string, vals [][]byte) (int, error) {
	var recs []*pb.IpnsEntry
	for _, v := range vals {
		e := new(pb.IpnsEntry)
		if err := proto.Unmarshal(v, e); err != nil {
			return -1, err
		}
		recs = append(recs, e)
	}

	return selectRecord(recs, vals)
}

// 有超过1条以上ipns记录，则返回最新的记录
func selectRecord(recs []*pb.IpnsEntry, vals [][]byte) (int, error) {
	switch len(recs) {
	case 0:
		return -1, errors.New("no usable records in given set")
	case 1:
		return 0, nil
	}

	var i int
	for j := 1; j < len(recs); j++ {
		cmp, err := Compare(recs[i], recs[j])
		if err != nil {
			return -1, err
		}
		if cmp == 0 {
			cmp = bytes.Compare(vals[i], vals[j])
		}
		if cmp < 0 {
			i = j
		}
	}

	return i, nil
}
```
