/*
  - Package wappalyzer
    @Author: zhizhuo
    @IDE：GoLand
    @File: wappalyzer.go
    @Date: 2025/4/17 下午1:59*
*/
package wappalyzer

import (
	"fmt"
	wappalyzer "github.com/projectdiscovery/wappalyzergo"
)

// Wappalyzer 技术识别器结构体
type Wappalyzer struct {
	client *wappalyzer.Wappalyze
}

// TypeWappalyzer 存储网站技术栈信息的结构体
type TypeWappalyzer struct {
	WebServers           []string `json:"web_servers"`            //WEB服务器
	ReverseProxies       []string `json:"reverse_proxies"`        //代理服务器
	JavaScriptFrameworks []string `json:"java_script_frameworks"` //JS框架
	JavaScriptLibraries  []string `json:"java_script_libraries"`  //JavaScript库
	WebFrameworks        []string `json:"web_frameworks"`         //WEB框架
	StaticSiteGenerator  []string `json:"static_site_generator"`  //静态站点生成器
	ProgrammingLanguages []string `json:"programming_languages"`  //开发语言
	Caching              []string `json:"caching"`                //站点缓存
	Security             []string `json:"security"`               //站点安全
	HostingPanels        []string `json:"hosting_panels"`         //主机面板
	Other                []string `json:"other"`                  //其他杂项
}

// NewWappalyzer 创建一个新的Wappalyzer实例
func NewWappalyzer() (*Wappalyzer, error) {
	client, err := wappalyzer.New()
	if err != nil {
		return nil, fmt.Errorf("初始化Wappalyzer失败: %w", err)
	}

	return &Wappalyzer{
		client: client,
	}, nil
}

// FormatData 将wappalyzer原始数据格式化为自定义结构
func (w *Wappalyzer) FormatData(data map[string]wappalyzer.AppInfo) *TypeWappalyzer {
	var result TypeWappalyzer

	// 创建类别映射表，简化分类逻辑
	categoryMap := map[string]*[]string{
		"Web servers":           &result.WebServers,
		"Web frameworks":        &result.WebFrameworks,
		"JavaScript frameworks": &result.JavaScriptFrameworks,
		"JavaScript libraries":  &result.JavaScriptLibraries,
		"Miscellaneous":         &result.Other,
		"Programming languages": &result.ProgrammingLanguages,
		"Security":              &result.Security,
		"Hosting panels":        &result.HostingPanels,
		"Caching":               &result.Caching,
		"Reverse proxies":       &result.ReverseProxies,
		"Static site generator": &result.StaticSiteGenerator,
	}

	// 遍历所有找到的技术
	for techName, info := range data {
		for _, category := range info.Categories {
			// 如果类别在映射表中存在，则添加技术名称到对应切片
			if slice, exists := categoryMap[category]; exists {
				*slice = append(*slice, techName)
			}
		}
	}

	return &result
}

// ConvertHeaders 将单值HTTP头转换为数组形式
func ConvertHeaders(headers map[string]string) map[string][]string {
	result := make(map[string][]string)
	for k, v := range headers {
		result[k] = []string{v}
	}
	return result
}

// GetWappalyzer 分析HTTP响应头和响应体，识别网站使用的技术栈
func (w *Wappalyzer) GetWappalyzer(respHeader map[string][]string, respData []byte) (*TypeWappalyzer, error) {
	if w == nil || w.client == nil {
		return nil, fmt.Errorf("wappalyzer实例未正确初始化")
	}
	fingerprintsWithCats := w.client.FingerprintWithInfo(respHeader, respData)
	return w.FormatData(fingerprintsWithCats), nil
}

// GetWappalyzerWithStringHeaders 针对单值HTTP头的便捷方法
func (w *Wappalyzer) GetWappalyzerWithStringHeaders(respHeader map[string]string, respData []byte) (*TypeWappalyzer, error) {
	// 将 map[string]string 转换为 map[string][]string
	headers := ConvertHeaders(respHeader)
	return w.GetWappalyzer(headers, respData)
}

// Analyze 是GetWappalyzer的别名方法，提供更简洁的调用方式
func (w *Wappalyzer) Analyze(respHeader map[string][]string, respData []byte) (*TypeWappalyzer, error) {
	return w.GetWappalyzer(respHeader, respData)
}
