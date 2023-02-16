-- name: GetResourceById :one
SELECT *
FROM resource
WHERE resource_id = pggen.arg('resource_id');

-- name: GetResources :many
SELECT *
FROM resource;

-- name: DeleteResourceById :exec
DELETE
FROM resource
WHERE resource_id = pggen.arg('resource_id');

-- name: InsertResource :exec
INSERT INTO resource(title,
                     description,
                     link)
VALUES(
       pggen.arg('title'),
       pggen.arg('description'),
       pggen.arg('link')
      );

-- name: UpdateResource :exec
UPDATE resource
SET title = pggen.arg('title'),
    description = pggen.arg('description'),
    link = pggen.arg('link'),
    updated_at = NOW()
WHERE resource_id = pggen.arg('resource_id');