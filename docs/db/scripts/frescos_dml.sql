-- Insert statements for populating the frescos database

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
-- LOCALIDADES
INSERT INTO `frescos`.`localities` (`id`, `locality`, `province`, `country`) VALUES
    (1, 'Bogotá', 'Cundinamarca', 'Colombia');

-- WAREHOUSES
INSERT INTO `frescos`.`warehouses` (`id`, `address`, `telephone`, `warehouse_code`, `minimum_capacity`, `minimum_temperature`, `locality_id`) VALUES
    (1, 'Cra 15 #45-23', '6013456789', 'WH-BOG-001', 50, 4, 1);

-- CARRIERS
INSERT INTO `frescos`.`carriers` (`id`, `name`, `address`, `telephone`, `locality_id`) VALUES
    (1, 'Transporte Ágil S.A.', 'Av. 68 #22-15', '6015566777', 1);

-- ORDER STATUS
INSERT INTO `frescos`.`order_status` (`id`, `name`, `description`) VALUES
                                                                       (1, 'Pendiente', 'La orden está pendiente de procesamiento'),
                                                                       (2, 'En tránsito', 'La orden ha sido despachada y está en tránsito'),
                                                                       (3, 'Entregada', 'La orden fue entregada satisfactoriamente');



INSERT INTO `frescos`.`purchase_orders` (
    `id`, `order_number`, `order_date`, `tracing_code`,
    `buyer_id`, `warehouse_id`, `carrier_id`, `order_status_id`
) VALUES
      (1, 'PO-20250715-001', '2025-07-15 10:00:00', 'TRC001', 1, 1, 1, 1),
      (2, 'PO-20250715-002', '2025-07-14 14:30:00', 'TRC002', 1, 1, 1, 2),
      (3, 'PO-20250713-003', '2025-07-13 09:15:00', 'TRC003', 1, 1, 1, 3);

INSERT INTO `frescos`.`sellers` (
    `id`, `name`, `address`, `telephone`, `locality_id`
) VALUES
      (1, 'Frutas El Valle', 'Calle 123 #45-67', '3123456789', 1),
      (2, 'Verduras La Cosecha', 'Cra 45 #12-34', '3139876543', 1);

INSERT INTO `frescos`.`product_type` (
    `id`, `description`, `name`
) VALUES
      (1, 'Frutas frescas y tropicales', 'Frutas'),
      (2, 'Hortalizas de hoja', 'Hortalizas');

INSERT INTO `frescos`.`products` (
    `id`, `product_code`, `description`, `width`, `height`, `length`, `net_weight`,
    `expiration_rate`, `recommended_freezing_temperature`, `product_type`, `seller_id`, `product_type_id`
) VALUES
      (1, 'PRD-001', 'Caja de fresas', 10.5, 5.0, 15.2, 1.2, 0.5, -2.0, 1, 1, 1),
      (2, 'PRD-002', 'Bolsa de espinacas', 8.0, 4.0, 12.0, 0.8, 0.3, -1.0, 2, 2, 2);

INSERT INTO `frescos`.`product_records` (
    `id`, `last_update`, `purchase_price`, `sale_price`, `product_id`
) VALUES
      (1, '2025-07-15 10:00:00', 20.50, 30.00, 1),
      (2, '2025-07-15 11:30:00', 10.75, 15.99, 2);