package postgres

import (
	"context"

	"github.com/readytotouch-yaaws/yaaws-go/internal/storage/postgres/dbs"
)

type BatchUpdateOnlineStorage struct {
	db *Database
}

func NewBatchUpdateOnlineStorage(db *Database) *BatchUpdateOnlineStorage {
	return &BatchUpdateOnlineStorage{db: db}
}

func (s *BatchUpdateOnlineStorage) BatchStore(ctx context.Context, pairs []UserOnlinePair) error {
	// user_online_hourly_stats
	{
		hourUserIDs, hourOnlineTimestamps := toHourlyStats(pairs)
		rowsAffected, err := s.db.Queries().UserOnlineHourlyStatsUpsert(ctx, dbs.UserOnlineHourlyStatsUpsertParams{
			UserIds:    hourUserIDs,
			CreatedAts: hourOnlineTimestamps,
		})
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
	}

	// user_online_daily_stats
	{
		dayUserIDs, dayOnlineTimestamps := toDailyStats(pairs)
		rowsAffected, err := s.db.Queries().UserOnlineDailyStatsUpsert(ctx, dbs.UserOnlineDailyStatsUpsertParams{
			UserIds:    dayUserIDs,
			CreatedAts: dayOnlineTimestamps,
		})
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			return nil
		}
	}

	// user_online_daily_count_stats
	{
		err := s.db.Queries().UserOnlineDailyCountStatsUpsert(ctx, toDailyMin(pairs))
		if err != nil {
			return err
		}
	}

	return nil
}
