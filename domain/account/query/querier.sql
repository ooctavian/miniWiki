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
    password = pggen.arg('password')
FROM session
WHERE
    account.account_id = pggen.arg('account_id')
    AND session.account_id = account.account_id
    AND session.session_id = pggen.arg('session_id');

-- name: GetAccount :one
SELECT account_id, email::TEXT, password
FROM account
WHERE email = pggen.arg('email')::domain_email;