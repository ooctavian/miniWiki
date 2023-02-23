CREATE OR REPLACE FUNCTION GetResourcesFilter
(
    fTitle TEXT DEFAULT NULL,
    fLink TEXT DEFAULT NULL,
    categories BIGINT[] DEFAULT NULL
) RETURNS SETOF resource
    LANGUAGE plpgsql AS $$
DECLARE
    fTitle TEXT := TRIM(fTitle);
    fLink TEXT := TRIM(fLink);
BEGIN
    RETURN QUERY
    SELECT *
    FROM resource
    WHERE (fTitle IS NULL OR title LIKE '%' || fTitle || '%')
      AND (fLink IS NULL OR link LIKE '%' || fLink || '%')
      AND (categories IS NULL OR resource.category_id = ANY(categories));
END;
$$;