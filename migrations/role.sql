create type role_code as enum ('user', 'receptionist', 'admin');

create table role
(
    code     role_code primary key,
    name_rus varchar(50) not null,
    name_kaz varchar(50) not null
);

insert into role (code, name_rus, name_kaz)
values ('user', 'Пользователь', 'Қолданушы'),
       ('receptionist', 'Принимающий', 'Қабылдаушы'),
       ('admin', 'Админ', 'Админ');