package models

const BUG = 1

const TASK = 2

const IMPROVEMENT = 3

const NEW_FEATURE = 4

const EPIC = 5

const RELEASE_TYPE = 6

const WARRANTY = 8

const RESEARCH = 9

type TaskType struct {
	Id      int32
	Name    string
	Color   string
	BgColor string
}

func NewTaskType(id int32, name string, color string, bgColor string) *TaskType {
	return &TaskType{
		Id:      id,
		Name:    name,
		Color:   color,
		BgColor: bgColor,
	}
}

func AllTaskTypes() map[int32]*TaskType {
	res := make(map[int32]*TaskType)
	res[BUG] = NewTaskType(BUG, "Баг", "#FFF1F0", "#FF5A4F")
	res[TASK] = NewTaskType(TASK, "Задача", "#EEF5FC", "#3787EB")
	res[IMPROVEMENT] = NewTaskType(IMPROVEMENT, "Улучшение", "#F1FBF8", "#32C997")
	res[NEW_FEATURE] = NewTaskType(NEW_FEATURE, "Новая функциональность", "#FFF8EC", "#FFA826")
	res[EPIC] = NewTaskType(EPIC, "Эпик", "#F0EEFF", "#6457FA")
	res[RELEASE_TYPE] = NewTaskType(RELEASE, "Релиз", "#FFF1EC", "#FF6E41")
	res[WARRANTY] = NewTaskType(WARRANTY, "Гарантия", "#EEF5FC", "#3787EB")
	res[RESEARCH] = NewTaskType(RESEARCH, "Исследование", "#F1FBF8", "#FF6E41")
	return res
}
