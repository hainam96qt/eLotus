/* name: CreateUser :exec */
INSERT INTO users (
    user_name,
    password
) VALUES (?,?);

/* name: GetUser :one */
SELECT * FROM users where user_name = ? LIMIT 1;