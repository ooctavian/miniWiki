-- name: GetResourceById :one
SELECT *
FROM resource
WHERE resource_id = pggen.arg('resource_id');

-- name: GetResources :many
SELECT *
FROM getresourcesfilter(ftitle := pggen.arg('title'),
    flink := pggen.arg('link'),
    categories := pggen.arg('categories'))
WHERE state = 'PUBLIC' OR
      author_id = pggen.arg('account_id');

-- name: DeleteResourceById :exec
DELETE
FROM resource
WHERE resource_id = pggen.arg('resource_id')
AND author_id = pggen.arg('account_id');

-- name: InsertResource :exec
INSERT INTO resource(title,
                     description,
                     link,
                     category_id,
                     author_id,
                     state)
VALUES(
       pggen.arg('title'),
       pggen.arg('description'),
       pggen.arg('link'),
       pggen.arg('category_id'),
       pggen.arg('author_id'),
       pggen.arg('state')
      );

-- name: UpdateResource :exec
UPDATE resource
SET title = pggen.arg('title'),
    description = pggen.arg('description'),
    link = pggen.arg('link'),
    category_id = pggen.arg('category_id'),
    state = pggen.arg('state'),
    updated_at = NOW()
WHERE resource_id = pggen.arg('resource_id');

--name: UpdateResourceImage :exec
UPDATE resource
SET image = pggen.arg('image_url')
WHERE resource_id = pggen.arg('resource_id')
AND author_id = pggen.arg('author_id');