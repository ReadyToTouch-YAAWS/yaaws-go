// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: vacancy_view_stats.sql

package dbs

import (
	"context"
	"time"
)

const userVacancyViewDailyStatsUpsert = `-- name: UserVacancyViewDailyStatsUpsert :execrows
INSERT INTO user_vacancy_view_daily_stats (vacancy_id, created_at, user_id)
VALUES ($1, $2, $3)
ON CONFLICT (vacancy_id, created_at, user_id) DO NOTHING
`

type UserVacancyViewDailyStatsUpsertParams struct {
	VacancyID int64
	CreatedAt time.Time
	UserID    int64
}

func (q *Queries) UserVacancyViewDailyStatsUpsert(ctx context.Context, arg UserVacancyViewDailyStatsUpsertParams) (int64, error) {
	result, err := q.exec(ctx, q.userVacancyViewDailyStatsUpsertStmt, userVacancyViewDailyStatsUpsert, arg.VacancyID, arg.CreatedAt, arg.UserID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const userVacancyViewHourlyStatsUpsert = `-- name: UserVacancyViewHourlyStatsUpsert :execrows
INSERT INTO user_vacancy_view_hourly_stats (vacancy_id, created_at, user_id)
VALUES ($1, $2, $3)
ON CONFLICT (vacancy_id, created_at, user_id) DO NOTHING
`

type UserVacancyViewHourlyStatsUpsertParams struct {
	VacancyID int64
	CreatedAt time.Time
	UserID    int64
}

func (q *Queries) UserVacancyViewHourlyStatsUpsert(ctx context.Context, arg UserVacancyViewHourlyStatsUpsertParams) (int64, error) {
	result, err := q.exec(ctx, q.userVacancyViewHourlyStatsUpsertStmt, userVacancyViewHourlyStatsUpsert, arg.VacancyID, arg.CreatedAt, arg.UserID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const vacancyViewDailyStatsUpsert = `-- name: VacancyViewDailyStatsUpsert :exec
INSERT INTO vacancy_view_daily_stats AS t (vacancy_id, created_at, user_count)
VALUES ($1, $2, 1)
ON CONFLICT (vacancy_id, created_at) DO UPDATE
    SET user_count = t.user_count + 1
`

type VacancyViewDailyStatsUpsertParams struct {
	VacancyID int64
	CreatedAt time.Time
}

func (q *Queries) VacancyViewDailyStatsUpsert(ctx context.Context, arg VacancyViewDailyStatsUpsertParams) error {
	_, err := q.exec(ctx, q.vacancyViewDailyStatsUpsertStmt, vacancyViewDailyStatsUpsert, arg.VacancyID, arg.CreatedAt)
	return err
}
