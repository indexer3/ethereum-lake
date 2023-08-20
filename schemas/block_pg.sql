create table if not exists blocks
(
    number           bigint not null primary key,
    "hash"             bytea,
    coinbase         bytea,
    parent_hash      bytea,
    receipt_hash     bytea,
    uncle_hash       bytea,
    mix_digest       bytea,
    "root"             bytea,
    bloom            bytea,
    nonce            bytea,
    extra            bytea,
    base_fee         numeric,
    gas_limit        bigint,
    gas_used         bigint,
    size             bigint,
    difficulty       numeric,
    total_difficulty numeric,
    timestamp        timestamp with time zone
);

alter table blocks
    owner to user;

create index if not exists idx_blocks_hash
    on blocks (hash);