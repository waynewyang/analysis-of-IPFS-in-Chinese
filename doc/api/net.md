## NETWORK COMMANDS

### id            Show info about IPFS peers


### bootstrap     Add or remove bootstrap peers
- ipfs bootstrap add
- ipfs bootstrap add default
- ipfs bootstrap list
- ipfs bootstrap rm
- ipfs bootstrap rm all

### swarm         Manage connections to the p2p network
>在ipfs网络中，swarm是一个组件，它用于打开、监听和维持与一个节点的连接。Swarm命令就是用来操作swarm组件的。

- ipfs swarm addrs 
>列出已知地址，debug时很有用  子命令local和listen用于列出本地和监听的地址。

- ipfs swarm addrs listen
- ipfs swarm addrs local
- ipfs swarm connect
- ipfs swarm disconnect
> 断开连接不是永久的，可随时重新连接。

- ipfs swarm ficidrlters 
> 设置过滤地址

- ipfs swarm filters add
- ipfs swarm filters rm
- ipfs swarm peers
> 列出已连接上的节点

### dht           Query the DHT for values or peers
- ipfs dht 
> 操作dht表

- ipfs dht findpeer
- ipfs dht findprovs
> 给定特定key的情况下，在DHT中查找可以提供特定值的节点。

- ipfs dht get
- ipfs dht provide
> 广播宣称，自己拥有这个key

- ipfs dht put  
>向dht表中增 加一个key-value对

- ipfs dht query  
> 查询与待查询节点最接近的peer

### ping          Measure the latency of a connection

### diag?          Print diagnostics
- ipfs diag cmds
- ipfs diag cmds clear
- ipfs diag cmds set-time
- ipfs diag sys

### bitswap
- ipfs bitswap
- ipfs bitswap ledger  与特定节点的交换账本
- ipfs bitswap reprovide  触发 reprovider,重新广播
- ipfs bitswap stat   交换状态
- ipfs bitswap wantlist   请求列表

### pubsub
- ipfs pubsub   //daemon，开发中
- ipfs pubsub ls
- ipfs pubsub peers
- ipfs pubsub pub
- ipfs pubsub sub
