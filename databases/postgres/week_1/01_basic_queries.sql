-- PostgreSQL Problem 1: Basic Queries
--
-- Tests: SELECT, INSERT, UPDATE, DELETE, WHERE clause, basic filtering

-- Setup: Create a practice database and table
-- Run these first:

CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    department VARCHAR(50),
    salary DECIMAL(10, 2),
    hire_date DATE
);

-- Insert sample data
INSERT INTO employees (name, department, salary, hire_date) VALUES
    ('Alice', 'Engineering', 95000.00, '2022-01-15'),
    ('Bob', 'Engineering', 85000.00, '2023-03-20'),
    ('Carol', 'Marketing', 75000.00, '2021-06-01'),
    ('David', 'Engineering', 105000.00, '2020-09-10'),
    ('Eve', 'Marketing', 80000.00, '2022-11-30');

-- =============================================================================
-- TASKS: Write queries for each task below
-- =============================================================================

-- Task 1: Select all employees
-- Your query:


-- Task 2: Select only name and salary columns
-- Your query:


-- Task 3: Select employees in Engineering department
-- Your query:


-- Task 4: Select employees with salary > 80000
-- Your query:


-- Task 5: Select employees hired after 2022-01-01
-- Your query:


-- Task 6: Update Bob's salary to 90000
-- Your query:


-- Task 7: Delete employees with salary < 80000
-- Your query:


-- Task 8: Count employees per department
-- Hint: Use GROUP BY
-- Your query:


-- =============================================================================
-- Questions to Answer:
-- 1. What does SERIAL do for the id column?
-- 2. What's the difference between WHERE and HAVING?
-- 3. What happens if you INSERT without specifying id?
-- =============================================================================
