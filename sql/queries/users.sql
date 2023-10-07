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

-- name: FindUserByEmail :many
SELECT id, is_email_verified, username 
FROM users 
WHERE email = $1;

-- name: FindUserByUsername :many
SELECT id, is_email_verified, username 
FROM users 
WHERE username = $1;