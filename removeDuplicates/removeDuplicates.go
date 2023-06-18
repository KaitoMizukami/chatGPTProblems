/*

与えられた文字列から重複している文字を取り除いた新しい文字列を作成する関数 removeDuplicates を作成してください。
関数のシグネチャは以下の通りです。

func removeDuplicates(s string) string {
    // ここにコードを追加してください

}

*/

package removeDuplicates

import "strings"

// func removeDuplicates(s string) string {
// 	var result string
// 	for _, ch := range s {
// 		if !strings.Contains(result, string(ch)) {
// 			result += string(ch)
// 		}
// 	}

// 	return result
// }

func removeDuplicates(s string) string {
	seen := make(map[rune]bool)

	// https://pkg.go.dev/strings#Builder
	// https://qiita.com/po3rin/items/2e406645e0b64e0339d3
	var builder strings.Builder

	for _, ch := range s {
		if !seen[ch] {
			seen[ch] = true
			builder.WriteRune(ch)
		}
	}

	return builder.String()
}
