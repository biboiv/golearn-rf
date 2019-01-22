package commonFunc

import (
	"fmt"
	"time"
)

func DdateTiDDMMYYY(dt time.Time) string {
	return fmt.Sprintf("%02d.%02d.%d", dt.Day(), dt.Month(), dt.Year())
}
func DdateTiDDMMYYYhhmm(dt time.Time) string {
	return fmt.Sprintf("%02d.%02d.%d %02d:%02d", dt.Day(), dt.Month(), dt.Year(), dt.Hour(), dt.Minute())
}
func DdateTiDDMMYYYhhmmss(dt time.Time) string {
	return fmt.Sprintf("%02d.%02d.%d %02d:%02d:%02d", dt.Day(), dt.Month(), dt.Year(), dt.Hour(), dt.Minute(), dt.Second())
}

//( "%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
//        t.Year(), t.Month(), t.Day(),
//        t.Hour(), t.Minute(), t.Second())
