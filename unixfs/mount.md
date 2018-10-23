- mount之后ipns空间可以做unix系统正常操作

```
wayne@wayne:/ipns/QmXAMaVdj4YyxfNze8ofhfveDPyUJ4Z4FWPWhPpugRb7RP$ ls
test  test1  wayne
wayne@wayne:/ipns/QmXAMaVdj4YyxfNze8ofhfveDPyUJ4Z4FWPWhPpugRb7RP$ echo  arsyun-test > arsyun-test
wayne@wayne:/ipns/QmXAMaVdj4YyxfNze8ofhfveDPyUJ4Z4FWPWhPpugRb7RP$ ls
arsyun-test  test  test1  wayne
wayne@wayne:/ipns/QmXAMaVdj4YyxfNze8ofhfveDPyUJ4Z4FWPWhPpugRb7RP$ cat arsyun-test 
arsyun-test
```

```
wayne@wayne:~/.ipfs/blocks$ ll -ltr
......
drwxr-xr-x   2 wayne wayne 4096 Oct 23 14:39 5K/
drwxr-xr-x   2 wayne wayne 4096 Oct 23 14:39 NW/
drwxr-xr-x   2 wayne wayne 4096 Oct 23 14:39 LM/
drwxr-xr-x   2 wayne wayne 4096 Oct 23 14:39 MB/

wayne@wayne:~/.ipfs/blocks$ cat NW/CIQI33T24OB6MU5PMMDV7PMLMZP3JSX4J5OJUL6HB2WTVEVZMTYFNWQ.data 
 arsyun-test

wayne@wayne:~/.ipfs/blocks$ cd ~/ipfs/datastores/cid-dskey/
wayne@wayne:~/ipfs/datastores/cid-dskey$ go build cid-dskey.go 
wayne@wayne:~/ipfs/datastores/cid-dskey$ ./cid-dskey  -d CIQI33T24OB6MU5PMMDV7PMLMZP3JSX4J5OJUL6HB2WTVEVZMTYFNWQ
QmXth7D91zk2QjsydWhr2hhzymdbsUnm9XYWWokoPMiNC1
wayne@wayne:~/ipfs/datastores/cid-dskey$ ipfs object get QmXth7D91zk2QjsydWhr2hhzymdbsUnm9XYWWokoPMiNC1
{"Links":[],"Data":"\u0008\u0000\u0012\u000carsyun-test\n\u0018\u000c"}
wayne@wayne:~/ipfs/datastores/cid-dskey$ ipfs cat QmXth7D91zk2QjsydWhr2hhzymdbsUnm9XYWWokoPMiNC1       
arsyun-test
wayne@wayne:~/ipfs/datastores/cid-dskey$ 
```



- mount命令说明
```
wayne@wayne:~/ipfs/unixfs$ ipfs mount --help
USAGE
  ipfs mount - Mounts IPFS to the filesystem (read-only).

SYNOPSIS
  ipfs mount [--ipfs-path=<ipfs-path> | -f] [--ipns-path=<ipns-path> | -n]

OPTIONS

  -f, --ipfs-path string - The path where IPFS should be mounted.
  -n, --ipns-path string - The path where IPNS should be mounted.

DESCRIPTION

  Mount IPFS at a read-only mountpoint on the OS. The default, /ipfs and /ipns,
  are set in the configuration file, but can be overriden by the options.
  All IPFS objects will be accessible under this directory. Note that the
  root will not be listable, as it is virtual. Access known paths directly.

  You may have to create /ipfs and /ipns before using 'ipfs mount':

  > sudo mkdir /ipfs /ipns
  > sudo chown $(whoami) /ipfs /ipns
  > ipfs daemon &
  > ipfs mount

  Example:

  # setup
  > mkdir foo
  > echo "baz" > foo/bar
  > ipfs add -r foo
  added QmWLdkp93sNxGRjnFHPaYg8tCQ35NBY3XPn6KiETd3Z4WR foo/bar
  added QmSh5e7S6fdcu75LAbXNZAFY2nGyZUJXyLCJDvn2zRkWyC foo
  > ipfs ls QmSh5e7S6fdcu75LAbXNZAFY2nGyZUJXyLCJDvn2zRkWyC
  QmWLdkp93sNxGRjnFHPaYg8tCQ35NBY3XPn6KiETd3Z4WR 12 bar
  > ipfs cat QmWLdkp93sNxGRjnFHPaYg8tCQ35NBY3XPn6KiETd3Z4WR
  baz

  # mount
  > ipfs daemon &
  > ipfs mount
  IPFS mounted at: /ipfs
  IPNS mounted at: /ipns
  > cd /ipfs/QmSh5e7S6fdcu75LAbXNZAFY2nGyZUJXyLCJDvn2zRkWyC
  > ls
  bar
  > cat bar
  baz
  > cat /ipfs/QmSh5e7S6fdcu75LAbXNZAFY2nGyZUJXyLCJDvn2zRkWyC/bar
  baz
  > cat /ipfs/QmWLdkp93sNxGRjnFHPaYg8tCQ35NBY3XPn6KiETd3Z4WR
  baz
```



