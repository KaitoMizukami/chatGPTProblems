/*

あなたは、与えられたURLからWebページのタイトルを取得するプログラムを作成する必要があります。

次の要件を満たすために、GetPageTitle関数を実装してください。

GetPageTitle関数は、URLを引数として受け取り、そのURLに対応するWebページのタイトルを返します。
タイトルの取得には、net/httpパッケージを使用してHTTP GETリクエストを送信し、<title>タグの内容を抽出します。
タイトルが見つからない場合や、リクエストが失敗した場合には、エラーメッセージを返します。

*/

package GetPageTitle

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// GetPageTitle は指定されたURLのWebページのタイトルを取得します。
func GetPageTitle(url string) (string, error) {
	// HTTP GETリクエストを送信してWebページのコンテンツを取得します。
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to send HTTP GET request: %w", err)
	}
	defer resp.Body.Close()

	// レスポンスのステータスコードをチェックします。
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// レスポンスのコンテンツをHTMLとして解析します。
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// タイトルを抽出します。
	title := extractTitle(doc)
	if title == "" {
		return "", errors.New("page title not found")
	}

	return title, nil
}

// extractTitle はHTMLドキュメントからタイトルを抽出します。
func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := extractTitle(c)
		if result != "" {
			return result
		}
	}

	return ""
}
