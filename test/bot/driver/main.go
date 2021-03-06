package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"traffic-dispatcher/model"

	"github.com/gorilla/websocket"
)

const (
	driverCount = 3
)

var (
	centerPoint = model.GeoLocation{
		Lng: 116.404,
		Lat: 39.915,
	}
	// apigw
	wsAddr = flag.String("addr", "localhost:8082", "http service address")
)

func oneClient(interrupt chan os.Signal, user model.User, point model.GeoLocation, idx int) {
	u := url.URL{Scheme: "ws", Host: *wsAddr, Path: "/geo/ws/lbs"}
	log.Printf("Client [%d] try to connect to server %s\n", idx, u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial err:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Printf("Read message err:%s\n", err.Error())
				return
			}
			log.Printf("Client [%d] Recevice message: %s\n", idx, message)
		}
	}()

	// 4秒上传一次
	ticker := time.NewTicker(time.Second * 4)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			wsMsg := model.WSMessage{
				Command: model.CmdReportGeo,
				Role:    model.ClientDriver,
				User:    user,
				Geo:     point,
			}
			if msg, err := json.Marshal(wsMsg); err == nil {
				err := c.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Printf("Write message err: %s\n", err.Error())
					return
				}
			} else {
				log.Fatal("Encode json err", err)
			}
		case <-interrupt:
			log.Println("Catch interrupt signal")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Printf("Write channel close: %s\n", err.Error())
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	// 监听和捕获信号量
	signal.Notify(interrupt, os.Interrupt)

	for i := 0; i < driverCount; i++ {
		go oneClient(
			interrupt,
			model.User{
				Name:  fmt.Sprintf("testuser_%d", i),
				UID:   fmt.Sprintf("testuid_%d", i),
				Token: "To_Define",
			},
			model.GeoLocation{
				Lat: centerPoint.Lat + float64(i),
				Lng: centerPoint.Lng + float64(i),
			},
			i)
	}

	for {
		select {
		case <-interrupt:
			return
		}
	}
}
