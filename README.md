# IPFS
## 目录
- [语言基础](#语言基础)
- [协议总览](#协议总览)
- [对外接口](#对外接口)
- [基础模块分析](#基础模块分析)
- [协议层分析](#协议层分析)
- [应用层数据结构](#应用层数据结构)
- [IPFS-Cluster](#ipfs-cluster)
- [协议当前对于我们来说存在的主要问题](#协议当前对于我们来说存在的主要问题)
- [我们的应用规划](#我们的应用规划)

### 语言基础
- [x] [go依赖包编译注意](/doc/compile.md)
- [x] [context用法](/go-basic/context)
- [x] [空struct用法](/go-basic/nullstruct)
- [x] [interface理解](/go-basic/interface)

- [回到目录](#目录)

### 协议总览
- [x] [IPFS项目进度]
	- [pm记录](https://github.com/ipfs/pm)
	- ![](https://img.shields.io/badge/status-wip-orange.svg?style=flat-square) - 规范制定中，不成熟。
	- ![](https://img.shields.io/badge/status-draft-yellow.svg?style=flat-square) - 草稿完成，很有可能做较大变更。
	- ![](https://img.shields.io/badge/status-reliable-green.svg?style=flat-square) -接近稳定，只会做较小变更。
	- ![](https://img.shields.io/badge/status-stable-brightgreen.svg?style=flat-square) - 规范稳定，本质不变，可能做小的改善。
	- ![](https://img.shields.io/badge/status-permanent-blue.svg?style=flat-square) - 完全稳定，永久不变。
- [x] [协议总览](https://github.com/ipfs/specs/tree/master/architecture)
![](/doc/image/mdag.waist.png) 

![](/doc/image/ip.waist.png)

- [x] [国内研究参考]
	- [x] [mornmist team](https://github.com/mornmist/Newcomer-Guide)
	- [x] [mornmist ipfs](https://github.com/mornmist/IPFS-For-Chinese)
- [x] [libp2p笔记](https://gitlab.com/waynewyang/libp2p)

- [回到目录](#目录)

### 对外接口
- [x] 参考
	- [x] [参考](https://ipfs.docs.apiary.io)
	- [x] [public-api](https://github.com/ipfs/specs/tree/master/public-api)
	- [ ] [js  interface-ipfs-core](https://github.com/ipfs/interface-ipfs-core)
- [x] API笔记
	- [x] [basic](/doc/api/basic.md)
	- [x] [data structure](/doc/api/datastructure.md)
	- [ ] [advance](/doc/api/adv.md)
	- [x] [network](/doc/api/net.md)
	- [x] [tool](/doc/api/tool.md)
	- [ ] [others](/doc/api/others.md) 

- [回到目录](#目录)

### libp2p接口使用
- [ ] 转发使用
	- [x] [命令](/doc/network.md)
	- [ ] 编码

- [回到目录](#目录)

### 基础模块分析
- [x] [multiformats](/doc/multiformat.md)
  - [x] [mafmt](/multiaddrs/mafmt)
  - [x] [addr-util](/multiaddrs/addr-util)
  - [x] [maddr-filter](/multiaddrs/maddr-filter)
  - [x] [multiaddr-filter](/multiaddrs/multiaddr-filter)
- [x] [cid](/doc/cid.md)
- [x] repo
	- [x] [参考 ](https://github.com/ipfs/specs/tree/master/repo) 
		- 文档未更新 
	- [x] [datastores](/datastores/README.md)
		- [x] [笔记](datastores/note.md)
		- [x] [示例](/datastores/example.md)
	- [ ] [细节源码分析](/datastores/源码分析.md)
- [ ] 其他
	- [ ] [protobuf](/protobuf)
	- [ ] [semver](/semver)
	- [ ] [IPFS数据导入导出/DEX](https://github.com/ipfs/specs/tree/master/dex)

- [回到目录](#目录)

### IPFS主流程源码概要分析
- [ ] ipfs init
- [ ] ipfs daemon
- [ ] ipfs add
- [ ] ipfs get
- [ ] cancel

- [回到目录](#目录)

### 协议层分析
- [ ] network
	- [x] [pnet](/go-libp2p-interface-pnet)
	- [ ] relay
	- [ ] 穿透
- [ ] routing
	- [x] DHT寻址原理
	- [x] 源码中修改K桶 a并发与k值提高效率
- [x] block exchange
	- [x] [block](/ipld/block)
	- [x] [总览](https://github.com/ipfs/specs/tree/master/bitswap)
	- [x] [笔记](/bitswap/bitswap)
	- [x] [接口](/bitswap/exchange)
	- [ ] [细节实现] 
- [x] merkledag
	- [x] DAG数据结构
		cid -> block -> node ->dag
	- [x] [ipld-format](/ipld/ipld-format)
	- [x] [ipld-cbor](/ipld/ipld-cbor)
	- [x] [go-merkledag](/ipld/go-merkledag)
	- [x] [dag-link数目大小限制](/ipld/dag-link数目大小限制)
	- [ ] 规范设计中的项目概览
		- [ ] [CAR](https://github.com/ipld/specs/blob/master/CAR.md)
- [x] [namesys](/ipns)
	- [x] [参考其他解释](https://www.jianshu.com/p/04d3e3cc9f1c)
	- [x] [go-ipfs-namesys](/ipns/go-ipfs-namesys)
		- [x] [dns](/ipns/go-ipfs-namesys/dns) 
		- [x] [go-ipns](/ipns/go-ipns)
		- [ ] [publish等其他细节代码分析]

- [回到目录](#目录)

### 应用层数据结构
- [x] [unixfs](/unixfs)
	- [x] mount
		- [x] 可使用mount操作将ipfs文件系统以只读的方式挂载在unix文件系统之上进行操作
		- [x] [mount 示例](/unixfs/mount.md)
	- [x] [files(mfs)参考](https://github.com/ipfs/interface-ipfs-core/blob/master/SPEC/FILES.md#mutable-file-system)
    	- [x] [操作示例](/unixfs/files_op.md) 
	- [ ] 其他
	- [ ] 源码分析

- [回到目录](#目录)

### IPFS-Cluster
- [x] 参考
	- [x] [源码](https://github.com/ipfs/ipfs-cluster)
	- [x] [文档](https://cluster.ipfs.io/documentation/)
- [x] [笔记](/ipfs-cluster)
	- [x] [使用示例](/ipfs-cluster/example)
	- [ ] 源码分析

- [回到目录](#目录)

### 协议当前对于我们来说存在的主要问题
- 只能指定单目录存储，如果FILECOIN不做处理，硬盘的换盘会比较麻烦
- libP2P穿透差

- [回到目录](#目录)

### 我们的应用规划
- [ ] 先河IPFS网盘
- [ ] 同步盘
- [ ] 共享圈
- [ ] 分布式鉴权云存储
- [ ] 分布式聊天

- [回到目录](#目录)