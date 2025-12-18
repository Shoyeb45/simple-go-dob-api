-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name, dob
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
  set name = $2,
  dob = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: UpdateUserPartial :one
UPDATE users
SET
  name = COALESCE($2, name),
  dob  = COALESCE($3, dob)
WHERE id = $1
RETURNING id, name, dob;



-- name: ListUsersPaginated :many
SELECT * 
FROM users
ORDER BY id
LIMIT $1 OFFSET $2;