create table user_queue_number
(
    id       serial primary key,
    username varchar(100) references "user" not null unique,
    queue_id int references queue           not null
);