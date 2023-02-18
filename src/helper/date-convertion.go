package helper

import (
	"fmt"
	"strings"
	"time"
)

func ConvertStringToDate(rawDate string) time.Time {
	splitDate := strings.Split(rawDate, "/")
	date := fmt.Sprintf("%s-%s-%s", splitDate[2], splitDate[1], splitDate[0])
	t, _ := time.Parse("2006-01-02", date)
	return t
}

func ConvertDateToString(time time.Time) string {
	strDate := time.Format("2006-01-02")
	splitDate := strings.Split(strDate, "-")
	return fmt.Sprintf("%s/%s/%s", splitDate[2], splitDate[1], splitDate[0])
}
