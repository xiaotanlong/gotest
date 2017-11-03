package StealMessage

import (
	"flag"
	"os"
	"log"
	"time"
	"fmt"
)
var logF = flag.String("log","StealMessage" + time.Now().Format("2006-01-02")+".log","Log File Name")
//log test.log Log File Name
func Logger(info string) {

	flag.Parse()//解析参数交付给logF
	outfile,err := os.OpenFile(*logF,os.O_CREATE|os.O_RDWR|os.O_APPEND,0666)
	if err != nil{
		fmt.Println("======================")
		fmt.Println(err)
		fmt.Println("======================")
	}
	log.SetOutput(outfile)//设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)//设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf(":%v \n", info) //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

}