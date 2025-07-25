-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0;
SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0;
SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE =
        'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema frescos
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `frescos`;

-- -----------------------------------------------------
-- Schema frescos
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `frescos` DEFAULT CHARACTER SET utf8mb4;
USE `frescos`;

-- -----------------------------------------------------
-- Table `frescos`.`buyers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`buyers`;

CREATE TABLE IF NOT EXISTS `frescos`.`buyers`
(
    `id`             INT    AUTO_INCREMENT      NOT NULL,
    `card_number_id` VARCHAR(64) NULL DEFAULT NULL,
    `first_name`     VARCHAR(64) NULL DEFAULT NULL,
    `last_name`      VARCHAR(64) NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`localities`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`localities`;

CREATE TABLE IF NOT EXISTS `frescos`.`localities`
(
    `id`       INT         NOT NULL,
    `locality` VARCHAR(64) NULL DEFAULT NULL,
    `province` VARCHAR(64) NULL DEFAULT NULL,
    `country`  VARCHAR(64) NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`carriers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`carriers`;

CREATE TABLE IF NOT EXISTS `frescos`.`carriers`
(
    `id`          INT AUTO_INCREMENT NOT NULL,
    `cid`         VARCHAR(64)  NULL DEFAULT NULL,
    `name`        VARCHAR(64)  NULL DEFAULT NULL,
    `address`     VARCHAR(128) NULL DEFAULT NULL,
    `telephone`   VARCHAR(16)  NULL DEFAULT NULL,
    `locality_id` INT          NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_carries_locality_idx` (`locality_id` ASC) VISIBLE,
    CONSTRAINT `fk_carries_locality`
        FOREIGN KEY (`locality_id`)
            REFERENCES `frescos`.`localities` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`warehouses`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`warehouses`;

CREATE TABLE IF NOT EXISTS `frescos`.`warehouses`
(
    `id`                  INT AUTO_INCREMENT NOT NULL,
    `address`             VARCHAR(128) NULL DEFAULT NULL,
    `telephone`           VARCHAR(16)  NULL DEFAULT NULL,
    `warehouse_code`      VARCHAR(32)  NULL DEFAULT NULL,
    `minimum_capacity`    INT          NULL DEFAULT NULL,
    `minimum_temperature` INT          NULL DEFAULT NULL,
    `locality_id`         INT          NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_warehouses_locality_idx` (`locality_id` ASC) VISIBLE,
    CONSTRAINT `fk_warehouses_locality`
        FOREIGN KEY (`locality_id`)
            REFERENCES `frescos`.`localities` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`employees`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`employees`;

CREATE TABLE IF NOT EXISTS `frescos`.`employees`
(
    `id`             INT         NOT NULL AUTO_INCREMENT,
    `card_number_id` VARCHAR(64) NULL DEFAULT NULL,
    `first_name`     VARCHAR(64) NULL DEFAULT NULL,
    `last_name`      VARCHAR(64) NULL DEFAULT NULL,
    `warehouse_id`   INT         NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_employees_warehouses_idx` (`warehouse_id` ASC) VISIBLE,
    CONSTRAINT `fk_employees_warehouses`
        FOREIGN KEY (`warehouse_id`)
            REFERENCES `frescos`.`warehouses` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`product_type`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`product_type`;

CREATE TABLE IF NOT EXISTS `frescos`.`product_type`
(
    `id`          INT          NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `name`        VARCHAR(64)  NOT NULL,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`sellers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`sellers`;

CREATE TABLE IF NOT EXISTS `frescos`.`sellers`
(
    `id`          INT     AUTO_INCREMENT     NOT NULL,
    `name`        VARCHAR(64)  NOT NULL,
    `address`     VARCHAR(128) NOT NULL,
    `telephone`   VARCHAR(16)  NOT NULL,
    `locality_id` INT          NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_sellers_locality_idx` (`locality_id` ASC) VISIBLE,
    CONSTRAINT `fk_sellers_locality`
        FOREIGN KEY (`locality_id`)
            REFERENCES `frescos`.`localities` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`products`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`products`;

CREATE TABLE IF NOT EXISTS `frescos`.`products`
(
    `id`                               INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `product_code`                     VARCHAR(32)                 NOT NULL,
    `description`                      VARCHAR(255)                NOT NULL,
    `width`                            DECIMAL(19, 2)              NOT NULL,
    `height`                           DECIMAL(19, 2)              NOT NULL,
    `length`                           DECIMAL(19, 2)              NOT NULL,
    `net_weight`                       DECIMAL(19, 2)              NOT NULL,
    `expiration_rate`                  DECIMAL(19, 2)              NOT NULL,
    `recommended_freezing_temperature` DECIMAL(19, 2)              NOT NULL,
    `freezing_rate`                    DECIMAL(19, 2)              NOT NULL,
    `product_type_id`                  INT                         NOT NULL,
    `seller_id`                        INT                         NULL DEFAULT NULL,

    PRIMARY KEY (`id`),
    INDEX `fk_products_sellers_idx` (`seller_id` ASC) VISIBLE,
    INDEX `fk_products_product_type_idx` (`product_type_id` ASC) VISIBLE,
    CONSTRAINT `fk_products_product_type`
        FOREIGN KEY (`product_type_id`)
            REFERENCES `frescos`.`product_type` (`id`),
    CONSTRAINT `fk_products_sellers1`
        FOREIGN KEY (`seller_id`)
            REFERENCES `frescos`.`sellers` (`id`)
            ON DELETE CASCADE
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`sections`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`sections`;

CREATE TABLE IF NOT EXISTS `frescos`.`sections`
(
    `id`                  INT AUTO_INCREMENT NOT NULL,
    `section_number`      VARCHAR(64)    NULL DEFAULT NULL,
    `current_capacity`    INT            NULL DEFAULT NULL,
    `current_temperature` DECIMAL(19, 2) NULL DEFAULT NULL,
    `maximum_capacity`    INT            NULL DEFAULT NULL,
    `minimum_capacity`    INT            NULL DEFAULT NULL,
    `minimum_temperature` DECIMAL(19, 2) NULL DEFAULT NULL,
    `warehouse_id`        INT            NOT NULL,
    `product_type_id`     INT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_sections_warehouses_idx` (`warehouse_id` ASC) VISIBLE,
    INDEX `fk_sections_product_type_idx` (`product_type_id` ASC) VISIBLE,
    CONSTRAINT `fk_sections_product_type`
        FOREIGN KEY (`product_type_id`)
            REFERENCES `frescos`.`product_type` (`id`),
    CONSTRAINT `fk_sections_warehouses1`
        FOREIGN KEY (`warehouse_id`)
            REFERENCES `frescos`.`warehouses` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`product_batches`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`product_batches`;

CREATE TABLE IF NOT EXISTS `frescos`.`product_batches`
(
    `id`                  INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `batch_number`        VARCHAR(32)    NOT NULL,
    `current_quantity`    INT            NOT NULL,
    `current_temperature` DECIMAL(19, 2) NOT NULL,
    `due_date`            DATE           NOT NULL,
    `initial_quantity`    INT            NOT NULL,
    `manufacturing_date`  DATE           NOT NULL,
    `manufacturing_hour`  TIME           NOT NULL,
    `minimum_temperature` DECIMAL(19, 2) NOT NULL,
    `section_id`          INT            NOT NULL,
    `product_id`          INT            UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uq_batch_number` (`batch_number` ASC) VISIBLE,
    INDEX `fk_product_batches_sections_idx` (`section_id` ASC) VISIBLE,
    INDEX `fk_product_batches_products_idx` (`product_id` ASC) VISIBLE,
    CONSTRAINT `fk_product_batches_products`
        FOREIGN KEY (`product_id`)
            REFERENCES `frescos`.`products` (`id`)
            ON DELETE CASCADE
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_product_batches_sections`
        FOREIGN KEY (`section_id`)
            REFERENCES `frescos`.`sections` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`inbound_orders`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`inbound_orders`;

CREATE TABLE IF NOT EXISTS `frescos`.`inbound_orders`
(
    `id`               INT UNSIGNED AUTO_INCREMENT NOT NULL,
    `order_date`       DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
    `order_number`     VARCHAR(64) NOT NULL,
    `employee_id`      INT         NOT NULL,
    `warehouse_id`     INT         NOT NULL,
    `product_batch_id` INT UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `uq_order_number` (`order_number` ASC) VISIBLE,
    INDEX `fk_inbound_orders_employees_idx` (`employee_id` ASC) VISIBLE,
    INDEX `fk_inbound_orders_warehouses_idx` (`warehouse_id` ASC) VISIBLE,
    INDEX `fk_inbound_orders_product_batches_idx` (`product_batch_id` ASC) VISIBLE,
    CONSTRAINT `fk_inbound_orders_employees`
        FOREIGN KEY (`employee_id`)
            REFERENCES `frescos`.`employees` (`id`),
    CONSTRAINT `fk_inbound_orders_product_batches`
        FOREIGN KEY (`product_batch_id`)
            REFERENCES `frescos`.`product_batches` (`id`)
            ON DELETE CASCADE
            ON UPDATE NO ACTION,
    CONSTRAINT `fk_inbound_orders_warehouses`
        FOREIGN KEY (`warehouse_id`)
            REFERENCES `frescos`.`warehouses` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`product_records`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`product_records`;

CREATE TABLE IF NOT EXISTS `frescos`.`product_records`
(
    `id`             INT  AUTO_INCREMENT NOT NULL,
    `last_update`    DATETIME(6)    NULL DEFAULT NULL,
    `purchase_price` DECIMAL(19, 2) NULL DEFAULT NULL,
    `sale_price`     DECIMAL(19, 2) NULL DEFAULT NULL,
    `product_id`     INT            UNSIGNED NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_product_records_products_idx` (`product_id` ASC) VISIBLE,
    CONSTRAINT `fk_product_records_products`
        FOREIGN KEY (`product_id`)
            REFERENCES `frescos`.`products` (`id`)
            ON DELETE CASCADE
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`order_status`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`order_status`;

CREATE TABLE IF NOT EXISTS `frescos`.`order_status`
(
    `id`          INT          NOT NULL,
    `name`        VARCHAR(64)  NULL DEFAULT NULL,
    `description` VARCHAR(255) NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`purchase_orders`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`purchase_orders`;

CREATE TABLE IF NOT EXISTS `frescos`.`purchase_orders`
(
    `id`              INT     AUTO_INCREMENT    NOT NULL,
    `order_number`    VARCHAR(64) Unique NULL DEFAULT NULL,
    `order_date`      DATETIME(6) NULL DEFAULT NULL,
    `tracing_code`    VARCHAR(64) NULL DEFAULT NULL,
    `buyer_id`        INT         NOT NULL,
    `warehouse_id`    INT         NOT NULL,
    `carrier_id`      INT         NOT NULL,
    `order_status_id` INT         NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_purchase_orders_buyers_idx` (`buyer_id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_warehouses_idx` (`warehouse_id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_carriers_idx` (`carrier_id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_order_status_idx` (`order_status_id` ASC) VISIBLE,
    CONSTRAINT `fk_purchase_orders_buyers`
        FOREIGN KEY (`buyer_id`)
            REFERENCES `frescos`.`buyers` (`id`),
    CONSTRAINT `fk_purchase_orders_carriers`
        FOREIGN KEY (`carrier_id`)
            REFERENCES `frescos`.`carriers` (`id`),
    CONSTRAINT `fk_purchase_orders_order_status`
        FOREIGN KEY (`order_status_id`)
            REFERENCES `frescos`.`order_status` (`id`),
    CONSTRAINT `fk_purchase_orders_warehouses`
        FOREIGN KEY (`warehouse_id`)
            REFERENCES `frescos`.`warehouses` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


-- -----------------------------------------------------
-- Table `frescos`.`order_details`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`order_details`;

CREATE TABLE IF NOT EXISTS `frescos`.`order_details`
(
    `id`                 INT       AUTO_INCREMENT NOT NULL,
    `quantity`           INT            NULL DEFAULT NULL,
    `clean_lines_status` VARCHAR(64)    NULL DEFAULT NULL,
    `temperature`        DECIMAL(19, 2) NULL DEFAULT NULL,
    `product_record_id`  INT            NOT NULL,
    `purchase_order_id`  INT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_order_details_product_records_idx` (`product_record_id` ASC) VISIBLE,
    INDEX `fk_order_details_purchase_orders_idx` (`purchase_order_id` ASC) VISIBLE,
    CONSTRAINT `fk_order_details_product_records`
        FOREIGN KEY (`product_record_id`)
            REFERENCES `frescos`.`product_records` (`id`),
    CONSTRAINT `fk_order_details_purchase_orders`
        FOREIGN KEY (`purchase_order_id`)
            REFERENCES `frescos`.`purchase_orders` (`id`)
)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb4;


SET SQL_MODE = @OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS;
