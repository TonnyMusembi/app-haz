CREATE TABLE customers (
 id BIGINT PRIMARY KEY AUTO_INCREMENT,
 full_name VARCHAR(255),
 phone VARCHAR(20),
 national_id VARCHAR(20),
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE loan_products (
 id BIGINT PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(100),
 interest_rate DECIMAL(5,2),
 duration_days INT,
 penalty_rate DECIMAL(5,2),
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE loans (
 id BIGINT PRIMARY KEY AUTO_INCREMENT,
 customer_id BIGINT,
 loan_product_id BIGINT,
 principal DECIMAL(12,2),
 interest DECIMAL(12,2),
 total_amount DECIMAL(12,2),
 paid_amount DECIMAL(12,2) DEFAULT 0,
 status ENUM('pending','approved','rejected','active','completed'),
 disbursed_at TIMESTAMP NULL,
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE loan_installments (
 id BIGINT PRIMARY KEY AUTO_INCREMENT,
 loan_id BIGINT,
 due_date DATE,
 amount DECIMAL(12,2),
 paid BOOLEAN DEFAULT FALSE
);

CREATE TABLE loan_payments (
 id BIGINT PRIMARY KEY AUTO_INCREMENT,
 loan_id BIGINT,
 amount DECIMAL(12,2),
 method VARCHAR(50),
 reference VARCHAR(100),
 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

# CREATE TABLE Users(
#   id BIGINT PRIMARY KEY AUTO_INCREMENT,
#   username VARCHAR(50),
#   password VARCHAR(255),
#   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
#   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
# );


CREATE TABLE IF NOT EXISTS users (
                                   id                  BIGINT AUTO_INCREMENT PRIMARY KEY,
                                   full_name           VARCHAR(255)        NOT NULL,
                                   email               VARCHAR(255)        NOT NULL UNIQUE,
                                   phone               VARCHAR(20)         NOT NULL UNIQUE,
                                   password_hash       VARCHAR(255)        NOT NULL,
                                   role                ENUM('user','admin') NOT NULL DEFAULT 'user',
                                   is_verified         BOOLEAN             NOT NULL DEFAULT FALSE,

  -- Auth fields
                                   reset_token         VARCHAR(255)        DEFAULT NULL,
                                   reset_token_expiry  DATETIME            DEFAULT NULL,
                                   last_login          DATETIME            DEFAULT NULL,

  -- Timestamps
                                   created_at          DATETIME            NOT NULL DEFAULT NOW(),
                                   updated_at          DATETIME            NOT NULL DEFAULT NOW() ON UPDATE NOW(),

  -- Indexes
                                   INDEX idx_users_email        (email),
                                   INDEX idx_users_phone        (phone),
                                   INDEX idx_users_reset_token  (reset_token),
                                   INDEX idx_users_role         (role)
);


CREATE TABLE contact_messages (
     id BIGINT AUTO_INCREMENT PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);