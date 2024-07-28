package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// 解析常见返回值格式
func parseApiResponseCommen(resp *http.Response) (map[string]interface{}, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var apiResponse ApiResponseCommon
	err := json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	if apiResponse.Code != 0 {
		return nil, fmt.Errorf("post response error: %s", apiResponse.Message)
	}

	return apiResponse.Data, nil
}

// PostApiResponseData 向api发送 application/x-www-form-urlencoded 格式的 POST 请求, 并获取返回值
func PostApiResponseData(client *http.Client, cookie string, apiUrl string, paramsMap map[string]interface{}) (map[string]interface{}, error) {
	params := url.Values{}
	for key, value := range paramsMap {
		params.Set(key, fmt.Sprintf("%v", value))
	}

	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := parseApiResponseCommen(resp)
	if err != nil {
		return nil, err
	}

	return data, nil
}
