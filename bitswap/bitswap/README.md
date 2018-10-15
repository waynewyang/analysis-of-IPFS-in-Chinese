# bitswap协议笔记

## 简介
- 基于消息的协议，消息包含:请求block或者响应wantlist
    - 收到wantlist消息，ipfs节点判断是否有该块，并发送
    - 收到blocks，ipfs节点应该发出cancel消息，通知其他节点已经不需要该块
- 重点在于实时性以及内存要求

## 子系统
- [概要图](https://cloud.githubusercontent.com/assets/1211152/21071077/4620387a-be4a-11e6-895c-aa8f2b06aa4e.png)
    请求模块：want-manager
    服务模块：decision engine

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
    调用Bitswap.GetBlock(cid)，将cid添加到wantlist，进行block请求

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

## 实现代码
gx/ipfs/Qmd8rU7X3VZzsgPnf2LSGUFu35zizYKajzXTRuHMUMqYJQ/go-bitswap
