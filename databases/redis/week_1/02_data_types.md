# Redis Problem 2: Data Types

**Tests:** Lists, Sets, Hashes - understanding when to use each data structure

## Tasks

### Part 1: Lists (ordered, duplicates allowed)
```redis
# Create a list of recent user actions
RPUSH user:1:actions "login" "view_page" "click_button" "logout"

# Your tasks:
# 1. Get all actions (LRANGE)
# 2. Get only the last 2 actions
# 3. Add "purchase" to the front of the list (LPUSH)
# 4. Remove and return the first action (LPOP)
# 5. Get the length of the list (LLEN)
```

Your commands and outputs:
```redis

```

### Part 2: Sets (unordered, no duplicates)
```redis
# Create sets of user interests
SADD user:1:interests "golang" "redis" "postgres" "docker"
SADD user:2:interests "python" "redis" "kubernetes" "docker"

# Your tasks:
# 1. Check if "golang" is in user:1's interests (SISMEMBER)
# 2. Get all interests for user:1 (SMEMBERS)
# 3. Find common interests between user:1 and user:2 (SINTER)
# 4. Find all unique interests across both users (SUNION)
# 5. Find interests user:1 has that user:2 doesn't (SDIFF)
```

Your commands and outputs:
```redis

```

### Part 3: Hashes (field-value pairs, like a mini object)
```redis
# Store user profile as a hash
HSET user:1:profile name "Alice" email "alice@example.com" age "28"

# Your tasks:
# 1. Get Alice's email (HGET)
# 2. Get all fields and values (HGETALL)
# 3. Update age to 29 (HSET)
# 4. Add a new field "city" with value "NYC"
# 5. Get only name and city (HMGET)
# 6. Check if "phone" field exists (HEXISTS)
```

Your commands and outputs:
```redis

```

## Questions to Answer
1. When would you use a List vs a Set?
2. Why use a Hash instead of storing JSON as a string value?
3. What's the time complexity of SADD, SISMEMBER, and SINTER?
