create table queue
(
    id                        serial primary key,
    name_rus                  varchar(100) unique not null,
    name_kaz                  varchar(100) unique not null,
    next_free_slot_number     int default 1       not null,
    responsible_user_username varchar(100) unique not null
        constraint queue_user_fk references "user"
);