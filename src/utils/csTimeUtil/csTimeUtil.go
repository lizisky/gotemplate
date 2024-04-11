package csTimeUtil

import (
	"time"

	"github.com/duke-git/lancet/datetime"
)

// @BeginOfDay 输入一个以毫秒为单位的时间，返回以毫秒为单位的这一天的起始时间
func BeginOfDay(timeslot uint64) int64 {
	beginOfDay := datetime.BeginOfDay(time.UnixMilli(int64(timeslot))).UnixMilli()
	return beginOfDay
}

// 获取今天的开始时间
func BeginOfToday() uint64 {
	return uint64(datetime.BeginOfDay(time.Now()).UnixMilli())
}

// @Weekday 输入一个以毫秒为单位的时间，返回这一天是星期几
func Weekday(timeslot uint64) uint {
	return uint((time.UnixMilli(int64(timeslot)).Weekday()))
}

// @GetDateTime 输入一个以毫秒为单位的时间，返回一个格式化的时间字符串
// format: YYYY-MM-DD HH:MM
func GetDateTime(timeslot uint64) string {
	// return time.UnixMilli(int64(timeslot)).Format("2006-01-02 15:04:05")
	return time.UnixMilli(int64(timeslot)).Format("2006-01-02 15:04")
}

// @GetMinutes 输入一个以毫秒为单位的时间，返回对应的分钟数
func GetMinutes(timeslot uint64) int {
	return time.UnixMilli(int64(timeslot)).Minute()
}
