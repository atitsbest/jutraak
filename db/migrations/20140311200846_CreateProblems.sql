
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE problems (
    id          serial NOT NULL,
    summary     varchar(100) NOT NULL,
    description text NOT NULL,
    tags        varchar(255),

    created_at  date NOT NULL,
    created_by  varchar(30) NOT NULL,
    lastchange_at date NOT NULL,
    lastchange_by varchar(30) NOT NULL,

    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
-- DROP TABLE problems;

