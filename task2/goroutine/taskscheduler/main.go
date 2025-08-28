package main

import (
	"fmt"
	"sync"
	"time"
)

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

type Task struct {
	Name string
	Func func()
}

type TaskResult struct {
	TaskID    int
	Task      Task
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Success   bool
	Error     error
}

type TaskScheduler struct {
	tasks      []Task
	wg         sync.WaitGroup
	mu         sync.Mutex
	concurrent int //并发数
	result     []TaskResult
}

func NewTaskScheduler(concurrent int) *TaskScheduler {
	return &TaskScheduler{tasks: make([]Task, 0),
		wg:         sync.WaitGroup{},
		mu:         sync.Mutex{},
		result:     make([]TaskResult, 0),
		concurrent: concurrent,
	}
}

// AddTasks /*
func (ts *TaskScheduler) AddTasks(tasks []Task) {
	ts.tasks = append(ts.tasks, tasks...)
}
func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}
func (ts *TaskScheduler) Run() {
	//判断任务数量
	taskCount := len(ts.tasks)
	if taskCount == 0 {
		fmt.Println("调度器没有任务")
		return
	}

	//根据预先定义的并发数 创建concurrent个协程
	//创建同样大小的channel用来分发任务

	queue := make(chan int, ts.concurrent)

	//创建任务执行结果信息
	ts.result = make([]TaskResult, taskCount)

	fmt.Printf("开始执行 %d个任务，总共 %d 个并发\n\n", taskCount, ts.concurrent)

	//启动并发
	for i := 0; i < ts.concurrent; i++ {
		ts.wg.Add(1)
		//启动goroutine 将channel和 并发id 传进去
		go ts.handleWork(queue, i)
	}

	//分发任务
	for i := 0; i < taskCount; i++ {
		queue <- i
	}
	close(queue)

	ts.wg.Wait()

}

func (ts *TaskScheduler) handleWork(queue chan int, workId int) {
	//线程开始执行内容，从channel中监听任务，拿到任务后执行
	defer ts.wg.Done()
	//从通道中持续监听数据，如果通道中没有数据，就会阻塞在这里，直到有数据到来或通道被关闭。
	for taskIndex := range queue {
		ts.executeTask(taskIndex, workId)
	}
}

func (ts *TaskScheduler) executeTask(index int, id int) {
	//执行单个任务，然后封装结果信息
	task := ts.tasks[index]
	//创建任务执行记录
	taskResult := TaskResult{TaskID: index, Task: task, StartTime: time.Now()}

	//此方法执行结束后执行此匿名函数
	defer func() {
		taskResult.EndTime = time.Now()
		taskResult.Duration = taskResult.EndTime.Sub(taskResult.StartTime)

		//判断是否有panic异常，只能在defer中使用
		if r := recover(); r != nil {
			taskResult.Success = false
			taskResult.Error = fmt.Errorf("error:%v", r)
		}

		//将taskResult保存到切片中，多个协程往ts.result中写数据,加锁
		ts.mu.Lock()
		ts.result[index] = taskResult
		ts.mu.Unlock()
		taskResult.Success = true
		fmt.Printf("Worker %d: 任务 %d--%s 完成，耗时: %v\n", id, index, task.Name, taskResult.Duration)
	}()

	task.Func() //执行任务
}

func main() {
	tasks := []Task{
		{"Task1", func() {
			time.Sleep(2 * time.Second)
		}},
		{"Task2", func() {
			time.Sleep(1 * time.Second)
		}},
		{"Task3", func() {
			time.Sleep(3 * time.Second)
		}},
		{"Task4", func() {
			panic("执行出错")
		}},
	}
	ts := NewTaskScheduler(2)
	ts.AddTasks(tasks)
	// 执行所有任务
	start := time.Now()
	ts.Run()
	fmt.Printf("调度器总运行时间: %v\n", time.Since(start))
	fmt.Println(ts.result)

}
