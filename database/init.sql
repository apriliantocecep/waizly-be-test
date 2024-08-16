CREATE DATABASE IF NOT EXISTS waizly;
USE waizly;

CREATE TABLE customers (
    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE users (
                       id INT UNSIGNED NOT NULL AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE taxes (
                       id INT UNSIGNED NOT NULL AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL,
                       rate DECIMAL(5, 2) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE currencies (
                            id INT UNSIGNED NOT NULL AUTO_INCREMENT,
                            code VARCHAR(255) NOT NULL,
                            name VARCHAR(255) NOT NULL,
                            exchange_rate DECIMAL(10, 2) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE items (
                       id INT UNSIGNED NOT NULL AUTO_INCREMENT,
                       name VARCHAR(255) NOT NULL,
                       type VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE invoices (
                          id INT UNSIGNED NOT NULL AUTO_INCREMENT,
                          subject VARCHAR(255) NOT NULL,
                          status VARCHAR(255) NOT NULL,
                          issue_date DATE NOT NULL,
                          due_date DATE NOT NULL,
                          tax_id INT UNSIGNED NOT NULL,
                          currency_id INT UNSIGNED NOT NULL,
                          customer_id INT UNSIGNED NOT NULL,
                          user_id INT UNSIGNED NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (id)
) ENGINE = InnoDB;

CREATE TABLE invoice_items (
                               id INT UNSIGNED NOT NULL AUTO_INCREMENT,
                               invoice_id INT UNSIGNED NOT NULL,
                               item_id INT UNSIGNED NOT NULL,
                               quantity DECIMAL(10, 2) UNSIGNED NOT NULL,
                               price DECIMAL(10, 2) UNSIGNED NOT NULL,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (id)
) ENGINE = InnoDB;