package timer

import "time"


//便于后续对 Time 对象做进一步的统一处理，因为后面会涉及时区的一些问题处理。
func GetNowTime() time.Time{
	return time.Now()
}

// 时间推算
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error){
	// 从字符串中解析出持续时间
	duration, err := time.ParseDuration(d)
	if err != nil{
		return time.Time{}, err
	}
	// 得到当前timer时间加上duration后的最终时间
	return currentTimer.Add(duration), nil
}
/*
预先知道准确的duration，不需要上面的方法，直接使用Add方法处理
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

...
timer.GetNowTime().Add(time.Second * 60)
*/