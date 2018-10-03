package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	var body string = `
		{
			"name":"ペリドット",
			"japanese_name":"かんらん石",
			"english_name":"peridot",
			"hardness":9.6,
			"price":8000,
			"producing_area":"ハワイ",
			"carat":8.23,
			"weight":19.6,
			"color":"緑",
			"memo":"形状:ラウンド",
			"scratched":true,
			"mining_date":"2017-08-23T00:00:00Z" 
		}
	`
	// bytesに変換
	buf := bytes.NewBufferString(body)

	// HTTP POSTリクエストを送信
	// レスポンスはhttp.Postt関数の戻り値として取得することができます。
	resp, err := http.Post("https://api.tech.gemcook.org/v1/gems", "application/json", buf)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}
