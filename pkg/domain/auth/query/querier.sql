-- name: CreateSession :exec
INSERT INTO
    session(
            session_id,
            account_id,
            ip_address,
            user_agent,
            expire_at
            )
VALUES (
        pggen.arg('session_id'),
        pggen.arg('account_id'),
        pggen.arg('ip_address'),
        pggen.arg('user_agent'),
        pggen.arg('expire_at')
       );

-- name: GetSession :one
SELECT *
FROM session
WHERE session_id = pggen.arg('session_id');

-- name: UpdateSessionId :exec
UPDATE session
SET session_id = pggen.arg('new_session_id'),
    expire_at = pggen.arg('expire_at')
WHERE session_id = pggen.arg('old_session_id');

-- name: DeleteSession :exec
DELETE
FROM session
WHERE session_id = pggen.arg('session_id');
