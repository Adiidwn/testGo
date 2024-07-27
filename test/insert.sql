CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
);

INSERT INTO users (id, name, email) VALUES
(1, 'John Doe', 'johndoe@example.com'),
(2, 'Jane Smith', 'janesmith@example.com'),
(3, 'Bob Johnson', 'bobjohnson@example.com');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    amount NUMERIC(10,2),
    created_at TIMESTAMP
);

INSERT INTO orders (id, user_id, amount, created_at) VALUES
(1, 1, 100.00, '2022-01-02 10:30:00'),
(2, 2, 50.00, '2022-01-03 09:00:00'),
(3, 1, 150.00, '2022-01-04 14:15:00'),
(4, 3, 200.00, '2022-01-05 17:45:00'),
(5, 2, 75.00, '2022-01-06 11:20:00');

CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255),
    order_id INT,
    price NUMERIC(10,2),
    quantity INT
);

INSERT INTO order_items (id, order_id, product_name, price, quantity) VALUES
(1, 1, 'T-Shirt', 25.00, 2),
(2, 1, 'Jeans', 50.00, 1),
(3, 2, 'Socks', 10.00, 5),
(4, 3, 'Shoes', 75.00, 2),
(5, 4, 'Jacket', 100.00, 1),
(6, 5, 'Sweater', 25.00, 3);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    email VARCHAR(255)
);

INSERT INTO users (name, email) 
VALUES 
    ('John Doe', 'john.doe@example.com'), 
    ('Jane Smith', 'jane.smith@example.com');

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT,
    amount NUMERIC(10,2),
    created_at TIMESTAMP
);

INSERT INTO orders (user_id, amount, created_at) 
VALUES 
    (1, 100.00, now()), 
    (2, 200.00, now()), 
    (1, 50.00, now());
