/* name: CreateUser :exec */
INSERT INTO users (
    user_name,
    password
) VALUES (?,?);

/* name: GetUser :one */
SELECT * FROM users where user_name = ? LIMIT 1;

/* name: CreateImage :exec */
INSERT INTO images (
    user_id,
    path,
    metadata
) VALUES (?,?,?);

/* name: GetImage :one */
SELECT * FROM images where id = ? LIMIT 1;
