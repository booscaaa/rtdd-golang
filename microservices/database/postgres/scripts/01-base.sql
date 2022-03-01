SET TIMEZONE TO 'America/Sao_Paulo';

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE EXTENSION IF NOT EXISTS unaccent;

CREATE TABLE action (
	id serial primary key not null,
	name varchar(60) not null
);

CREATE TABLE permission (
	id serial primary key not null,
	name varchar(60) not null
);

CREATE TABLE account_group (
	id serial primary key not null,
	name varchar(60) not null
);

CREATE TABLE account_group_permission (
	id serial primary key not null,
	account_group_id integer not null,
	permission_id integer not null,
	action_id integer not null,

	CONSTRAINT account_group_id_fk
	FOREIGN KEY (account_group_id) REFERENCES account_group (id),

	CONSTRAINT permission_id_fk
	FOREIGN KEY (permission_id) REFERENCES permission (id),

	CONSTRAINT action_id_fk
	FOREIGN KEY (action_id) REFERENCES action (id)
);

CREATE TABLE account (
	id serial primary key not null,
	login varchar(60) unique not null,
	email varchar(100) unique,
	name varchar(150) not null,
	password varchar(300) not null,
	created_date timestamp default current_timestamp,
	account_group_id integer not null,
	active bool not null default true,
	

	CONSTRAINT account_group_id_fk
	FOREIGN KEY (account_group_id) REFERENCES account_group (id)
);

CREATE TABLE auth (
	id serial primary key not null,
	type varchar(60) not null,
	hash varchar(600) not null,
	token varchar(600) not null,
	account_id integer not null,
	revoked boolean default false,
	created_date timestamp default current_timestamp,

	CONSTRAINT account_id_fk
	FOREIGN KEY (account_id) REFERENCES account (id)
);

INSERT INTO account_group(name) VALUES ('ADMINISTRATOR');

INSERT INTO account(
	login, 
	email, 
	name,
	password, 
	account_group_id
) VALUES (
	'admin', 
	'admin@admin.com.br', 
	'ADMINISTRATOR',
	crypt('admin', gen_salt('bf', 10)),
	1
);