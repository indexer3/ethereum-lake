-- Postgres Creation 
create table if not exists transactions
(
    "hash"         bytea not null primary key,
    tx_index       bigint,
    "from"       bytea,
    "to"         bytea,
    method_id        bytea,
    "value"        numeric,
    "type"         bigint,
    gas          bigint,
    gas_price    numeric,
    block_hash   bytea,
    block_number bigint,
    timestamp    timestamp with time zone
);

alter table transactions
    owner to user;

create index if not exists idx_transactions_from_block_number
    on transactions ("to", block_number);

create index if not exists idx_transactions_from_block_number
    on transactions ("from", block_number);

create index if not exists idx_transactions_timestamp
    on transactions (timestamp);

create index if not exists idx_transactions_order
    on transactions (block_number);




