package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stianeikeland/go-rpio"
)

func main() {

	fmt.Println("opening gpio")
	err := rpio.Open() 
	if err != nil {
		panic(fmt.Sprint("unable to open gpio", err.Error()))
	}

	defer rpio.Close()

	c, _, err := websocket.DefaultDialer.Dial("ws://192.168.1.142:8080/count", nil)
	defer c.Close()
	if err != nil {
		log.Fatal("dial:", err)
	}
	white := rpio.Pin(25)
	white.Output()
	white.High()

	for {
		_, message, err := c.ReadMessage()
		white.Low()
		tmr := time.NewTimer(time.Second / 5)
		if err != nil {
			log.Println("read:", err)
			return
		}
		go func() {
			<-tmr.C
			white.High()
		}()
		
		str := string(message)
		log.Println(str)
	}
}
