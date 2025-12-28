-- name: GetCounter :one
SELECT * FROM counters WHERE user_id = ? LIMIT 1;

-- name: IncrementCounter :one
UPDATE counters SET count = count + 1 WHERE user_id = ? RETURNING count;

-- name: CreateCounter :exec
INSERT INTO counters (user_id, count) VALUES (?, 1);

-- Self-Hosted users

-- name: GetUser :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users WHERE name = ? LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (id, name, display_name) VALUES (?, ?, ?) RETURNING *;

-- name: CreateCredential :exec
INSERT INTO credentials (id, public_key, attestation_type, transport, flags, authenticator, user_id) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: UpdateCredential :exec
UPDATE credentials
SET public_key = ?,
attestation_type = ?,
transport = ?,
flags = ?,
authenticator = ?
WHERE id = sqlc.arg(by_id);

-- name: DeleteCredential :exec
DELETE FROM credentials WHERE id = ? AND user_id = ?;

-- name: GetCredentialsByUser :many
SELECT * FROM credentials WHERE user_id = ?;

-- Datapoints

-- name: GetDatapoint :one
SELECT * FROM datapoints WHERE id = ? LIMIT 1;

-- name: ListDatapointsByUser :many
SELECT * FROM datapoints WHERE user_id = ? ORDER BY created_at DESC;

-- name: CreateDatapoint :one
INSERT INTO datapoints (user_id, name) VALUES (?, ?) RETURNING *;

-- name: UpdateDatapointName :one
UPDATE datapoints SET name = ? WHERE id = ? RETURNING *;

-- name: DeleteDatapoint :exec
DELETE FROM datapoints WHERE id = ?;

-- Dataentries

-- name: GetDataentry :one
SELECT * FROM dataentries WHERE id = ? LIMIT 1;

-- name: ListDataentriesByDatapoint :many
SELECT * FROM dataentries WHERE datapoint_id = ? ORDER BY created_at DESC;

-- name: CreateDataentry :one
INSERT INTO dataentries (datapoint_id, type, text_value, int_value) VALUES (?, ?, ?, ?) RETURNING *;

-- name: UpdateDataentry :one
UPDATE dataentries SET type = ?, text_value = ?, int_value = ? WHERE id = ? RETURNING *;

-- name: DeleteDataentry :exec
DELETE FROM dataentries WHERE id = ?;
