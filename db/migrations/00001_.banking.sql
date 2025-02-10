-- +goose Up
CREATE TABLE "customers" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    city VARCHAR(15) NOT NULL,
    zip_code VARCHAR(10) NOT NULL,
    status BIGINT NOT NULL DEFAULT 1
);
INSERT INTO customers VALUES
(100, 'Ruth', '1983-09-11', 'Meru', '60600', 1),
(101, 'James', '1985-12-01', 'Embu', '59786', 1),
(102, 'Dominic', '1996-08-11', 'Nairobi', '63452', 1),
(103, 'David', '1997-05-17', 'Nakuru', '26743', 1),
(104, 'Doris', '1998-11-28', 'Kisumu', '147832', 1);

CREATE TABLE "accounts" (
    id SERIAL PRIMARY KEY,
    customer_id BIGINT REFERENCES customers(id),
    pin VARCHAR NOT NULL,
    account_type VARCHAR(10) NOT NULL,
    amount FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT clock_timestamp(),
    status BIGINT NOT NULL DEFAULT '1'
);

INSERT INTO accounts VALUES
(1, 100, '1025', 'Saving', 500000, '2025-01-22 10:25:09', 1 ),
(2, 101, '3078', 'Deposits', 34578, '2025-01-05 03:09:09', 1 ),
(3, 102, '5608', 'Saving', 320000, '2025-01-10 09:34:12', 0 ),
(4, 103, '3905', 'Business', 134500, '2024-12-13 01:45:50', 1 ),
(5, 104, '2189', 'Saving', 310500, '2025-01-22 02:18:27', 1 );

CREATE TABLE "transactions" (
    id SERIAL PRIMARY KEY,
    customer_id BIGINT REFERENCES customers(id),
    account_id BIGINT NOT NULL REFERENCES accounts(id),
    amount FLOAT NOT NULL,
    transaction_type VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL  DEFAULT clock_timestamp());

INSERT INTO transactions VALUES
(2000, 1, 40000, 'deposit', '2025-01-22 10:56:09'),
(2001, 2, 450000, 'deposit', '2025-01-06 3:00:34'),
(2002, 3, 3400, 'withdrawal', '2025-01-11 04:02:34'),
(2003, 4, 234500, 'deposit', '2025-01-22 08:57:23'),
(2004, 5, 21000, 'deposit', '2025-01-20 03:21:04');

-- +goose Down
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transactions
