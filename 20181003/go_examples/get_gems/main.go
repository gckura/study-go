package main

// fmtパッケージ 標準出力に出力する
// io/ioutilパッケージ 入出力関連の関数が定義されている
// net/httpパッケージ HTTPを扱うパッケージで、HTTPクライアントとHTTPサーバーを実装するために必要な機能が提供されている

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// HTTP GETリクエストを送信
	// レスポンスはhttp.Get関数の戻り値として取得することができます。
	resp, err := http.Get("https://api.tech.gemcook.org/v1/gems")

	if err != nil {
		fmt.Println(err)
		return
	}

	// レスポンスを受け取ったら必ずBodyをCloseするのが決まり
	defer resp.Body.Close()

	// b = buffer のこと
	// レスポンスから内容を取得する(resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(b))
	}
}
