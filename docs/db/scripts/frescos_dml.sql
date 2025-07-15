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

INSERT INTO product_records (id, last_update, purchase_price, sale_price, products_id)
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