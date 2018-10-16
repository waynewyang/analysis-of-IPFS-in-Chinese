# unixfs
- [x] [官网参考](https://github.com/ipfs/specs/tree/master/unixfs)

## 概述
    unixfs主要目的是将ipfs系统抽象为通用的unix文件系统格式
    unixfs抽象——文件和目录的主要方式是人们设想文件在互联网。在ipfs unixfs是一种数据结构,是ipfs的unix文件。我们需要一个单独的数据结构来携带信息,如
    
    - 对象是否代表一个文件或目录。
    - 总大小,减去索引的开销
