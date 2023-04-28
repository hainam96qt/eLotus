create table users (
    id int AUTO_INCREMENT primary key,
    user_name VARCHAR(255),
    password VARCHAR(255)
);

create unique index user_name_idx on users(user_name);

create table images (
    id int AUTO_INCREMENT primary key,
    user_id int,
    path VARCHAR(255),
    metadata JSON
);

ALTER TABLE images
    ADD FOREIGN KEY (user_id) REFERENCES users(id);