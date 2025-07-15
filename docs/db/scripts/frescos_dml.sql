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