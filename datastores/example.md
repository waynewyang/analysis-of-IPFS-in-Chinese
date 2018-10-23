# repo资源示例

## blocks存储内容

### ipfs init之后block内容

```
└── CIQN5PVU4ECEKNEVXWRYKCPTVMIPHP5AWWKOCGCPXVNMF7ZWU6UG2HI.data
└── CIQKKLBWAIBQZOIS5X7E32LQAL6236OUKZTMHPQSFIXPWXNZHQOV7JQ.data
└── CIQIVZGLOUFY5L4TEYD5WLSCRDDEAW2TNVZSH3OBM5UKNOREHYSY7RA.data
└── CIQDOZU3EAGXWK3PLVFOFOZOAE5USX3XM6I5CSHSQGTML2BAGN7MB5I.data
└── CIQJFGRQHQ45VCQLM7AJNF2GF5UHUAGGHC6LLAH6VYDEKLQMD4QLILY.data
└── CIQPHMHGQLLZXC32FQQW2YVM4KGFORVFJAQYY55VK3WJGLZ2MS4RJNQ.data
└── CIQBIQXZ4NWWDXUSIYSCX7RE6EBXHMGENZNMUDEMGNKMGT2K6LLUL5Y.data
└── CIQJBQD2O6K4CGJVCCTJNUP57QHR4SKHZ74OIITBBGLOMCO3ZOLWLGA.data
└── CIQDWKPBHXLJ3XVELRJZA2SYY7OGCSX6FRSIZS2VQQPVKOA2Z4VXN2I.data
└── CIQBT4N7PS5IZ5IG2ZOUGKFK27IE33WKGJNDW2TY3LSBNQ34R6OVOOQ.data
└── CIQKNNRB2NFYXUZDJ2UWNMSKYLGTKUYDRQTJCDI7JTUDFH6YOYNUPMA.data
└── CIQOHMGEIKMPYHAUTL57JSEZN64SIJ5OIHSGJG4TJSSJLGI3PBJLQVI.data
└── CIQBED3K6YA5I3QQWLJOCHWXDRK5EXZQILBCKAPEDUJENZ5B5HJ5R3A.data
└── CIQL3XIOKVDAW5KQF6NNWGFFYAHEQP63TJOVZHAEO7XZBD7KQOCSSHY.data
└── CIQFTFEEHEDF6KLBT32BFAGLXEZL4UWFNWM4LFTLMXQBCERZ6CMLX3Y.data

QmdL9t1YP99v4a2wyXFYAQJtbD9zKnPrugFLQWXBXb82sn
QmZTR5bcpQD7cFgTorqxZDYaew1Wqgfbd2ud9QqGPAkK2V  //about
QmXgqKTbzdh83pQtKFb19SpMCpDDcKR2ujqk3pKph9aCNF  //quick-start
QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv  //文档目录
QmYCvbfNbCwFR45HiNP45rwJgvatpiW38D961L5qAhUM5Y  //about
QmejvEPop4D7YUadeGqYWmZxHhLc4JBUCzJJHWMzdcMe2y  //ping
QmPhk6cJkRcFfZCdYam4c9MKYjFG9V29LswUnbrFNhtk2S
QmY5heUM5qgRubMDD1og9fhCPA6QdkMp3QCwd4s7gJsyE7  //help
QmSKboVigcD3AY4kLsob117KJcMHvMUu6vNFqk1PQzYUpp  
QmQ5vhrL7uv6tuoN9KeVBwd4PwfQkXdVVmDLUZuTNxqgvm  //security-notes
QmZZRTyhDpL5Jgift1cHbAhexeE1m2Hw8x8g7rTcPahDvo  // 包含2S pp
QmdfTbBqBPQ7VNxZEYEj14VmRuZBkqFbiwReogJgS1zR1n  //空文件
QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB  //readme
Qmb7oGTxge7amSArtJsGUAqswY8y1G7m5QNjV57Nj5sEHU  // 包含2sn pp
QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn  //空目录

wayne@wayne:~/ipfs/datastores$ ipfs object get QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv
{"Links":[
{"Name":"about","Hash":"QmZTR5bcpQD7cFgTorqxZDYaew1Wqgfbd2ud9QqGPAkK2V","Size":1688},
{"Name":"contact","Hash":"QmYCvbfNbCwFR45HiNP45rwJgvatpiW38D961L5qAhUM5Y","Size":200},
{"Name":"help","Hash":"QmY5heUM5qgRubMDD1og9fhCPA6QdkMp3QCwd4s7gJsyE7","Size":322},
{"Name":"ping","Hash":"QmejvEPop4D7YUadeGqYWmZxHhLc4JBUCzJJHWMzdcMe2y","Size":12},
{"Name":"quick-start","Hash":"QmXgqKTbzdh83pQtKFb19SpMCpDDcKR2ujqk3pKph9aCNF","Size":1692},
{"Name":"readme","Hash":"QmPZ9gcCEpqKTo6aq61g2nXGUhM4iCL3ewB6LDXZCtioEB","Size":1102},
{"Name":"security-notes","Hash":"QmQ5vhrL7uv6tuoN9KeVBwd4PwfQkXdVVmDLUZuTNxqgvm","Size":1173}],
"Data":"\u0008\u0001"}
```


### ipfs add之后block内容

```
wayne@wayne:~/.ipfs/blocks$ ipfs add ~/testarsyun
added QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP testarsyun

CIQAPMCNS6C7TQDW4YDOPH7P72BOIIZRLHUCTQHHLXFEF4LXDMS5HHA ： QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP
CIQCHWY75CY6PSCMWOTWMIXNIANHSDUAJTWMGTU36WKW4NNXQ24DRKQ ： QmQkcmqfxQALC5rowikNwZKUhqEZvFc2GSizDBvq41Y59K 目录的link
	wayne@wayne:~/ipfs/datastores$ ipfs object get QmQkcmqfxQALC5rowikNwZKUhqEZvFc2GSizDBvq41Y59K 
	{"Links":[{"Name":"testarsyun","Hash":"QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP","Size":21}],"Data":"\u0008\u0001"}
	
	wayne@wayne:~$ cp testarsyun testtest/
	wayne@wayne:~$ ipfs add testtest/ -r
	added QmNrfVBkdiwh3GSYxdts4HrVRqzPDruLNBjGweQK8UKHGP testtest/testarsyun
	added QmQkcmqfxQALC5rowikNwZKUhqEZvFc2GSizDBvq41Y59K testtest

CIQIUWBCIFKG6NJBFKDYU62A6NI3ZSUGYMKRKH7L4V3UAUMBCSRZIAQ ： QmXegvBnh3bzX8otHR8VX3ZPKbDetMyeaGCfVwgrw5FUQH
CIQHUWN6N5NP7LW6ZESQPVF6G76HVRM6I4FMGFKRWGOZ6POVW27SR6I :  QmWaFpS4sFUSjSVm44PB4PvfmqJCqtoBeTEei7C6YBAFSp
```

### ipfs daemon启动之前，是可以对本地资源进行add get

## leveldb

### daemon之后的内容dht记录
dht记录  /provides/内容key/peer的key，value：time、ttl
```
key:/providers/CIQPH6HPVAE2NE7DTFJ7WZHKI47TLSC6VC7GNSRZIPQHZ5CQTPCKFNY/CIQMMSWJFI7AZM2NZPA2BDYVQQWNXJYPTIZ2EMJ342RHHVENJPF4SII,value:嘄樬韹踹*
key:/providers/CIQPHEPB4ETGOPSLL7JJTS24TUDB2MBZ5S3M4XT5KKAJH563YEXEGDI/CIQOOYDVBQWUZKBAF5GFVJWFX6UMNYQJ4YREBLY5AAZJGOZQATXYXQI,value:娜崯嚊踹*
key:/providers/CIQPMCYWJAWAOLKPKLX6EO3YMMKX5AC4SODCSBWG6TXA7KXASTCKW2Q/CIQOOYDVBQWUZKBAF5GFVJWFX6UMNYQJ4YREBLY5AAZJGOZQATXYXQI,value:湥姛嚊踹*
key:/providers/CIQPMQL3DUV65LJIBPZLYTPEBNTDIEEDHMAG5A7U3DBX6K7ZHXJFNBQ/CIQK7OG7SAI66E3UNOXL27QNCWCHIDMLBDGQQNCKN5VC5MR2LPMETPI,value:炥泴踹*
key:/providers/CIQPRSWQIECOAMYK23SIFWZ23FXB3BWLGAFPVIKBVUDXT7VIOLBCE5Y/CIQOOYDVBQWUZKBAF5GFVJWFX6UMNYQJ4YREBLY5AAZJGOZQATXYXQI,value:垱霝嚊踹*
key:/providers/CIQPXCBJJC2CTOXJN5X5ZSXCMSKA2RIXH37GNHBGMUDDNJ7GDMWG2VA/CIQOOYDVBQWUZKBAF5GFVJWFX6UMNYQJ4YREBLY5AAZJGOZQATXYXQI,value:蛴儫嚊踹*
```

```
key：
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQPH6HPVAE2NE7DTFJ7WZHKI47TLSC6VC7GNSRZIPQHZ5CQTPCKFNY
Qmem1wW5FygVnqPiAhs8LFVcEg64GgNF19KYsqKFbQrwrn
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQPHEPB4ETGOPSLL7JJTS24TUDB2MBZ5S3M4XT5KKAJH563YEXEGDI
QmejSoJgZ5w8W8V6MtCcAVjG8FSHQEuk1Sn9RxnBSPpm7i
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQPMCYWJAWAOLKPKLX6EO3YMMKX5AC4SODCSBWG6TXA7KXASTCKW2Q
Qmeu6orTBmucUtp7Y3fKFKm9Y1S4Aupv7S97A9oTJsg4tR
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQPMQL3DUV65LJIBPZLYTPEBNTDIEEDHMAG5A7U3DBX6K7ZHXJFNBQ
Qmeuvv1mmVGjycAAAyiB2rprPPw5vCpZiEzTohMZccK2Vj
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQPRSWQIECOAMYK23SIFWZ23FXB3BWLGAFPVIKBVUDXT7VIOLBCE5Y
Qmf5qBrE1Sr6fNdEyWt158oozVAPpAQQ5tuvs6g8t3eKDp
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQPXCBJJC2CTOXJN5X5ZSXCMSKA2RIXH37GNHBGMUDDNJ7GDMWG2VA
QmfGXTq6Vc5tcJPbrT4v79CniNzu2YhtD75yMCvgz5nNdH

peer：
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d  CIQMMSWJFI7AZM2NZPA2BDYVQQWNXJYPTIZ2EMJ342RHHVENJPF4SII
QmbghaHL1Zj1ViGCr5ScyYV8PNGS97Ge7eVMRvsbLYjDoe
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQOOYDVBQWUZKBAF5GFVJWFX6UMNYQJ4YREBLY5AAZJGOZQATXYXQI
QmdurCG977gsZyVQi9p2nbZeVzu93jH6TUukkzQERCrFS8
wayne@wayne:~/ipfs/datastores$ ./cidblock  -d CIQK7OG7SAI66E3UNOXL27QNCWCHIDMLBDGQQNCKN5VC5MR2LPMETPI
QmaAbZQVFavUub1cvsXP8tfbk7p5i2cRVvimWv5La2E9U8
```

### ipns 记录：
- 发布记录
```
wayne@wayne:~/ipfs/datastores$ ipfs key gen --type=rsa --size=2048 arsyuntestkey
Qmcph9BKFc6jgXTvTGdipyjYpXMwEpH4jZwkz9t5sPvCbF
wayne@wayne:~/ipfs/datastores$  ipfs name publish --key=arsyuntestkey /ipfs/QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv
Published to Qmcph9BKFc6jgXTvTGdipyjYpXMwEpH4jZwkz9t5sPvCbF: /ipfs/QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uV
```

- leveldb记录：包含一个默认的ipns节点以及一个新增的arsyuntestkey发布的节点
```
key:/ipns/CIQNOMVBHL7GBYW2FHXLJVN25GQ2HUEFLODC77HB7ETORFHQUX6XTNQ,value:
4/ipfs/QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv€}皖?h?篼圿碪d/V?6匍,Y 蟂蠺M?笧倮垨鉁N貼j葮%9?7制?垴竷n髾<???愲嘕Cz祭?栆诸诟c惧豎熹?螎u???	轊濬q$蛉?/埦1e敖?认?c婓	?[Zf極雗Aq( u#尶悧韑?e?	7*f釔e偪?1J鯼菱玒聓,=哋纕Z乀輖?>摎骤p?NW7皕咵q8垘%鰄5俋i?雡uc殹靲紱o>? "2018-10-24T01:14:47.665651388Z( 
key:/ipns/CIQOOYDVBQWUZKBAF5GFVJWFX6UMNYQJ4YREBLY5AAZJGOZQATXYXQI,value:
4/ipfs/QmUNLLsPACCz1vLxQVkXqqLX5R1X345qqfHbsf67hvA3Nn€!2厾XCCf嫷譺	./J?q>=撍秡?Nh砞X冇肘4埒罶獭唚?嗽XU屼?脓=#┬_氃?e谅;??8??樳炧|[DY??C盏'襌?鐋?题綁?匦3贎鳤Y?C鹯?旟凂"?&3?	咔欼瞮?戕烇5.(錬/齦d鱛槾璷虪氘w祣鴿i?堲浐笌颩_鴐I鹂傖#晩8?瘹k搞谇⒊滶0/?Y懧?贿e? "2018-10-23T23:17:08.427618356Z( 
```