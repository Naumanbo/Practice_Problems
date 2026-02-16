-- Tests: Window functions, ROW_NUMBER, RANK, DENSE_RANK, OVER, PARTITION BY
-- Week 3: Window Functions
-- Window functions perform calculations across rows related to the current row
-- WITHOUT collapsing rows like GROUP BY does.
-- Setup: Using same tables from week_2 (customers, products, orders)
-- Run week_2/01_aggregations.sql first to create tables

-- =============================================================================
-- EXERCISES
-- =============================================================================

-- 1. Add a row number to each order, ordered by order_date
-- Expected columns: order_id, customer_id, order_date, row_num
-- Expected: row numbers 1 through 10
-- YOUR QUERY:


-- 2. Rank products by price (highest first)
-- Expected columns: product_name, price, price_rank
-- Use RANK() - ties get same rank, next rank skipped
-- Expected: Laptop=1, Desk=2, Chair=3, Mouse=4, Notebook=5
-- YOUR QUERY:


-- 3. For each customer, rank their orders by date
-- Expected columns: customer_name, order_date, order_rank
-- Use ROW_NUMBER() with PARTITION BY customer_id
-- YOUR QUERY:


-- 4. Show each product's price alongside the average price of its category
-- Expected columns: product_name, category, price, category_avg_price
-- Use AVG() as window function with PARTITION BY category
-- Expected: Electronics avg ~514.99, Furniture avg ~199.99, Office avg ~4.99
-- YOUR QUERY:


-- 5. For each order, show the running total of quantity (ordered by date)
-- Expected columns: order_date, quantity, running_total
-- Use SUM() with ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW
-- YOUR QUERY:


-- 6. Show each product with price difference from most expensive in its category
-- Expected columns: product_name, category, price, max_in_category, price_diff
-- YOUR QUERY:


-- 7. Use DENSE_RANK to rank customers by total spending
-- Expected columns: customer_name, total_spent, spending_rank
-- Hint: Use a subquery or CTE first
-- YOUR QUERY:


-- 8. Show each order with previous order date (LAG) and next order date (LEAD)
-- Expected columns: order_id, order_date, prev_order_date, next_order_date
-- YOUR QUERY:


-- 9. For each category, show what percentage of category revenue each product represents
-- Expected columns: product_name, category, product_revenue, category_revenue, pct
-- YOUR QUERY:


-- 10. Find first and last order date for each customer using window functions
-- Expected columns: customer_name, order_date, first_order, last_order
-- Use FIRST_VALUE() and LAST_VALUE()
-- YOUR QUERY:


-- =============================================================================
-- BONUS: RANK vs DENSE_RANK vs ROW_NUMBER
-- =============================================================================

-- 11. Show all three ranking functions side by side for product prices
-- First add: INSERT INTO products (name, category, price) VALUES ('Shelf', 'Furniture', 149.99);
-- Expected columns: product_name, price, row_num, rank, dense_rank
-- YOUR QUERY:


-- 12. Find the second most expensive product using window function
-- YOUR QUERY:


-- =============================================================================
-- Questions:
-- 1. Difference between RANK(), DENSE_RANK(), and ROW_NUMBER()?
-- 2. How is a window function different from GROUP BY?
-- 3. What does PARTITION BY do in a window function?
-- 4. When would you use LAG/LEAD vs a self-join?
-- =============================================================================
