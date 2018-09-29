# multiformat

- **特点**
>多格式项目是一系列协议，旨在面向未来的系统。这样做主要是通过自我描述来增强格式值。这允许互操作性，协议敏捷性
>用固定的格式来表达不同的对象

- **Self-description**
>自描述的，数据本身和它的价值评估策略必须区分开来，数据本身必须是自描述的，除了安全上的考虑，数据本身不能强加任何限制策略。只有这样，整个数据存储架构才能适应多方面的需求。

- **参考网址**

>[官网](https://multiformats.io/)
>[github](https://github.com/multiformats)

multiformat类型| 解释 | code |举例
- |- |- |- 
multibase |编解码| [multibase.csv](https://github.com/multiformats/multibase/blob/master/multibase.csv )|在CID v1中，multibase是最外层的编码，默认为z，base58btc 
multicodec | 数据格式|[multicodec.csv](https://github.com/multiformats/multicodec/blob/master/table.csv)  |在CID v1中，multicodec默认为0x55，raw类型
multihash | 哈希方法| [multihash.csv](https://github.com/multiformats/multihash/blob/master/hashtable.csv) |cid中默认的哈希方法，代码0x12，sha2-256
multiaddr |将表达网络地址的域名、ip、端口、协议编号等全部统一格式、路径（比如http/s）| 见下面表格|ipfs 网络使用的均为multiaddr
multiaddr-dns |解析 /dns4, /dns6, and /dnsaddr multiaddrs. |  NA|见下表
multiaddr-net |在标准net库基础上基于multiaddr封装一层接口 | NA |
multistream |一个conn对应一个stream，在连接之上封装 |  string类型|[使用示例](https://gitlab.com/waynewyang/ipfs/tree/master/multistream)
~~multigram~~|||

***
### multiaddr Code表
multiaddr类型 | code
- |
P_IP4|0x0004
P_TCP|0x0006
P_UDP|0x0111
P_DCCP|0x0021
P_IP6|0x0029
P_IP6ZONE|0x002A
P_QUIC|0x01CC
P_SCTP |0x0084
P_UDT| 0x012D
P_UTP |0x012E
P_UNIX |0x0190
P_P2P|0x01A5
P_IPFS|0x01A5
P_HTTP|0x01E0
P_HTTPS|0x01BB
P_ONION|0x01BC
dns4|54
dns6|55
dnsaddr|56

### multiaddr-dns示例
```
>madns /dnsaddr/ipfs.io/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/ip4/104.236.151.122/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/ip6/2604:a880:1:20::1d9:6001/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/ip6/fc3d:9a4e:3c96:2fd2:1afa:18fe:8dd2:b602/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/dns4/jupiter.i.ipfs.io/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
/dns6/jupiter.i.ipfs.io/tcp/4001/ipfs/QmSoLju6m7xTh3DuokvT3886QRYqxAzb1kShaanJgW36yx
```