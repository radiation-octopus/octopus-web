package utils

import (
	"time"
)

func SetTime(hour, min, second int) (d time.Duration) {
	now := time.Now()
	setTime := time.Date(now.Year(), now.Month(), now.Day(), hour, min, second, 0, now.Location())
	d = setTime.Sub(now)
	if d > 0 {
		return
	}
	return d + time.Hour*24
}

func ScheduleTask() {
	// 每天7时0分触发更新
	t := time.NewTimer(SetTime(7, 0, 0))
	defer t.Stop()
	for {
		select {
		case <-t.C:
			t.Reset(time.Hour * 24)
			// 定时任务函数
		}
	}
}

//func WriteFileLine(path string, str string) {
//	f, err := os.Open(path)
//	if f == nil {
//		split := strings.Split(path, "/")
//		filePath := strings.Join(split[:len(split)-1], "/")
//		err = os.MkdirAll(filePath, os.ModePerm)
//		f, err = os.Create(path)
//	}
//	if err != nil {
//		panic(err)
//	}
//	r := bufio.NewWriter(f)
//	str = str + " \n"
//	_, err = r.WriteString(str)
//	if err != nil {
//		panic(err)
//	}
//	defer f.Close()
//}
