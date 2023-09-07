create table if not exists erc20_tokens
(
    network text not null,
    contract_address bytea not null,
    symbol  text,
    decimal numeric not null,
    "name"  text not null,
    price   numeric not null,
    updated_at  timestamp with time zone,
    primary key (network, contract_address)
);

alter table erc20_tokens
    owner to user;
