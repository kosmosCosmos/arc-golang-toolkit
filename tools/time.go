package tools

import (
	"fmt"
	"github.com/araddon/dateparse"
	"time"
)

func ParseTime(RunMode, Start, End string) (time.Time, time.Time, error) {
	now := time.Now()
	var StartTime, EndTime time.Time
	switch RunMode {
	case "daily":
		StartTime = now.AddDate(0, 0, -1)
		EndTime = now
		return StartTime, EndTime, nil
	case "weekly":
		StartTime = now.AddDate(0, 0, -7)
		EndTime = now
		return StartTime, EndTime, nil
	case "monthly":
		StartTime = now.AddDate(0, 0, -30)
		EndTime = now
		return StartTime, EndTime, nil
	case "custom":
		// 自定义时间区间，需要解析 Start 和 End 参数
		if Start == "" || End == "" {
			return StartTime, EndTime, fmt.Errorf("for custom mode, please specify both start and end time")
		}

		var err error
		StartTime, err = dateparse.ParseLocal(Start)
		if err != nil {
			return StartTime, EndTime, fmt.Errorf("invalid start time: %w", err)
		}

		EndTime, err = dateparse.ParseLocal(End)
		if err != nil {
			return StartTime, EndTime, fmt.Errorf("invalid end time: %w", err)
		}

		if StartTime.After(EndTime) {
			return StartTime, EndTime, fmt.Errorf("start time (%v) is after end time (%v)", StartTime, EndTime)
		}

	default:
		return StartTime, EndTime, fmt.Errorf("invalid RunMode: %s", RunMode)
	}

	return StartTime, EndTime, nil
}

func IsRecentTime(timeStr string, year, month, days int) bool {
	replyTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return false
	}

	cutoffTime := time.Now().AddDate(-year, -month, -days) // 6 months ago
	return replyTime.After(cutoffTime)
}
