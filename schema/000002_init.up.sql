CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE IF NOT EXISTS order_main_info
(
--     id                 serial       not null unique,
    order_uid          varchar(19)  not null
        constraint orders_uid_pkey
            primary key,
    track_number       varchar(255) not null,
    entry              varchar(255),
    locale             varchar(5)   not null,
    internal_signature varchar(255),
    customer_id        varchar(255) not null,
    delivery_service   varchar(255) not null,
    shardkey           varchar(255) not null,
    sm_id              integer      not null,
    date_created       varchar(20)  not null,
    oof_shard          varchar(255) not null
);


CREATE TABLE delivery
(
--     id       serial                                                not null unique,
    delivery_id varchar(19) references order_main_info (order_uid) on delete cascade not null,
    name        varchar(255)                                                         not null,
    phone       varchar(15),
    zip         varchar(30)                                                          not null,
    city        varchar(255)                                                         not null,
    address     varchar(255)                                                         not null,
    region      varchar(255),
    email       varchar(255)
);

CREATE TABLE payment
(
--     id            serial                                                               not null unique,
    payment_id    varchar(19) references order_main_info (order_uid) on delete cascade not null,
    transaction   varchar(255),
    request_id    varchar(255),
    currency      varchar(10),
    provider      varchar(255),
    amount        integer,
    payment_dt    integer,
    bank          varchar(255),
    delivery_cost integer,
    goods_total   integer,
    custom_fee    integer
);

CREATE TABLE items
(
--     id           serial                                                               not null unique,
    item_id      varchar(19) references order_main_info (order_uid) on delete cascade not null,
    chrt_id      integer,
    track_number varchar(255),
    price        integer,
    rid          varchar(255),
    name         varchar(255),
    sale         integer,
    size         varchar(255),
    total_price  integer,
    nm_id        integer,
    brand        varchar(255),
    status       integer

);
