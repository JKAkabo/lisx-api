-- migrate:up
create table user_permission (
    user_id int references "user",
    resource varchar(100) not null ,
    can_create bool not null ,
    can_read bool not null ,
    can_update bool not null ,
    can_delete bool not null,
    primary key (user_id, resource)
);

-- migrate:down
drop table user_permission;
