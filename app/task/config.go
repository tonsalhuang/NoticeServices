package task

import (
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/os/glog"
)

type Entity struct {
	FuncName string
	Param    []string
	Run      func()
}

var taskList = garray.New()

//增加Task方法
func Add(task Entity) {
	if task.FuncName == "" {
		return
	}
	if task.Run == nil {
		return
	}
	taskList.Append(task)
}

//检查方法名是否存在
func GetByName(funcName string) *Entity {
	var result *Entity
	for _, item := range taskList.Slice() {
		task := item.(Entity)
		if task.FuncName == funcName {
			result = &task
			break
		}
	}
	return result
}

//EditParams 修改参数
func EditParams(funcName string, params []string) {
	for index, item := range taskList.Slice() {
		task := item.(Entity)
		if task.FuncName == funcName {
			task.Param = params
			if err := taskList.Set(index, task); err != nil {
				glog.Error(err)
			}
			break
		}
	}
}
