package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

const (
	ledCnt = 144
)

func main() {
	fmt.Println("connecting...")
	c, err := net.Dial("tcp", "192.168.219.147:8888")
	chk(err)
	defer c.Close()

	fillOneColor(c, 0xff0000)
	time.Sleep(time.Second)
	fillOneColor(c, 0x00ff00)
	time.Sleep(time.Second)
	fillOneColor(c, 0x0000ff)
	time.Sleep(time.Second)
	// 2 turns
	for i := 0; i < ledCnt*2; i++ {
		fillRainbow(c, i)
		time.Sleep(50 * time.Millisecond)
	}
	turnOff(c)
	log.Println("all done")
}

func fillRainbow(c net.Conn, start int) {
	// log.Printf("fillRainbow: start %d", start)
	clear := false
	m := NeoPixel{
		Clear: &clear,
	}
	for i := 0; i < ledCnt; i++ {
		idx := uint32((i + start) % ledCnt)
		pixelHue := float64(i) / ledCnt // TODO: rotate rainbow
		hsv := &HSV{
			H: pixelHue,
			S: 1,
			V: 1,
		}
		rgb := colorToUint32(hsv)
		led := LED{
			Index: &idx,
			Color: &rgb,
		}
		m.Strip = append(m.Strip, &led)
	}
	b, err := proto.Marshal(&m)
	chk(err)
	sendPb(c, b)
}

func fillOneColor(c net.Conn, color uint32) {
	log.Printf("fillOneColor, 0x%08x", color)
	turnOff(c)

	clear := false
	for i := 0; i < ledCnt; i++ {
		m := NeoPixel{
			Clear: &clear,
		}
		index := uint32(i)
		color := uint32(color)
		led := LED{
			Index: &index,
			Color: &color,
		}
		m.Strip = append(m.Strip, &led)
		b, err := proto.Marshal(&m)
		chk(err)
		sendPb(c, b)
		time.Sleep(10 * time.Millisecond)
	}
}

func turnOff(c net.Conn) {
	clear := true
	m := NeoPixel{
		Clear: &clear,
	}
	b, err := proto.Marshal(&m)
	chk(err)
	sendPb(c, b)
}

func sendPb(c net.Conn, b []byte) {
	lengthBuff := make([]byte, 4)
	binary.LittleEndian.PutUint32(lengthBuff, uint32(len(b)))

	// log.Println(lengthBuff, b)
	c.Write(lengthBuff)
	c.Write(b)
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
