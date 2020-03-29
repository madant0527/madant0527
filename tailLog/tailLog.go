package tailLog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

//专门从日志文件中收集日志的模块

var (
	tailObj = new(tail.Tail)
)

func Init(fileName string) (err error) {
	config := tail.Config{

		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}

	tailObj, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("taile file err, err=", err)
		return
	}

	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
	//var (
	//	line *tail.Line
	//	ok bool
	//)
	//for{
	//	line,ok=<-tailObj.Lines
	//	if !ok{
	//		fmt.Printf("tail file close reopen,file name =%v\n",)
	//		time.Sleep(time.Second)
	//		continue
	//	}
	//
	//}
}
