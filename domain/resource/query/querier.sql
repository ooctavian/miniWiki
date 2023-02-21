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
                     link,
                     category_id)
VALUES(
       pggen.arg('title'),
       pggen.arg('description'),
       pggen.arg('link'),
       pggen.arg('category_id')
      );

-- name: UpdateResource :exec
UPDATE resource
SET title = pggen.arg('title'),
    description = pggen.arg('description'),
    link = pggen.arg('link'),
    category_id = pggen.arg('category_id'),
    updated_at = NOW()
WHERE resource_id = pggen.arg('resource_id');