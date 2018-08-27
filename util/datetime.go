package util

import (
	"github.com/auth-core/logging"

	"bytes"
	"time"
)

var log = logging.MustGetLogger("auth-core")

// ParseStringToTimeWithFormat will return time.Time if timeString match with format and nil error
// otherwise will return zero Time and specific error
// format example
// 		dd-MM-yyyy HH:mm:ss for "15-01-2017 13:01:00"
//		yyyy MMMM dd HH:mm:ss for "2017 January 15 13:01:00"
//		dd MM yyyy for (15 January 2017)
func ParseStringToTimeWithFormat(format string, timeString string) (time.Time, error) {
	return time.Parse(getFormat(format), timeString)
}

// FormatTime will return string from time with inputted format
// format example
// 		dd-MM-yyyy HH:mm:ss for "15-01-2017 13:01:00"
//		yyyy MMMM dd HH:mm:ss for "2017 January 15 13:01:00"
//		dd MM yyyy for (15 January 2017)
func FormatTime(format string, time time.Time) string {
	return time.Format(getFormat(format))
}

// RFC3339ToTime will return time.Time if timeString match with time.RFC3339 and nil error
// Otherwise return zero time and specific error
func RFC3339ToTime(timeString string) (time.Time, error) {
	return ParseStringToTimeWithFormat(time.RFC3339, timeString)
}

// DDMMyyyyHHmmssToTime will return time.Time if timeString match with dd-MM-yyyy HH:mm:ss and nil error
// Otherwise return zero time and specific error
func DDMMyyyyHHmmssToTime(timeString string) (time.Time, error) {
	return ParseStringToTimeWithFormat("dd-MM-yyyy HH:mm:ss", timeString)
}

// DDMMyyyyToISO8601 will return string form timeString in ISO8601 format
func DDMMyyyyToISO8601(timeString string) (string, error) {
	t, err := ParseStringToTimeWithFormat("dd-MM-yyyy", timeString)
	if err != nil {
		log.Error(logging.INTERNAL, err)
		return "", err
	}
	return t.Format(time.RFC3339), nil
}

// DDMMyyyyHHmmssToISO8601 will return string form timeString in ISO8601 format
func DDMMyyyyHHmmssToISO8601(timeString string) (string, error) {
	t, err := ParseStringToTimeWithFormat("dd-MM-yyyy HH:mm:ss", timeString)
	if err != nil {
		log.Error(logging.INTERNAL, err)
		return "", err
	}
	return t.Format(time.RFC3339), nil
}

// ToddMMyyyyHHmmss will return string from time with dd-MM-yyyy HH:mm:ss format
func ToddMMyyyyHHmmss(time time.Time) string {
	return FormatTime("dd-MM-yyyy HH:mm:ss", time)
}

// ToddMMyyyy will return string from time with dd-MM-yyyy format
func ToddMMyyyy(time time.Time) string {
	return FormatTime("dd-MM-yyyy", time)
}

// ToyyyyMMdd will return string from time with dd-MM-yyyy format
func ToyyyyMMdd(time time.Time) string {
	return FormatTime("yyyy-MM-dd", time)
}

// ToHHmmss will return string from time with HH:mm:ss format
func ToHHmmss(time time.Time) string {
	return FormatTime("HH:mm:ss", time)
}

func getFormat(format string) string {
	var buffer bytes.Buffer
	before := rune('!')
	count := 0
	i := 0
	for _, r := range format {
		i++
		if before == '!' {
			before = r
			count = 1
			continue
		}

		if before != r {
			appendToBuffer(before, count, &buffer)
			before = r
			count = 1
			continue
		}
		count++
		if i == len(format) {
			appendToBuffer(before, count, &buffer)
		}
	}
	return buffer.String()
}

func appendToBuffer(before rune, count int, buffer *bytes.Buffer) {
	if before == 'y' {
		if count >= 4 {
			buffer.WriteString("2006")
		} else {
			buffer.WriteString("06")
		}
	} else if before == 'M' {
		if count >= 4 {
			buffer.WriteString("January")
		} else if count == 3 {
			buffer.WriteString("Jan")
		} else if count == 2 {
			buffer.WriteString("01")
		} else {
			buffer.WriteString("1")
		}
	} else if before == 'D' || before == 'd' {
		if count >= 4 {
			buffer.WriteString("Monday")
		} else if count == 3 {
			buffer.WriteString("Mon")
		} else if count == 2 {
			buffer.WriteString("02")
		} else {
			buffer.WriteString("2")
		}
	} else if before == 'H' || before == 'h' {
		if count == 2 {
			buffer.WriteString("15")
		} else {
			buffer.WriteString("3")
		}
	} else if before == 'm' {
		if count == 2 {
			buffer.WriteString("04")
		} else {
			buffer.WriteString("4")
		}
	} else if before == 's' {
		if count == 2 {
			buffer.WriteString("05")
		} else {
			buffer.WriteString("5")
		}
	} else {
		for j := 0; j < count; j++ {
			buffer.WriteString(string(before))
		}
	}

}
