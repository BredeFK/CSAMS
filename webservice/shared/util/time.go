package util

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// DatetimeLocalToRFC3339 converts a string from datetime-local HTML input-field to time.Time object
func DatetimeLocalToRFC3339(str string) (time.Time, error) {
	// TODO (Svein): Move this to a utils.go or something
	if str == "" {
		return time.Time{}, errors.New("error: could not parse empty datetime-string")
	}
	if len(str) < 16 {
		return time.Time{}, errors.New("cannot convert a string less then 16 characters: DatetimeLocalToRFC3339()")
	}
	year := str[0:4]
	month := str[5:7]
	day := str[8:10]
	hour := str[11:13]
	min := str[14:16]

	value := fmt.Sprintf("%s-%s-%sT%s:%s:00Z", year, month, day, hour, min)
	return time.Parse(time.RFC3339, value)
}

// GoToHTMLDatetimeLocal converts time.Time object to 'datetime-local' in HTML
func GoToHTMLDatetimeLocal(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	day := fmt.Sprintf("%d", t.Day())
	month := fmt.Sprintf("%d", t.Month())
	year := fmt.Sprintf("%d", t.Year())
	hour := fmt.Sprintf("%d", t.Hour())
	minute := fmt.Sprintf("%d", t.Minute())

	if t.Day() < 10 {
		day = "0" + day
	}

	if t.Month() < 10 {
		month = "0" + month
	}

	if t.Hour() < 10 {
		hour = "0" + hour
	}

	if t.Minute() < 10 {
		minute = "0" + minute
	}

	return fmt.Sprintf("%s-%s-%sT%s:%s", year, month, day, hour, minute)
}

// GetTimeInCorrectTimeZone returns the time in "TIME_ZONE" time
func GetTimeInCorrectTimeZone() time.Time {
	//init the loc
	loc, err := time.LoadLocation(os.Getenv("TIME_ZONE"))
	if err != nil {
		panic(err.Error())
	}
	return time.Now().In(loc)
}

// ConvertTimeStampToString converts date to string for inserting in db
func ConvertTimeStampToString(timestamp time.Time) string {

	// ex: 2019-03-13 10:14:40
	return timestamp.Format("2006-01-02 15:04:05")
}
