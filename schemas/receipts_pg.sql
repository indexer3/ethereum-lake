create table if not exists logs
(
    transaction_hash  bytea not null,
    "index" bigint not null,
    "address"           bytea,
    topic_first       bytea,
    topic_second      bytea,
    topic_third       bytea,
    topic_fourth      bytea,
    "data"              bytea,
    removed           boolean,
    block_hash        bytea,
    block_number      numeric,
    transaction_index bigint,
    timestamp         timestamp with time zone,
    primary key (transaction_hash, index)
);


alter table logs
    owner to user;

create index if not exists idx_logs_hash
    on logs (transaction_hash);

create index if not exists idx_logs_timestamp
    on logs (timestamp);

create index if not exists idx_logs_topic_third
    on logs (topic_third);

create index if not exists idx_logs_topic_fourth
    on logs (topic_fourth);

create index if not exists idx_logs_topic_first
    on logs (topic_first);

create index if not exists idx_logs_topic_second
    on logs (topic_second);

create index if not exists idx_logs_order
    on logs (block_number, transaction_index, index);

create index if not exists idx_logs_address
    on logs (address);

