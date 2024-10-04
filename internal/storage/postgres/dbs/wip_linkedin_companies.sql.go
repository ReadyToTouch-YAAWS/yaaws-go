// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: wip_linkedin_companies.sql

package dbs

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

const wipLinkedInCompaniesGetByVanityName = `-- name: WipLinkedInCompaniesGetByVanityName :one
SELECT id
FROM wip_linkedin_companies
WHERE vanity_name = $1
`

func (q *Queries) WipLinkedInCompaniesGetByVanityName(ctx context.Context, vanityName string) (int64, error) {
	row := q.queryRow(ctx, q.wipLinkedInCompaniesGetByVanityNameStmt, wipLinkedInCompaniesGetByVanityName, vanityName)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const wipLinkedInCompaniesNew = `-- name: WipLinkedInCompaniesNew :exec
INSERT INTO wip_linkedin_companies (id, vanity_name, name, payload, created_at, created_by)
VALUES ($1, $2, $3, $4, $5, $6)
`

type WipLinkedInCompaniesNewParams struct {
	ID         int64
	VanityName string
	Name       string
	Payload    json.RawMessage
	CreatedAt  time.Time
	CreatedBy  int64
}

func (q *Queries) WipLinkedInCompaniesNew(ctx context.Context, arg WipLinkedInCompaniesNewParams) error {
	_, err := q.exec(ctx, q.wipLinkedInCompaniesNewStmt, wipLinkedInCompaniesNew,
		arg.ID,
		arg.VanityName,
		arg.Name,
		arg.Payload,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	return err
}

const wipLinkedInCompanyRequestHistoryCount = `-- name: WipLinkedInCompanyRequestHistoryCount :one
SELECT COUNT(*) AS total
FROM wip_linkedin_company_request_history
WHERE created_by = $1
`

func (q *Queries) WipLinkedInCompanyRequestHistoryCount(ctx context.Context, createdBy int64) (int64, error) {
	row := q.queryRow(ctx, q.wipLinkedInCompanyRequestHistoryCountStmt, wipLinkedInCompanyRequestHistoryCount, createdBy)
	var total int64
	err := row.Scan(&total)
	return total, err
}

const wipLinkedInCompanyRequestHistoryExistsVanityName = `-- name: WipLinkedInCompanyRequestHistoryExistsVanityName :one
SELECT EXISTS(
    SELECT
    FROM wip_linkedin_company_request_history
    WHERE vanity_name = $1
)
`

func (q *Queries) WipLinkedInCompanyRequestHistoryExistsVanityName(ctx context.Context, vanityName string) (bool, error) {
	row := q.queryRow(ctx, q.wipLinkedInCompanyRequestHistoryExistsVanityNameStmt, wipLinkedInCompanyRequestHistoryExistsVanityName, vanityName)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const wipLinkedInCompanyRequestHistoryNew = `-- name: WipLinkedInCompanyRequestHistoryNew :exec
INSERT INTO wip_linkedin_company_request_history (vanity_name, linkedin_company_id, created_at, created_by)
VALUES ($1, $2, $3, $4)
`

type WipLinkedInCompanyRequestHistoryNewParams struct {
	VanityName string
	ID         sql.NullInt64
	CreatedAt  time.Time
	CreatedBy  int64
}

func (q *Queries) WipLinkedInCompanyRequestHistoryNew(ctx context.Context, arg WipLinkedInCompanyRequestHistoryNewParams) error {
	_, err := q.exec(ctx, q.wipLinkedInCompanyRequestHistoryNewStmt, wipLinkedInCompanyRequestHistoryNew,
		arg.VanityName,
		arg.ID,
		arg.CreatedAt,
		arg.CreatedBy,
	)
	return err
}

const wipUserLinkedInCompanies = `-- name: WipUserLinkedInCompanies :many
SELECT wlc.id,
       wlc.vanity_name,
       wlc.name
FROM wip_user_to_linkedin_companies wu2lc
         INNER JOIN wip_linkedin_companies wlc ON wu2lc.linkedin_company_id = wlc.id
WHERE wu2lc.user_id = $1
  AND wu2lc.active = TRUE
ORDER BY wu2lc.updated_at DESC
`

type WipUserLinkedInCompaniesRow struct {
	ID         int64
	VanityName string
	Name       string
}

func (q *Queries) WipUserLinkedInCompanies(ctx context.Context, userID int64) ([]WipUserLinkedInCompaniesRow, error) {
	rows, err := q.query(ctx, q.wipUserLinkedInCompaniesStmt, wipUserLinkedInCompanies, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []WipUserLinkedInCompaniesRow
	for rows.Next() {
		var i WipUserLinkedInCompaniesRow
		if err := rows.Scan(&i.ID, &i.VanityName, &i.Name); err != nil {
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

const wipUserToLinkedInCompaniesAdd = `-- name: WipUserToLinkedInCompaniesAdd :exec
INSERT INTO wip_user_to_linkedin_companies AS t (user_id, linkedin_company_id, active, created_at, created_by,
                                                 updated_at, updated_by)
VALUES ($1, $2, TRUE, $3, $1,
        $3, $1)
ON CONFLICT (user_id, linkedin_company_id) DO UPDATE
    SET active     = TRUE,
        updated_at = excluded.updated_at,
        updated_by = excluded.updated_by
`

type WipUserToLinkedInCompaniesAddParams struct {
	CreatedBy         int64
	LinkedinCompanyID int64
	CreatedAt         time.Time
}

func (q *Queries) WipUserToLinkedInCompaniesAdd(ctx context.Context, arg WipUserToLinkedInCompaniesAddParams) error {
	_, err := q.exec(ctx, q.wipUserToLinkedInCompaniesAddStmt, wipUserToLinkedInCompaniesAdd, arg.CreatedBy, arg.LinkedinCompanyID, arg.CreatedAt)
	return err
}

const wipUserToLinkedInCompaniesDelete = `-- name: WipUserToLinkedInCompaniesDelete :exec
UPDATE wip_user_to_linkedin_companies
SET active     = FALSE,
    updated_at = $1,
    updated_by = $2
WHERE user_id = $3
  AND linkedin_company_id = $4
`

type WipUserToLinkedInCompaniesDeleteParams struct {
	UpdatedAt         time.Time
	UpdatedBy         int64
	UserID            int64
	LinkedinCompanyID int64
}

func (q *Queries) WipUserToLinkedInCompaniesDelete(ctx context.Context, arg WipUserToLinkedInCompaniesDeleteParams) error {
	_, err := q.exec(ctx, q.wipUserToLinkedInCompaniesDeleteStmt, wipUserToLinkedInCompaniesDelete,
		arg.UpdatedAt,
		arg.UpdatedBy,
		arg.UserID,
		arg.LinkedinCompanyID,
	)
	return err
}
