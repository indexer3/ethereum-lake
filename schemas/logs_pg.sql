create table if not exists logs
(
    transaction_hash  bytea not null,
    "index" bigint not null,
    "address"         bytea,
    topic_zero        bytea, 
    topic_first       bytea,
    topic_second      bytea,
    topic_third       bytea,
    "data"              bytea,
    removed           boolean,
    block_hash        bytea,
    block_number      numeric,
    transaction_index bigint,
    timestamp         timestamp with time zone,
    primary key (transaction_hash, "index")
);

comment on table logs is 'Returns a box2d encapsulating the given Geometry.';

alter table logs
    owner to user;

create index if not exists idx_logs_hash
    on logs (transaction_hash);

create index if not exists idx_logs_timestamp
    on logs (timestamp);

create index if not exists idx_logs_topic_zero
    on logs (topic_zero);

create index if not exists idx_logs_topic_first
    on logs (topic_first);

create index if not exists idx_logs_topic_second
    on logs (topic_second);

create index if not exists idx_logs_topic_third
    on logs (topic_third);

create index if not exists idx_logs_order
    on logs (block_number, transaction_index, "index");

create index if not exists idx_logs_address
    on logs ("address");
