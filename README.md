# IPFS分析笔记——Wayne
![](/study.png) 
## 目录
- IPFS协议总览
	- [IPFS定义及参考技术](#ipfs定义及参考技术)
	- [协议总览](#协议总览)
	- [对外接口](#对外接口)
- 基础
	- [语言基础](#语言基础)
	- [基础模块分析](#基础模块分析)
		- multiformat
		- cid
		- repo—看看IPFS持久化的数据是什么样的
- IPFS协议分析
	- [协议层分析](#协议层分析)
	- [IPFS应用层数据结构](#应用层数据结构)
	- [debug方法](#协议debug方法)
- IPFS主要业务逻辑分析
	- [IPFS主流程概要分析](#ipfs主流程概要分析)
		- init、daemon、add、get、cancel
- 应用
	- 官方应用
		- [IPFS-Cluster](#ipfs-cluster)
		- [ipfs-companion](#ipfs-companion)
		- [ipfs-desktop](#ipfs-desktop)
	- 第三方应用
	- 先河IPFS应用
		- [协议当前对于我们来说存在的主要问题](#协议当前对于我们来说存在的主要问题)
		- [我们的应用规划](#我们的应用规划)

### IPFS定义及参考技术
- 定义
	> 白皮书：IPFS是一种内容可寻址、版本化、点对点超媒体的分布式文件系统

	> 官网：IPFS是分布式Web,一种点对点超媒体协议，使网络更快，更安全，更开放。
	
- 参考技术
	- ice协议——网络连通性
	> 解释：解决peers之间的连通性

	- DHT技术，KAD、coral算法——分布式路由
	> 解释：解决peers之间的路由、内容路由（内容寻址）、以及分布式哈希key-value存储
	
	> 重要特征：内容寻址、或者说哈希寻址
	
	- bt——交换技术
	> 解释：解决peers之间的内容交换
	
	> 重要特征：P2P分布式、更快、更开放
	
	- git——数据结构
	> 解释：MerkleDAG借鉴git的数据存储方式，解决数据的防篡改、去重
	
	> 重要特征：防篡改、去重
	
	- sfs——命名系统
	> 解决哈希难以记忆、以及由于DAG特性导致的动态数据变更导致的更新成本问题
	
	> 重要特征：版本化（通过ipns的发布记录）、ipns解析（人类更容易识别操作）、dns解析

	- unixfs系统
	> 特征：unixfs挂载特性，挂载的ipfs类似一个文件系统，可以使用unix方式访问全球哈希。

```
第一步：在PEER A上，在IPFS添加another.txt文件
D:\go-ipfs>ipfs add another.txt
 29 B / 29 B [===================================================
added QmamQ2prV7FTfFm1eJc5o6QRA2BAkUJAbc5JCrLpu9dY2z another.txt

第二步：在peer B上，挂载ipfs文件系统（全球性的），直接通过unix的方式，cat第一步的哈希，可以查询到文件内容
wayne@wayne:~$ ipfs mount
IPFS mounted at: /ipfs
IPNS mounted at: /ipns
wayne@wayne:~/go/src$ cat /ipfs/QmamQ2prV7FTfFm1eJc5o6QRA2BAkUJAbc5JCrLpu9dY2z
"This is new another file" 
```

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
	- MerkleDAG vs IP 瘦腰图
		- ipfs通过MerkleDAG访问资源、http通过IP访问资源
		- 向下提供支撑，向上提供应用

![](/doc/image/mdag.waist.png) 

![](/doc/image/ip.waist.png)

- [x] [国内研究参考]
	- [x] 董天一团队
		- [x] [mornmist team](https://github.com/mornmist/Newcomer-Guide)
		- [x] [mornmist ipfs](https://github.com/mornmist/IPFS-For-Chinese)
		- [x] [《IPFS与区块链：原理与实践》 董天一书籍目录](http://www.ipfstalk.org/?t/197.html)
	- [x] [西二旗李老师-简书](https://www.jianshu.com/u/832753b872c5)

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

### 语言基础
- [x] [go依赖包编译注意](/doc/compile.md)
- [x] [context用法](/go-basic/context)
- [x] [空struct用法](/go-basic/nullstruct)
- [x] [interface理解](/go-basic/interface)

- [回到目录](#目录)

### 基础模块分析
- [x] [multiformats](/doc/multiformat.md)
- [x] 基于multiaddr的一些支撑包（网络层）
  - [x] [mafmt](/multiaddrs/mafmt)
  - [x] [addr-util](/multiaddrs/addr-util)
  - [x] [maddr-filter](/multiaddrs/maddr-filter)
  - [x] [multiaddr-filter](/multiaddrs/multiaddr-filter)
- [x] [cid](/doc/cid.md)
- [x] repo
	- [x] [参考 ](https://github.com/ipfs/specs/tree/master/repo) 
		- 文档未更新 
	- [x] [持久化数据分析](/datastores/README.md)
		- [x] [笔记](datastores/note.md)
		- [x] [示例](/datastores/example.md)
	- [ ] [细节源码分析](/datastores/源码分析.md)
- [ ] 其他
	- [ ] [protobuf](/protobuf)
	- [ ] [semver](/semver)
	- [ ] [IPFS数据导入导出/DEX](https://github.com/ipfs/specs/tree/master/dex)

- [回到目录](#目录)

### IPFS主流程概要分析
- [ ] ipfs init
- [ ] ipfs daemon
- [ ] ipfs add
- [ ] ipfs get
- [ ] cancel

- [回到目录](#目录)

### 协议层分析
- [x] [network](/network)
	- [x] [pnet](/go-libp2p-interface-pnet)
	- [x] [relay](/doc/network.md)
	- [ ] 穿透，待分析，目前端口转发OK，穿透比较差
- [x] [routing](/routing)
	- [x] [DHT寻址原理](/routing/dht.md)
	- [x] [源码中修改K桶 a并发与k值提高效率](/routing/修改并发属.md)
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
- [x] [namesys](/namesys)
	- [x] [dns](/namesys/dns) 
	- [x] [ipns](/namesys/ipns)

- [回到目录](#目录)

### 应用层数据结构
- [x] [unixfs](/unixfs)
	- [x] mount
		- [x] 挂载在unix文件系统之上进行操作
		- [x] [mount 示例](/unixfs/mount.md)
	- [x] [files(mfs)参考](https://github.com/ipfs/interface-ipfs-core/blob/master/SPEC/FILES.md#mutable-file-system)
    	- [x] [操作示例](/unixfs/files_op.md) 
	- [ ] 其他
	- [ ] 源码分析

- [回到目录](#目录)
	
### 协议debug方法
> ipfs log tail
- [回到目录](#目录)

### IPFS-Cluster
- [x] 参考
  - [x] [源码](https://github.com/ipfs/ipfs-cluster)
  - [x] [文档](https://cluster.ipfs.io/documentation/)
- [x] [笔记](/ipfs-cluster)
  - [x] [使用示例](/ipfs-cluster/example)
  - [ ] 源码分析
- [回到目录](#目录)

### ipfs-companion
- [网址](https://github.com/ipfs-shipyard/ipfs-companion)

### ipfs-desktop
- [网址](https://github.com/ipfs-shipyard/ipfs-desktop)

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