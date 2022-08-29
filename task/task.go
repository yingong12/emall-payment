// 异步任务
package task

import (
	"context"
	"emall/providers"
	"emall/utils"
	"fmt"
	"log"
	"sync"
	"time"
)

var globalMu = sync.RWMutex{}

const (
	TaskStateRegitered  = iota //已注册
	TaskStateProcessing        //正在处理
	TaskStateCompleted         //成功
	TaskStateFailed            //失败
)

type Task struct {
	TaskID   string `json:"taskID"`
	OrderID  string `json:"orderID"`
	KickTime string `json:"kickTime"`
	State    int8   `json:"state"` //状态 0-pending 1-completed 2-failed
	Next     *Task
}

const slotLen = 3600

var timeSlot []*Task = make([]*Task, slotLen)

// 当前时间片位置
var idx = 0

func Start(ctx context.Context, cancel context.CancelFunc) func() {
	wg := &sync.WaitGroup{}
	go cleanRound()
	go func() {
		//TODO:这里需要保证所有task都被做完才退出
		//! 只是标记了状态，没有真正去删除节点,内存会爆
		for {
			select {
			case <-ctx.Done():
				//clean
				return
			default:
				wg.Add(1)
				go func() {
					defer func() {
						wg.Done()
					}()
					tt := timeSlot[idx]
					log.Println("doing task ", idx)
					if err := doSlotTasks(tt); err != nil {
						//retry
						doSlotTasks(tt)
					}
					//标记task
					markTask(tt, TaskStateCompleted)
				}()
				//1秒执行一次
				time.Sleep(time.Second)
				idx = (idx + 1) % 3600
			}
		}
	}()
	//
	return func() {
		cancel()
		log.Println("Stoping asyn tasks...")
		wg.Wait()
	}
}

// 20分钟过期
const timeGap = 10

// task加入list
func AddTask(orderID string) {
	taskID := utils.GenerateTaskID()
	task := &Task{
		TaskID:  taskID,
		OrderID: orderID,
		State:   TaskStateRegitered,
	}
	slot := (idx + timeGap) % slotLen
	defer func() {
		fmt.Printf("Task registered at slot:%d taskID:%s orderID:%s\n", slot, taskID, orderID)
	}()
	head := timeSlot[slot]
	//为空
	if head == nil {
		timeSlot[slot] = task
		return
	}
	//add task to head
	globalMu.RLock()
	nxt := head.Next
	head.Next = task
	task.Next = nxt
	globalMu.RUnlock()
}

// cleaning interval
const CleanInterval = time.Second * 10

// cleanRound 删除标记为Completed的task
func cleanRound() {
	// 不用安全退出，因为cleaner导致数据丢失
	globalMu.Lock()
	log.Print("Cleaner doing his job!")
	total := 0
	wg := &sync.WaitGroup{}
	for _, v := range timeSlot {
		wg.Add(1)
		//*并发回收,提高效率
		go func(head *Task) {
			prev := &Task{
				Next: head,
			}
			for head != nil {
				if head.State != TaskStateCompleted {
					head = head.Next
					continue
				}
				prev.Next = head.Next
				nxt := head.Next
				prev = head
				head = nxt
				total++
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	// 尾递归，不会造成stack overflow
	log.Printf("Cleaner job done! Collected %d completed taskes\n", total)
	globalMu.Unlock()
	time.Sleep(CleanInterval)
	cleanRound()
}

func markTask(task *Task, state int) {
	globalMu.RLock()
	for ; task != nil; task = task.Next {
		fmt.Println("marked ", task.TaskID)
		task.State = TaskStateCompleted
	}
	globalMu.RUnlock()
}

func doSlotTasks(task *Task) (err error) {
	//遍历链表，取出orderID 更新为过期
	globalMu.RLock()
	defer func() {
		globalMu.RUnlock()
	}()
	ordrIDList := []string{}
	for ; task != nil; task = task.Next {
		//已完成了的任务，略过
		if task.State == TaskStateCompleted {
			continue
		}
		ordrIDList = append(ordrIDList, task.OrderID)
		task.State = TaskStateProcessing
	}
	if len(ordrIDList) == 0 {
		return
	}
	//update DB
	tx := providers.DBconnector.Table("t_order").Where("order_id in ?", ordrIDList).Update("state", utils.OrderStateExpired)
	err = tx.Error
	return
}
