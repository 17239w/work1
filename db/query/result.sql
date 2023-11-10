-- name: CreateResult :one
INSERT INTO results(
    voltage,
    point_number,
    test_id,
    temperature,
    humidity
)VALUES(
    $1,$2,$3,$4,$5
)RETURNING *;

-- name: ListResults :many
SELECT * FROM results
WHERE devices_id =$1
ORDER BY id
LIMIT $2
OFFSET $3;
