-- name: CreateProfile :one
INSERT INTO profiles (
    id,
    username,
    firstname,
    lastname,
    avatar
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: UpdateProfile :one
UPDATE profiles SET
    username = $1,
    firstname = $2,
    lastname = $3,
    avatar = $4
WHERE id = $5
RETURNING *;

-- name: DeleteProfile :exec
DELETE FROM profiles
WHERE id = $1;

-- name: GetProfile :one
SELECT * FROM profiles
WHERE id = $1 LIMIT 1;

-- name: GetProfileByUsername :one
SELECT * FROM profiles
WHERE username = $1 LIMIT 1;

-- name: GetManyProfiles :many
SELECT * FROM profiles
WHERE id IN(sqlc.arg('ids')::text[])
LIMIT sqlc.arg('limit');

-- name: UsernameTaken :one
select exists(select 1 from profiles where username=$1) AS "exists";
