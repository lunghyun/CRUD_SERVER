-- name: CreateUser :exec
INSERT INTO users (name, age) VALUES (?, ?);

-- name: GetAllUsers :many
SELECT id, name, age FROM users;

-- name: UpdateUserAge :execresult
UPDATE users SET age = ? WHERE name = ?;

-- name: DeleteUserByName :execresult
DELETE FROM users WHERE name = ?;
