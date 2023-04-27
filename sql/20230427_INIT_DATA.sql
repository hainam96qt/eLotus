create table users (
    id int AUTO_INCREMENT primary key,
    user_name VARCHAR(255),
    password VARCHAR(255)
);

create unique index user_name_idx on users(user_name);