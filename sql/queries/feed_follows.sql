-- name: CreateFeedFollow :one
INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING
    id,
    created_at,
    updated_at,
    feed_id,
    user_id,
    (SELECT feeds.name FROM feeds WHERE feeds.id = feed_follows.feed_id) AS feed_name,
    (SELECT users.name FROM users WHERE users.id = feed_follows.user_id) AS user_name;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.id,
    feed_follows.created_at,
    feed_follows.updated_at,
    feed_follows.feed_id,
    feed_follows.user_id,
    (SELECT feeds.name FROM feeds WHERE feeds.id = feed_follows.feed_id) AS feed_name,
    (SELECT users.name FROM users WHERE users.id = feed_follows.user_id) AS user_name
FROM feed_follows
WHERE feed_follows.user_id = $1;


-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
WHERE feed_id = $1 AND user_id = $2;

