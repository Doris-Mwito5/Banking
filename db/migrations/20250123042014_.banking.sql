-- +goose Up
CREATE TABLE "customers" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    date_of_birth DATE NOT NULL,
    city VARCHAR(15) NOT NULL,
    zip_code VARCHAR(10) NOT NULL,
    status BIGINT NOT NULL DEFAULT '1',
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
    account_type VARCHAR(10) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT clock_timestamp(),
    status BIGINT NOT NULL DEFAULT '1'
);

INSERT INTO accounts VALUES
(01, 100, 'Saving', '2025-01-22 10:25:09', 1 ),
(02, 101, 'School_Deposits', '2025-01-05 03:09:09', 1 ),
(03, 102, 'Saving', '2025-01-10 09:34:12', 0 ),
(04, 103, 'Business', '2024-12-13 01:45:50', 1 ),
(05, 104, 'Saving', '2025-01-22 02:18:27', 1 );

CREATE TABLE "transactions" (
    id SERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id),
    amount FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL  DEFAULT clock_timestamp(),
    updated_at TIMESTAMP NOT NULL  DEFAULT clock_timestamp()
);

INSERT INTO transactions VALUES
(2000, 01, 40000, '2025-01-22 10:56:09'),
(2001, 02, 450000, '2025-01-06 3:00:34'),
(2002, 03, 3400, '2025-01-11 04:02:34'),
(2003, 04, 234500, '2025-01-22 08:57:23'),
(2004, 05, 21000, "2025-01-20 03:21:04");

-- +goose Down
DROP TABLE IF EXISTS customers;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS transactions
