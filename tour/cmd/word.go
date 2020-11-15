package cmd

import (
	"GoProgrammingJourney/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

const (
	// 全转大写
	ModeUpper = iota + 1
	// 全转小写
	ModeLower
	// 下划线转大写驼峰
	ModeUnderscoreToUpperCamelCase
	// 下划线转小写驼峰
	ModeUnderscoreToLowerCamelCase
	// 驼峰转下划线
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词转换, 模式如下: ",
	"1: 全部单词转为大写",
	"2: 全部单词转为小写",
	"3: 下划线单词转为大写驼峰单词",
	"4: 下划线单词转为小写驼峰单词",
	"5: 驼峰单词转为下划线单词",
}, "\n")

var str string
var mode int8
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式, 请执行help word查看帮助文档")
		}

		log.Printf("输出结果: %s", content)
	},
}

func init(){
	// 1. 需要绑定的变量
	// 2. 接收该参数的完整命令
	// 3. 对应的短标识
	// 4. 默认值
	// 5. 使用说明
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}