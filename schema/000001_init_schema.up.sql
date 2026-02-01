CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    Username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);
CREATE TABLE targets (
    id serial not null unique,
    url varchar(255) not null,
    status varchar(255) default 'pending',
    title VARCHAR(255) NOT NULL,
    user_id int references users (id) on delete cascade not null,
    last_check timestamp
);