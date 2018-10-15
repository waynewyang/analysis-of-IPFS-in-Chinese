## 格式
###  修饰符
-  required : 　不可以增加或删除的字段，必须初始化；
-  optional : 　 可选字段，可删除，可以不初始化；
-  repeated : 　可重复字段， 对应到java文件里，生成的是List。

### 编码规则
> protobuf之所以小且快，就是因为使用变长的编码规则，只保存有用的信息，节省了大量空间。

-  Base-128变长编码
	- 每个字节使用低7位表示数字，除了最后一个字节，其他字节的最高位都设置为1；
	- 采用Little-Endian字节序。（小端表示，地址增长顺序与值增长顺序相同）

-  ZigZag
	- 左移1位（n << 1），无论正数还是负数，相当于乘以2；对于正数，若大于Integer.MAX_VALUE/2（1076741823），则会发生溢出，导致左移1位后为负数
	- 右移31位（n >> 31），对于正数，则返回0x00000000；对于负数，则返回0xffffffff

### IPLD的protobuf定义示例
```
 18 // An IPFS MerkleDAG Link
 19 message PBLink {
 20 
 21   // multihash of the target object
 22   optional bytes Hash = 1;
 23     
 24   // utf string name. should be unique per object
 25   optional string Name = 2;
 26     
 27   // cumulative size of target object 
 28   optional uint64 Tsize = 3;
 29 }
 30 
 31 // An IPFS MerkleDAG Node
 32 message PBNode {
 33 
 34   // refs to other objects
 35   repeated PBLink Links = 2;
 36 
 37   // opaque user data
 38   optional bytes Data = 1;
 39 }
 ```
