-- PostgreSQL Problem 2: JOINs
--
-- Tests: INNER JOIN, LEFT JOIN, RIGHT JOIN, understanding relationships

-- Setup: Create related tables

CREATE TABLE IF NOT EXISTS departments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    budget DECIMAL(12, 2)
);

CREATE TABLE IF NOT EXISTS employees_v2 (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    department_id INT REFERENCES departments(id),
    salary DECIMAL(10, 2)
);

-- Insert sample data
INSERT INTO departments (name, budget) VALUES
    ('Engineering', 500000.00),
    ('Marketing', 200000.00),
    ('Sales', 300000.00),
    ('HR', 150000.00);  -- Note: HR has no employees

INSERT INTO employees_v2 (name, department_id, salary) VALUES
    ('Alice', 1, 95000.00),
    ('Bob', 1, 85000.00),
    ('Carol', 2, 75000.00),
    ('David', 1, 105000.00),
    ('Eve', 2, 80000.00),
    ('Frank', 3, 70000.00),
    ('Grace', NULL, 60000.00);  -- Note: Grace has no department

-- =============================================================================
-- TASKS
-- =============================================================================

-- Task 1: INNER JOIN - Get employees with their department names
-- Only shows rows where there's a match in both tables
-- Your query:


-- Task 2: LEFT JOIN - Get ALL employees, even those without a department
-- Your query:


-- Task 3: RIGHT JOIN - Get ALL departments, even those without employees
-- Your query:


-- Task 4: Find total salary per department (include department name)
-- Your query:


-- Task 5: Find departments where total salary exceeds their budget
-- Hint: Use HAVING with GROUP BY
-- Your query:


-- Task 6: List employees in Engineering department using a subquery
-- Your query:


-- Task 7: Find the department with the highest average salary
-- Your query:


-- =============================================================================
-- Questions to Answer:
-- 1. What's the difference between INNER JOIN and LEFT JOIN?
-- 2. Why does Grace appear in LEFT JOIN but not INNER JOIN?
-- 3. Why does HR appear in RIGHT JOIN but not INNER JOIN?
-- 4. When would you use a subquery vs a JOIN?
-- =============================================================================
