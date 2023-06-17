/*

与えられたURLがHTTPかHTTPSのいずれかのスキームで始まるかどうかを判定する関数 checkHTTPScheme を作成してください。
関数のシグネチャは以下の通りです。

func checkHTTPScheme(url string) (bool, error) {
    // ここにコードを追加してください

}

*/

package checkHTTPScheme

import (
	"errors"
	"fmt"
	"net/url"
)

func checkHTTPScheme(rowUrl string) (bool, error) {
	if len(rowUrl) == 0 {
		return false, errors.New("URL is empty")
	}

	// URL構造で返してくれる
	// https://pkg.go.dev/net/url@go1.20.5#URL
	url, err := url.Parse(rowUrl)
	if err != nil {
		return false, fmt.Errorf("can not parse url: %v", rowUrl)
	}
	if url.Scheme == "http" || url.Scheme == "https" {
		return true, nil
	} else {
		return false, nil
	}
}
