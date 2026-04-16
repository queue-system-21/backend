create table "user"
(
    username  varchar(100) primary key,
    password  varchar(100) not null,
    role_code role_code    not null default 'user'
        constraint user_role_fk references role
);