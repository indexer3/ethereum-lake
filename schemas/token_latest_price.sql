-- Postgres Creation 

create table if not exists token_latest_price 
(
    token_address bytea not null primary key,
    network text not null, 
    symbol test,
    decimal numeric not null,
    name text not null,
    price numeric not null, 
    updated_at timestamp with time zone
);


alter table token_latest_price
    owner to user;

create index if not exists idx_token_address
    on token_latest_price (token_address);