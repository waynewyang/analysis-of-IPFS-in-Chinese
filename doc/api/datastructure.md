## DATA STRUCTURE COMMANDS

### block         
>Interact with raw blocks in the datastore
>操作raw blocks

- ipfs block get
- ipfs block put
- ipfs block rm
- ipfs block stat

### object        
>Interact with raw dag nodes

- ipfs object data
>输出ipfs对象的裸数据

- ipfs object diff 
>显示两个ipfs对象的差异

- ipfs object get 
>读取名称为`<key>`的DAG节点并进行序列化

- ipfs object links 
>输出指定对象指向的链接

- ipfs object new 
>使用给定的ipfs模版创建一个新的对象
```
wayne@wayne:~$ ipfs object new
QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n
wayne@wayne:~$ ipfs object get QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n
{"Links":[],"Data":""}
```

- ipfs object patch
> 基于现有对象创建一个新的DAG对象

    1. ipfs object patch add-link
    2. ipfs object patch append-data
    3. ipfs object patch rm-link
    4. ipfs object patch set-data

- ipfs object put 
>将输入保存为DAG对象，并输出显示生成的密钥

- ipfs object stat 
>读取名称为`<key>`的DAG节点旳统计信息

### files         
>Interact with objects as if they were a unix filesystem
>IPFS内存文件系统操作集
- ipfs file ls
- ipfs files chcid
- ipfs files cp
- ipfs files flush
- ipfs files ls
- ipfs files mkdir
- ipfs files mv
- ipfs files read
- ipfs files rm
- ipfs files stat
- ipfs files write

### dag           
>Interact with IPLD documents (experimental)

- ipfs dag get  
>获取json格式的dag，与ipfs ls区别，更低层次，ipfs ls是其一个特例

- ipfs dag put
- ipfs dag resolve
