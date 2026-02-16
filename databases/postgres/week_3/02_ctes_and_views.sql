-- Tests: CTEs (Common Table Expressions), Views, Recursive CTEs
-- Week 3: CTEs and Views
-- CTEs make complex queries readable. Views save reusable query logic.
-- Setup: Using same tables from week_2
-- Run week_2/01_aggregations.sql first to create tables

-- =============================================================================
-- PART 1: Common Table Expressions (WITH)
-- =============================================================================

-- 1. Use a CTE to find customers who spent above average
-- CTE: total spending per customer
-- Main query: filter where spending > average
-- Expected columns: customer_name, total_spent
-- YOUR QUERY:


-- 2. Use multiple CTEs to find most popular product in each category
-- CTE 1: total quantity per product
-- CTE 2: rank within category
-- Main: select rank 1
-- Expected columns: category, product_name, total_ordered
-- YOUR QUERY:


-- 3. Find customers who ordered in both March and April 2024
-- Expected columns: customer_name
-- YOUR QUERY:


-- 4. CTE: For each city, show customer count and total revenue
-- Expected columns: city, customer_count, total_revenue
-- YOUR QUERY:


-- =============================================================================
-- PART 2: Views
-- =============================================================================

-- 5. Create view customer_summary: name, city, total orders, total spending
-- YOUR QUERY:


-- 6. Query the view for New York customers
-- YOUR QUERY:


-- 7. Create view product_performance: name, category, qty sold, revenue
-- YOUR QUERY:


-- 8. Using product_performance, find best-selling product by revenue
-- YOUR QUERY:


-- 9. Create view order_details joining all three tables
-- Columns: customer name, product name, category, quantity, total price, order date
-- YOUR QUERY:


-- 10. Use order_details to find orders over $500
-- YOUR QUERY:


-- =============================================================================
-- PART 3: Recursive CTEs
-- =============================================================================

-- Setup for recursive CTE exercises:
-- CREATE TABLE employee_hierarchy (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(100),
--     manager_id INT REFERENCES employee_hierarchy(id),
--     title VARCHAR(100)
-- );
-- INSERT INTO employee_hierarchy (name, manager_id, title) VALUES
--     ('CEO', NULL, 'Chief Executive Officer'),
--     ('VP Engineering', 1, 'Vice President'),
--     ('VP Marketing', 1, 'Vice President'),
--     ('Senior Dev', 2, 'Senior Developer'),
--     ('Junior Dev', 4, 'Junior Developer'),
--     ('Marketing Lead', 3, 'Lead'),
--     ('Content Writer', 6, 'Writer');

-- 11. Recursive CTE: show full management chain
-- Expected columns: employee_name, title, manager_name, level
-- Level 0 = CEO, Level 1 = VPs, etc.
-- YOUR QUERY:


-- 12. Generate numbers 1-10 using recursive CTE (pure SQL)
-- YOUR QUERY:


-- =============================================================================
-- BONUS
-- =============================================================================

-- 13. Create materialized view of monthly revenue
-- YOUR QUERY:


-- 14. Combine CTEs + window functions: customer orders with running total
-- Expected columns: customer_name, order_date, order_total, running_total
-- YOUR QUERY:


-- =============================================================================
-- Questions:
-- 1. CTE vs subquery?
-- 2. View vs CTE?
-- 3. View vs materialized view?
-- 4. Performance implications of recursive CTEs?
-- =============================================================================
