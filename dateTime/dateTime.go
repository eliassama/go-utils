package dateTime

import (
	"errors"
	"time"
)

// GetDateUnixScope 获取某一天的的开始时间戳和结束时间戳（按秒）
//
//goland:noinspection GoUnusedExportedFunction
func GetDateUnixScope(date string) (start int64, end int64, err error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return 0, 0, err
	}

	unix, err := time.ParseInLocation("2006-01-02", date, loc)
	if err != nil {
		return 0, 0, err
	}

	start = time.Date(unix.Year(), unix.Month(), unix.Day(), 0, 0, 0, 0, unix.Location()).Unix()
	end = time.Date(unix.Year(), unix.Month(), unix.Day(), 23, 59, 59, 999, unix.Location()).Unix()

	return
}

// GetDateScope 获取两个日期之间经过/包含的日期，开始时间必须小于等于结束时间。
// 返回的值中，第一个值是开始日期，最后一个是结束日期。如果开始日期等于结束日期，则只返回一个日期。
//
//goland:noinspection GoUnusedExportedFunction
func GetDateScope(start string, end string) (dateScope []string, err error) {
	var loc *time.Location
	var startUnix time.Time
	var endUnix time.Time

	loc, err = time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return
	}

	startUnix, err = time.ParseInLocation("2006-01-02", start, loc)
	if err != nil {
		return
	}

	endUnix, err = time.ParseInLocation("2006-01-02", end, loc)
	if err != nil {
		return
	}

	if startUnix.Equal(endUnix) {
		return []string{startUnix.Format("2006-01-02")}, nil
	}

	if !startUnix.Before(endUnix) {
		return nil, errors.New("startDate Must before endDate")
	}

	for {

		dateScope = append(dateScope, startUnix.Format("2006-01-02"))

		if startUnix.Equal(endUnix) {
			break
		}

		startUnix = startUnix.AddDate(0, 0, 1)
	}

	return
}
