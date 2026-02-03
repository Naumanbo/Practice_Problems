-- Tests: Aggregate functions, GROUP BY, HAVING
--
-- Week 2: Aggregations and Grouping
-- Practice with COUNT, SUM, AVG, MIN, MAX, GROUP BY, HAVING

-- Setup: Create and populate sample tables
DROP TABLE IF EXISTS orders CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS customers CASCADE;

CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    city VARCHAR(50),
    signup_date DATE
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(50),
    price DECIMAL(10, 2)
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES customers(id),
    product_id INT REFERENCES products(id),
    quantity INT,
    order_date DATE
);

-- Sample data
INSERT INTO customers (name, city, signup_date) VALUES
    ('Alice', 'New York', '2024-01-15'),
    ('Bob', 'Los Angeles', '2024-02-20'),
    ('Charlie', 'New York', '2024-03-10'),
    ('Diana', 'Chicago', '2024-01-05'),
    ('Eve', 'Los Angeles', '2024-04-01');

INSERT INTO products (name, category, price) VALUES
    ('Laptop', 'Electronics', 999.99),
    ('Mouse', 'Electronics', 29.99),
    ('Desk', 'Furniture', 249.99),
    ('Chair', 'Furniture', 149.99),
    ('Notebook', 'Office', 4.99);

INSERT INTO orders (customer_id, product_id, quantity, order_date) VALUES
    (1, 1, 1, '2024-03-01'),
    (1, 2, 2, '2024-03-01'),
    (2, 3, 1, '2024-03-05'),
    (2, 4, 2, '2024-03-05'),
    (3, 1, 1, '2024-03-10'),
    (3, 5, 10, '2024-03-10'),
    (4, 2, 5, '2024-03-15'),
    (4, 5, 20, '2024-03-15'),
    (1, 4, 1, '2024-03-20'),
    (5, 1, 2, '2024-04-01');

-- =============================================================================
-- EXERCISES: Write queries for each problem
-- =============================================================================

-- 1. Count total number of orders
-- Expected: 10
-- YOUR QUERY:


-- 2. Find the total quantity of all items ordered
-- Expected: 46
-- YOUR QUERY:


-- 3. Find the average price of all products
-- Expected: ~286.99
-- YOUR QUERY:


-- 4. Find the cheapest and most expensive product prices
-- Expected: min=4.99, max=999.99
-- YOUR QUERY:


-- 5. Count how many customers are in each city
-- Expected: New York: 2, Los Angeles: 2, Chicago: 1
-- YOUR QUERY:


-- 6. Find total revenue (price * quantity) for each product category
-- Hint: JOIN products with orders
-- YOUR QUERY:


-- 7. Find customers who have placed more than 1 order
-- Hint: Use GROUP BY and HAVING
-- YOUR QUERY:


-- 8. Find the average order quantity per customer
-- YOUR QUERY:


-- 9. Find categories where total revenue exceeds $500
-- Hint: Use HAVING with aggregate
-- YOUR QUERY:


-- 10. Find the date with the highest total order quantity
-- YOUR QUERY:


-- =============================================================================
-- BONUS: More complex aggregations
-- =============================================================================

-- 11. For each city, find the total spending of customers from that city
-- (sum of price * quantity for all orders by customers in each city)
-- YOUR QUERY:


-- 12. Find customers who have ordered every product category
-- Hint: COUNT(DISTINCT category) = total categories
-- YOUR QUERY:
