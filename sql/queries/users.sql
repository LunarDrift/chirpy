-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
     gen_random_uuid(),
     now(),
     now(),
     $1,
     $2
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByRefreshToken :one
SELECT * FROM users
JOIN refresh_tokens ON users.id = refresh_tokens.user_id
WHERE refresh_tokens.token = $1
AND refresh_tokens.expires_at > NOW()
AND refresh_tokens.revoked_at IS NULL;

-- name: UpdateUser :one
UPDATE users
SET email = $1, hashed_password = $2, updated_at = NOW()
WHERE id = $3
RETURNING *;

-- name: UpgradeUserToChirpyRed :execresult
UPDATE users
SET is_chirpy_red = true, updated_at = NOW()
WHERE id = $1;
