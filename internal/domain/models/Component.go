package models

const DESIGN = 1
const MARKUP = 2
const DEV = 3
const TEST_C = 4
const CONTENT = 5
const MANAGEMENT = 6
const FRONTEND = 7
const ADMINISTERING = 8
const ANALYSIS = 9
const COPYRIGHT = 10

type Component struct {
	Id    int32
	Name  string
	Color string
}

func NewComponent(id int32, name string, color string) *Component {
	return &Component{
		Id:    id,
		Name:  name,
		Color: color,
	}
}

func AllComponents() map[int32]*Component {
	res := make(map[int32]*Component)
	res[DESIGN] = NewComponent(DESIGN, "Дизайн", "#EA5471")
	res[MARKUP] = NewComponent(MARKUP, "Вёрстка", "#FFA826")
	res[DEV] = NewComponent(DEV, "Разработка", "#3787EB")
	res[TEST_C] = NewComponent(TEST_C, "Тестирование", "#18BACE")
	res[CONTENT] = NewComponent(CONTENT, "Контент", "#FF5A4F")
	res[MANAGEMENT] = NewComponent(MANAGEMENT, "Менеджмент", "#FF6E41")
	res[FRONTEND] = NewComponent(FRONTEND, "Фронтенд", "#32C997")
	res[ADMINISTERING] = NewComponent(ADMINISTERING, "Администрирование", "#4761EE")
	res[ANALYSIS] = NewComponent(ANALYSIS, "Аналитика", "#4C33FF")
	res[COPYRIGHT] = NewComponent(COPYRIGHT, "Копирайт", "#13B4ED")
	return res
}
