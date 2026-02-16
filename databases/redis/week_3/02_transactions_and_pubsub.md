# Redis Week 3: Transactions & Pub/Sub

## Tests: MULTI, EXEC, DISCARD, WATCH, PUBLISH, SUBSCRIBE, PSUBSCRIBE

---

## Part 1: Transactions (MULTI/EXEC)

### Exercise 1: Basic Transaction

```redis
# TODO: Execute as a transaction:
# 1. SET account:alice:balance 1000
# 2. SET account:bob:balance 500
# Use MULTI to start, then EXEC to execute


# TODO: Verify both values were set
```

### Exercise 2: Transfer Money

```redis
# Transfer $200 from Alice to Bob atomically

# TODO: Start transaction
# DECRBY account:alice:balance 200
# INCRBY account:bob:balance 200
# Execute


# TODO: Verify balances
# Expected: Alice=800, Bob=700
```

### Exercise 3: DISCARD (Abort Transaction)

```redis
# TODO: Start transaction
# Queue: SET important:data "bad_value"
# Then DISCARD (abort before EXEC)


# TODO: Verify important:data was NOT changed
```

### Exercise 4: WATCH (Optimistic Locking)

```redis
# WATCH makes a transaction fail if watched keys change before EXEC

# TODO: SET inventory:item:1 to 10
# WATCH inventory:item:1

# In SEPARATE terminal: SET inventory:item:1 5

# Back in first terminal:
# MULTI -> DECRBY inventory:item:1 1 -> EXEC
# Expected: nil (transaction aborted due to watched key change)


# TODO: Try again without interference - should succeed
```

### Exercise 5: Multi-Type Transaction

```redis
# TODO: In single transaction:
# 1. HSET user:1 name "Alice" status "active"
# 2. LPUSH user:1:log "login at 2024-03-01"
# 3. SADD active:users "user:1"
# 4. ZADD user:scores 100 "user:1"
# Execute all at once


# TODO: Verify each data structure was updated
```

---

## Part 2: Pub/Sub (requires 2 terminals)

### Exercise 6: Basic Pub/Sub

```redis
# Terminal 1 (Subscriber):
# SUBSCRIBE notifications

# Terminal 2 (Publisher):
# PUBLISH notifications "Hello, World!"
# PUBLISH notifications "Second message"

# Observe messages in Terminal 1
```

### Exercise 7: Pattern Subscriptions

```redis
# Terminal 1: PSUBSCRIBE user:*:events

# Terminal 2:
# PUBLISH user:1:events "login"
# PUBLISH user:2:events "purchase"
# PUBLISH user:1:messages "hello"

# Which does Terminal 1 receive?
# Expected: first two (match pattern), NOT third (doesn't match)
```

### Exercise 8: Multi-Channel

```redis
# Terminal 1: SUBSCRIBE alerts errors warnings

# Terminal 2:
# PUBLISH alerts "CPU high"
# PUBLISH errors "DB connection failed"
# PUBLISH warnings "Disk 80% full"
# PUBLISH info "All normal"

# Expected: first three only (not "info" - not subscribed)
```

---

## Part 3: Combining Concepts

### Exercise 9: Inventory System

```redis
# Setup: SET stock:widget 100, SET stock:gadget 50

# TODO: Purchase transaction:
# DECRBY stock:widget 2
# DECRBY stock:gadget 1
# RPUSH purchases "widget:2,gadget:1"
# Execute atomically


# TODO: Verify stock levels and purchase log
# Expected: widget=98, gadget=49, purchases has 1 entry
```

### Exercise 10: Limitations

```redis
# PUBLISH cannot be inside MULTI/EXEC (it runs immediately, not queued)

# Question: How would you handle: update data AND notify subscribers?
# Think about this trade-off and write your answer below:
#
# YOUR ANSWER:
```

---

## Key Commands Reference

| Command | Description |
|---------|-------------|
| `MULTI` | Start a transaction |
| `EXEC` | Execute all queued commands |
| `DISCARD` | Abort the transaction |
| `WATCH key` | Watch key for changes (optimistic lock) |
| `UNWATCH` | Cancel all watches |
| `PUBLISH channel message` | Send message to channel |
| `SUBSCRIBE channel [channel...]` | Listen to channels |
| `PSUBSCRIBE pattern [pattern...]` | Listen to channels matching pattern |
| `UNSUBSCRIBE` | Stop listening |

---

## Questions
1. Are Redis transactions truly atomic like SQL transactions?
2. What happens if one command in a MULTI/EXEC block fails?
3. Why can't you read values inside MULTI/EXEC and use them?
4. SUBSCRIBE vs PSUBSCRIBE?
5. Does Pub/Sub guarantee delivery if subscriber is offline?
6. WATCH vs Lua script for atomicity?
