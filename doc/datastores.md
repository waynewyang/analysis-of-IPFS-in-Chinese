# datastore（待分析完善）

## 参考
[fsrepo](https://github.com/ipfs/go-ipfs/blob/master/repo/fsrepo/fsrepo.go)

## 原始块block存储
- [flatfs](https://github.com/ipfs/go-ds-flatfs)
- 路径 /path/to/ipfsrepo/blocks
1. linux默认路径~/.ipfs/blocks
2. 存储的最终block格式待分析？

## 键值存储
- [leveldb](https://github.com/ipfs/go-ds-leveldb)
- 存储的具体内容待分析？

- 非init阶段存储

```key:/providers/AFKREIAIVK25IDCFTUY3P3UAKRVX3PZJBBNHEGDTMBTDZ3VUOI34FGRJ64/CIQFGA2QEH2DE7WXGKGLRVSX4Y23Y6L2VNYE33HJZCB6EENNFSCMO3Q, value::鞖檬櫵呇*```

- init阶段存储
```寰幆閬嶅巻鏁版嵁
key:/F5UXA3TTF4JCBMC4HJRBNOMOVBZRBQSWFONFZ3KEXTW2UIMVANT5L4JCBWLWIAIN, value:
(/ipns/ 癨:b箮╯耉+歕鞤柬??g振"
梔
?
4/ipfs/QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn€??>r諅JJnapM?茂ZZ
頓喓s#?淇?"嘈捕席_G.穏?峗?D叡€縶|?/@b(驠O慖?擗凄4D?龄J呋[）sJ纥q觅Iv闝匹s闹?嶾箸??L叟a窘劤Q捦軗1u?7Т蛋dV譕CV?Ш9剙料黃?5l颻*/?灋H缶霣?溏??俘腪??s?禉Z箵?[j氨
4巣+錁閲膻}惆??@厱焱鼞湬??? "2018-09-15T07:34:30.276935115Z( :? ?0?"0
	*咹嗺
 ? 0?
? 掸魣-襯~trQZ?
8毱j╈j衈喆|癚?})复@z?m?? MW殡┑,6\耜>樜0鼴鹥掏?`?3?擁??疺k?&?6煴J?唑?菿??搚(\伅Z+>W?偨邭瘹炧O帼
刐?た8颫糺CQ認褷令5???枭衔?~揩欣?蠾Ys|浮X謈;??踦y聱鞥爟曠暅_<V礟?L"KRM嵁d"8?A錄橹峌輦p'赍}践J? 
key:/F5YGWLYSECYFYOTCC24Y5KDTCDBFMK42LTWUJPHNVIQZKA3H2XYSEDMXMQAQ2, value:
&/pk/ 癨:b箮╯耉+歕鞤柬??g振"
梔
? ?0?"0
	*咹嗺
 ? 0?
? 掸魣-襯~trQZ?
8毱j╈j衈喆|癚?})复@z?m?? MW殡┑,6\耜>樜0鼴鹥掏?`?3?擁??疺k?&?6煴J?唑?菿??搚(\伅Z+>W?偨邭瘹炧O帼
刐?た8颫糺CQ認褷令5???枭衔?~揩欣?蠾Ys|浮X謈;??踦y聱鞥爟曠暅_<V礟?L"KRM嵁d"8?A錄橹峌輦p'赍}践J? 
key:/ipns/CIQLAXB2MILLTDVIOMIMEVRLTJOO2RF45WVCDFIDM7K7CIQNS5SACDI, value:
4/ipfs/QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn€??>r諅JJnapM?茂ZZ
頓喓s#?淇?"嘈捕席_G.穏?峗?D叡€縶|?/@b(驠O慖?擗凄4D?龄J呋[）sJ纥q觅Iv闝匹s闹?嶾箸??L叟a窘劤Q捦軗1u?7Т蛋dV譕CV?Ш9剙料黃?5l颻*/?灋H缶霣?溏??俘腪??s?禉Z箵?[j氨
4巣+錁閲膻}惆??@厱焱鼞湬??? "2018-09-15T07:34:30.276935115Z( 
key:/local/filesroot, value: Y攧9_)a烎€斯2綬舖櫯杒e?9饦伙
key:/local/pins, value: Χ!覭嬘#Nゝ睯峦5S??L?熦vG?
```