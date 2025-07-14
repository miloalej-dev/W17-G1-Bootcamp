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
CREATE SCHEMA IF NOT EXISTS `frescos` DEFAULT CHARACTER SET utf8mb3;
USE `frescos`;

-- -----------------------------------------------------
-- Table `frescos`.`buyers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`buyers`;

CREATE TABLE IF NOT EXISTS `frescos`.`buyers`
(
    `id`             INT         NOT NULL,
    `card_number_id` VARCHAR(64) NULL DEFAULT NULL,
    `first_name`     VARCHAR(64) NULL DEFAULT NULL,
    `last_name`      VARCHAR(64) NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


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
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`carriers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`carriers`;

CREATE TABLE IF NOT EXISTS `frescos`.`carriers`
(
    `id`          INT          NOT NULL,
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
    DEFAULT CHARACTER SET = utf8mb3;


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
    INDEX `fk_warehouses_locality1_idx` (`locality_id` ASC) VISIBLE,
    CONSTRAINT `fk_warehouses_locality1`
    FOREIGN KEY (`locality_id`)
    REFERENCES `frescos`.`localities` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`employees`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`employees`;

CREATE TABLE IF NOT EXISTS `frescos`.`employees`
(
    `id`             INT         NOT NULL,
    `card_number_id` VARCHAR(64) NULL DEFAULT NULL,
    `first_name`     VARCHAR(64) NULL DEFAULT NULL,
    `last_name`      VARCHAR(64) NULL DEFAULT NULL,
    `warehouses_id`  INT         NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_employees_warehouses1_idx` (`warehouses_id` ASC) VISIBLE,
    CONSTRAINT `fk_employees_warehouses1`
    FOREIGN KEY (`warehouses_id`)
    REFERENCES `frescos`.`warehouses` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`sellers`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`sellers`;

CREATE TABLE IF NOT EXISTS `frescos`.`sellers`
(
    `id`          INT          NOT NULL,
    `name`        VARCHAR(64)  NULL DEFAULT NULL,
    `address`     VARCHAR(128) NULL DEFAULT NULL,
    `telephone`   VARCHAR(16)  NULL DEFAULT NULL,
    `locality_id` INT          NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_sellers_locality1_idx` (`locality_id` ASC) VISIBLE,
    CONSTRAINT `fk_sellers_locality1`
    FOREIGN KEY (`locality_id`)
    REFERENCES `frescos`.`localities` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`product_type`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`product_type`;

CREATE TABLE IF NOT EXISTS `frescos`.`product_type`
(
    `id`          INT          NOT NULL,
    `description` VARCHAR(255) NULL,
    `name`        VARCHAR(64)  NULL,
    PRIMARY KEY (`id`)
    )
    ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `frescos`.`products`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`products`;

CREATE TABLE IF NOT EXISTS `frescos`.`products`
(
    `id`                               INT            NOT NULL,
    `product_code`                     VARCHAR(32)    NULL DEFAULT NULL,
    `description`                      VARCHAR(255)   NULL DEFAULT NULL,
    `width`                            DECIMAL(19, 2) NULL DEFAULT NULL,
    `height`                           DECIMAL(19, 2) NULL DEFAULT NULL,
    `length`                           DECIMAL(19, 2) NULL DEFAULT NULL,
    `net_weight`                       DECIMAL(19, 2) NULL DEFAULT NULL,
    `expiration_rate`                  DECIMAL(19, 2) NULL DEFAULT NULL,
    `recommended_freezing_temperature` DECIMAL(19, 2) NULL DEFAULT NULL,
    `product_type`                     INT            NULL DEFAULT NULL,
    `sellers_id`                       INT            NOT NULL,
    `product_type_id`                  INT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_products_sellers1_idx` (`sellers_id` ASC) VISIBLE,
    INDEX `fk_products_product_type1_idx` (`product_type_id` ASC) VISIBLE,
    CONSTRAINT `fk_products_sellers1`
    FOREIGN KEY (`sellers_id`)
    REFERENCES `frescos`.`sellers` (`id`),
    CONSTRAINT `fk_products_product_type1`
    FOREIGN KEY (`product_type_id`)
    REFERENCES `frescos`.`product_type` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`sections`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`sections`;

CREATE TABLE IF NOT EXISTS `frescos`.`sections`
(
    `id`                  INT            NOT NULL,
    `section_number`      VARCHAR(64)    NULL DEFAULT NULL,
    `current_capacity`    INT            NULL DEFAULT NULL,
    `current_temperature` DECIMAL(19, 2) NULL DEFAULT NULL,
    `maximum_capacity`    INT            NULL DEFAULT NULL,
    `minimum_capacity`    INT            NULL DEFAULT NULL,
    `minimum_temperature` DECIMAL(19, 2) NULL DEFAULT NULL,
    `product_type_id`     INT            NULL DEFAULT NULL,
    `warehouses_id`       INT            NOT NULL,
    `product_type_id1`    INT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_sections_warehouses1_idx` (`warehouses_id` ASC) VISIBLE,
    INDEX `fk_sections_product_type1_idx` (`product_type_id1` ASC) VISIBLE,
    CONSTRAINT `fk_sections_warehouses1`
    FOREIGN KEY (`warehouses_id`)
    REFERENCES `frescos`.`warehouses` (`id`),
    CONSTRAINT `fk_sections_product_type1`
    FOREIGN KEY (`product_type_id1`)
    REFERENCES `frescos`.`product_type` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`product_batches`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`product_batches`;

CREATE TABLE IF NOT EXISTS `frescos`.`product_batches`
(
    `id`                  INT            NOT NULL,
    `batch_number`        VARCHAR(32)    NULL DEFAULT NULL,
    `current_quantity`    INT            NULL DEFAULT NULL,
    `current_temperature` DECIMAL(19, 2) NULL DEFAULT NULL,
    `due_date`            DATE           NULL DEFAULT NULL,
    `initial_quantity`    INT            NULL DEFAULT NULL,
    `manufacturing_date`  DATE           NULL DEFAULT NULL,
    `manufacturing_hour`  TIME           NULL DEFAULT NULL,
    `minumum_temperature` DECIMAL(19, 2) NULL DEFAULT NULL,
    `sections_id`         INT            NOT NULL,
    `products_id`         INT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_product_batches_sections1_idx` (`sections_id` ASC) VISIBLE,
    INDEX `fk_product_batches_products1_idx` (`products_id` ASC) VISIBLE,
    CONSTRAINT `fk_product_batches_products1`
    FOREIGN KEY (`products_id`)
    REFERENCES `frescos`.`products` (`id`),
    CONSTRAINT `fk_product_batches_sections1`
    FOREIGN KEY (`sections_id`)
    REFERENCES `frescos`.`sections` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`inbound_orders`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`inbound_orders`;

CREATE TABLE IF NOT EXISTS `frescos`.`inbound_orders`
(
    `id`                 INT         NOT NULL,
    `order_date`         DATETIME(6) NULL DEFAULT NULL,
    `order_number`       VARCHAR(64) NULL DEFAULT NULL,
    `employees_id`       INT         NOT NULL,
    `warehouses_id`      INT         NOT NULL,
    `product_batches_id` INT         NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_inbound_orders_employees1_idx` (`employees_id` ASC) VISIBLE,
    INDEX `fk_inbound_orders_warehouses1_idx` (`warehouses_id` ASC) VISIBLE,
    INDEX `fk_inbound_orders_product_batches1_idx` (`product_batches_id` ASC) VISIBLE,
    CONSTRAINT `fk_inbound_orders_employees1`
    FOREIGN KEY (`employees_id`)
    REFERENCES `frescos`.`employees` (`id`),
    CONSTRAINT `fk_inbound_orders_product_batches1`
    FOREIGN KEY (`product_batches_id`)
    REFERENCES `frescos`.`product_batches` (`id`),
    CONSTRAINT `fk_inbound_orders_warehouses1`
    FOREIGN KEY (`warehouses_id`)
    REFERENCES `frescos`.`warehouses` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`product_records`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`product_records`;

CREATE TABLE IF NOT EXISTS `frescos`.`product_records`
(
    `id`             INT            NOT NULL,
    `last_update`    DATETIME(6)    NULL DEFAULT NULL,
    `purchase_price` DECIMAL(19, 2) NULL DEFAULT NULL,
    `sale_price`     DECIMAL(19, 2) NULL DEFAULT NULL,
    `products_id`    INT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_product_records_products1_idx` (`products_id` ASC) VISIBLE,
    CONSTRAINT `fk_product_records_products1`
    FOREIGN KEY (`products_id`)
    REFERENCES `frescos`.`products` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


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
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`purchase_orders`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`purchase_orders`;

CREATE TABLE IF NOT EXISTS `frescos`.`purchase_orders`
(
    `id`              INT         NOT NULL,
    `order_number`    VARCHAR(64) NULL DEFAULT NULL,
    `order_date`      DATETIME(6) NULL DEFAULT NULL,
    `tracing_code`    VARCHAR(64) NULL DEFAULT NULL,
    `buyers_id`       INT         NOT NULL,
    `warehouses_id`   INT         NOT NULL,
    `carriers_id`     INT         NOT NULL,
    `order_status_id` INT         NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_purchase_orders_buyers1_idx` (`buyers_id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_warehouses1_idx` (`warehouses_id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_carriers1_idx` (`carriers_id` ASC) VISIBLE,
    INDEX `fk_purchase_orders_order_status1_idx` (`order_status_id` ASC) VISIBLE,
    CONSTRAINT `fk_purchase_orders_buyers1`
    FOREIGN KEY (`buyers_id`)
    REFERENCES `frescos`.`buyers` (`id`),
    CONSTRAINT `fk_purchase_orders_carriers1`
    FOREIGN KEY (`carriers_id`)
    REFERENCES `frescos`.`carriers` (`id`),
    CONSTRAINT `fk_purchase_orders_order_status1`
    FOREIGN KEY (`order_status_id`)
    REFERENCES `frescos`.`order_status` (`id`),
    CONSTRAINT `fk_purchase_orders_warehouses1`
    FOREIGN KEY (`warehouses_id`)
    REFERENCES `frescos`.`warehouses` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


-- -----------------------------------------------------
-- Table `frescos`.`order_details`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `frescos`.`order_details`;

CREATE TABLE IF NOT EXISTS `frescos`.`order_details`
(
    `id`                 INT            NOT NULL,
    `quantity`           INT            NULL DEFAULT NULL,
    `clean_lines_status` VARCHAR(64)    NULL DEFAULT NULL,
    `temperature`        DECIMAL(19, 2) NULL DEFAULT NULL,
    `product_records_id` INT            NOT NULL,
    `purchase_orders_id` INT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_order_details_product_records_idx` (`product_records_id` ASC) VISIBLE,
    INDEX `fk_order_details_purchase_orders_idx` (`purchase_orders_id` ASC) VISIBLE,
    CONSTRAINT `fk_order_details_product_records`
    FOREIGN KEY (`product_records_id`)
    REFERENCES `frescos`.`product_records` (`id`),
    CONSTRAINT `fk_order_details_purchase_orders`
    FOREIGN KEY (`purchase_orders_id`)
    REFERENCES `frescos`.`purchase_orders` (`id`)
    )
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8mb3;


SET SQL_MODE = @OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS;
