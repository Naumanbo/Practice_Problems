# Redis Week 2: Lists and Sets

## Tests: LPUSH, RPUSH, LPOP, RPOP, LRANGE, SADD, SMEMBERS, SINTER, SUNION

---

## Part 1: Lists

Redis lists are ordered collections of strings. They're implemented as linked lists, so adding elements to head or tail is O(1).

### Exercise 1: Basic List Operations

```redis
# TODO: Create a list called "tasks" with these items (in order): "email", "code", "review"
# Use RPUSH to add to the right (end)


# TODO: Add "meeting" to the beginning of the list


# TODO: Get all elements of the list
# Expected: meeting, email, code, review


# TODO: Get only the first 2 elements


# TODO: Remove and return the last element (RPOP)


# TODO: Get the length of the list

```

### Exercise 2: Building a Queue (FIFO)

Implement a simple task queue where tasks are added to the back and processed from the front.

```redis
# TODO: Create a queue "job_queue" with jobs: "job1", "job2", "job3"


# TODO: Process (remove from front) the next job
# Expected: job1


# TODO: Add a new high-priority job "urgent" to the front


# TODO: Show the current queue state

```

### Exercise 3: Building a Stack (LIFO)

```redis
# TODO: Create a stack "undo_history" with actions: "create", "edit", "delete"
# Hint: Use LPUSH so most recent is at front


# TODO: Pop the most recent action (undo it)
# Expected: delete


# TODO: Show remaining history

```

---

## Part 2: Sets

Redis sets are unordered collections of unique strings.

### Exercise 4: Basic Set Operations

```redis
# TODO: Create a set "languages" with: "python", "go", "rust", "javascript"


# TODO: Try adding "python" again - what happens?
# Note the return value


# TODO: Check if "go" is in the set


# TODO: Check if "java" is in the set


# TODO: Get all members of the set


# TODO: Get the count of members


# TODO: Remove "javascript" from the set

```

### Exercise 5: Set Operations (Union, Intersection, Difference)

```redis
# TODO: Create set "alice_skills": python, go, sql, docker


# TODO: Create set "bob_skills": go, rust, sql, kubernetes


# TODO: Find skills they BOTH have (intersection)
# Expected: go, sql


# TODO: Find ALL skills between them (union)


# TODO: Find skills Alice has that Bob doesn't (difference)
# Expected: python, docker


# TODO: Find skills Bob has that Alice doesn't
# Expected: rust, kubernetes

```

### Exercise 6: Real-World Example - Tagging System

```redis
# TODO: Create sets for blog post tags
# post:1:tags -> "redis", "database", "nosql"
# post:2:tags -> "redis", "caching", "performance"
# post:3:tags -> "postgresql", "database", "sql"


# TODO: Find all posts tagged with "redis"
# Hint: Use SMEMBERS on each, or think about inverse index


# TODO: Find tags common to post:1 and post:2


# TODO: Find all unique tags across all posts

```

---

## Part 3: Combining Lists and Sets

### Exercise 7: Unique Recent Items

Build a "recently viewed" feature that:
- Keeps only last 5 items
- No duplicates

```redis
# TODO: User views product "A"
# 1. Remove "A" if it exists (LREM)
# 2. Add "A" to front (LPUSH)
# 3. Trim to 5 items (LTRIM)


# TODO: User views: B, C, A, D, E, F (in that order)
# Write commands for each view


# TODO: Show final list
# Expected: F, E, D, A, C (most recent first, A moved up, B pushed out)

```

---

## Answers Reference

Run these to verify your solutions:

```redis
# Exercise 1 verification
LRANGE tasks 0 -1

# Exercise 4 verification
SMEMBERS languages

# Exercise 5 verification
SINTER alice_skills bob_skills
```
