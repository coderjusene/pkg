package timeutil

import (
	"fmt"
	"math"
	"net/http"
	"time"
)


const timeLayout = "2006-01-02 15:04:05"

var (
	tz *time.Location
)

func init() {
	var err error
	if tz, err = time.LoadLocation("Asia/Shanghai"); err != nil {
		panic(err)
	}

	// 默认设置中国时区
	time.Local = tz
}

// 更改时区
func ChangeTZ(zone string) {
	var err error
	if tz, err = time.LoadLocation(zone); err != nil {
		panic(err)
	}

	// 更改时区
	time.Local = tz
}

// RFC3339TotimeLayout convert rfc3339 value to standard time layout
// 2021-11-08T15:04:05Z => 2020-11-08 08:18:46
func RFC3339TotimeLayout(value string) (string, error) {
	t, err := time.Parse(time.RFC3339, value)
	fmt.Println(t)
	if err != nil {
		return "", err
	}

	return t.In(tz).Format(timeLayout), nil
}

// LayoutString 格式化时间
// 返回 "2006-01-02 15:04:05" 格式的时间
func NowLayoutString() string {
	ts := time.Now()
	return ts.In(tz).Format(timeLayout)
}

// ParseInLocation 格式化时间
func ParseInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(timeLayout, date, tz)
}

// LayoutStringToUnix 返回 unix 时间戳
// 2020-01-24 21:11:11 => 1579871471
func LayoutStringToUnix(layoutstring string) (int64, error) {
	tamp, err := ParseInLocation(layoutstring)
	if err != nil {
		return 0, err
	}

	return tamp.Unix(), nil
}


// GMTLayoutString 格式化时间
// 返回 "Mon, 02 Jan 2006 15:04:05 GMT" 格式的时间
func GMTLayoutString() string {
	return time.Now().In(tz).Format(http.TimeFormat)
}

// ParseGMTInLocation 格式化时间
func ParseGMTInLocation(date string) (time.Time, error) {
	return time.ParseInLocation(http.TimeFormat, date, tz)
}


// SubInLocation 计算时间差
func SubInLocation(ts time.Time) float64 {
	return math.Abs(time.Now().In(tz).Sub(ts).Seconds())
}


