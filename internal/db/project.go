package db

import (
	"context"
)

type ListProjectResult struct {
	Name       string  `json:"name"`
	Slug       string  `json:"slug"`
	LogoPath   string  `json:"logo_path"`
	Id         int64   `json:"id"`
	IsArchived int     `json:"is_archived"`
	Begin      *string `json:"begin"`
	End        *string `json:"end"`
	UserCount  int64   `json:"user_count"`
	IsFavorite int64   `json:"is_favorite"`
}

const projectListQuery = "select `files`.`path` as `logo_path`," +
	" `projects`.`id` as Id," +
	" `projects`.`name` as Name," +
	" `projects`.`slug` as Slug," +
	" `projects`.`begin` as Begin," +
	" `projects`.`is_archived` as IsArchived," +
	" `projects`.`end` as End , " +
	"(select count(1) from project_user where project_id = projects.id) as UserCount, " +
	"count(contain_favorites.id) as IsFavorite " +
	"from `projects` " +
	"inner join `files` on `projects`.`logo_id` = `files`.`id` " +
	"inner join `project_user` on `projects`.`id` = `project_user`.`project_id` " +
	"left join `contain_favorites` on `projects`.`id` = contain_favorites.contain_favorite_id and contain_favorites.user_id = ? " +
	"where `projects`.`deleted_at` is null group by `projects`.`id`"

func (q *Queries) GetListProjects(ctx context.Context, userId int) ([]ListProjectResult, error) {
	rows, err := q.db.QueryContext(ctx, projectListQuery, userId)
	if err != nil {
		return nil, err
	}

	var items []ListProjectResult
	for rows.Next() {
		var item ListProjectResult
		rows.Scan(&item.LogoPath, &item.Id, &item.Name, &item.Slug, &item.Begin, &item.IsArchived, &item.End, &item.UserCount, &item.IsFavorite)
		if item.IsFavorite > 0 {
			item.IsFavorite = 1
		}
		items = append(items, item)
	}
	rows.Close()

	return items, err
}
