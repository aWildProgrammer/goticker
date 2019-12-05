package goticker

/**
 * 通用函数
 */

import "time"

// 输入时间日期得出时间戳
// @param dateTime := "2019-12-04 17:56:00"
func strtotime(dateTime string) time.Time {
    timeLayout := "2006-01-02 15:04:05"
    loc, _ := time.LoadLocation("Local")
    theTime, _ := time.ParseInLocation(timeLayout, dateTime, loc)
    return theTime
}
