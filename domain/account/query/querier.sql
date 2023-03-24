-- name: CreateAccount :exec
INSERT INTO
    account(email, password, active)
VALUES(
       pggen.arg('email'),
       pggen.arg('password'),
       true
       );

--name: UpdateAccount :exec
UPDATE account
SET email = pggen.arg('email'),
    password = pggen.arg('password'),
    updated_at = NOW()
WHERE
    account.account_id = pggen.arg('account_id');

-- name: GetAccount :one
SELECT account_id, email::TEXT, password, active
FROM account
WHERE email = pggen.arg('email')::domain_email;

-- name: GetAccountById :one
SELECT email::TEXT, password, updated_at, created_at
FROM account
WHERE account_id = pggen.arg('account_id');

-- name: GetAccountStatus :one
SELECT active
FROM account
WHERE account_id = pggen.arg('account_id');

-- name: UpdateAccountStatus :exec
UPDATE account
SET active = pggen.arg('status')
WHERE account_id = pggen.arg('account_id');
