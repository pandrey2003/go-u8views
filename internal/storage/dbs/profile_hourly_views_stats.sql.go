// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: profile_hourly_views_stats.sql

package dbs

import (
	"context"
	"time"

	"github.com/lib/pq"
)

const profileHourlyViewsStats = `-- name: ProfileHourlyViewsStats :many
SELECT user_id,
       SUM(CASE WHEN time >= $1 THEN count ELSE 0 END)::BIGINT  AS day_count,
       SUM(CASE WHEN time >= $2 THEN count ELSE 0 END)::BIGINT AS week_count,
       SUM(count)::BIGINT                                         AS month_count
FROM profile_hourly_views_stats
WHERE user_id = ANY ($3::BIGINT[])
  AND time >= $4
GROUP BY user_id
`

type ProfileHourlyViewsStatsParams struct {
	Day     time.Time
	Week    time.Time
	UserIds []int64
	Month   time.Time
}

type ProfileHourlyViewsStatsRow struct {
	UserID     int64
	DayCount   int64
	WeekCount  int64
	MonthCount int64
}

func (q *Queries) ProfileHourlyViewsStats(ctx context.Context, arg ProfileHourlyViewsStatsParams) ([]ProfileHourlyViewsStatsRow, error) {
	rows, err := q.query(ctx, q.profileHourlyViewsStatsStmt, profileHourlyViewsStats,
		arg.Day,
		arg.Week,
		pq.Array(arg.UserIds),
		arg.Month,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProfileHourlyViewsStatsRow
	for rows.Next() {
		var i ProfileHourlyViewsStatsRow
		if err := rows.Scan(
			&i.UserID,
			&i.DayCount,
			&i.WeekCount,
			&i.MonthCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const profileHourlyViewsStatsByHour = `-- name: ProfileHourlyViewsStatsByHour :many
SELECT g.time                          AS time,
       COALESCE(phvs.count, 0)::BIGINT AS count
FROM (
    SELECT time::TIMESTAMP
    FROM generate_series(
        $1::TIMESTAMP,
        $2::TIMESTAMP,
        '1 HOUR'::INTERVAL
    ) AS time
) AS g
    LEFT JOIN (
        SELECT time,
               count
        FROM profile_hourly_views_stats
        WHERE user_id = $3
          AND time >= $1::TIMESTAMP
    ) AS phvs ON (g.time = phvs.time)
ORDER BY g.time
`

type ProfileHourlyViewsStatsByHourParams struct {
	From   time.Time
	To     time.Time
	UserID int64
}

type ProfileHourlyViewsStatsByHourRow struct {
	Time  time.Time
	Count int64
}

func (q *Queries) ProfileHourlyViewsStatsByHour(ctx context.Context, arg ProfileHourlyViewsStatsByHourParams) ([]ProfileHourlyViewsStatsByHourRow, error) {
	rows, err := q.query(ctx, q.profileHourlyViewsStatsByHourStmt, profileHourlyViewsStatsByHour, arg.From, arg.To, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ProfileHourlyViewsStatsByHourRow
	for rows.Next() {
		var i ProfileHourlyViewsStatsByHourRow
		if err := rows.Scan(&i.Time, &i.Count); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const profileHourlyViewsStatsUpsert = `-- name: ProfileHourlyViewsStatsUpsert :exec
INSERT INTO profile_hourly_views_stats (user_id, time, count)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, time) DO UPDATE
    SET count = profile_hourly_views_stats.count + $3
`

type ProfileHourlyViewsStatsUpsertParams struct {
	UserID int64
	Time   time.Time
	Count  int64
}

func (q *Queries) ProfileHourlyViewsStatsUpsert(ctx context.Context, arg ProfileHourlyViewsStatsUpsertParams) error {
	_, err := q.exec(ctx, q.profileHourlyViewsStatsUpsertStmt, profileHourlyViewsStatsUpsert, arg.UserID, arg.Time, arg.Count)
	return err
}
