-- +goose Up
-- +goose StatementBegin
CREATE TABLE Users(
    Id UUID PRIMARY KEY,
    FirstName TEXT NOT NULL,
    LastName TEXT NULL,
    Email TEXT NOT NULL UNIQUE,
    Username TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL,
    CreatedAt TimeStamp NOT NULL DEFAULT CURRENT_DATE,
    UpdatedAt TimeStamp NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users;
-- +goose StatementEnd
