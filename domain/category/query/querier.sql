-- name: GetCategoryById :one
SELECT *
FROM category
WHERE category_id = pggen.arg('category_id');

-- name: GetCategories :many
SELECT *
FROM category;

-- name: DeleteCategoryById :exec
DELETE
FROM category
WHERE category_id = pggen.arg('category_id');

-- name: InsertCategory :exec
INSERT INTO category(title)
VALUES(
       pggen.arg('title')
       );

-- name: InsertSubCategory :exec
INSERT INTO category(title, parent_id)
VALUES(
       pggen.arg('title'),
       pggen.arg('parent_id')
       );


-- name: UpdateCategory :exec
UPDATE category
SET title = pggen.arg('title'),
    updated_at = NOW()
WHERE category_id = pggen.arg('category_id');

-- name: UpdateSubCategory :exec
UPDATE category
SET title = pggen.arg('title'),
    parent_id = pggen.arg('parent_id'),
    updated_at = NOW()
WHERE category_id = pggen.arg('category_id');