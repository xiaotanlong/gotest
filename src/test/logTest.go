package main

import (
	"flag"
	"os"
	"fmt"
	"log"
)

var logF = flag.String("log","test.log","Log File Name");
func main() {
	flag.Parse()//解析参数交付给logF
	outfile,err := os.OpenFile(*logF,os.O_CREATE|os.O_RDWR|os.O_APPEND,0666);
	if err != nil{
		fmt.Println(*outfile,"open failed")
	}
	log.SetOutput(outfile)//设置log的输出文件，不设置log输出默认为stdout
	log.SetFlags(log.Ldate|log.Ltime|log.Llongfile)//设置答应日志每一行前的标志信息，这里设置了日期，打印时间，当前go文件的文件名

	//write log
	log.Printf("test out:%v \n", "test log") //向日志文件打印日志，可以看到在你设置的输出文件中有输出内容了

}
