
-- Insert statements for populating the frescos database

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
