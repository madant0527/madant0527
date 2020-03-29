package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer //声明一个全局的连接kafka的生产者
)

//专门往kafka里写日志的模块
func Init(addrs []string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		fmt.Println("Init kafka err,err:", err)
		return
	}
	return
}

func SendToKafka(topic, data string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(data)
	pid,offset,err:=client.SendMessage(msg)
	fmt.Printf("msg:%v\n",msg.Value)
	if err!=nil{
		fmt.Println("send msg to kafka err,err=",err)
		return
	}
	fmt.Printf("pid=%v,offset=%v,\n",pid,offset)
}


