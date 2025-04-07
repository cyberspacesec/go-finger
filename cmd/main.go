/*
  - Package main
    @Author: zhizhuo
    @IDE：GoLand
    @File: main.go
    @Date: 2025/3/10 下午2:11*
*/
package main

import (
	"fmt"
	"gxx/cmd/cli"
	"gxx/utils/logger"
	"gxx/utils/output"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	color.Green(cli.Banner)
	options, err := cli.NewCmdOptions()
	if err != nil {
		// 在初始化logger之前的错误使用默认logger
		logger.Error(err.Error())
		os.Exit(1)
	}

	// 配置日志级别
	logLevel := 1 // 默认INFO级别
	if options.Debug {
		logLevel = 4 // DEBUG级别
		color.Blue("DEBUG模式已开启")
	}
	// 初始化日志系统，设置日志保存目录、最大文件数和日志级别
	logger.InitLogger("logs", 5, logLevel, options.NoFileLog)
	
	// 如果禁用了文件日志，显示通知
	if options.NoFileLog {
		color.Blue("文件日志记录功能已禁用")
	}
	
	// 配置输出文件啊
	if options.Output == "" {
		options.Output = "result_" + fmt.Sprintf("%d", time.Now().Unix()) + ".txt"
	}

	logger.Info(fmt.Sprintf("输出文件：%s", options.Output))
	
	// 从文件扩展名确定输出格式
	outputFormat := "txt" // 默认为txt格式
	if options.Output != "" {
		ext := strings.ToLower(filepath.Ext(options.Output))
		if ext == ".csv" {
			outputFormat = "csv"
		}
	}
	
	// 初始化输出文件
	if err := output.InitOutput(options.Output, outputFormat); err != nil {
		logger.Error("初始化输出文件失败: %v", err)
		os.Exit(1)
	}
	defer output.Close()

	// 显示开始扫描的信息
	startTime := time.Now()
	fmt.Println(color.CyanString("─────────────────────────────────────────────────────"))
	fmt.Println(color.YellowString(" 开始扫描，请耐心等待..."))
	fmt.Println(color.CyanString("─────────────────────────────────────────────────────"))
	
	// 执行主程序
	cli.Run(options)

	// 显示扫描完成的信息
	elapsed := time.Since(startTime)
	fmt.Println(color.CyanString("─────────────────────────────────────────────────────"))
	fmt.Printf("%s 耗时: %s\n", 
		color.GreenString("扫描完成!"), 
		color.YellowString("%s", elapsed.Round(time.Second)))
	fmt.Printf("结果已保存至: %s\n", color.CyanString(options.Output))
	fmt.Println(color.CyanString("─────────────────────────────────────────────────────"))

	// 确保所有日志都被写入
	time.Sleep(2 * time.Second)
	logger.Success("扫描完成")
}
