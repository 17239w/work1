-- name: CreateTest :one
INSERT INTO tests(
    devices_id,
    test_name,
    gear,
    percentage,
    lower_limit,
    upper_limit,
    test_data
) VALUES(
    $1,$2,$3,$4,$5,$6,$7
)RETURNING *;

-- name: ListTests :many
SELECT * FROM tests
WHERE devices_id =$1
ORDER BY id
LIMIT $2
OFFSET $3;

