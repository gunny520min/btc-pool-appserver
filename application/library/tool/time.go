package tool

import (
	"strconv"
	"time"
)

const (
	UTCDatetimeDay = "2006-01-02"
	UTCDatetime    = "2006-01-02 15:04:05"
	UTCDay         = "20060102"
	UTCDay00       = "2006010200"
	UTCDayHour     = "2006010215"
	UTCHourMinute  = "15:04:05"
)

// 获取昨日日期 20061203
func UtcYesterdayIntFormat() uint {
	daystr := time.Now().UTC().AddDate(0, 0, -1).Format(UTCDay)

	dayint, _ := strconv.Atoi(daystr)
	return uint(dayint)
}

// 获取今日日期 20061203
func UtcTodayIntFormat() uint {
	daystr := time.Now().UTC().Format(UTCDay)

	dayint, _ := strconv.Atoi(daystr)
	return uint(dayint)
}

// 获取某天凌晨时间
func UtcDayStartTime(day uint) (time.Time, error) {
	dayTime, err := time.ParseInLocation(UTCDay, strconv.FormatUint(uint64(day), 10), time.UTC)
	if err != nil {
		return dayTime, err
	}
	return time.Date(dayTime.Year(), dayTime.Month(), dayTime.Day(), 0, 0, 0, 0, time.UTC), nil
}

// 获取某天凌晨时间
func UtcDayEndTime(day uint) (time.Time, error) {
	dayStart, err := UtcDayStartTime(day)
	if err != nil {
		return dayStart, err
	}

	return dayStart.Add(time.Second * 86399), nil
}

// 获取今日凌晨时间
func UtcTodayStartTime() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

// 获取明日凌晨时间
func UtcTomorrowStartTime() time.Time {
	return UtcTodayStartTime().AddDate(0, 0, 1)
}
