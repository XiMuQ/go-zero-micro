package utils

import (
	"strconv"
	"strings"
)

// Substring 截取指定位置前的字符串
func Substring(str string, pos int) string {
	s1 := []rune(str)
	newStr := string(s1[:pos])
	return newStr
}

// ArrayToString 将数组转为字符串，sql使用
func ArrayToString(array []int64) string {
	strArray := make([]string, len(array))
	for i, data := range array {
		strArray[i] = strconv.FormatInt(data, 10)
	}
	return strings.Join(strArray, ",")
}

// SplitStringToInt 分割字符串到两个数字
// e. "11-22" => 11,22
func SplitStringToInt(from string) (int, int) {
	if from == "" {
		return 0, 0
	}

	var v1 int
	var v2 int

	parts := strings.SplitN(from, "-", 2)
	if len(parts) == 1 {
		v1, _ = strconv.Atoi(parts[0])
	} else {
		v1, _ = strconv.Atoi(parts[0])
		v2, _ = strconv.Atoi(parts[1])
	}

	return v1, v2
}

// Contains 判断是否包含 不区分大小写
func Contains(s, substr string) bool {
	//统一转为小写进行比较
	s = strings.ToLower(s)
	substr = strings.ToLower(substr)
	return strings.Contains(s, substr)
}
