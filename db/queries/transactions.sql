-- name: CreateTransaction :one
INSERT INTO transactions (
        user_id,
        kind,
        amount,
        currency_code,
        category_id,
        description
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAllUserTransactions :many
SELECT *
FROM transactions
WHERE user_id = $1;