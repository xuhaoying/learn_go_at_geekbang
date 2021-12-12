package main

import (
	"fmt"
	"learn_go_at_geekbang/week02/src"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	id := rand.Intn(10) // id为10以内随机
	dao := src.NewDAO(src.NewDB())
	controller := src.NewController(dao)
	username, err := controller.FindUserNameByID(id)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
	fmt.Printf("User[%v]'s name is [%v]\n", id, username)
}
