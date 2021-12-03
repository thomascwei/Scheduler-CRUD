-- name: ListSchedule :many
SELECT *
FROM schedules
ORDER BY id;

-- name: ListRepeatWeekdays :many
SELECT schedule_id, weekday
FROM repeat_weekday
ORDER BY id;

-- name: ListRepeatDays :many
SELECT schedule_id, day
FROM repeat_day
ORDER BY id;

-- name: ListRepeatMonths :many
SELECT schedule_id, month
FROM repeat_month
ORDER BY id;

-- name: GetSchedule :one
SELECT *
FROM schedules
WHERE id = ?
LIMIT 1;

-- name: GetDays :many
SELECT *
FROM repeat_day
WHERE schedule_id = ?;

-- name: GetWeekdays :many
SELECT *
FROM repeat_weekday
WHERE schedule_id = ?;

-- name: GetMonths :many
SELECT *
FROM repeat_month
WHERE schedule_id = ?;

-- name: GetCommands :many
SELECT *
FROM commands;

-- name: CreateSchedule :execresult
INSERT INTO schedules (time_type_id, interval_day, interval_seconds, at_time, start_time, end_time,
                       command_id, name, start_date, end_date, enable, `repeat`)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);


-- name: UpdateSchedule :exec
UPDATE schedules
SET time_type_id=?,
    interval_day=?,
    interval_seconds=?,
    at_time=?,
    start_time=?,
    end_time=?,
    command_id=?,
    name=?,
    start_date=?,
    end_date=?,
    enable=?,
    `repeat`=?
WHERE id = ?;

-- name: CreateRepeatDays :exec
INSERT INTO repeat_day(schedule_id, day)
VALUES (?, ?);

-- name: CreateRepeatMonth :exec
INSERT INTO repeat_month(schedule_id, month)
VALUES (?, ?);

-- name: CreateRepeatWeekdays :exec
INSERT INTO repeat_weekday(schedule_id, weekday)
VALUES (?, ?);

-- name: CreateCommand :execresult
INSERT INTO commands(command)
VALUES (?);

-- name: DeleteCommand :exec
DELETE
FROM commands
WHERE id = ?;

-- name: UpdateCommand :exec
UPDATE commands
set command=?
WHERE id = ?;

-- name: DeleteSchedule :exec
DELETE
FROM schedules
WHERE id = ?;

-- name: DeleteRepeatWeekday :exec
DELETE
FROM repeat_weekday
WHERE schedule_id = ?;

-- name: DeleteRepeatMonth :exec
DELETE
FROM repeat_month
WHERE schedule_id = ?;

-- name: DeleteRepeatDay :exec
DELETE
FROM repeat_day
WHERE schedule_id = ?;



