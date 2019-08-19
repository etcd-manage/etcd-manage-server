package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/* 自定义时间格式 */
var timeFormart = "2006-01-02 15:04:05" // time.RFC3339

// JSONTime 时间格式别名
type JSONTime time.Time

// UnmarshalJSON 字节转为JSONTime对象
func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.UTC)
	*t = JSONTime(now)
	return
}

// MarshalJSON 将时间对象转为字节
func (t JSONTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).Local().AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

// String 格式化为文本
func (t JSONTime) String() string {
	return time.Time(t).Format(timeFormart)
}

// Format 格式化函数
func (t JSONTime) Format(format string) string {
	return time.Time(t).Format(timeFormart)
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
