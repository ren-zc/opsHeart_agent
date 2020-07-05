package call_http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"opsHeart_agent/common"
	"time"
)

var C *http.Client

func init() {
	C = &http.Client{
		Timeout: 10 * time.Second,
	}
}

func HttpGet(path string, para string) (int, []byte, error) {
	urlWithPara := fmt.Sprintf("http://%s%s?%s", common.UsedServer, path, para)
	//fmt.Printf("%s", urlWithPara)
	req, _ := http.NewRequest("GET", urlWithPara, nil)

	// set basic auth
	req.SetBasicAuth(common.UUID, common.SelfToken)

	resp, err := C.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, b, nil
}

func HttpPost(path string, para []byte) (int, []byte, error) {
	url := fmt.Sprintf("http://%s%s", common.UsedServer, path)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(para))

	// set basic auth
	req.SetBasicAuth(common.UUID, common.SelfToken)

	resp, err := C.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	return resp.StatusCode, b, nil
}
