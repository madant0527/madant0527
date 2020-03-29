package conf
import (
	"fmt"
	"gopkg.in/ini.v1"
)

type ConfigApp struct {
	Kafka   `ini:"kafka"`
	TailLog `ini:"tailLog"`
}

type Kafka struct {
	Addr    string `ini:"addr"`
	Topic string `ini:"topic"`
}

type TailLog struct {
	Path string `ini:"path"`
}

var ConfApp = new(ConfigApp)

func Init() {
	err := ini.StrictMapTo(ConfApp, "./conf/confApp.ini")
	if err != nil {
		fmt.Printf( "ini confApp err, err=%v\n", err)
		return
	}
	fmt.Println("addr=", ConfApp.Kafka.Addr)
	fmt.Println("topic=", ConfApp.Kafka.Topic)
	fmt.Println("path=", ConfApp.TailLog.Path)
}