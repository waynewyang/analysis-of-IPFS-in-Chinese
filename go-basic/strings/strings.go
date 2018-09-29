package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	s := strings.TrimRight("/127.0.0.1/4001/tcp/", "/")
	sp := strings.Split(s, "/")
	sp = sp[1:]
	fmt.Println(s)
	fmt.Println(sp)
	fmt.Println(len(sp))
	i := net.ParseIP("1.1.1.1").To4()
	for _, val := range i {
		fmt.Println(val)
	}

	stest := "test.test.tset.tst"
	j := []byte(stest)
	for _, val1 := range j {
		fmt.Printf("%c ", val1)
	}
	z := string(j)
	fmt.Println(z)
}
