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

-- BUYERS
INSERT INTO `frescos`.`buyers` (`id`, `card_number_id`, `first_name`, `last_name`) VALUES
    (1, 1, 'alejo', 'CM');


INSERT INTO `frescos`.`purchase_orders` (
    `id`, `order_number`, `order_date`, `tracing_code`,
    `buyers_id`, `warehouses_id`, `carriers_id`, `order_status_id`
) VALUES
      (1, 'PO-20250715-001', '2025-07-15 10:00:00', 'TRC001', 1, 1, 1, 1),
      (2, 'PO-20250715-002', '2025-07-14 14:30:00', 'TRC002', 1, 1, 1, 2),
      (3, 'PO-20250713-003', '2025-07-13 09:15:00', 'TRC003', 1, 1, 1, 3);