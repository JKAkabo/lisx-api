-- migrate:up
create table api_key (
    id serial primary key ,
    name varchar(25) not null unique ,
    key varchar(32) not null unique ,
    user_id int references "user"
);

-- migrate:down
drop table api_key;
