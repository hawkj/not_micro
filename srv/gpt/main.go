package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiKey = "YOUR_API_KEY"

func main() {
	// 构造请求的数据
	reqData := map[string]interface{}{
		"prompt":      "What is the meaning of life?",
		"max_tokens":  10,
		"temperature": 0.5,
		"n":           1,
		"stop":        "\n",
	}

	// 转换请求数据为 JSON 格式
	reqJson, err := json.Marshal(reqData)
	if err != nil {
		fmt.Println("Failed to marshal request data: ", err)
		return
	}

	// 构造 API 请求
	apiUrl := "https://api.openai.com/v1/engines/davinci-codex/completions"
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(reqJson))
	if err != nil {
		fmt.Println("Failed to create API request: ", err)
		return
	}

	// 设置 API 访问密钥
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// 发送 API 请求
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send API request: ", err)
		return
	}
	defer resp.Body.Close()

	// 解析 API 响应
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read API response: ", err)
		return
	}

	var respData map[string]interface{}
	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		fmt.Println("Failed to unmarshal API response: ", err)
		return
	}

	// 提取生成的答案
	answer := respData["choices"].([]interface{})[0].(map[string]interface{})["text"].(string)

	fmt.Println("Answer:", answer)
}
