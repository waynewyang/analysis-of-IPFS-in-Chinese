package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"

	ms "github.com/multiformats/go-multistream"
)

// 一个conn对应一个stream
// This example creates a multistream muxer, adds handlers for the protocols
// "/cats" and "/docs" and exposes it on a localhost:8765. It then opens connections
// to that port, selects the protocols and tests that the handlers are working.
func main() {
	//实例化一个multistreammuter
	mux := ms.NewMultistreamMuxer()

	//注册流协议为cats的钩子函数。
	mux.AddHandler("/cats", func(proto string, rwc io.ReadWriteCloser) error {
		//往rwc内写入
		fmt.Fprintln(rwc, proto, ": HELLO I LIKE CATS")
		return rwc.Close()
	})
	//注册流协议为dogs的钩子函数。
	mux.AddHandler("/dogs", func(proto string, rwc io.ReadWriteCloser) error {
		//往rwc内写入
		fmt.Fprintln(rwc, proto, ": HELLO I LIKE DOGS")
		return rwc.Close()
	})
	//侦听8765端口
	list, err := net.Listen("tcp", ":8765")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			//等待连接，其中con实现了 io.ReadWriteCloser接口
			con, err := list.Accept()
			if err != nil {
				panic(err)
			}

			//通过conn传递过来协议字符串（协议），这里做协商,并会执行钩子函数
			//1 协商过程，con串与之前注册的mux比对，是否mux有对应的处理服务
			//2 如果有，执行钩子函数，可以看到其上的钩子函数就是将协议号及固定打印，写入con中
			go mux.Handle(con)
		}
	}()

	// The Muxer is ready, let's test it
	conn, err := net.Dial("tcp", ":8765")
	if err != nil {
		panic(err)
	}

	// Create a new multistream to talk to the muxer
	// which will negotiate that we want to talk with /cats
	// 建立一条stream,自动去协商
	mstream := ms.NewMSSelect(conn, "/cats")
	cats, err := ioutil.ReadAll(mstream)
	if err != nil {
		panic(err)
	}
	fmt.Printf("1%s", cats)
	mstream.Close()

	// A different way of talking to the muxer
	// is to manually selecting the protocol ourselves
	conn, err = net.Dial("tcp", ":8765")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	err = ms.SelectProtoOrFail("/dogs", conn)
	if err != nil {
		panic(err)
	}
	dogs, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("2%s", dogs)
	conn.Close()
}
