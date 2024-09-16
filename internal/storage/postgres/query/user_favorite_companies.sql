-- name: UserFavoriteCompaniesUpsert :exec
INSERT INTO user_favorite_companies AS t (user_id, company_id, favorite, created_at, updated_at)
VALUES (@user_id, @company_id, @favorite, @created_at, @updated_at)
ON CONFLICT (user_id, company_id)
    DO UPDATE
    SET favorite   = excluded.favorite,
        updated_at = excluded.updated_at
WHERE t.favorite <> excluded.favorite;

-- name: UserFavoriteCompanies :many
SELECT company_id
FROM user_favorite_companies
WHERE user_id = @user_id
  AND favorite = TRUE;

-- name: UserFavoriteCompaniesStats :one
SELECT COUNT(*)                                                             AS count,
       COUNT(*) FILTER ( WHERE created_at::DATE >= sqlc.arg('from')::DATE ) AS count_since
FROM user_favorite_companies
WHERE company_id = @company_id
  AND favorite = TRUE;
