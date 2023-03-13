-- name: GetProfile :one
SELECT *
FROM profile
WHERE account_id = pggen.arg('account_id');

-- name: CreateProfile :exec
INSERT INTO profile(account_id, name)
VALUES (
        pggen.arg('account_id'),
        pggen.arg('name')
        );

-- name: UpdateAlias :exec
UPDATE profile
SET alias = pggen.arg('alias'),
    updated_at = NOW()
WHERE account_id = pggen.arg('account_id');

-- name: UpdateName :exec
UPDATE profile
SET name = pggen.arg('name'),
    updated_at = NOW()
WHERE account_id = pggen.arg('account_id');

-- name: UpdateProfilePicture :exec
UPDATE profile
SET picture_url = pggen.arg('picture_url'),
    updated_at = NOW()
WHERE account_id = pggen.arg('account_id');

