-- DDL
DROP DATABASE IF EXISTS product_db;

CREATE DATABASE product_db;

USE product_db;

-- Table: warehouse
CREATE TABLE warehouse (
    `id`      int          NOT NULL AUTO_INCREMENT,
    `name`    varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_warehouse_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Table: product
CREATE TABLE product (
    `id`            int           NOT NULL AUTO_INCREMENT,
    `name`          varchar(255)  NOT NULL,
    `type`          varchar(255)  NOT NULL,
    `count`         int           NOT NULL,
    `price`         decimal(10,2) NOT NULL,
    `warehouse_id`  int           NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_product_name` (`name`),
    KEY `idx_product_warehouse_id` (`warehouse_id`),
    CONSTRAINT `fk_product_warehouse_id` FOREIGN KEY (`warehouse_id`) REFERENCES `warehouse` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;