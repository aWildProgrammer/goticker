# goticker
一个超轻量级的计划任务，可用于周期性及定时周期性任务，替代go自代的 time.Ticker 机制。

由于 go 自带的 time.Ticker 比较弱鸡，仅支持 chan 回写，所以封装在业务里需要大量的工作及协程处理，在性能上消耗很大，所以这里写一个轻量级的分享。


```code


import ."fmt"
import "github.com/aWildProgrammer/goticker"

func main() {

	task := goticker.New(100)
	ch2 := make(chan bool, 0)
	ch4 := make(chan bool, 0)
	id1 := task.AddTaskCallBackFunc(test, 3, "任务111111") // 每间隔3秒执行一次 test 函数
	id2 := task.AddTaskCallBackChannel(ch2, 5) // 每间隔5秒向ch2回写 bool true
	id3 := task.AddCycleTaskCallBackFunc(test, "16:30:45", "任务333333") // 每天 16:18:55 执行一次 test，如果在建立该任务时当前时间已经超过该指定时间，则任务推迟至明天同一时间执行
	id4 := task.AddCycleTaskCallBackChannel(ch4, "16:30:49") // 每天 16:18:55 向ch4回写 bool true，如果在建立该任务时当前时间已经超过该指定时间，则任务推迟至明天同一时间执行

	Println(id1, id2, id3, id4)
	for {
		select {
		case <- ch2:
			Println("任务222222")
		case <- ch4:
			Println("任务444444")
		}
	}
}

func test(args interface{}) {
	Println(args)
}

```
