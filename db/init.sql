create database db;
\c db;
create table urls (
    id serial primary key,
    orig_url text,
    short_id text
);
