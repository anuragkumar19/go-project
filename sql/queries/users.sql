-- name: CreateUser :many
INSERT INTO users (
  name, username, email, password, otp, otp_expiry
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;

-- name: UpdateUser :many
UPDATE users 
SET name = $2, username = $3, email = $4, password = $5, otp = $6, otp_expiry = $7
WHERE id = $1
RETURNING id
;

-- name: ForgotPassword :many
UPDATE users 
SET otp = $2, otp_expiry = $3
WHERE id = $1
RETURNING id
;

-- name: ResetPassword :many
UPDATE users 
SET password = $2, otp = $3, otp_expiry = $4
WHERE id = $1
RETURNING id
;

-- name: UpdateName :exec
UPDATE users 
SET name = $2
WHERE id = $1;

-- name: UpdateUsername :exec
UPDATE users 
SET username = $2
WHERE id = $1;

-- name: UpdatePassword :exec
UPDATE users 
SET password = $2
WHERE id = $1;

-- name: UpdateAvatar :exec
UPDATE users 
SET avatar = $2
WHERE id = $1;

-- name: VerifyUser :many
UPDATE users 
SET is_email_verified = $2, otp = $3, otp_expiry = $4
WHERE id = $1
RETURNING id
;

-- name: FindUserByEmail :many
SELECT id, is_email_verified, username,otp , otp_expiry 
FROM users 
WHERE email = $1;

-- name: FindUserByUsername :many
SELECT id, is_email_verified, username 
FROM users 
WHERE username = $1;

-- name: LoginQuery :many
SELECT id, password, email, is_email_verified, username
FROM users
WHERE email = $1
OR username = $1;

-- name: GetUserById :many
SELECT id, name, username, email, is_email_verified, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByUsernamePublic :many
SELECT id, name, username, avatar, created_at FROM users
WHERE username = $1 AND is_email_verified = true;

-- name: GetUserByIDPublic :many
SELECT id, name, username, avatar, created_at FROM users
WHERE id = $1 AND is_email_verified = true;

-- name: SearchUserPublic :many
SELECT id, name, username, avatar FROM users 
WHERE is_email_verified = true AND (LOWER(username) LIKE $1 OR LOWER(name) LIKE $1)
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;