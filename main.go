package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("HelloWorld")
	c, err := net.Dial("tcp", "192.168.219.147:8888")
	chk(err)
	defer c.Close()

	clear := true
	m := NeoPixel{
		Clear: &clear,
	}

	for i := 0; i < 144; i++ {
		index := uint32(i)
		color := uint32(0xff4500)
		led := LED{
			Index: &index,
			Color: &color,
		}
		m.Strip = append(m.Strip, &led)
	}

	b, err := proto.Marshal(&m)
	chk(err)
	log.Printf("%d", len(b))

	c.Write(b)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
