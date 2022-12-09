package utility

import (
	"fmt"

	"github.com/gogf/gf/text/gregex"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 如果回复时间在今天零点之后，则把时间显示改为a小时b分钟之前
// 否则显示为年-月-日 时：分：秒
func TimeFormatDivide24Hour(inputTime *gtime.Time) (show string, err error) {

	if inputTime == nil {
		return
	}
	//获取今日零点时间戳
	todayZeroTime, err := gtime.StrToTime(gtime.Now().Format("Y-m-d") + " 00:00:00")

	if inputTime.Timestamp() >= todayZeroTime.Timestamp() {
		pastTime := gtime.Now().Timestamp() - inputTime.Timestamp()
		hour := pastTime / 3600
		minute := (pastTime - hour*3600) / 60
		if hour == minute && minute == 0 {
			show = "刚刚"
		} else if hour > 0 {
			show = fmt.Sprintf("%d小时%d分钟前", hour, minute)
		} else {
			show = fmt.Sprintf("%d分钟前", minute)
		}
	} else {
		show = inputTime.String()
	}

	return
}

func InSlice(needle uint, haystack []uint) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func ToTemplateMap(m interface{}) g.Map {
	return g.Map{
		"List": m,
	}
}

// ReplaceWarp 因为textarea 无法换行，所以需要把换行符替换为<br>
func ReplaceWarp(content string) (r string, err error) {
	var replaceRegex = `(\n\r|\r\n|\r|\n)`
	r, err = gregex.ReplaceString(replaceRegex, "<br/>", content)
	return
}
