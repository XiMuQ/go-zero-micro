package utils

import (
	"fmt"
	"time"
)

const (
	//YearFormat 年份格式化模板
	YearFormat string = "2006"

	//MonthFormat 月份格式化模板
	MonthFormat string = "2006-01"

	//DateFormat 日期格式化模板
	DateFormat string = "2006-01-02"

	//DateHourFormat 日期小时格式化模板
	DateHourFormat string = "2006-01-02 15"

	//DateMinFormat 日期小时格式化模板
	DateMinFormat string = "2006-01-02 15:04"

	//DateTimeFormat 日期时间格式化模板
	DateTimeFormat string = "2006-01-02 15:04:05"

	//DateTimeFormatString 日期时间字符串格式化模板
	DateTimeFormatString string = "Mon, 02 Jan 2006 15:04:05 MST"

	DateTimeFormatString2 string = "Mon, 02-Jan-2006 15:04:05 MST"

	DateTimeFormatString3 string = "02 Jan 2006 15:04:05 MST"
)

// BeforeTime 获取指定数值前的时间 dateType：second、minute、hour、day、month、year
func BeforeTime(curTime string, dateType string, interval int) string {
	beforeTime := ""
	layout := "2006-01-02 15:04:05"
	interval = -interval
	var before time.Time
	start, _ := time.ParseInLocation(layout, curTime, time.Local)
	if dateType == "second" {
		before = start.Add(time.Duration(interval) * time.Second)
	} else if dateType == "minute" {
		before = start.Add(time.Duration(interval*60) * time.Second)
	} else if dateType == "hour" {
		before = start.Add(time.Duration(interval*60*60) * time.Second)
	} else if dateType == "day" {
		before = start.AddDate(0, 0, interval)
	} else if dateType == "month" {
		before = start.AddDate(0, interval, 0)
	} else if dateType == "year" {
		before = start.AddDate(interval, 0, 0)
	}
	beforeTime = before.Format(layout)
	return beforeTime
}

// TimeList 根据开始时间和结束时间获取时间段集合：dateType：second、minute、hour、day、month、year
func TimeList(startDate string, endDate string, dateType string, interval int) []string {
	// layout 日期时间格式
	layout := DateTimeFormat
	endTime, _ := time.ParseInLocation(layout, endDate, time.Local)
	startTime, _ := time.ParseInLocation(layout, startDate, time.Local)

	array := make([]string, 0)
	if dateType == "second" {
		for endTime.After(startTime) {
			array = append(array, startTime.Format(layout))
			startTime = startTime.Add(time.Duration(interval) * time.Second)
		}
		array = append(array, endDate)
	} else if dateType == "minute" {
		for endTime.After(startTime) {
			array = append(array, startTime.Format(layout))
			startTime = startTime.Add(time.Duration(interval*60) * time.Second)
		}
		end := endTime.Format(layout)
		if array[len(array)-1] != end {
			array = append(array, end)
		}
	} else if dateType == "hour" {
		for endTime.After(startTime) {
			array = append(array, startTime.Format(layout))
			startTime = startTime.Add(time.Duration(interval*60*60) * time.Second)
		}
		end := endTime.Format(DateMinFormat)
		if array[len(array)-1] != end {
			array = append(array, end)
		}
	} else if dateType == "day" {
		for endTime.After(startTime) {
			date := Substring(startTime.Format(layout), 10)
			array = append(array, date)
			startTime = startTime.AddDate(0, 0, interval)
		}
		end := Substring(endTime.Format(layout), 10)
		if array[len(array)-1] != end {
			array = append(array, end)
		}
	} else if dateType == "month" {
		for endTime.After(startTime) {
			date := Substring(startTime.Format(layout), 7)
			array = append(array, date)
			startTime = startTime.AddDate(0, 1, 0)
		}
		end := Substring(endTime.Format(layout), 7)
		if array[len(array)-1] != end {
			array = append(array, end)
		}
	} else if dateType == "year" {
		for endTime.After(startTime) {
			date := Substring(startTime.Format(layout), 4)
			array = append(array, date)
			startTime = startTime.AddDate(interval, 0, 0)
		}
		end := Substring(endTime.Format(layout), 4)
		if array[len(array)-1] != end {
			array = append(array, end)
		}
	}
	return array
}

// TimeStampToString 时间戳转字符串
func TimeStampToString(longtime int64) string {
	//time1 := time.Unix(longtime, 0).Add(time.Duration(-8*60*60) * time.Second)
	//timeStr := time1.Format(DateTimeFormat)
	//return timeStr
	timeStr := time.Unix(longtime, 0).Format(DateTimeFormat)
	return timeStr
}

// TimeStampToTime 时间戳转时间
func TimeStampToTime(longtime int64) time.Time {
	return time.Unix(longtime, 0)
}

// TimeDiff 时间差：单位 小时/分钟
func TimeDiff(startTime, endTime, timeType int64) int64 {
	// 将时间戳转换为时间类型
	time1 := time.Unix(startTime, 0)
	time2 := time.Unix(endTime, 0)
	// 计算时间差
	duration := time2.Sub(time1)
	// 获取小时数
	hours := int64(duration.Hours()) //默认是小时
	if timeType > 0 {
		hours = int64(duration.Minutes()) //分钟数
	}
	return hours
}

// TimeStampToWeekday 时间戳转星期几
func TimeStampToWeekday(timeStamp int64) int64 {
	t := time.Unix(timeStamp, 0)
	// 获取星期几
	weekday := t.Weekday()
	var whichDay int64 = 7 //默认星期日
	if int64(weekday) != 0 {
		whichDay = int64(weekday)
	}
	return whichDay
}

// TimeStampAtRange 判断时间戳是否在指定时间范围内
func TimeStampAtRange(timeStamp int64, startHour, startMin, endHour, endMin int) bool {
	captureTime := time.Unix(timeStamp, 0)

	isAtRange := false
	if endHour >= startHour { //同一天
		startTime, endTime := GetTime(captureTime, startHour, startMin, endHour, endMin)
		// 判断当前时间是否在工作时间范围内
		if captureTime.After(startTime) && captureTime.Before(endTime) {
			isAtRange = true
		}
	} else { //非同一天时两个时间段
		startHour1 := 00
		startMin1 := 00
		endHour1 := endHour
		endMin1 := endMin

		startHour2 := startHour
		startMin2 := startMin
		endHour2 := 23
		endMin2 := 59

		startTime1, endTime1 := GetTime(captureTime, startHour1, startMin1, endHour1, endMin1)
		startTime2, endTime2 := GetTime(captureTime, startHour2, startMin2, endHour2, endMin2)

		// 判断当前时间是否在工作时间范围内
		if captureTime.After(startTime1) && captureTime.Before(endTime1) {
			isAtRange = true
		}
		// 判断当前时间是否在工作时间范围内
		if captureTime.After(startTime2) && captureTime.Before(endTime2) {
			isAtRange = true
		}
	}
	return isAtRange
}

// GetTime 获取探针捕获当天所在的开始时间和结束时间
func GetTime(captureTime time.Time, startHour, startMin, endHour, endMin int) (startTime, endTime time.Time) {
	// 构建开始时间和结束时间
	startTime = time.Date(captureTime.Year(), captureTime.Month(), captureTime.Day(), startHour, startMin, 0, 0, captureTime.Location())
	endTime = time.Date(captureTime.Year(), captureTime.Month(), captureTime.Day(), endHour, endMin, 0, 0, captureTime.Location())
	return startTime, endTime
}

// TraverseTime 根据开始日期和结束日期获取时间段集合
func TraverseTime(startDate string, endDate string, interval string) []string {
	// layout 日期时间格式
	// 首先将startDate和endDate转换为time类型
	endTime, _ := time.ParseInLocation(DateFormat, endDate, time.Local)
	startTime, _ := time.ParseInLocation(DateFormat, startDate, time.Local)

	if interval == "day" {
		// dateArray 开始以及结束日期中的所有日期
		dateArray := make([]string, 0)
		// After 判断时间先后
		for endTime.After(startTime) {
			dateArray = append(dateArray, startTime.Format(DateFormat))
			startTime = startTime.AddDate(0, 0, 1)
		}
		dateArray = append(dateArray, endDate)
		return dateArray
	}
	if interval == "month" {
		//monthArray 开始以及结束日期中的所有日期
		monthArray := make([]string, 0)
		for endTime.After(startTime) {
			monthArray = append(monthArray, startTime.Format("2006-01"))
			startTime = startTime.AddDate(0, 1, 0)
		}
		monthArray = append(monthArray, endTime.Format("2006-01"))
		return monthArray
	}

	fmt.Println("Wrong interval")
	return nil
}
