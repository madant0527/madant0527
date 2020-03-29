package main

import (
	"fmt"
	"goStudy/go-log-agent/conf"
	"goStudy/go-log-agent/kafka"
	"goStudy/go-log-agent/tailLog"

	"time"
)

func run() {

	for {
		select {
		//读取日志
		case line := <-tailLog.ReadChan():
			//发送到kafka
			kafka.SendToKafka(conf.ConfApp.Kafka.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}


func main()  {
	conf.Init()

	err:= kafka.Init([]string{conf.ConfApp.Addr})
	if err!=nil{
		fmt.Println("init kafka err! err=",err)
		return
	}
	fmt.Println("init kafka success!")
	err2:= tailLog.Init(conf.ConfApp.TailLog.Path)
	if err2!=nil{
		fmt.Println("init tailLog err! err=",err)
		return
	}
	fmt.Println("init tailLog success!")
	run()
}

