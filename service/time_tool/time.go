package time_tool

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"strconv"
	"time"
)

func ConvertTime(wf *aw.Workflow, cli *cli.Context) {
	// 如果未传递参数，则展示当前时间
	if cli.Args().Len() == 0 {
		logrus.Println("arg size is zero")
		AddTimeItem(wf, time.Now())
		wf.SendFeedback()
		return
	}

	if cli.Args().Len() == 1 {
		argStr := cli.Args().First()
		arg, err := strconv.ParseInt(argStr, 10, 64)
		if err != nil {
			title := "请输入正确的时间戳格式"
			wf.NewItem(title).Arg(title).Valid(true)
		} else {
			// 秒的处理
			if len(argStr) == 10 {
				AddTimeItem(wf, time.Unix(int64(arg), 0))
			} else {
				// 毫秒的处理
				AddTimeItem(wf, time.Unix(int64(arg/1000), 0))
			}
		}
	} else if cli.Args().Len() == 2 {
		argStr := fmt.Sprintf("%s %s", cli.Args().Get(0), cli.Args().Get(1))
		res, err := time.Parse("2006-01-02 15:04:05", argStr)
		if err != nil {
			wf.NewItem("请输入2006-01-02 15:04:05格式").Arg("请输入2006-01-02 15:04:05格式").Valid(true)
			wf.SendFeedback()
			return
		}
		AddTimeItem(wf, res)
	} else {
		wf.NewItem("请输入2006-01-02 15:04:05 或者时间戳的格式").Arg("请输入2006-01-02 15:04:05 或者时间戳的格式").Valid(true)
	}

	wf.SendFeedback()
}

func AddTimeItem(wf *aw.Workflow, now time.Time) {
	second := strconv.Itoa(int(now.Unix()))
	millSecond := strconv.Itoa(int(now.UnixMilli()))
	wf.NewItem(second).Arg(second).Subtitle("秒").Valid(true)
	wf.NewItem(millSecond).Arg(millSecond).Subtitle("毫秒").Valid(true)

	timeFormat := now.Format("2006-01-02 15:04:05")
	wf.NewItem(timeFormat).Arg(timeFormat).Var("result", timeFormat).Valid(true)
}
