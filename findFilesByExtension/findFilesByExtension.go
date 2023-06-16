/*

与えられたディレクトリ内の全てのファイル（サブディレクトリ内のファイルを含む）を再帰的に探索し、
指定された拡張子を持つファイルのリストを返す関数 findFilesByExtension を作成してください。
関数のシグネチャは以下の通りです。

func findFilesByExtension(dirPath, extension string) ([]string, error) {
    // ここにコードを追加してください
}

*/

package findFilesByExtension

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func findFilesByExtension(dirPath, extension string) ([]string, error) {
	// dirPathが存在するかのチェック
	info, err := os.Stat(dirPath)
	if err != nil {
		return nil, fmt.Errorf("%v is not exist", dirPath)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%v is not a directory", dirPath)
	}

	var files []string

	err = filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("can not access path: %v", path)
		}
		if !info.IsDir() && filepath.Ext(path) == extension {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
