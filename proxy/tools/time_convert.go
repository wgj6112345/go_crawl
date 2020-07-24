package tools

import (
	"fmt"
	"strings"
	"time"
)

func TimeConvertWithSlash(timeStr string) time.Time {
	splits := strings.Split(timeStr, "/")
	month := splits[1]
	if len(month) < 2 {
		month = fmt.Sprintf("0%s", month)
	}

	monthStr := fmt.Sprintf("%s/%s/%s", splits[0], month, splits[2])
	result, _ := time.ParseInLocation("2006/01/02 15:04:05", monthStr, time.Local)
	return result
}

func TimeConvertWithBar(timeStr string) time.Time {
	result, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	return result
}
