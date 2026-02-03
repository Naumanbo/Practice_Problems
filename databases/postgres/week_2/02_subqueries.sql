-- Tests: Subqueries, IN, EXISTS, scalar subqueries, correlated subqueries
--
-- Week 2: Subqueries
-- Practice with different types of subqueries

-- Using same tables from 01_aggregations.sql
-- Run that file first to create tables, or copy the CREATE/INSERT statements

-- =============================================================================
-- EXERCISES: Write queries using subqueries
-- =============================================================================

-- 1. Find all customers from cities that have more than 1 customer
-- Hint: Use subquery with GROUP BY and HAVING in WHERE clause
-- YOUR QUERY:


-- 2. Find products that have never been ordered
-- Hint: Use NOT IN or NOT EXISTS
-- YOUR QUERY:


-- 3. Find customers who have ordered the most expensive product
-- Hint: Use subquery to find max price, then find who ordered it
-- YOUR QUERY:


-- 4. Find orders where quantity is above average quantity
-- Hint: Scalar subquery for average
-- YOUR QUERY:


-- 5. For each product, show its price and the average price of all products
-- Hint: Scalar subquery in SELECT
-- Expected columns: product_name, price, avg_price
-- YOUR QUERY:


-- 6. Find customers who have spent more than the average customer spending
-- Hint: Calculate total spending per customer, then compare to average
-- YOUR QUERY:


-- 7. Find products that cost more than any product in the 'Office' category
-- Hint: Use > ALL or > (SELECT MAX...)
-- YOUR QUERY:


-- 8. Find customers who have ordered from every product category
-- Hint: Compare count of distinct categories ordered vs total categories
-- YOUR QUERY:


-- =============================================================================
-- CORRELATED SUBQUERIES
-- =============================================================================

-- 9. For each customer, find their most recent order date
-- Hint: Correlated subquery references outer query's customer
-- YOUR QUERY:


-- 10. Find products where the price is above the average price in their category
-- Hint: Correlated subquery calculates avg per category
-- YOUR QUERY:


-- 11. Find customers who have placed an order in the same city as 'Alice'
-- (excluding Alice herself)
-- YOUR QUERY:


-- =============================================================================
-- EXISTS vs IN
-- =============================================================================

-- 12. Rewrite problem 2 (products never ordered) using EXISTS
-- YOUR QUERY:


-- 13. Find customers who have ordered products from the 'Electronics' category
-- Write two versions: one with IN, one with EXISTS
-- YOUR QUERY (IN version):

-- YOUR QUERY (EXISTS version):
