create table "user"
(
    id       serial primary key,
    username varchar(100) not null,
    password varchar(100) not null
);