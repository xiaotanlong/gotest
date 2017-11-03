package StealMessage

import (
	"net/http"
	"io/ioutil"
)

func HttpGet(url string) (content string,code int) {
	resp,err1 := http.Get(url)
	if err1 != nil {
		code = -100
		return
	}
	defer resp.Body.Close();
	data, err2 := ioutil.ReadAll(resp.Body)
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