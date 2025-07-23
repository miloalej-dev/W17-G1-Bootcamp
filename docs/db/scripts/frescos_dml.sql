-- Select the database to use
USE `frescos`;

SET NAMES utf8mb4;
SET CHARACTER SET utf8mb4;

-- Insert statements for countries
INSERT INTO `frescos`.`countries` (`id`, `country`) VALUES
(1, 'Argentina'),
(2, 'Colombia'),
(3, 'México');

-- Insert statements for provinces
INSERT INTO `frescos`.`provinces` (`id`, `province`, `country_id`) VALUES
(1, 'Buenos Aires', 1),
(2, 'Córdoba', 1),
(3, 'Santa Fe', 1),
(4, 'Mendoza', 1),
(5, 'Atlántico', 2),
(6, 'Bolívar', 2),
(7, 'Antioquia', 2),
(8, 'Cundinamarca', 2),
(9, 'Valle del Cauca', 2),
(10, 'Jalisco', 3),
(11, 'Nuevo León', 3),
(12, 'Ciudad de México', 3),
(13, 'Puebla', 3),
(14, 'Baja California', 3),
(15, 'San Salvador', 3);

-- Insert statements of localities
INSERT INTO `frescos`.`localities` (`id`, `locality`, `province_id`) VALUES
(1, 'Buenos Aires', 1),      -- Buenos Aires, Argentina
(2, 'Córdoba', 2),           -- Córdoba, Argentina
(3, 'Rosario', 3),           -- Santa Fe, Argentina
(4, 'Mendoza', 4),           -- Mendoza, Argentina
(5, 'La Plata', 1),          -- Buenos Aires, Argentina
(6, 'Barranquilla', 5),      -- Atlántico, Colombia
(7, 'Cartagena', 6),         -- Bolívar, Colombia
(8, 'Medellín', 7),          -- Antioquia, Colombia
(9, 'Bogotá', 8),            -- Cundinamarca, Colombia
(10, 'Cali', 9),             -- Valle del Cauca, Colombia
(11, 'Guadalajara', 10),     -- Jalisco, México
(12, 'Monterrey', 11),       -- Nuevo León, México
(13, 'Ciudad de México', 12), -- Ciudad de México, México
(14, 'Puebla', 13),          -- Puebla, México
(15, 'Tijuana', 14),         -- Baja California, México
(16, 'Quilmes', 1),          -- Buenos Aires, Argentina
(17, 'Santo Domingo', 9),    -- Valle del Cauca, Colombia
(18, 'San Salvador', 15),    -- San Salvador, México
(19, 'Santa Fe', 3);         -- Santa Fe, Argentina

-- Insert statements of sellers
INSERT INTO `sellers` (`id`, `name`, `address`, `telephone`, `locality_id`) VALUES
(1, 'Compañía A', 'Calle Falsa 123, Buenos Aires', '1122334455', 1),
(2, 'Compañía B', 'Avenida Siempre Viva 742, Córdoba', '2233445566', 2),
(3, 'Compañía C', 'Paseo de la Reforma 456, Rosario', '3344556677', 3),
(4, 'Compañía D', 'Calle de la Paz 789, Mendoza', '4455667788', 4),
(5, 'Compañía E', 'Calle Independencia 321, La Plata', '5566778899', 5),
(6, 'Compañía F', 'Carrera 70 #20-5, Barranquilla', '6677889900', 6),
(7, 'Compañía G', 'Calle Real 25, Cartagena', '7788990011', 7),
(8, 'Compañía H', 'Calle 50 #30-10, Medellín', '8899001122', 8),
(9, 'Compañía I', 'Carrera 7 #32-12, Bogotá', '9900112233', 9),
(10, 'Compañía J', 'Avenida 6 #25-4, Cali', '1012233445', 10),
(11, 'Compañía K', 'Calle 12 #2-3, Guadalajara', '1212345678', 11),
(12, 'Compañía L', 'Avenida Juárez 59, Monterrey', '1312456789', 12),
(13, 'Compañía M', 'Paseo de la Reforma 123, Ciudad de México', '1412567890', 13),
(14, 'Compañía N', 'Calle de la Paz 456, Puebla', '1512678901', 14),
(15, 'Compañía O', 'Calle del Sol 789, Tijuana', '1612789012', 15),
(16, 'Compañía P', 'Avenida Libertad 123, Quilmes', '1712890123', 16),
(17, 'Compañía Q', 'Carrera 5 #10-20, Santo Domingo', '1812901234', 17),
(18, 'Compañía R', 'Calle 9 #15-30, San Salvador', '1913012345', 18),
(19, 'Compañía S', 'Avenida Cero 78, Santa Fe', '2023123456', 19),
(20, 'Compañía T', 'Calle América 42, Cúcuta', '2123234567', 19),
(21, 'Compañía U', 'Calle de los Vendedores 101, Buenos Aires', '2223345678', 1),
(22, 'Compañía V', 'Avenida del Río 102, Córdoba', '2324456789', 2),
(23, 'Compañía W', 'Paseo del Mercado 103, Rosario', '2425567890', 3),
(24, 'Compañía X', 'Calle Verde 104, Mendoza', '2526678901', 4),
(25, 'Compañía Y', 'Calle del Lago 105, La Plata', '2627789012', 5),
(26, 'Compañía Z', 'Carrera 11 #50-1, Barranquilla', '2728890123', 6),
(27, 'Compañía AA', 'Calle del Mar 106, Cartagena', '2829901234', 7),
(28, 'Compañía AB', 'Calle 9 #69-1, Medellín', '2930012345', 8),
(29, 'Compañía AC', 'Carrera 15 #73-12, Bogotá', '3041123456', 9),
(30, 'Compañía AD', 'Avenida 1 #20-1, Cali', '3152233445', 10),
(31, 'Compañía AE', 'Calle de Arte 107, Guadalajara', '3263345678', 11),
(32, 'Compañía AF', 'Avenida de los Mártires 108, Monterrey', '3374456789', 12),
(33, 'Compañía AG', 'Calle del Sol 109, Ciudad de México', '3485567890', 13),
(34, 'Compañía AH', 'Calle de la Libertad 110, Puebla', '3596678901', 14),
(35, 'Compañía AI', 'Calle de la Paz 111, Tijuana', '3607789012', 15),
(36, 'Compañía AJ', 'Calle Templanza 112, Quilmes', '3718890123', 16),
(37, 'Compañía AK', 'Calle de la Esperanza 113, Santo Domingo', '3829901234', 17),
(38, 'Compañía AL', 'Calle del Amor 114, San Salvador', '3930012345', 18),
(39, 'Compañía AM', 'Avenida de la Amistad 115, Santa Fe', '4041123456', 19),
(40, 'Compañía AN', 'Calle del Encuentro 116, Cúcuta', '4152234567', 19),
(41, 'Compañía AO', 'Calle de la Montaña 117, Buenos Aires', '4263345678', 1),
(42, 'Compañía AP', 'Avenida de la Alegría 118, Córdoba', '4374456789', 2),
(43, 'Compañía AQ', 'Paseo de los Ríos 119, Rosario', '4485567890', 3),
(44, 'Compañía AR', 'Calle del Horizonte 120, Mendoza', '4596678901', 4),
(45, 'Compañía AS', 'Calle del Recuerdo 121, La Plata', '4707789012', 5),
(46, 'Compañía AT', 'Carrera 20 #85-3, Barranquilla', '4818890123', 6),
(47, 'Compañía AU', 'Calle del Manantial 122, Cartagena', '4929901234', 7),
(48, 'Compañía AV', 'Calle de los Talentos 123, Medellín', '5030012345', 8),
(49, 'Compañía AW', 'Carrera 10 #5-8, Bogotá', '5141123456', 9),
(50, 'Compañía AX', 'Avenida del Futuro 124, Cali', '5252234567', 10);

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

-- Step 2: Populate the child table 'products'.
INSERT INTO `products` (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate , product_type_id)
VALUES
( 'JKL012', 'Gourmet truffle mashed potatoes', 7.25, 55.19, 133.35, 3.83, 7.51, -13.47, -3.78, 1),
( 'QRS345', 'Farm-fresh kale', 50.35, 106.70, 15.00, 1.95, 6.60, -14.85, -1.24, 2),
( 'QRS123', 'Organic arugula and beet salad', 1.59, 27.04, 72.92, 1.76, 2.97, -15.89, -2.38, 3),
( 'STU789', 'Organic strawberries', 28.42, 49.17, 69.23, 3.82, 9.08, -31.93, 0.35, 4),
( 'ABC123', 'Hand-rolled sushi combo', 113.96, 75.86, 151.19, 6.74, 2.50, -42.64, -7.91, 5),
( 'YZA567', 'Handcrafted gluten-free bread', 151.32, 39.27, 105.51, 1.28, 9.44, -40.94, -1.53, 6);



INSERT INTO buyers (id, card_number_id, first_name, last_name)
VALUES
(1, '428-62-7504', 'Gracie', 'Hatter'),
(2, '721-99-3742', 'Tabbitha', 'Cucuzza'),
(3, '299-04-0115', 'Rhonda', 'Houseman'),
(4, '428-04-0662', 'Sharia', 'O''Brogane'),
(5, '702-09-4957', 'Filmore', 'O'' Culligan'),
(6, '123-17-1836', 'Alick', 'Dabnot'),
(7, '730-44-0280', 'Selby', 'Gregson'),
(8, '162-07-6620', 'Lucius', 'Durdle'),
(9, '750-60-2271', 'Abbye', 'Wedmore'),
(10, '620-25-5585', 'Genny', 'Mothersole');

INSERT INTO product_records (id, last_update, purchase_price, sale_price, product_id)
VALUES
(1, '2024-06-11 16:45:00.123456', 3.99, 29.99, 1),
(2, '2024-06-12 16:45:00.123456', 22.99, 2.79, 2),
(3, '2024-06-13 16:45:00.123456', 59.99, 2.49, 3),
(4, '2024-06-14 16:45:00.123456', 59.99, 29.99, 4),
(5, '2024-06-15 16:45:00.123456', 2.49, 69.99, 5),
(6, '2024-06-16 16:45:00.123456', 6.99, 39.99, 6),
(7, '2024-06-17 16:45:00.123456', 5.49, 29.99, 1),
(8, '2024-06-18 16:45:00.123456', 18.99, 6.49, 2),
(9, '2024-06-19 16:45:00.123456', 5.29, 5.49, 3),
(10, '2024-06-20 16:45:00.123456', 64.99, 7.99, 4);

INSERT INTO warehouses
(id, warehouse_code, address, telephone, minimum_capacity, minimum_temperature, locality_id)
VALUES
    (1, '49349-189', 'Room 1780', '209-196-8436', 18, -4, 1),
    (2, '49349-790', 'PO Box 60689', '286-543-7343', 100, 52, 1),
    (3, '48951-7027', 'PO Box 40683', '323-380-2538', 20, 47, 1),
    (4, '52125-405', 'PO Box 76971', '904-142-2437', 82, 7, 1),
    (5, '0074-3333', 'Apt 1487', '559-200-1497', 80, -3, 1),
    (6, '10671-023', '4th Floor', '177-904-1618', 70, -4, 1),
    (7, '36987-3249', 'Apt 641', '558-424-2815', 24, -7, 1),
    (8, '34690-8001', 'Suite 90', '108-953-2113', 37, 20, 1),
    (9, '58281-561', 'Apt 263', '464-599-1731', 17, 24, 1),
    (10, '65643-336', '17th Floor', '110-222-2797', 60, 66, 1),
    (11, '0944-8503', 'Room 551', '586-176-1501', 52, -8, 11),
    (12, '68094-106', 'PO Box 97201', '794-740-7182', 66, 54, 1),
    (13, '59630-780', 'Apt 1966', '462-468-5531', 70, -9, 1),
    (14, '55154-7716', '6th Floor', '789-241-4571', 71, 36, 1),
    (15, '66129-105', 'Suite 92', '108-233-7993', 92, 49, 1),
    (16, '41163-690', 'Apt 107', '830-926-4604', 74, 16, 1),
    (17, '37012-647', 'Suite 64', '716-955-5236', 3, -4, 1),
    (18, '16571-101', '18th Floor', '592-836-0118', 74, -1, 1),
    (19, '54738-963', '18th Floor', '579-229-6699', 22, -3, 1),
    (20, '42865-307', '6th Floor', '520-862-2960', 45, 98, 1);

-- Insert statement for employee
INSERT INTO `employees` (`id`, `card_number_id`, `first_name`, `last_name`, `warehouse_id`) VALUES
(1, 'C0001', 'John', 'Doe', 1),
(2, 'C0002', 'Jane', 'Smith', 1),
(3, 'C0003', 'Alice', 'Johnson', 2),
(4, 'C0004', 'Bob', 'Brown', 2),
(5, 'C0005', 'Charlie', 'Davis', 3),
(6, 'C0006', 'Diana', 'Miller', 3),
(7, 'C0007', 'Evan', 'Wilson', 4),
(8, 'C0008', 'Fiona', 'Taylor', 4),
(9, 'C0009', 'George', 'Anderson', 5),
(10, 'C0010', 'Hannah', 'Thomas', 5),
(11, 'C0011', 'Isaac', 'Jackson', 1),
(12, 'C0012', 'Julia', 'White', 1),
(13, 'C0013', 'Kyle', 'Harris', 2),
(14, 'C0014', 'Laura', 'Martin', 2),
(15, 'C0015', 'Michael', 'Thompson', 3),
(16, 'C0016', 'Nina', 'Garcia', 3),
(17, 'C0017', 'Oscar', 'Martinez', 4),
(18, 'C0018', 'Pam', 'Robinson', 4),
(19, 'C0019', 'Quinn', 'Clark', 5),
(20, 'C0020', 'Ryan', 'Rodriguez', 5),
(21, 'C0021', 'Sophia', 'Lewis', 1),
(22, 'C0022', 'Tom', 'Lee', 1),
(23, 'C0023', 'Uma', 'Walker', 2),
(24, 'C0024', 'Victor', 'Hall', 2),
(25, 'C0025', 'Wendy', 'Allen', 3),
(26, 'C0026', 'Xander', 'Young', 3),
(27, 'C0027', 'Yara', 'Hernandez', 4),
(28, 'C0028', 'Zach', 'King', 4),
(29, 'C0029', 'Amy', 'Wright', 5),
(30, 'C0030', 'Brian', 'Scott', 1),
(31, 'C0031', 'Clara', 'Adams', 1),
(32, 'C0032', 'Daniel', 'Baker', 2),
(33, 'C0033', 'Ella', 'Gonzalez', 2),
(34, 'C0034', 'Frank', 'Nelson', 3),
(35, 'C0035', 'Grace', 'Carter', 3),
(36, 'C0036', 'Henry', 'Mitchell', 4),
(37, 'C0037', 'Ivy', 'Perez', 4),
(38, 'C0038', 'Jack', 'Roberts', 5),
(39, 'C0039', 'Kathy', 'Turner', 5),
(40, 'C0040', 'Liam', 'Phillips', 1),
(41, 'C0041', 'Mia', 'Campbell', 1),
(42, 'C0042', 'Noah', 'Parker', 2),
(43, 'C0043', 'Olivia', 'Evans', 2),
(44, 'C0044', 'Paul', 'Edwards', 3),
(45, 'C0045', 'Quincy', 'Collins', 3),
(46, 'C0046', 'Rose', 'Stewart', 4),
(47, 'C0047', 'Sam', 'Sanchez', 4),
(48, 'C0048', 'Tina', 'Morris', 5),
(49, 'C0049', 'Ulysses', 'Rogers', 5),
(50, 'C0050', 'Vera', 'Reed', 1);



INSERT INTO carriers (id, cid, name, address, telephone, locality_id) VALUES
(1, 'CID#1', 'Thoughtstorm', 'PO Box 22954', '288-738-3936', 1),
(2, 'CID#2', 'Browsezoom', 'Suite 35', '940-137-0407', 1),
(3, 'CID#3', 'Skipstorm', 'Suite 74', '840-873-2923', 1),
(4, 'CID#4', 'Avamm', 'Apt 640', '847-917-8225', 1),
(5, 'CID#5', 'Eire', 'PO Box 56757', '384-527-6343', 1),
(6, 'CID#6', 'JumpXS', '2nd Floor', '551-728-3925', 1),
(7, 'CID#7', 'Mycat', 'Room 1041', '341-238-3077', 1),
(8, 'CID#8', 'Tagopia', 'Apt 277', '568-723-2205', 1),
(9, 'CID#9', 'Edgetag', '19th Floor', '811-992-3510', 1),
(10, 'CID#10', 'Edgeify', 'Apt 450', '680-505-0391', 1),
(11, 'CID#11', 'Meevee', '16th Floor', '181-434-3837', 1),
(12, 'CID#12', 'Mydeo', 'PO Box 4575', '193-669-2263', 1),
(13, 'CID#13', 'Dynava', 'Apt 197', '538-860-0102', 1),
(14, 'CID#14', 'Miboo', '2nd Floor', '915-701-5966', 1),
(15, 'CID#15', 'Yoveo', 'Apt 776', '958-369-9070', 1),
(16, 'CID#16', 'Tambee', 'Apt 107', '359-918-1051', 1),
(17, 'CID#17', 'Skidoo', 'Apt 1600', '838-491-7980', 1),
(18, 'CID#18', 'Yabox', 'PO Box 96159', '922-135-9305', 1),
(19, 'CID#19', 'Meedoo', 'PO Box 56809', '971-217-9192', 1),
(20, 'CID#20', 'Jaxnation', 'Apt 802', '747-275-6029', 1);

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



INSERT INTO `frescos`.`sections` (
    `section_number`, `current_capacity`, `current_temperature`,
    `maximum_capacity`, `minimum_capacity`, `minimum_temperature`,
    `product_type_id`, `warehouse_id`)

VALUES
    ('1', 30, 4.00, 50, 10, 2.00, 1, 1),
    ('2', 45, 3.50, 60, 15, 2.00, 1, 1),
    ('3', 25, 5.00, 40, 10, 3.00, 1, 1),
    ('4', 50, 4.50, 70, 20, 2.50, 1, 1),
    ('5', 35, 3.80, 50, 15, 1.80, 1, 1),
    ('6', 40, 4.20, 55, 18, 2.20, 1, 2),
    ('7', 48, 3.70, 60, 20, 2.00, 1, 2),
    ('8', 20, 5.10, 35, 8, 3.50, 1, 2),
    ('9', 32, 4.60, 50, 12, 2.30, 1, 2),
    ('10', 38, 3.90, 55, 14, 1.90, 1, 2);


INSERT INTO `product_batches` (batch_number, current_quantity, current_temperature,due_date,initial_quantity,manufacturing_date,manufacturing_hour,minimum_temperature,section_id,product_id)
VALUES
    ( 1, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,1,1),
    ( 2, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,1,1),
    ( 3, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,1,1),
    ( 4, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,2,1),
    ( 5, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,2,1),
    ( 6, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,3,1),
    ( 7, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,4,1),
    ( 8, 200,20,"2022-04-04",1000
    ,"2020-04-04",10,5,5,1);

insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (1, '2012-07-16', '263-93-6778', 6, 10, 6);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (2, '2003-11-29', '613-86-9402', 10, 6, 8);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (3, '2018-12-12', '295-32-0720', 9, 5, 8);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (4, '2016-09-22', '426-37-0179', 5, 6, 7);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (5, '2015-09-13', '222-43-4329', 10, 6, 5);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (6, '2012-02-21', '850-02-8615', 9, 6, 5);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (7, '2016-09-17', '430-41-2235', 7, 10, 2);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (8, '2009-01-19', '188-43-1616', 9, 6, 8);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (9, '2008-06-10', '752-18-3288', 3, 2, 1);
insert into inbound_orders (id, order_date, order_number, employee_id, warehouse_id, product_batch_id) values (10, '2006-10-09', '284-71-5152', 10, 2, 3);