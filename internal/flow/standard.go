package flow

import "gokanban/internal/domain/models"

type FlowInterface interface {
	name() string
	slug() string
	PossibleTaskNextStages(task *models.Task, user *models.User, project *models.Project) []int32
	PossibleProjectStages() []int32
}

type StandardFlow struct {
}

func NewStandardFlow() *StandardFlow {
	return &StandardFlow{}
}

func (f *StandardFlow) name() string {
	return "Стандартный"
}

func (f *StandardFlow) slug() string {
	return "standard"
}

func (f *StandardFlow) PossibleTaskNextStages(task *models.Task, user *models.User, project *models.Project) []int32 {
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

func (f *StandardFlow) PossibleProjectStages() []int32 {
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

//public function possibleTaskNextStages(object $task, User $user): array
//{
//$stageId = $task->stage_id;
////Из беклога только в новые.
//if ($stageId == Stage::BACKLOG) {
//return $user->is_admin || $user->hasRole($task->project_id, Role::MANAGER) ? [Stage::NEW] : [];
//}
//
//if ($user->is_admin || $task->project->shouldTreatUserAsAdmin($user)) {
//// Админы могут перетаскивать задачи как угодно
//return array_merge($this->possibleProjectStages(), [Stage::BACKLOG]);
//}
//
//// ТАБЛИЦА ПЕРЕХОДОВ ДЛЯ БАГОВ
//$stages = [];
//if ($task->task_type_id === TaskType::BUG) {
//switch ($stageId) {
//case Stage::NEW:
//$stages = [Stage::WORK];
//break;
//case Stage::WORK:
//$stages = [Stage::READY_TO_TEST, Stage::REVIEW];
//break;
//case Stage::REVIEW:
//$stages = [Stage::READY_TO_TEST];
//break;
//case Stage::READY_TO_TEST:
//$stages = [Stage::TEST];
//break;
//case Stage::TEST:
//$stages = [Stage::NEW, Stage::COMPLETED];
//break;
//}
//
//if ($user->is_admin || $user->hasRole($task->project_id, Role::MANAGER)) {
//return array_merge($stages, [Stage::BACKLOG]);
//}
//
//return $stages;
//}
//
//if($task->task_type_id === TaskType::EPIC) {
//if($user->is_admin || $user->hasRole($task->project_id, Role::MANAGER)) {
//return array_merge($this->possibleProjectStages(), [Stage::BACKLOG]);
//}
//
//return [];
//}
//
//$stages = [];
//
//switch ($task->component_id) {
//// КОМПОНЕНТ ДИЗАЙН
//case Component::DESIGN:
//switch ($stageId) {
//case Stage::NEW:
//$stages = [Stage::WORK];
//break;
//case Stage::WORK:
//$stages = [Stage::DONE];
//break;
//case Stage::DONE:
//$stages = [Stage::WORK, Stage::COMPLETED];
//break;
//case Stage::READY_TO_TEST:
//$stages = [Stage::TEST];
//break;
//case Stage::TEST:
//$stages = [Stage::COMPLETED];
//break;
//}
//break;
//
//// КОМПОНЕНТ АДМИНИСТРИРОВАНИЕ
//case Component::ADMINISTERING:
//// КОМПОНЕНТ ВЁРСТКА
//case Component::MARKUP:
//// КОМПОНЕНТ ФРОНТЕНД-РАЗРАБОТКА
//case Component::FRONTEND:
//// КОМПОНЕНТ РАЗРАБОТКА
//case Component::DEV:
//// Аналитика
//case Component::ANALYSIS:
//// Копирайт
//case Component::COPYRIGHT:
//// КОМПОНЕНТ КОНТЕНТ
//case Component::CONTENT:
//switch ($stageId) {
//case Stage::NEW:
//$stages = [Stage::WORK];
//break;
//case Stage::WORK:
//$stages = [Stage::DONE];
//break;
//case Stage::DONE:
//$stages = [Stage::WORK, Stage::READY_TO_TEST, Stage::REVIEW];
//break;
//case Stage::REVIEW:
//$stages = [Stage::WORK, Stage::READY_TO_TEST];
//break;
//case Stage::READY_TO_TEST:
//$stages = [Stage::TEST];
//break;
//case Stage::TEST:
//$stages = [Stage::COMPLETED];
//break;
//}
//break;
//
//// КОМПОНЕНТ ТЕСТИРОВАНИЕ
//case Component::TEST:
//switch ($stageId) {
//case Stage::NEW:
//$stages = [Stage::TEST];
//break;
//case Stage::TEST:
//$stages = [Stage::COMPLETED];
//break;
//}
//
//// КОМПОНЕНТ МЕНЕДЖМЕНТ
//case Component::MANAGEMENT:
//switch ($stageId) {
//case Stage::NEW:
//$stages = [Stage::WORK];
//break;
//case Stage::WORK:
//$stages = [Stage::COMPLETED];
//break;
//}
//}
//
//if ($user->is_admin || $user->hasRole($task->project_id, Role::MANAGER)) {
//return array_merge($stages, [Stage::BACKLOG]);
//}
//
//////        $project = $task->project;
////        if (!empty($stages)
////            && $project->perm_user_to_rft
////            && in_array($stageId, [Stage::WORK, Stage::DONE, Stage::REVIEW])
////        ) {
////            $stages[] = Stage::READY_TO_TEST;
////
////            $doneCount = $this->tasksCount[$project->id][Stage::DONE] ?? null;
////            if ($doneCount === null) {
////                $doneCount = $project->tasks()->where('stage_id', Stage::DONE)->count();
////                $this->tasksCount[$project->id][Stage::DONE] = $doneCount;
////            }
////
////            if ($doneCount === 0) {
////                $stages = array_diff($stages, [Stage::DONE]);
////            }
////        }
//
//return $stages;
//}
