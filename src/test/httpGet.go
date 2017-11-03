package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"regexp"
	"strings"
	"os"
)

func main() {
	content,_ := httpGet("https://www.cnblogs.com/news/");
	//fmt.Println(code)
	//fmt.Println(content)

	var (
		//图片正则表达式
		imageItemExp = regexp.MustCompile(`//images2017\.cnblogs\.com/news/[0123456789]*/[0123456789]*/[0123456789/-]*\.jpg`)
		//帖子路径正则表达式
		threadItemExp = regexp.MustCompile(`"//news.cnblogs.com/n/[0123456789]{1,}`)
	)

	//解析正则表达式，返回参数tds是一个数组
	tds := threadItemExp.FindAllStringSubmatch(content, 10000)
	var tdStr = make([]string, 0)
	//fmt.Println(len(tds))
	//去掉引号，并放到一维数组中
	for _, t := range tds {
		var n = strings.Replace(t[0], "\"", "", -1)
		tdStr = append(tdStr, n)
		//fmt.Println(n)
	}

		var threads = make([]ThreadItem, 0)
		//组装帖子结构体
		for _, t := range tdStr {
			threads = append(threads, ThreadItem{url: "http:" + t})
			//fmt.Println("http:" + t)
			content2,_ := httpGet("http:" + t)
			//fmt.Println(code2)
			fmt.Println(content2)
			tds2 := imageItemExp.FindAllStringSubmatch(content2, 10000)
			fmt.Println(tds2)
			/*if len(tds2) > 0 {
				for _,value := range tds2{

					//download("http:" + value[0])
					//fmt.Println("http:" + value[0])
				}
			}*/
		}


}
//下载图片
func download(url string) {
	fmt.Println(url)
	var impPathre = regexp.MustCompile(`[\/]{1,}([0123456789]*-[0123456789]*-[0123456789]*)+\.jpg$`)
	imgResponse, _ := http.Get(url)
	defer imgResponse.Body.Close()
	imgByte, _ := ioutil.ReadAll(imgResponse.Body)

	path := "D:/filee/"
	impPath := impPathre.FindAllStringSubmatch(url, 10000)
	fmt.Println(impPath)
	pInfo, pErr := os.Stat(path)
	if pErr != nil || pInfo.IsDir() == false {
		errDir := os.Mkdir(path, os.ModePerm)


		if errDir != nil {
			fmt.Println(errDir)
			os.Exit(-1)
		}
	}

	fn := path + string(impPath[0][0])
	_, fErr := os.Stat(fn)

	var fh *os.File
	if fErr != nil {
		fh, _ = os.Create(fn)
	} else {
		fh, _ = os.Open(fn)
	}
	defer fh.Close()
	fh.Write(imgByte)
}

func httpGet(url string) (content string,code int) {
	resp,err1 := http.Get(url);//
	if err1 != nil {
		code = -100
		return
	}
	defer resp.Body.Close();
	data, err2 := ioutil.ReadAll(resp.Body);
	if err2 != nil {
		code = -200
		return
	}
	code = resp.StatusCode
	content = string(data)
	return
}

//ThreadItem 帖子数据
type ThreadItem struct {
	url     string   //帖子路径
	content string   //帖子内容
	imgs    []string //帖子图片
}