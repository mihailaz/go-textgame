package player

import (
	"strings"
)

type Task interface {
	Notifiable
	Description() string
	IsDone() bool
}

type TaskList []*Task

func (l *TaskList) ShowTasks() string {
	var taskNames []string
	for _, task := range *l {
		if !(*task).IsDone() {
			taskNames = append(taskNames, (*task).Description())
		}
	}
	if len(taskNames) == 0 {
		return "все выполнено"
	}
	return "надо " + strings.Join(taskNames, " и ")
}

func (l *TaskList) Notify(state *Player) {
	for _, task := range *l {
		(*task).Notify(state)
	}
}

type TakeItemTask struct {
	TaskImpl
	ItemName string
}

func (t *TakeItemTask) Notify(state *Player) {
	_, idx, _ := state.FindItem(t.ItemName)
	t._isDone = idx >= 0
}

type GoLocationTask struct {
	TaskImpl
	LocationName string
}

func (t *GoLocationTask) Notify(state *Player) {
	t._isDone = state.CurrentLocation.Name() == t.LocationName
}

type Notifiable interface {
	Notify(*Player)
}

type TaskImpl struct {
	_description string
	_isDone      bool
}

func (t *TaskImpl) Description() string {
	return t._description
}
func (t *TaskImpl) IsDone() bool {
	return t._isDone
}

func Create() TaskList {
	takeItem := &TakeItemTask{TaskImpl{_description: "собрать рюкзак"}, "рюкзак"}
	goLocation := &GoLocationTask{TaskImpl{_description: "идти в универ"}, "универ"}
	var takeItemTask Task = takeItem
	var goLocationTask Task = goLocation
	return TaskList{&takeItemTask, &goLocationTask}
}
