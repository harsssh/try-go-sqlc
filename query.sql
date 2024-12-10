# name: GetUserPostWithComments :many
SELECT users.id    AS 'user_id',
       users.username,
       posts.id    AS 'post_id',
       posts.title,
       comments.id AS 'comment_id',
       comments.body
FROM users
         JOIN posts ON users.id = posts.user_id
         LEFT JOIN comments ON posts.id = comments.post_id
WHERE users.id = sqlc.arg(user_id) AND posts.id = sqlc.arg(post_id);

# name: GetUsersPostsCount :many
SELECT users.id,
       users.username,
       COUNT(posts.id) AS posts_count
FROM users
         LEFT JOIN posts ON users.id = posts.user_id
GROUP BY users.id;

# name: GetUsersByIDs :many
SELECT id,
       username
FROM users
WHERE id IN (sqlc.slice('ids'));