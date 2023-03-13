-- name: GetCategoryById :one
SELECT *
FROM category
WHERE category_id = pggen.arg('category_id');

-- name: GetCategories :many
SELECT *
FROM category
WHERE author_id = pggen.arg('author_id');

-- name: DeleteCategoryById :exec
DELETE
FROM category
WHERE category_id = pggen.arg('category_id')
AND author_id = pggen.arg('author_id');

-- name: InsertCategory :exec
INSERT INTO category(title, parent_id, author_id)
VALUES(pggen.arg('title'),
       NULLIF(pggen.arg('parent_id'),0),
       pggen.arg('author_id'));

-- name: UpdateCategory :exec
UPDATE category
SET title = COALESCE(NULLIF(pggen.arg('title'),''),title),
    parent_id = COALESCE(NULLIF(pggen.arg('parent_id'),0),parent_id),
    updated_at = NOW()
WHERE category_id = pggen.arg('category_id')
        AND author_id = pggen.arg('author_id');