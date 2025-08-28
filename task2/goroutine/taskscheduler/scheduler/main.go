package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 定义任务类型
type Task func()

// TaskResult 任务执行结果
type TaskResult struct {
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Success   bool
	Error     error
}

// Scheduler 任务调度器
type Scheduler struct {
	tasks      []Task
	results    []TaskResult
	concurrent int // 最大并发数
	mu         sync.Mutex
	wg         sync.WaitGroup
}

// NewScheduler 创建新的调度器
func NewScheduler(concurrent int) *Scheduler {
	return &Scheduler{
		tasks:      make([]Task, 0),
		results:    make([]TaskResult, 0),
		concurrent: concurrent,
	}
}

// AddTask 添加任务
func (s *Scheduler) AddTask(task Task) {
	s.tasks = append(s.tasks, task)
}

// AddTasks 批量添加任务
func (s *Scheduler) AddTasks(tasks []Task) {
	s.tasks = append(s.tasks, tasks...)
}

// Run 执行所有任务
func (s *Scheduler) Run() {
	totalTasks := len(s.tasks)
	if totalTasks == 0 {
		fmt.Println("没有任务需要执行")
		return
	}

	// 创建任务通道
	taskChan := make(chan int, s.concurrent)
	s.results = make([]TaskResult, totalTasks)

	fmt.Printf("开始执行 %d 个任务，最大并发数: %d\n", totalTasks, s.concurrent)

	// 启动工作协程
	for i := 0; i < s.concurrent; i++ {
		s.wg.Add(1)
		go s.worker(taskChan, i)
	}

	// 分发任务
	for i := 0; i < totalTasks; i++ {
		taskChan <- i
	}

	close(taskChan)
	s.wg.Wait()

	s.printResults()
}

// worker 工作协程
func (s *Scheduler) worker(taskChan chan int, workerID int) {
	defer s.wg.Done()

	//从通道中持续监听数据，如果通道中没有数据，就会阻塞在这里，直到有数据到来或通道被关闭。
	for taskID := range taskChan {
		s.executeTask(taskID, workerID)
	}
}

// executeTask 执行单个任务
func (s *Scheduler) executeTask(taskID int, workerID int) {
	task := s.tasks[taskID]
	result := TaskResult{
		TaskID:    taskID,
		StartTime: time.Now(),
	}

	defer func() {
		result.EndTime = time.Now()
		result.Duration = result.EndTime.Sub(result.StartTime)

		// 捕获panic
		if r := recover(); r != nil {
			result.Success = false
			result.Error = fmt.Errorf("panic: %v", r)
		}

		// 保存结果
		s.mu.Lock()
		s.results[taskID] = result
		s.mu.Unlock()

		fmt.Printf("Worker %d: 任务 %d 完成，耗时: %v\n", workerID, taskID, result.Duration)
	}()

	// 执行任务
	result.Success = true
	task()
}

// printResults 打印执行结果
func (s *Scheduler) printResults() {
	fmt.Println("\n=== 任务执行结果统计 ===")

	var totalDuration time.Duration
	successCount := 0

	for i, result := range s.results {
		status := "成功"
		if !result.Success {
			status = "失败"
		} else {
			successCount++
			totalDuration += result.Duration
		}

		fmt.Printf("任务 %d: %s, 耗时: %v", i, status, result.Duration)
		if result.Error != nil {
			fmt.Printf(", 错误: %v", result.Error)
		}
		fmt.Println()
	}

	fmt.Printf("\n总计: %d/%d 任务成功完成\n", successCount, len(s.tasks))
	fmt.Printf("总耗时: %v\n", totalDuration)
	if successCount > 0 {
		fmt.Printf("平均耗时: %v\n", totalDuration/time.Duration(successCount))
	}
}

// GetResults 获取执行结果
func (s *Scheduler) GetResults() []TaskResult {
	return s.results
}

func main() {
	// 创建调度器，最大并发数为3
	scheduler := NewScheduler(3)

	// 定义一些示例任务
	tasks := []Task{
		func() {
			time.Sleep(1 * time.Second)
			fmt.Println("任务1执行完成")
		},
		func() {
			time.Sleep(2 * time.Second)
			fmt.Println("任务2执行完成")
		},
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务3执行完成")
		},
		func() {
			time.Sleep(3 * time.Second)
			fmt.Println("任务4执行完成")
		},
		func() {
			// 模拟一个会失败的任务
			panic("模拟任务执行失败")
		},
		func() {
			time.Sleep(800 * time.Millisecond)
			fmt.Println("任务6执行完成")
		},
	}

	// 添加任务
	scheduler.AddTasks(tasks)

	// 执行所有任务
	start := time.Now()
	scheduler.Run()
	fmt.Printf("调度器总运行时间: %v\n", time.Since(start))
}
