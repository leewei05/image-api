CREATE DATABASE db with ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8' CONNECTION LIMIT = -1 template=template0;
ALTER DATABASE db SET timezone TO 'UTC';

\connect db

CREATE TABLE materials
(
	id SERIAL PRIMARY KEY,
	name character varying(1000) not null,
	url  character varying(1000) not null,
	price integer not null,

	create_time timestamp with time zone default current_timestamp,
	update_time timestamp with time zone default current_timestamp

);

INSERT INTO materials(name,url,price)
values
('test1','gs://image-api/gg.png',100);
