package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}

// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg: msg,
	}
}

// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func (g Greeter) Greet() Message {
	return g.Message
}

func startGin() {
	//engine := gin.New()
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.String(200, "hello gin")
	})
	engine.Run(":9000")
}

// 使用wire前
func beforeWire() {
	message := NewMessage("hello world")
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}

// 使用wire后
//func afterWire() {
//	event := InitializeEvent("hello_world")
//	event.Start()
//
//}

func main() {
	startGin()
	beforeWire()
}
