package notification

import (
	"fmt"
	"time"

	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
)

// var topic = "go.micro.msg.order"

// Publish : 创建一个发布者, 并每秒钟给主题发送一次信息
func Publish(topic string) {
	// 创建一个每秒钟执行的定时器
	tick := time.NewTicker(time.Second)
	i := 0
	// 定时器开始执行
	for range tick.C {
		// 创建一个消息
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d:%s", i, time.Now().String())),
		}
		// 打印 broker
		log.Info(broker.String())
		// 发布消息
		if err := broker.Publish(topic, msg); err != nil {
			log.Info("[pub] Message publication failed: %v", err)
		} else {
			fmt.Println("[pub] Message published: ", string(msg.Body))
		}
		i++
	}
}