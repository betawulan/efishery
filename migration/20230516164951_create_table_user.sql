-- +goose Up
-- +goose StatementBegin
create table user (
    id int auto_increment not null,
    phone varchar(13) unique not null,
    name varchar(100) not null,
    role varchar(5) not null,
    password varchar(4) not null,
    created_at timestamp(3) default current_timestamp(3),
    primary key (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
