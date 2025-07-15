-- Insert statements for populating the frescos database

-- Step 1: Populate the parent table 'product_type' first.
-- Note: A placeholder description is used.
INSERT INTO `product_type` (id, name, description)
VALUES
(1, 'Fruits', 'Placeholder Type Description'),
(2, 'Red Meat', 'Placeholder Type Description'),
(3, 'Grain', 'Placeholder Type Description'),
(4, 'Candy', 'Placeholder Type Description'),
(5, 'Canned food', 'Placeholder Type Description'),
(6, 'Vegetables', 'Placeholder Type Description');
INSERT INTO `frescos`.`localities` (
    `id`, `locality`, `province`, `country`
) VALUES
    (1, 'Bogotá', 'Cundinamarca', 'Colombia');

INSERT INTO `frescos`.`warehouses` (
    `id`, `address`, `telephone`, `warehouse_code`,
    `minimum_capacity`, `minimum_temperature`, `locality_id`
) VALUES
      (1, 'Carrera 15 # 45-23, Bogotá', '6013456789', 'WH-BOG-001', 50, 4, 1),
      (2, 'Carrera 50 # 12-80, Bogotá', '6019876543', 'WH-BOG-002', 60, 5, 1);
INSERT INTO `frescos`.`product_type` (
    `id`, `description`, `name`
) VALUES
    (1, 'Productos refrigerados que requieren cadena de frío, como lácteos y carnes.', 'Refrigerados');



INSERT INTO `frescos`.`sections` (
    `section_number`, `current_capacity`, `current_temperature`,
    `maximum_capacity`, `minimum_capacity`, `minimum_temperature`,
    `product_type_id`, `warehouses_id`
) VALUES
      ('SEC-001', 30, 4.00, 50, 10, 2.00, 1, 1),
      ('SEC-002', 45, 3.50, 60, 15, 2.00, 1, 1),
      ('SEC-003', 25, 5.00, 40, 10, 3.00, 1, 1),
      ('SEC-004', 50, 4.50, 70, 20, 2.50, 1, 1),
      ('SEC-005', 35, 3.80, 50, 15, 1.80, 1, 1),
      ('SEC-006', 40, 4.20, 55, 18, 2.20, 1, 2),
      ('SEC-007', 48, 3.70, 60, 20, 2.00, 1, 2),
      ('SEC-008', 20, 5.10, 35, 8, 3.50, 1, 2),
      ('SEC-009', 32, 4.60, 50, 12, 2.30, 1, 2),
      ('SEC-010', 38, 3.90, 55, 14, 1.90, 1, 2);


USE frescos;

-- Step 2: Populate the child table 'products'.
INSERT INTO `products` (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id)
VALUES
( 'JKL012', 'Gourmet truffle mashed potatoes', 7.25, 55.19, 133.35, 3.83, 7.51, -13.47, -3.78, 1),
( 'QRS345', 'Farm-fresh kale', 50.35, 106.70, 15.00, 1.95, 6.60, -14.85, -1.24, 2),
( 'QRS123', 'Organic arugula and beet salad', 1.59, 27.04, 72.92, 1.76, 2.97, -15.89, -2.38, 3),
( 'STU789', 'Organic strawberries', 28.42, 49.17, 69.23, 3.82, 9.08, -31.93, 0.35, 4),
( 'ABC123', 'Hand-rolled sushi combo', 113.96, 75.86, 151.19, 6.74, 2.50, -42.64, -7.91, 5),
( 'YZA567', 'Handcrafted gluten-free bread', 151.32, 39.27, 105.51, 1.28, 9.44, -40.94, -1.53, 6);



insert into buyers (id, card_number_id, first_name, last_name) values (1, '428-62-7504', 'Gracie', 'Hatter');
insert into buyers (id, card_number_id, first_name, last_name) values (2, '721-99-3742', 'Tabbitha', 'Cucuzza');
insert into buyers (id, card_number_id, first_name, last_name) values (3, '299-04-0115', 'Rhonda', 'Houseman');
insert into buyers (id, card_number_id, first_name, last_name) values (4, '428-04-0662', 'Sharia', 'O''Brogane');
insert into buyers (id, card_number_id, first_name, last_name) values (5, '702-09-4957', 'Filmore', 'O'' Culligan');
insert into buyers (id, card_number_id, first_name, last_name) values (6, '123-17-1836', 'Alick', 'Dabnot');
insert into buyers (id, card_number_id, first_name, last_name) values (7, '730-44-0280', 'Selby', 'Gregson');
insert into buyers (id, card_number_id, first_name, last_name) values (8, '162-07-6620', 'Lucius', 'Durdle');
insert into buyers (id, card_number_id, first_name, last_name) values (9, '750-60-2271', 'Abbye', 'Wedmore');
insert into buyers (id, card_number_id, first_name, last_name) values (10, '620-25-5585', 'Genny', 'Mothersole');