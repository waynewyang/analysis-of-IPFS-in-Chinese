## IPFS

### GO语言基础
- [x] [go依赖包编译注意](https://gitlab.com/waynewyang/ipfs/blob/master/doc/compile.md)
- [x] [context用法](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/context)
- [x] [空struct用法](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/nullstruct)
- [x] [interface理解](https://gitlab.com/waynewyang/ipfs/tree/master/go-basic/interface)

### IPFS协议总览
- [x] [IPFS项目进度]()
	- [pm记录](https://github.com/ipfs/pm)
	- ![](https://img.shields.io/badge/status-wip-orange.svg?style=flat-square) - 规范制定中，不成熟。
	- ![](https://img.shields.io/badge/status-draft-yellow.svg?style=flat-square) - 草稿完成，很有可能做较大变更。
	- ![](https://img.shields.io/badge/status-reliable-green.svg?style=flat-square) -接近稳定，只会做较小变更。
	- ![](https://img.shields.io/badge/status-stable-brightgreen.svg?style=flat-square) - 规范稳定，本质不变，可能做小的改善。
	- ![](https://img.shields.io/badge/status-permanent-blue.svg?style=flat-square) - 完全稳定，永久不变。
- [x] [协议总览](https://github.com/ipfs/specs/tree/master/architecture)
- [x] [merkledag瘦腰图](https://github.com/ipfs/specs/blob/master/merkledag/mdag.waist.png) 与 [IP瘦腰图](https://github.com/ipfs/specs/blob/master/merkledag/ip.waist.png)
- [x] [国内研究参考]
	- [x] [mornmist team](https://github.com/mornmist/Newcomer-Guide)
	- [x] [mornmist ipfs](https://github.com/mornmist/IPFS-For-Chinese)
- [x] [libp2p笔记](https://gitlab.com/waynewyang/libp2p)

### IPFS API熟悉
- [x] 参考
	- [x] [参考](https://ipfs.docs.apiary.io)
	- [x] [public-api](https://github.com/ipfs/specs/tree/master/public-api)
	- [ ] [js  interface-ipfs-core](https://github.com/ipfs/interface-ipfs-core)
- [x] API笔记
	- [x] [basic](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/basic.md)
	- [x] [data structure](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/datastructure.md)
	- [x] [advance](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/adv.md)
	- [x] [network](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/net.md)
	- [x] [tool](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/tool.md)
	- [x] [others](https://gitlab.com/waynewyang/ipfs/blob/master/doc/api/others.md) 

### libp2p接口使用
- [ ] 转发使用
	- [x] [命令](https://gitlab.com/waynewyang/ipfs/blob/master/doc/network.md)
	- [ ] 编码

### 基础模块分析
- [x] [multiformats](https://gitlab.com/waynewyang/ipfs/blob/master/doc/multiformat.md)
  - [x] [mafmt](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/mafmt)
  - [x] [addr-util](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/addr-util)
  - [x] [maddr-filter](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/maddr-filter)
  - [x] [multiaddr-filter](https://gitlab.com/waynewyang/ipfs/blob/master/multiaddrs/multiaddr-filter)
- [x] [cid](https://gitlab.com/waynewyang/ipfs/blob/master/doc/cid.md)

- [ ] repo
	- [ ] [参考 ](https://github.com/ipfs/specs/tree/master/repo) 
	- [ ] [datastores](https://gitlab.com/waynewyang/ipfs/blob/master/doc/datastores.md)
  		- [ ] leveldb
		- [ ] go-ds-flatfs
- [ ] [protobuf](https://gitlab.com/waynewyang/ipfs/blob/master/protobuf)
  - [ ] io(gogo-protobuf/io)
- [ ] [semver](https://gitlab.com/waynewyang/ipfs/blob/master/semver)
- [ ] 未完成部分
  - [ ] [IPFS数据导入导出/DEX](https://github.com/ipfs/specs/tree/master/dex)

### IPFS主流程源码概要分析
- [ ] ipfs init
- [ ] ipfs daemon
- [ ] ipfs add
- [ ] ipfs get
- [ ] cancel

### IPFS协议层
- [ ] network
	- [x] [pnet](https://gitlab.com/waynewyang/ipfs/blob/master/go-libp2p-interface-pnet)
	- [ ] relay
	- [ ] 穿透
- [ ] routing
	- [x] DHT寻址原理
	- [x] 源码中修改K桶 a并发与k值提高效率
- [x] block exchange
	- [x] [总览](https://github.com/ipfs/specs/tree/master/bitswap)
	- [x] [笔记](https://gitlab.com/waynewyang/ipfs/tree/master/bitswap/bitswap)
	- [x] [接口](https://gitlab.com/waynewyang/ipfs/tree/master/bitswap/exchange)
	- [ ] [细节实现] 
- [x] merkledag
	- [x] [block](https://gitlab.com/waynewyang/ipfs/tree/master/ipld/block)
	- [x] [ipld-format](https://gitlab.com/waynewyang/ipfs/tree/master/ipld/ipld-format)
	- [x] [ipld-cbor](https://gitlab.com/waynewyang/ipfs/tree/master/ipld/ipld-cbor)
	- [x] [go-merkledag](https://gitlab.com/waynewyang/ipfs/tree/master/ipld/go-merkledag)
	- [x] [dag-link数目大小限制](https://gitlab.com/waynewyang/ipfs/tree/master/ipld/dag-link数目大小限制)
	- [ ] 规范设计中的项目概览
		- [ ] [CAR](https://github.com/ipld/specs/blob/master/CAR.md)
- [x] namesys
	- [x] [参考其他解释](https://www.jianshu.com/p/04d3e3cc9f1c)
	- [x] [笔记](https://gitlab.com/waynewyang/ipfs/tree/master/ipns/go-ipns)
	- [x] [go-ipns](https://gitlab.com/waynewyang/ipfs/tree/master/ipns/go-ipns)
	- [x] [go-ipfs-namesys](https://gitlab.com/waynewyang/ipfs/tree/master/ipns/go-ipfs-namesys)
		- [x] [dns](https://gitlab.com/waynewyang/ipfs/tree/master/ipns/go-ipfs-namesys/dns) 
		- [ ] [publish等其他细节代码分析]

### IPFS应用层及数据结构
- [x] [unixfs](https://gitlab.com/waynewyang/ipfs/tree/master/unixfs)
  - [x] 挂载操作
    - [x] 可使用mount操作将ipfs文件系统挂载在unix文件系统之上
    - [x] 在unix系统中，可提供给用户传统文件的操作方式 
    - [x] [file(mfs)参考](https://github.com/ipfs/interface-ipfs-core/blob/master/SPEC/FILES.md#mutable-file-system)
    - [x] [mfs使用](https://gitlab.com/waynewyang/ipfs/blob/master/doc/image/ipfs_files.png) 
  - [ ] 源码分析

### IPFS Cluster
- [ ] [参考](https://github.com/ipfs/ipfs-cluster)
- [ ] 使用

### IPFS&libP2P协议对于我们来说存在的主要问题
- 只能指定单目录存储，如果FILECOIN不做处理，硬盘的换盘会比较麻烦
- libP2P穿透差

### 基于IPFS应用规划
- [ ] 先河网盘
- [ ] 共享圈
- [ ] 分布式鉴权云存储
- [ ] 分布式聊天
- [ ] VIPFS
