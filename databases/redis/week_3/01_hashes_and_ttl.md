# Redis Week 3: Hashes Advanced & TTL Patterns

## Tests: HSET, HGET, HMGET, HINCRBY, EXPIRE, PEXPIRE, TTL, PERSIST

---

## Part 1: Advanced Hash Operations

### Exercise 1: User Session Management

```redis
# TODO: Create session hash "session:abc123" with fields:
# user_id: "42", login_time: "1700000000", ip: "192.168.1.1", role: "admin"


# TODO: Get just user_id and role (HMGET)


# TODO: Set expiration of 30 minutes (1800 seconds)


# TODO: Check remaining TTL


# TODO: Add field "last_active" with value "1700000100"
# Does adding a field reset the TTL?


# TODO: Verify TTL is still counting down
```

### Exercise 2: Atomic Counters in Hashes

```redis
# TODO: Create "stats:page:/home" with: views: 0, unique_visitors: 0, avg_load_time_ms: 250


# TODO: Increment views by 1 (HINCRBY)


# TODO: Increment views by 100


# TODO: Decrement avg_load_time_ms by 10 (HINCRBY with negative)


# TODO: HINCRBYFLOAT to add 0.5 to avg_load_time_ms


# TODO: Get all fields to verify
# Expected: views=101, unique_visitors=0, avg_load_time_ms=240.5
```

### Exercise 3: Object Comparison

```redis
# TODO: Create two product hashes:
# product:1 -> name: "Laptop", price: 999, stock: 50, category: "electronics"
# product:2 -> name: "Mouse", price: 29, stock: 200, category: "electronics"


# TODO: Get price of both products


# TODO: Decrease stock of product:1 by 1


# TODO: Check if product:1 has "discount" field (HEXISTS)
# Expected: 0 (false)


# TODO: Get number of fields in product:1 (HLEN)
# Expected: 4


# TODO: Get only field names (HKEYS) and only values (HVALS)
```

---

## Part 2: TTL Patterns

### Exercise 4: Cache with Expiration

```redis
# TODO: Create cache entry "cache:user:42:profile" as hash
# Fields: name: "Alice", email: "alice@test.com" | TTL: 300 seconds


# TODO: Check TTL


# TODO: Refresh cache: update email and reset TTL to 300


# TODO: Remove expiration entirely (PERSIST)


# TODO: Verify TTL is now -1
```

### Exercise 5: Rate Limiter with Hash

```redis
# Max 5 requests per minute per user

# TODO: User makes request at timestamp 1700000000
# HSET rate:user:42 "req:1" "1700000000"
# EXPIRE rate:user:42 60


# TODO: 4 more requests at +10, +20, +30, +40


# TODO: Check HLEN (expected: 5)


# TODO: 6th request at +50 - should be blocked (HLEN >= 5)
```

### Exercise 6: Per-Field TTL Limitation

```redis
# Redis TTL is per-KEY, not per-field

# TODO: Create hash "user:42:tokens" with access_token and refresh_token

# Question: Can you set different TTLs per field? (Answer: No)

# TODO: Workaround using separate keys:
# SET token:access:42 "abc123" EX 3600     (1 hour)
# SET token:refresh:42 "xyz789" EX 604800  (7 days)


# TODO: Verify different TTLs
```

---

## Part 3: Shopping Cart Pattern

### Exercise 7: Cart Operations

```redis
# TODO: Create cart "cart:user:42"
# product:1 qty 2, product:2 qty 1, product:3 qty 3


# TODO: Add 1 more of product:1 (HINCRBY)


# TODO: Remove product:2 (HDEL)


# TODO: Get all items (HGETALL)
# Expected: product:1=3, product:3=3


# TODO: Count distinct items (HLEN)
# Expected: 2
```

---

## Questions
1. Why use Hash vs separate string keys for related data?
2. Time complexity of HGETALL? When is it a problem?
3. Why can't you set TTL on individual hash fields?
4. When choose separate keys over a hash?
