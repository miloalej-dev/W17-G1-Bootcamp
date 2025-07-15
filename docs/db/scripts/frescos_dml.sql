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
    ,"2020-04-04",10,5,1,1);

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
