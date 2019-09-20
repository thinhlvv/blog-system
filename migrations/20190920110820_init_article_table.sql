-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS article(
	id INT UNSIGNED AUTO_INCREMENT,
	title VARCHAR(255) NOT NULL,
  content TEXT NOT NULL,
  author VARCHAR(255) NOT NULL,

  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at DATETIME DEFAULT NULL,
	PRIMARY KEY (id)
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS article;
-- +goose StatementEnd
