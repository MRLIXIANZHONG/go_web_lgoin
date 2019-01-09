package poptime
import (
	"time"
)
func FormatUnixTimeAsStr(ctime int64) string {//时间搓转换成时间函数
	str_time := time.Unix(ctime, 0).Format("2006-01-02 15:04:05")//这个日期必须是这个
	return str_time
}
