/*

問題：
与えられたディレクトリ内の全てのファイル（サブディレクトリ内のファイルを含む）の合計サイズを計算する関数 calculateTotalSize
を作成してください。関数のシグネチャは以下の通りです。

func calculateTotalSize(dirPath string) (int64, error) {
    // ここにコードを追加してください
}

*/

package calculateTotalSize

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func calculateTotalSize(dirPath string) (int64, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return 0, fmt.Errorf("failed to access directory: %v", dirPath)
	}

	if !info.IsDir() {
		return 0, fmt.Errorf("%s is not a directory", dirPath)
	}

	var totalFileSize int64

	err = filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %q: %v", path, err)
		}

		if !info.IsDir() {
			totalFileSize += info.Size()
		}

		return nil
	})
	return totalFileSize, err
}
