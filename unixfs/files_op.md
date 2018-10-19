# 目录
- [mfs操作](#mfs操作)
- [发布ipns](#发布ipns)

## mfs操作
```
wayne@wayne:~$ mkdir test
wayne@wayne:~$ cd test/
wayne@wayne:~/test$ echo aa > a
wayne@wayne:~/test$ ls
a
wayne@wayne:~/test$ echo bb > b
wayne@wayne:~/test$ cat b
bb
wayne@wayne:~/test$ cat a
aa
wayne@wayne:~/test$ ls
a  b
wayne@wayne:~/test$ cd ..
wayne@wayne:~$ ipfs add test -r
added QmQafgrzPNCmLSp32XDeZsbTYpu24zUNAqhdGoqxh6W96W test/a
added QmfFP7dbBUzgs1f5YLaTE4B6QHW5ntQt6wXLwoH3obH4Sp test/b
added QmQSBQxfhGUaRWqaAEgJJt1FBFWaNwjwQzz8nKTuDHvGUs test
 6 B / 6 B [===========================================================================================] 100.00%
 
wayne@wayne:~$ ipfs files mkdir /mfs
wayne@wayne:~$ ipfs files chcid /mfs
wayne@wayne:~$ ipfs files cp /ipfs/QmQSBQxfhGUaRWqaAEgJJt1FBFWaNwjwQzz8nKTuDHvGUs /mfs/root 
wayne@wayne:~$ ipfs file ls QmQSBQxfhGUaRWqaAEgJJt1FBFWaNwjwQzz8nKTuDHvGUs
a
b
wayne@wayne:~$ ipfs files read /mfs/root/a
aa
wayne@wayne:~$ ipfs files read /mfs/root/b
bb

wayne@wayne:~$ echo "hello world" | ipfs files write --create /mfs/root/c
wayne@wayne:~$ ipfs files read /mfs/root/c
hello world
wayne@wayne:~$ ipfs file ls QmQSBQxfhGUaRWqaAEgJJt1FBFWaNwjwQzz8nKTuDHvGUs       
a
b

wayne@wayne:~$ ipfs files ls /mfs/root/  
a
b
c


wayne@wayne:~$ ipfs files flush /mfs/root/
wayne@wayne:~$ ipfs file ls QmQSBQxfhGUaRWqaAEgJJt1FBFWaNwjwQzz8nKTuDHvGUs
a
b
wayne@wayne:~$ ipfs files ls /mfs/root/  
a
b
c
wayne@wayne:~$ ipfs files stat /mfs/root/
QmUcuL8qVARvjMrCDhLwcXkJmkYwoDoVR5hmjB2ipq8TrY
Size: 0
CumulativeSize: 225
ChildBlocks: 3
Type: directory
wayne@wayne:~$ ipfs files rm /mfs/rootc
Error: file does not exist
wayne@wayne:~$ ipfs files rm /mfs/root/c
wayne@wayne:~$ ipfs files stat /mfs/root/  
QmQSBQxfhGUaRWqaAEgJJt1FBFWaNwjwQzz8nKTuDHvGUs
Size: 0
CumulativeSize: 112
ChildBlocks: 2
Type: directory
wayne@wayne:~$ echo "hello world" | ipfs files write --create /mfs/root/c
wayne@wayne:~$ ipfs files stat /mfs/root
QmUcuL8qVARvjMrCDhLwcXkJmkYwoDoVR5hmjB2ipq8TrY
Size: 0
CumulativeSize: 225
ChildBlocks: 3
Type: directory
wayne@wayne:~$ ipfs file ls QmQSBQxfhGUaRWqaAEgJJt1FBFWaNwjwQzz8nKTuDHvGUs
a
b
wayne@wayne:~$ ipfs file ls QmUcuL8qVARvjMrCDhLwcXkJmkYwoDoVR5hmjB2ipq8TrY
a
b
c
wayne@wayne:~$ ipfs files stat /mfs/root/a
QmQafgrzPNCmLSp32XDeZsbTYpu24zUNAqhdGoqxh6W96W
Size: 3
CumulativeSize: 11
ChildBlocks: 0
Type: file
wayne@wayne:~$ ipfs files stat /mfs/root/b
QmfFP7dbBUzgs1f5YLaTE4B6QHW5ntQt6wXLwoH3obH4Sp
Size: 3
CumulativeSize: 11
ChildBlocks: 0
Type: file
```
## 发布ipns

```
wayne@wayne:~/ipfs$ ipfs files stat /mfs/root
QmUcuL8qVARvjMrCDhLwcXkJmkYwoDoVR5hmjB2ipq8TrY
Size: 0
CumulativeSize: 225
ChildBlocks: 3
Type: directory
wayne@wayne:~/ipfs$  ipfs name publish --key=arsyuntestkey /ipfs/QmUcuL8qVARvjMrCDhLwcXkJmkYwoDoVR5hmjB2ipq8TrY
Published to QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF: /ipfs/QmUcuL8qVARvjMrCDhLwcXkJmkYwoDoVR5hmjB2ipq8TrY
wayne@wayne:~/ipfs$ ipfs get /ipns/QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF
Saving file(s) to QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF
 225 B / 225 B [====================================================================================] 100.00% 0s
wayne@wayne:~/ipfs$ cd QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF/
wayne@wayne:~/ipfs/QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF$ ls
a  b  c
wayne@wayne:~/ipfs/QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF$ cat a
aa
wayne@wayne:~/ipfs/QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF$ cat b
bb
wayne@wayne:~/ipfs/QmQYVKeRkbE9aE21ysH1qUQJyQ6nHKvFKh98Kg8wfwRaYF$ cat c
hello world
```
