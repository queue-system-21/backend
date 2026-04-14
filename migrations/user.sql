create table "user"
(
    id       serial primary key,
    username varchar(100) not null unique,
    password varchar(100) not null
);