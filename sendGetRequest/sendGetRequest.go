package sendGetRequest

import (
	"io"
	"net/http"
)

func SendGetRequest(url string) (int, []byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return res.StatusCode, nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, nil, err
	}

	// ここに書くとio.ReadAllでエラー起きたら、Close()が呼び出されないかも
	// defer res.Body.Close()

	return res.StatusCode, body, nil
}
