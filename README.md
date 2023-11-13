# work1

/query中的sql文件，对应的postgresql实现

```sql
-- name: CreateDevice :one
INSERT INTO devices(
    device_name,
    device_manufacturer,
    device_origin,
    production_date,
    testing_date,
    device_model
) VALUES (
    '检测仪', '厂商Z', '深圳', '2019-01-20T09:10:50Z', '2023-02-24T09:10:50Z', '型号Z'
) RETURNING *;
-- name: ListDevice :one
SELECT * FROM devices
WHERE id=2 LIMIT 1;
-- name: ListRecords :many
SELECT *
FROM devices
JOIN results ON devices.id = results.devices_id
JOIN tests ON devices.id = tests.devices_id
WHERE devices.id = 2 AND results.id =  tests.id 


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
    2,'PT2',5,100,19,100,50
)RETURNING *;
-- name: ListTests :many
SELECT * FROM tests
WHERE devices_id =2
ORDER BY id


-- name: CreateResult :one
INSERT INTO results(
    test_id,
    devices_id,
    voltage,
    point_number,
    temperature,
    humidity
)VALUES(
    5,2,20,1005,20,30
)RETURNING *;
-- name: ListResults :many
SELECT * FROM results
WHERE devices_id =2
ORDER BY id
```