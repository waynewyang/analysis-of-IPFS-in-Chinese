# bitswap协议笔记

[参考github规范](https://github.com/ipfs/specs/tree/master/bitswap)

## 简介
- 基于消息的协议，消息包含:请求block或者响应wantlist
    - 收到wantlist消息，ipfs节点判断是否有该块，并发送
    - 收到blocks，ipfs节点应该发出cancel消息，通知其他节点已经不需要该块

## 子系统
![](/bitswap/bitswap.png)
​    请求模块：want-manager，    服务模块：decision engine

- 相关模块
    cid、peer、block
    Message：bitswap消息,包含wantlists与blocks
    Entry：包含于Message中，是一个wantlist实例，包含cid、priority、cancel（设置为unwantlist）
    Ledger:交换账本
```
message Message {
        message Wantlist {
                message Entry {
                        optional string block = 1; // the block key
                        optional int32 priority = 2; // the priority (normalized). default to 1
                        optional bool cancel = 3;  // whether this revokes an entry
                }

                repeated Entry entries = 1; // a list of wantlist entries
                optional bool full = 2;     // whether this is the full wantlist. default to false
        }

        optional Wantlist wantlist = 1;
        repeated bytes blocks = 2;
}

```

- Want-Manager
    调用Bitswap.GetBlock(cid)，将cid添加到wantlist，调用路有层进行block请求

- Decision Engine
    收到Message之后，（wanlist）为每一个cid创建一个task，添加到任务队列中，直至源节点请求的block被压入消息队列，任务完成。
    主要数据结构：peer request queue (PRQ)，

- Message Queue
    Wantlist manager,当发生add/del cid行为时，需要更新所有节点的messagequeue
    Decision engine:当拥有源节点请求的block时，需要对应更新源节点的messagequeue

## 执行细节
- 消息队列的合并去重
- sender
    消息发送
- listener
    消息接收
- event事件
    - bitswap.addedBlock(block)
        - see if any peers want this block, and send it,广播本地变化?
    - bitswap.getBlock(key, cb)
        - add to wantlist,增加请求
        - maybe send wantlist updates to peers
    - bitswap.cancelGet(key),删除请求
        - so that can send wantlist cancels
    - bitswap.receivedMessage(msg)
        - process the wantlist changes
        - process the blocks
    - bitswap.peerConnected(peer)
        - add peer to peer set + send them wantlist (maybe)
    - bitswap.peerDisconnected(peer)
        - remove peer from peer set
## 对外接口
- BitSwapNetwork
```
// BitSwapNetwork provides network connectivity for BitSwap sessions
type BitSwapNetwork interface {

	// SendMessage sends a BitSwap message to a peer.
	//消息包含wantlists与blocks
	// send wantlist 会调用exchange接口的GetBlocks()进行获取块
	// send blocks 响应其他的blocks请求，调用网络层的发送
	SendMessage(
		context.Context,
		peer.ID,
		bsmsg.BitSwapMessage) error

	// SetDelegate registers the Reciver to handle messages received from the
	// network.
	//接收，见type Receiver interface
	SetDelegate(Receiver)

	ConnectTo(context.Context, peer.ID) error

	NewMessageSender(context.Context, peer.ID) (MessageSender, error)

	ConnectionManager() ifconnmgr.ConnManager

	Routing  //内容路由，接收到block的时候会调用路有层的Provide方法
}

type MessageSender interface {
	SendMsg(context.Context, bsmsg.BitSwapMessage) error
	Close() error
	Reset() error
}

// 接收到wantlist的逻辑处理，判断本地是否有，没有的话调用路由层findprovide获取
// 接收到需要的block的时候，cancel wanlist条目
// Implement Receiver to receive messages from the BitSwapNetwork
type Receiver interface {
	ReceiveMessage(
		ctx context.Context,
		sender peer.ID,
		incoming bsmsg.BitSwapMessage)

	ReceiveError(error)

	// Connected/Disconnected warns bitswap about peer connections
	PeerConnected(peer.ID)
	PeerDisconnected(peer.ID)
}

// 最终调用到路由层的接口，接口方法名称一致
type Routing interface {
	// FindProvidersAsync returns a channel of providers for the given key
	FindProvidersAsync(context.Context, *cid.Cid, int) <-chan peer.ID

	// Provide provides the key to the network
	Provide(context.Context, *cid.Cid) error
}
```

- exchange 在bitswap中实现
```
// Package exchange defines the IPFS exchange interface
package exchange

import (
	"context"
	"io"

	blocks "gx/ipfs/QmWAzSEoqZ6xU6pu8yL8e5WaMb7wtbfbhhN4p1DknUPtr3/go-block-format"
	cid "gx/ipfs/QmZFbDTY9jfSBms2MchvYM9oYRbAF19K7Pby47yDBfpPrb/go-cid"
)

// Interface defines the functionality of the IPFS block exchange protocol.
type Interface interface { // type Exchanger interface
	Fetcher

	// TODO Should callers be concerned with whether the block was made
	// available on the network?
	HasBlock(blocks.Block) error

	IsOnline() bool

	io.Closer
}

// Fetcher is an object that can be used to retrieve blocks
type Fetcher interface {
	// GetBlock returns the block associated with a given key.
	GetBlock(context.Context, *cid.Cid) (blocks.Block, error)
	GetBlocks(context.Context, []*cid.Cid) (<-chan blocks.Block, error)
	//调用路由层获取
}

// SessionExchange is an exchange.Interface which supports
// sessions.
type SessionExchange interface {
	Interface
	NewSession(context.Context) Interface
}
```
