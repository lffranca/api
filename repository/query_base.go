package repository

const (
	queryBase = `
	create table if not exists accounts (id text not null primary key, status text);
create table if not exists transactions (
  id text not null primary key,
  id_account text,
  description text,
  value numeric
);
insert into accounts (id, status)
values ('3b1b1c0c-5c68-4352-b937-e3c68b6b1b16', 'ATIVA'),
  (
    '53sb1c0c-5c68-4352-b937-e3c62316b1cd',
    'INATIVA'
  ) on conflict(id) do
update
set status = excluded.status;
insert into transactions (id, id_account, description, value)
values (
    '3b1b1c0c-5c68-4352-b937-e3c68b6b1123',
    '3b1b1c0c-5c68-4352-b937-e3c68b6b1b16',
    'Apple Store',
    199.5
  ),
  (
    '3qw31c0c-5c68-4352-b937-e3c68b6b123s',
    '3b1b1c0c-5c68-4352-b937-e3c68b6b1b16',
    'Amazon Prime',
    9.9
  ),
  (
    'c203b91a-91a4-41d2-8583-86401c0fb1e4',
    '53sb1c0c-5c68-4352-b937-e3c62316b1cd',
    'Netflix',
    27.5
  ) on conflict(id) do
update
set description = excluded.description;
	`
)
