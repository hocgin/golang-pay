package net

import "testing"

func TestGetUrl(t *testing.T) {
	params := make(map[string]string)
	params["a"] = "1"
	params["b"] = "2"

	url := GetUrl("https://www.baidu.com", params)
	if url != "https://www.baidu.com?a=1&b=2" {
		t.Errorf("测试失败")
	}
}

func TestGetUrlEncode(t *testing.T) {
	params := make(map[string]string)
	params["a"] = "1"
	params["b"] = "你好"

	url := GetUrlEncode("https://www.baidu.com?a=1&b=%E4%BD%A0%E5%A5%BD", params)
	if url != "https://www.baidu.com?a=1&b=2" {
		t.Errorf("测试失败")
	}
}
