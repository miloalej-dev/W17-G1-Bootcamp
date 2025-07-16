-- Select the database to use
USE `frescos`;

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

-- Step 2: Populate the child table 'products'.
INSERT INTO `products` (product_code, description, width, height, length, net_weight, expiration_rate, recommended_freezing_temperature, freezing_rate, product_type_id)
VALUES
( 'JKL012', 'Gourmet truffle mashed potatoes', 7.25, 55.19, 133.35, 3.83, 7.51, -13.47, -3.78, 1),
( 'QRS345', 'Farm-fresh kale', 50.35, 106.70, 15.00, 1.95, 6.60, -14.85, -1.24, 2),
( 'QRS123', 'Organic arugula and beet salad', 1.59, 27.04, 72.92, 1.76, 2.97, -15.89, -2.38, 3),
( 'STU789', 'Organic strawberries', 28.42, 49.17, 69.23, 3.82, 9.08, -31.93, 0.35, 4),
( 'ABC123', 'Hand-rolled sushi combo', 113.96, 75.86, 151.19, 6.74, 2.50, -42.64, -7.91, 5),
( 'YZA567', 'Handcrafted gluten-free bread', 151.32, 39.27, 105.51, 1.28, 9.44, -40.94, -1.53, 6);

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