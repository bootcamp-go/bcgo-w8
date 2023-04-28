USE product_db;

-- DML Operation
INSERT INTO warehouse (name, address) VALUES
('Fresh Products', 'Av Not Known, 1234');

INSERT INTO product (name, type, count, price, warehouse_id) VALUES
('Strawberry', 'Fruit', 10, 100.00, 1),
('Pepsi', 'Drinks', 20, 200.00, 1);