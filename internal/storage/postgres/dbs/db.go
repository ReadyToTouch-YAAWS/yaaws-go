// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dbs

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.companyTotalViewsStmt, err = db.PrepareContext(ctx, companyTotalViews); err != nil {
		return nil, fmt.Errorf("error preparing query CompanyTotalViews: %w", err)
	}
	if q.companyViewDailyStatsStmt, err = db.PrepareContext(ctx, companyViewDailyStats); err != nil {
		return nil, fmt.Errorf("error preparing query CompanyViewDailyStats: %w", err)
	}
	if q.companyViewDailyStatsUpsertStmt, err = db.PrepareContext(ctx, companyViewDailyStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query CompanyViewDailyStatsUpsert: %w", err)
	}
	if q.featureViewDailyStatsStmt, err = db.PrepareContext(ctx, featureViewDailyStats); err != nil {
		return nil, fmt.Errorf("error preparing query FeatureViewDailyStats: %w", err)
	}
	if q.featureViewDailyStatsUpsertStmt, err = db.PrepareContext(ctx, featureViewDailyStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query FeatureViewDailyStatsUpsert: %w", err)
	}
	if q.featureViewStatsStmt, err = db.PrepareContext(ctx, featureViewStats); err != nil {
		return nil, fmt.Errorf("error preparing query FeatureViewStats: %w", err)
	}
	if q.featureViewStatsUpsertStmt, err = db.PrepareContext(ctx, featureViewStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query FeatureViewStatsUpsert: %w", err)
	}
	if q.socialUserProfilesStmt, err = db.PrepareContext(ctx, socialUserProfiles); err != nil {
		return nil, fmt.Errorf("error preparing query SocialUserProfiles: %w", err)
	}
	if q.socialUserProfilesByUserStmt, err = db.PrepareContext(ctx, socialUserProfilesByUser); err != nil {
		return nil, fmt.Errorf("error preparing query SocialUserProfilesByUser: %w", err)
	}
	if q.userFavoriteCompaniesStmt, err = db.PrepareContext(ctx, userFavoriteCompanies); err != nil {
		return nil, fmt.Errorf("error preparing query UserFavoriteCompanies: %w", err)
	}
	if q.userFavoriteCompaniesStatsStmt, err = db.PrepareContext(ctx, userFavoriteCompaniesStats); err != nil {
		return nil, fmt.Errorf("error preparing query UserFavoriteCompaniesStats: %w", err)
	}
	if q.userFavoriteCompaniesUpsertStmt, err = db.PrepareContext(ctx, userFavoriteCompaniesUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserFavoriteCompaniesUpsert: %w", err)
	}
	if q.userFavoriteVacanciesStmt, err = db.PrepareContext(ctx, userFavoriteVacancies); err != nil {
		return nil, fmt.Errorf("error preparing query UserFavoriteVacancies: %w", err)
	}
	if q.userFavoriteVacanciesUpsertStmt, err = db.PrepareContext(ctx, userFavoriteVacanciesUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserFavoriteVacanciesUpsert: %w", err)
	}
	if q.userFeatureWaitlistStmt, err = db.PrepareContext(ctx, userFeatureWaitlist); err != nil {
		return nil, fmt.Errorf("error preparing query UserFeatureWaitlist: %w", err)
	}
	if q.userFeatureWaitlistDailyStatsStmt, err = db.PrepareContext(ctx, userFeatureWaitlistDailyStats); err != nil {
		return nil, fmt.Errorf("error preparing query UserFeatureWaitlistDailyStats: %w", err)
	}
	if q.userFeatureWaitlistStatsStmt, err = db.PrepareContext(ctx, userFeatureWaitlistStats); err != nil {
		return nil, fmt.Errorf("error preparing query UserFeatureWaitlistStats: %w", err)
	}
	if q.userFeatureWaitlistStatsUpsertStmt, err = db.PrepareContext(ctx, userFeatureWaitlistStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserFeatureWaitlistStatsUpsert: %w", err)
	}
	if q.userFeatureWaitlistUpsertStmt, err = db.PrepareContext(ctx, userFeatureWaitlistUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserFeatureWaitlistUpsert: %w", err)
	}
	if q.userOnlineDailyCountStatsStmt, err = db.PrepareContext(ctx, userOnlineDailyCountStats); err != nil {
		return nil, fmt.Errorf("error preparing query UserOnlineDailyCountStats: %w", err)
	}
	if q.userOnlineDailyCountStatsUpsertStmt, err = db.PrepareContext(ctx, userOnlineDailyCountStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserOnlineDailyCountStatsUpsert: %w", err)
	}
	if q.userOnlineDailyStatsUpsertStmt, err = db.PrepareContext(ctx, userOnlineDailyStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserOnlineDailyStatsUpsert: %w", err)
	}
	if q.userOnlineHourlyStatsStmt, err = db.PrepareContext(ctx, userOnlineHourlyStats); err != nil {
		return nil, fmt.Errorf("error preparing query UserOnlineHourlyStats: %w", err)
	}
	if q.userOnlineHourlyStatsUpsertStmt, err = db.PrepareContext(ctx, userOnlineHourlyStatsUpsert); err != nil {
		return nil, fmt.Errorf("error preparing query UserOnlineHourlyStatsUpsert: %w", err)
	}
	if q.userRegistrationDailyCountStatsStmt, err = db.PrepareContext(ctx, userRegistrationDailyCountStats); err != nil {
		return nil, fmt.Errorf("error preparing query UserRegistrationDailyCountStats: %w", err)
	}
	if q.userSocialProfileChangeHistoryNewStmt, err = db.PrepareContext(ctx, userSocialProfileChangeHistoryNew); err != nil {
		return nil, fmt.Errorf("error preparing query UserSocialProfileChangeHistoryNew: %w", err)
	}
	if q.userSocialProfileGetByIDStmt, err = db.PrepareContext(ctx, userSocialProfileGetByID); err != nil {
		return nil, fmt.Errorf("error preparing query UserSocialProfileGetByID: %w", err)
	}
	if q.userSocialProfileGetUserByEmailStmt, err = db.PrepareContext(ctx, userSocialProfileGetUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query UserSocialProfileGetUserByEmail: %w", err)
	}
	if q.userSocialProfileNewStmt, err = db.PrepareContext(ctx, userSocialProfileNew); err != nil {
		return nil, fmt.Errorf("error preparing query UserSocialProfileNew: %w", err)
	}
	if q.userSocialProfileUpdateStmt, err = db.PrepareContext(ctx, userSocialProfileUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query UserSocialProfileUpdate: %w", err)
	}
	if q.usersNewStmt, err = db.PrepareContext(ctx, usersNew); err != nil {
		return nil, fmt.Errorf("error preparing query UsersNew: %w", err)
	}
	if q.usersUpdateStmt, err = db.PrepareContext(ctx, usersUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query UsersUpdate: %w", err)
	}
	if q.wipLinkedInCompaniesGetByVanityNameStmt, err = db.PrepareContext(ctx, wipLinkedInCompaniesGetByVanityName); err != nil {
		return nil, fmt.Errorf("error preparing query WipLinkedInCompaniesGetByVanityName: %w", err)
	}
	if q.wipLinkedInCompaniesNewStmt, err = db.PrepareContext(ctx, wipLinkedInCompaniesNew); err != nil {
		return nil, fmt.Errorf("error preparing query WipLinkedInCompaniesNew: %w", err)
	}
	if q.wipLinkedInCompanyRequestHistoryCountStmt, err = db.PrepareContext(ctx, wipLinkedInCompanyRequestHistoryCount); err != nil {
		return nil, fmt.Errorf("error preparing query WipLinkedInCompanyRequestHistoryCount: %w", err)
	}
	if q.wipLinkedInCompanyRequestHistoryExistsVanityNameStmt, err = db.PrepareContext(ctx, wipLinkedInCompanyRequestHistoryExistsVanityName); err != nil {
		return nil, fmt.Errorf("error preparing query WipLinkedInCompanyRequestHistoryExistsVanityName: %w", err)
	}
	if q.wipLinkedInCompanyRequestHistoryNewStmt, err = db.PrepareContext(ctx, wipLinkedInCompanyRequestHistoryNew); err != nil {
		return nil, fmt.Errorf("error preparing query WipLinkedInCompanyRequestHistoryNew: %w", err)
	}
	if q.wipUserLinkedInCompaniesStmt, err = db.PrepareContext(ctx, wipUserLinkedInCompanies); err != nil {
		return nil, fmt.Errorf("error preparing query WipUserLinkedInCompanies: %w", err)
	}
	if q.wipUserToLinkedInCompaniesAddStmt, err = db.PrepareContext(ctx, wipUserToLinkedInCompaniesAdd); err != nil {
		return nil, fmt.Errorf("error preparing query WipUserToLinkedInCompaniesAdd: %w", err)
	}
	if q.wipUserToLinkedInCompaniesDeleteStmt, err = db.PrepareContext(ctx, wipUserToLinkedInCompaniesDelete); err != nil {
		return nil, fmt.Errorf("error preparing query WipUserToLinkedInCompaniesDelete: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.companyTotalViewsStmt != nil {
		if cerr := q.companyTotalViewsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing companyTotalViewsStmt: %w", cerr)
		}
	}
	if q.companyViewDailyStatsStmt != nil {
		if cerr := q.companyViewDailyStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing companyViewDailyStatsStmt: %w", cerr)
		}
	}
	if q.companyViewDailyStatsUpsertStmt != nil {
		if cerr := q.companyViewDailyStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing companyViewDailyStatsUpsertStmt: %w", cerr)
		}
	}
	if q.featureViewDailyStatsStmt != nil {
		if cerr := q.featureViewDailyStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing featureViewDailyStatsStmt: %w", cerr)
		}
	}
	if q.featureViewDailyStatsUpsertStmt != nil {
		if cerr := q.featureViewDailyStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing featureViewDailyStatsUpsertStmt: %w", cerr)
		}
	}
	if q.featureViewStatsStmt != nil {
		if cerr := q.featureViewStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing featureViewStatsStmt: %w", cerr)
		}
	}
	if q.featureViewStatsUpsertStmt != nil {
		if cerr := q.featureViewStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing featureViewStatsUpsertStmt: %w", cerr)
		}
	}
	if q.socialUserProfilesStmt != nil {
		if cerr := q.socialUserProfilesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing socialUserProfilesStmt: %w", cerr)
		}
	}
	if q.socialUserProfilesByUserStmt != nil {
		if cerr := q.socialUserProfilesByUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing socialUserProfilesByUserStmt: %w", cerr)
		}
	}
	if q.userFavoriteCompaniesStmt != nil {
		if cerr := q.userFavoriteCompaniesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFavoriteCompaniesStmt: %w", cerr)
		}
	}
	if q.userFavoriteCompaniesStatsStmt != nil {
		if cerr := q.userFavoriteCompaniesStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFavoriteCompaniesStatsStmt: %w", cerr)
		}
	}
	if q.userFavoriteCompaniesUpsertStmt != nil {
		if cerr := q.userFavoriteCompaniesUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFavoriteCompaniesUpsertStmt: %w", cerr)
		}
	}
	if q.userFavoriteVacanciesStmt != nil {
		if cerr := q.userFavoriteVacanciesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFavoriteVacanciesStmt: %w", cerr)
		}
	}
	if q.userFavoriteVacanciesUpsertStmt != nil {
		if cerr := q.userFavoriteVacanciesUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFavoriteVacanciesUpsertStmt: %w", cerr)
		}
	}
	if q.userFeatureWaitlistStmt != nil {
		if cerr := q.userFeatureWaitlistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFeatureWaitlistStmt: %w", cerr)
		}
	}
	if q.userFeatureWaitlistDailyStatsStmt != nil {
		if cerr := q.userFeatureWaitlistDailyStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFeatureWaitlistDailyStatsStmt: %w", cerr)
		}
	}
	if q.userFeatureWaitlistStatsStmt != nil {
		if cerr := q.userFeatureWaitlistStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFeatureWaitlistStatsStmt: %w", cerr)
		}
	}
	if q.userFeatureWaitlistStatsUpsertStmt != nil {
		if cerr := q.userFeatureWaitlistStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFeatureWaitlistStatsUpsertStmt: %w", cerr)
		}
	}
	if q.userFeatureWaitlistUpsertStmt != nil {
		if cerr := q.userFeatureWaitlistUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userFeatureWaitlistUpsertStmt: %w", cerr)
		}
	}
	if q.userOnlineDailyCountStatsStmt != nil {
		if cerr := q.userOnlineDailyCountStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userOnlineDailyCountStatsStmt: %w", cerr)
		}
	}
	if q.userOnlineDailyCountStatsUpsertStmt != nil {
		if cerr := q.userOnlineDailyCountStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userOnlineDailyCountStatsUpsertStmt: %w", cerr)
		}
	}
	if q.userOnlineDailyStatsUpsertStmt != nil {
		if cerr := q.userOnlineDailyStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userOnlineDailyStatsUpsertStmt: %w", cerr)
		}
	}
	if q.userOnlineHourlyStatsStmt != nil {
		if cerr := q.userOnlineHourlyStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userOnlineHourlyStatsStmt: %w", cerr)
		}
	}
	if q.userOnlineHourlyStatsUpsertStmt != nil {
		if cerr := q.userOnlineHourlyStatsUpsertStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userOnlineHourlyStatsUpsertStmt: %w", cerr)
		}
	}
	if q.userRegistrationDailyCountStatsStmt != nil {
		if cerr := q.userRegistrationDailyCountStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userRegistrationDailyCountStatsStmt: %w", cerr)
		}
	}
	if q.userSocialProfileChangeHistoryNewStmt != nil {
		if cerr := q.userSocialProfileChangeHistoryNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userSocialProfileChangeHistoryNewStmt: %w", cerr)
		}
	}
	if q.userSocialProfileGetByIDStmt != nil {
		if cerr := q.userSocialProfileGetByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userSocialProfileGetByIDStmt: %w", cerr)
		}
	}
	if q.userSocialProfileGetUserByEmailStmt != nil {
		if cerr := q.userSocialProfileGetUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userSocialProfileGetUserByEmailStmt: %w", cerr)
		}
	}
	if q.userSocialProfileNewStmt != nil {
		if cerr := q.userSocialProfileNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userSocialProfileNewStmt: %w", cerr)
		}
	}
	if q.userSocialProfileUpdateStmt != nil {
		if cerr := q.userSocialProfileUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing userSocialProfileUpdateStmt: %w", cerr)
		}
	}
	if q.usersNewStmt != nil {
		if cerr := q.usersNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing usersNewStmt: %w", cerr)
		}
	}
	if q.usersUpdateStmt != nil {
		if cerr := q.usersUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing usersUpdateStmt: %w", cerr)
		}
	}
	if q.wipLinkedInCompaniesGetByVanityNameStmt != nil {
		if cerr := q.wipLinkedInCompaniesGetByVanityNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipLinkedInCompaniesGetByVanityNameStmt: %w", cerr)
		}
	}
	if q.wipLinkedInCompaniesNewStmt != nil {
		if cerr := q.wipLinkedInCompaniesNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipLinkedInCompaniesNewStmt: %w", cerr)
		}
	}
	if q.wipLinkedInCompanyRequestHistoryCountStmt != nil {
		if cerr := q.wipLinkedInCompanyRequestHistoryCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipLinkedInCompanyRequestHistoryCountStmt: %w", cerr)
		}
	}
	if q.wipLinkedInCompanyRequestHistoryExistsVanityNameStmt != nil {
		if cerr := q.wipLinkedInCompanyRequestHistoryExistsVanityNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipLinkedInCompanyRequestHistoryExistsVanityNameStmt: %w", cerr)
		}
	}
	if q.wipLinkedInCompanyRequestHistoryNewStmt != nil {
		if cerr := q.wipLinkedInCompanyRequestHistoryNewStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipLinkedInCompanyRequestHistoryNewStmt: %w", cerr)
		}
	}
	if q.wipUserLinkedInCompaniesStmt != nil {
		if cerr := q.wipUserLinkedInCompaniesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipUserLinkedInCompaniesStmt: %w", cerr)
		}
	}
	if q.wipUserToLinkedInCompaniesAddStmt != nil {
		if cerr := q.wipUserToLinkedInCompaniesAddStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipUserToLinkedInCompaniesAddStmt: %w", cerr)
		}
	}
	if q.wipUserToLinkedInCompaniesDeleteStmt != nil {
		if cerr := q.wipUserToLinkedInCompaniesDeleteStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing wipUserToLinkedInCompaniesDeleteStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                                                   DBTX
	tx                                                   *sql.Tx
	companyTotalViewsStmt                                *sql.Stmt
	companyViewDailyStatsStmt                            *sql.Stmt
	companyViewDailyStatsUpsertStmt                      *sql.Stmt
	featureViewDailyStatsStmt                            *sql.Stmt
	featureViewDailyStatsUpsertStmt                      *sql.Stmt
	featureViewStatsStmt                                 *sql.Stmt
	featureViewStatsUpsertStmt                           *sql.Stmt
	socialUserProfilesStmt                               *sql.Stmt
	socialUserProfilesByUserStmt                         *sql.Stmt
	userFavoriteCompaniesStmt                            *sql.Stmt
	userFavoriteCompaniesStatsStmt                       *sql.Stmt
	userFavoriteCompaniesUpsertStmt                      *sql.Stmt
	userFavoriteVacanciesStmt                            *sql.Stmt
	userFavoriteVacanciesUpsertStmt                      *sql.Stmt
	userFeatureWaitlistStmt                              *sql.Stmt
	userFeatureWaitlistDailyStatsStmt                    *sql.Stmt
	userFeatureWaitlistStatsStmt                         *sql.Stmt
	userFeatureWaitlistStatsUpsertStmt                   *sql.Stmt
	userFeatureWaitlistUpsertStmt                        *sql.Stmt
	userOnlineDailyCountStatsStmt                        *sql.Stmt
	userOnlineDailyCountStatsUpsertStmt                  *sql.Stmt
	userOnlineDailyStatsUpsertStmt                       *sql.Stmt
	userOnlineHourlyStatsStmt                            *sql.Stmt
	userOnlineHourlyStatsUpsertStmt                      *sql.Stmt
	userRegistrationDailyCountStatsStmt                  *sql.Stmt
	userSocialProfileChangeHistoryNewStmt                *sql.Stmt
	userSocialProfileGetByIDStmt                         *sql.Stmt
	userSocialProfileGetUserByEmailStmt                  *sql.Stmt
	userSocialProfileNewStmt                             *sql.Stmt
	userSocialProfileUpdateStmt                          *sql.Stmt
	usersNewStmt                                         *sql.Stmt
	usersUpdateStmt                                      *sql.Stmt
	wipLinkedInCompaniesGetByVanityNameStmt              *sql.Stmt
	wipLinkedInCompaniesNewStmt                          *sql.Stmt
	wipLinkedInCompanyRequestHistoryCountStmt            *sql.Stmt
	wipLinkedInCompanyRequestHistoryExistsVanityNameStmt *sql.Stmt
	wipLinkedInCompanyRequestHistoryNewStmt              *sql.Stmt
	wipUserLinkedInCompaniesStmt                         *sql.Stmt
	wipUserToLinkedInCompaniesAddStmt                    *sql.Stmt
	wipUserToLinkedInCompaniesDeleteStmt                 *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                                                   tx,
		tx:                                                   tx,
		companyTotalViewsStmt:                                q.companyTotalViewsStmt,
		companyViewDailyStatsStmt:                            q.companyViewDailyStatsStmt,
		companyViewDailyStatsUpsertStmt:                      q.companyViewDailyStatsUpsertStmt,
		featureViewDailyStatsStmt:                            q.featureViewDailyStatsStmt,
		featureViewDailyStatsUpsertStmt:                      q.featureViewDailyStatsUpsertStmt,
		featureViewStatsStmt:                                 q.featureViewStatsStmt,
		featureViewStatsUpsertStmt:                           q.featureViewStatsUpsertStmt,
		socialUserProfilesStmt:                               q.socialUserProfilesStmt,
		socialUserProfilesByUserStmt:                         q.socialUserProfilesByUserStmt,
		userFavoriteCompaniesStmt:                            q.userFavoriteCompaniesStmt,
		userFavoriteCompaniesStatsStmt:                       q.userFavoriteCompaniesStatsStmt,
		userFavoriteCompaniesUpsertStmt:                      q.userFavoriteCompaniesUpsertStmt,
		userFavoriteVacanciesStmt:                            q.userFavoriteVacanciesStmt,
		userFavoriteVacanciesUpsertStmt:                      q.userFavoriteVacanciesUpsertStmt,
		userFeatureWaitlistStmt:                              q.userFeatureWaitlistStmt,
		userFeatureWaitlistDailyStatsStmt:                    q.userFeatureWaitlistDailyStatsStmt,
		userFeatureWaitlistStatsStmt:                         q.userFeatureWaitlistStatsStmt,
		userFeatureWaitlistStatsUpsertStmt:                   q.userFeatureWaitlistStatsUpsertStmt,
		userFeatureWaitlistUpsertStmt:                        q.userFeatureWaitlistUpsertStmt,
		userOnlineDailyCountStatsStmt:                        q.userOnlineDailyCountStatsStmt,
		userOnlineDailyCountStatsUpsertStmt:                  q.userOnlineDailyCountStatsUpsertStmt,
		userOnlineDailyStatsUpsertStmt:                       q.userOnlineDailyStatsUpsertStmt,
		userOnlineHourlyStatsStmt:                            q.userOnlineHourlyStatsStmt,
		userOnlineHourlyStatsUpsertStmt:                      q.userOnlineHourlyStatsUpsertStmt,
		userRegistrationDailyCountStatsStmt:                  q.userRegistrationDailyCountStatsStmt,
		userSocialProfileChangeHistoryNewStmt:                q.userSocialProfileChangeHistoryNewStmt,
		userSocialProfileGetByIDStmt:                         q.userSocialProfileGetByIDStmt,
		userSocialProfileGetUserByEmailStmt:                  q.userSocialProfileGetUserByEmailStmt,
		userSocialProfileNewStmt:                             q.userSocialProfileNewStmt,
		userSocialProfileUpdateStmt:                          q.userSocialProfileUpdateStmt,
		usersNewStmt:                                         q.usersNewStmt,
		usersUpdateStmt:                                      q.usersUpdateStmt,
		wipLinkedInCompaniesGetByVanityNameStmt:              q.wipLinkedInCompaniesGetByVanityNameStmt,
		wipLinkedInCompaniesNewStmt:                          q.wipLinkedInCompaniesNewStmt,
		wipLinkedInCompanyRequestHistoryCountStmt:            q.wipLinkedInCompanyRequestHistoryCountStmt,
		wipLinkedInCompanyRequestHistoryExistsVanityNameStmt: q.wipLinkedInCompanyRequestHistoryExistsVanityNameStmt,
		wipLinkedInCompanyRequestHistoryNewStmt:              q.wipLinkedInCompanyRequestHistoryNewStmt,
		wipUserLinkedInCompaniesStmt:                         q.wipUserLinkedInCompaniesStmt,
		wipUserToLinkedInCompaniesAddStmt:                    q.wipUserToLinkedInCompaniesAddStmt,
		wipUserToLinkedInCompaniesDeleteStmt:                 q.wipUserToLinkedInCompaniesDeleteStmt,
	}
}
