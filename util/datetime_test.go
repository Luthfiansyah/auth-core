package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToTimeWithFormat(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)
}

func TestParseStringToTimeWithFormat2(t *testing.T) {
	timeString := "2017-01-21 20:17:50"

	time, err := ParseStringToTimeWithFormat("yyyy-MM-dd HH:mm:ss", timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)
}

func TestParseStringToTimeWithFormat3(t *testing.T) {
	timeString := "2017-01-21"

	time, err := ParseStringToTimeWithFormat("yyyy-MM-dd", timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)
}

func TestParseStringToTimeWithFormat4(t *testing.T) {
	timeString := "21-01-2017 20:17:50"

	tx, err := ParseStringToTimeWithFormat("dd-MM-yyyy HH:mm:ss", timeString)
	loc, _ := time.LoadLocation("Asia/Jakarta")

	assert.Nil(t, err)
	assert.NotNil(t, tx)

	fmt.Printf("original       : %s\n", timeString)
	fmt.Printf("time format    : %s\n", tx)
	fmt.Printf("RFC3339 format : %s\n", tx.In(loc).Format(time.RFC3339))
}

func TestFormatTime(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)

	s := FormatTime("yyyy-MM-dd HH:mm:ss", time)
	assert.Equal(t, "2017-01-21 20:17:50", s)
	fmt.Println(s)
}

func TestFormatTime2(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)

	s := FormatTime("yyyy-MM-dd", time)
	assert.Equal(t, "2017-01-21", s)
	fmt.Println(s)
}

func TestFormatTime3(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)

	s := FormatTime("HH:mm:ss", time)
	assert.Equal(t, "20:17:50", s)
	fmt.Println(s)
}

func TestRFC3339ToTime(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := RFC3339ToTime(timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)
}

func TestDDMMyyyyHHmmssToTime(t *testing.T) {
	timeString := "21-01-2017 20:17:50"

	time, err := DDMMyyyyHHmmssToTime(timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)
}

func TestToddMMyyyyHHmmss(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)

	s := ToddMMyyyyHHmmss(time)
	assert.Equal(t, "21-01-2017 20:17:50", s)
	fmt.Println(s)
}

func TestToddMMyyyy(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)

	s := ToddMMyyyy(time)
	assert.Equal(t, "21-01-2017", s)
	fmt.Println(s)
}

func TestToyyyyMMdd(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)

	s := ToyyyyMMdd(time)
	assert.Equal(t, "2017-01-21", s)
	fmt.Println(s)
}

func TestToHHmmss(t *testing.T) {
	timeString := "2017-01-21T20:17:50.476+07:00"

	time, err := ParseStringToTimeWithFormat(time.RFC3339, timeString)

	assert.Nil(t, err)
	assert.NotNil(t, time)

	fmt.Println(time)

	s := ToHHmmss(time)
	assert.Equal(t, "20:17:50", s)
	fmt.Println(s)
}

func TestDDMMyyyyToISO8601(t *testing.T) {
	timeString := "21-01-2017"

	s, err := DDMMyyyyToISO8601(timeString)

	assert.Nil(t, err)
	assert.NotEqual(t, "", s)

	fmt.Println(s)
}
