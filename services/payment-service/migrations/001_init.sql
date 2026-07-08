CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE Order IF NOT EXISTS(
    order_id VARCHAR(255) PRIMARY KEY DEFAULT gen_random_uuid(),
    UserId VARCHAR(255) NOT NULL,
    amount NUMERIC(10,2) NOT NULL DEFAULT 0,
    created_at DateTime DEFAULT NOW()
);

CREATE TABLE LecturerBalance IF NOT EXISTS(
    balance_id VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    lecturer_id VARCHAR(255) PRIMARY KEY NOT NULL,
    total_withdraw NUMERIC(10,2) NOT NULL DEFAULT 0,
    total_earnings NUMERIC(10, 2) NOT NULL DEFAULT 0,
    UpdateAt DateTime NOT NULL DEFAULT NOW()    
);

CREATE TABLE OrderItem IF NOT EXISTS(
    order_item_id VARCHAR(255) NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    order_id VARCHAR(255) NOT NULL,
    course_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    lecturer_id VARCHAR(255) NOT NULL
);

CREATE TABLE Withdrawal IF NOT EXISTS(
    withdraw_id VARCHAR(255) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    lecturer_id VARCHAR(255) NOT NULL,
    amount NUMERIC(10,2) NOT NULL DEFAULT 0,
    withdraw_at DateTime DEFAULT NOW()
);

CREATE INDEX idx_withdraw ON Withdrawal(withdraw_id, lecturer_id) IF NOT EXISTS;
CREATE UNIQUE INDEX idx_order_course ON OrderItem (order_id, course_id) IF NOT EXISTS;
CREATE UNIQUE INDEX idx_balance_lecturer ON LecturerBalance(lecturer_id, balance_id) IF NOT EXISTS;