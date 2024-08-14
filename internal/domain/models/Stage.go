package models

const NEW = 1

const WORK = 2

const DONE = 3

const REVIEW = 4

const READY_TO_TEST = 5

const TEST = 6

const COMPLETED = 7

const RELEASE = 8

const END = 100 // финальное состояние задачи, задача в таком статусе не отображается на доске
const BACKLOG = 101

const EMPLOYMENT_SEARCH_LIST = 1001

const EMPLOYMENT_SEND_MAIL = 1002

const EMPLOYMENT_NOT_ANSWER = 1003

const EMPLOYMENT_FEEDBACK = 1004

const EMPLOYMENT_INTERVIEW_PLANNED = 1005

const EMPLOYMENT_INTERVIEW = 1006

const EMPLOYMENT_DECLINE = 1007

const EMPLOYMENT_ACCEPT = 1008

const EMPLOYMENT_INTERNSHIP = 1009

type Stage struct {
	Id   int32
	Name string
}

func NewStage(id int32, name string) *Stage {
	return &Stage{
		Id:   id,
		Name: name,
	}
}

func AllStages() map[int32]*Stage {
	res := make(map[int32]*Stage)
	res[NEW] = NewStage(NEW, "Новые")
	res[WORK] = NewStage(WORK, "В работе")
	res[DONE] = NewStage(DONE, "Выполнены")
	res[REVIEW] = NewStage(REVIEW, "Ревью")
	res[READY_TO_TEST] = NewStage(READY_TO_TEST, "Готовы к тестированию")
	res[TEST] = NewStage(TEST, "В тестировании")
	res[COMPLETED] = NewStage(COMPLETED, "Решены")
	res[RELEASE] = NewStage(RELEASE, "Релиз")
	res[END] = NewStage(END, "Завершена")
	res[BACKLOG] = NewStage(BACKLOG, "Бэклог")

	res[EMPLOYMENT_SEARCH_LIST] = NewStage(EMPLOYMENT_SEARCH_LIST, "Список")
	res[EMPLOYMENT_SEND_MAIL] = NewStage(EMPLOYMENT_SEND_MAIL, "Отправлено письмо")
	res[EMPLOYMENT_NOT_ANSWER] = NewStage(EMPLOYMENT_NOT_ANSWER, "Нет ответа")
	res[EMPLOYMENT_FEEDBACK] = NewStage(EMPLOYMENT_FEEDBACK, "Получена обратная связь")
	res[EMPLOYMENT_INTERVIEW_PLANNED] = NewStage(EMPLOYMENT_INTERVIEW_PLANNED, "Назначено время собеседования")
	res[EMPLOYMENT_INTERVIEW] = NewStage(EMPLOYMENT_INTERVIEW, "Провели собеседование")
	res[EMPLOYMENT_DECLINE] = NewStage(EMPLOYMENT_DECLINE, "Отказался")
	res[EMPLOYMENT_ACCEPT] = NewStage(EMPLOYMENT_ACCEPT, "Принят")
	res[EMPLOYMENT_INTERNSHIP] = NewStage(EMPLOYMENT_INTERNSHIP, "Стажировка")

	return res
}
