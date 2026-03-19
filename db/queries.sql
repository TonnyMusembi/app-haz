-- name: CreateCustomer :execresult
INSERT INTO customers (full_name, phone, national_id)
VALUES (?, ?, ?);

-- name: GetCustomerByID :one
SELECT * FROM customers WHERE id = ? LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customers ORDER BY id DESC LIMIT ? OFFSET ?;

-- name: CreateLoanProduct :execresult
INSERT INTO loan_products (name, interest_rate, duration_days, penalty_rate)
VALUES (?, ?, ?, ?);

-- name: ListLoanProducts :many
SELECT * FROM loan_products;

-- name: CreateLoan :execresult
INSERT INTO loans (
 customer_id, loan_product_id, principal, interest, total_amount, status
) VALUES (?, ?, ?, ?, ?, 'pending');

-- name: GetLoanByID :one
SELECT * FROM loans WHERE id = ? LIMIT 1;

-- name: ListLoans :many
SELECT * FROM loans ORDER BY id DESC LIMIT ? OFFSET ?;

-- name: UpdateLoanStatus :exec
UPDATE loans SET status = ? WHERE id = ?;

-- name: UpdateLoanPaidAmount :exec
UPDATE loans SET paid_amount = paid_amount + ? WHERE id = ?;


-- name: CreateInstallment :exec
INSERT INTO loan_installments (loan_id, due_date, amount)
VALUES (?, ?, ?);

-- name: GetInstallmentsByLoan :many
SELECT * FROM loan_installments WHERE loan_id = ?;

-- name: CreatePayment :execresult
INSERT INTO loan_payments (loan_id, amount, method, reference)
VALUES (?, ?, ?, ?);

-- name: CreateUser :execresult
INSERT INTO users (full_name, email, phone, password_hash, role)
VALUES (?, ?, ?, ?, 'user');

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = ? LIMIT 1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = ? LIMIT 1;

-- name: UpdateLastLogin :exec
UPDATE users
SET last_login = NOW()
WHERE id = ?;

-- name: UpdatePassword :exec
UPDATE users
SET password_hash = ?, updated_at = NOW()
WHERE id = ?;

-- name: SetResetToken :exec
UPDATE users
SET reset_token = ?,
    reset_token_expiry = ?,
    updated_at = NOW()
WHERE email = ?;

-- name: GetUserByResetToken :one
SELECT * FROM users
WHERE reset_token = ?
  AND reset_token_expiry > NOW()
LIMIT 1;

-- name: ClearResetToken :exec
UPDATE users
SET reset_token = NULL,
    reset_token_expiry = NULL,
    updated_at = NOW()
WHERE id = ?;

-- name: EmailExists :one
SELECT COUNT(*) FROM users WHERE email = ?;

-- name: PhoneExists :one
SELECT COUNT(*) FROM users WHERE phone = ?;