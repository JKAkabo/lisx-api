-- migrate:up
create table "user"
(
    id                    serial primary key,
    first_name            varchar(25) not null,
    last_name             varchar(25) not null,
    type                  varchar(10) not null,
    username              varchar(15) unique,
    password              varchar(60),
    force_password_change bool default false
);

-- migrate:down
drop table "user";