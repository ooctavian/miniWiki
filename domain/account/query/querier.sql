-- name: CreateAccount :exec
INSERT INTO
    account(email, password)
VALUES(
       pggen.arg('email'),
       pggen.arg('password')
       );

--name: UpdateAccount :exec
UPDATE account
SET email = pggen.arg('email'),
    password = pggen.arg('password'),
    updated_at = NOW()
FROM session
WHERE
    account.account_id = pggen.arg('account_id');

-- name: GetAccount :one
SELECT account_id, email::TEXT, password
FROM account
WHERE email = pggen.arg('email')::domain_email;

-- name: GetAccountById :one
SELECT email::TEXT, password, updated_at, created_at
FROM account
WHERE account_id = pggen.arg('account_id');