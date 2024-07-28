-- name: InsertUser :one
INSERT INTO Users(
    Id,
    FirstName,
    LastName,
    Email,
    Username,
    Password
)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5
)

RETURNING *;