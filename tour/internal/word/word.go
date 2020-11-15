package word

import (
	"strings"
	"unicode"
)

/**
转成大写
*/
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

/**
转成小写
*/
func ToLower(s string) string {
	return strings.ToLower(s)
}

/**
下划线转大写
*/
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)

	return strings.Replace(s, " ", "", -1)
}

/**
下划线转小写(同UnderscoreToUpperCamelCase, 只是首字母小写了)
*/
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

/**
驼峰转下划线
*/
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		// 如果当前字符是大写, 则追加一个 _ 再转小写
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}
