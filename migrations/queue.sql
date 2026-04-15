create table queue
(
    id                        serial primary key,
    name                      varchar(100)        not null,
    next_free_slot_number     int default 1,
    responsible_user_username varchar(100) unique not null
        constraint queue_user_fk references "user"
);