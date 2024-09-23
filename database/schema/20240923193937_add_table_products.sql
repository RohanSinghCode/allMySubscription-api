-- +goose Up
-- +goose StatementBegin
CREATE TABLE Products(
        Id UUID PRIMARY KEY,
        Name TEXT NOT NULL,
        Logo TEXT NULL,
        ApiKey TEXT NULL,
        CreatedAt TimeStamp NOT NULL DEFAULT CURRENT_DATE,
        UpdatedAt TimeStamp NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Products
-- +goose StatementEnd
