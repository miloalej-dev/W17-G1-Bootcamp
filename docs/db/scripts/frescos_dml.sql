
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


-- Insert statements for populating the frescos.sections table

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

