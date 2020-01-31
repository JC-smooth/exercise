package main

import (
	"github.com/streadway/amqp"
	"log"
	"strconv"
)

// 用于报错处理
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}

func main(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 打开Channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明Queue
	q, err := ch.QueueDeclare(
		"test", // 队列名
		false, // 持久化
		false,   // 不用时是否删除队列
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
		)
	failOnError(err, "Failed to declare a queue")

	// // 需要发送的文本
	// body := "hello"
	for i := 1; i <= 10; i++ {
		body := strconv.Itoa(i)
		// 发布消息
		err = ch.Publish(
			"",  // 交换机名称
			q.Name,  // 路由名称
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType:     "text/plain",  // 文本格式
				Body:            []byte(body),  // 内容
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}

}