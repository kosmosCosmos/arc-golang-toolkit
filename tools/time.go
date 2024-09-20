package tools

import (
	"fmt"
	"github.com/araddon/dateparse"
	"log"
	"strings"
	"time"
)

func parseDateTime(dateStr string) (time.Time, error) {
	originalDateStr := dateStr
	dateStr = strings.TrimSpace(dateStr)

	// 处理特殊情况
	if strings.Contains(dateStr, "今天") {
		// 今天
		today := time.Now().Format("2006-01-02")
		dateStr = strings.Replace(dateStr, "今天", today, 1)
	} else if strings.Contains(dateStr, "昨天") {
		// 昨天
		yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		dateStr = strings.Replace(dateStr, "昨天", yesterday, 1)
	} else if strings.Contains(dateStr, "前天") {
		// 前天
		dayBeforeYesterday := time.Now().AddDate(0, 0, -2).Format("2006-01-02")
		dateStr = strings.Replace(dateStr, "前天", dayBeforeYesterday, 1)
	} else if strings.Contains(dateStr, "月") && strings.Contains(dateStr, "日") {
		// 将 "9月5日 14:20" 转换为 "2023-09-05 14:20"
		dateStr = strings.Replace(dateStr, "月", "-", 1)
		dateStr = strings.Replace(dateStr, "日", "", 1)
		currentYear := time.Now().Year()
		dateStr = fmt.Sprintf("%d-%s", currentYear, dateStr)
	} else if len(dateStr) <= 5 && strings.Contains(dateStr, ":") {
		// 只有时间，假设为今天
		today := time.Now().Format("2006-01-02")
		dateStr = fmt.Sprintf("%s %s", today, dateStr)
	}

	// 使用 dateparse 解析日期字符串
	parsedTime, err := dateparse.ParseIn(dateStr, time.Local)
	if err != nil {
		log.Printf("Failed to parse date '%s' (original: '%s'): %v", dateStr, originalDateStr, err)
		return time.Time{}, fmt.Errorf("无法解析日期时间: %v", err)
	}
	log.Printf("Successfully parsed date '%s' (original: '%s') as %v", dateStr, originalDateStr, parsedTime)
	return parsedTime, nil
}

func IsRecentTime(dateStr string) bool {
	parsedTime, err := parseDateTime(dateStr)
	if err != nil {
		log.Printf("Failed to parse date '%s': %v", dateStr, err)
		return false // 无法解析日期，跳过该条目
	}
	threshold := time.Now().Add(-1 * time.Hour) // 设定时间阈值
	isRecent := parsedTime.After(threshold)
	log.Printf("Time: %s, Is recent: %t", parsedTime.Format("2006-01-02 15:04"), isRecent)
	return isRecent
}
