package flow

import "gokanban/internal/domain/models"

type DemoFlow struct {
}

func NewDemoFlow() *DemoFlow {
	return &DemoFlow{}
}

func (f *DemoFlow) name() string {
	return "Демо"
}

func (f *DemoFlow) slug() string {
	return "demo"
}

func (f *DemoFlow) PossibleTaskNextStages(task *models.Task, user *models.User, project *models.Project) []int32 {
	var res []int32
	stageId := task.StageId

	if stageId == models.BACKLOG {
		if user.IsAdmin || user.HasRole(1, project) {
			return []int32{
				models.NEW,
			}
		} else {
			return []int32{}
		}
	}
	if user.IsAdmin || project.ShouldTreatUserAsAdmin(user) {
		res = f.PossibleProjectStages()
		res = append(res, models.BACKLOG)

		return res
	}

	if task.TaskTypeId == models.BUG {
		switch stageId {
		case models.NEW:
			res = []int32{
				models.WORK,
			}
			break
		case models.WORK:
			res = []int32{
				models.READY_TO_TEST,
				models.REVIEW,
			}
			break
		case models.REVIEW:
			res = []int32{
				models.READY_TO_TEST,
			}
			break

		case models.READY_TO_TEST:
			res = []int32{
				models.TEST,
			}
			break
		case models.TEST:
			res = []int32{
				models.NEW,
				models.COMPLETED,
			}
			break
		}
	}
	if user.IsAdmin || user.HasRole(1, project) {
		res = append(res, models.BACKLOG)
		return res
	}

	if task.TaskTypeId == models.EPIC {
		if user.IsAdmin || user.HasRole(1, project) {
			res = f.PossibleProjectStages()
			res = append(res, models.BACKLOG)

			return res
		}
	}

	return res
}

func (f *DemoFlow) PossibleProjectStages() []int32 {
	return []int32{
		models.NEW,
		models.WORK,
		models.DONE,
		models.READY_TO_TEST,
		models.TEST,
		models.COMPLETED,
		models.REVIEW,
	}
}
