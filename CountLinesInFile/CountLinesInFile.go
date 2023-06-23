/*

次の要件を満たすために、GetPageTitle関数を実装してください。

GetPageTitle関数は、URLを引数として受け取り、そのURLに対応するWebページのタイトルを返します。
タイトルの取得には、net/httpパッケージを使用してHTTP GETリクエストを送信し、<title>タグの内容を抽出します。
タイトルが見つからない場合や、リクエストが失敗した場合には、エラーメッセージを返します。

*/

package CountLinesInFile

import (
	"bufio"
	"fmt"
	"os"
)

func CountLinesInFile(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %w", err)
	}

	return lineCount, nil
}
