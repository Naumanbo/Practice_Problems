# Redis Week 2: Sorted Sets

## Tests: ZADD, ZRANGE, ZREVRANGE, ZRANK, ZSCORE, ZINCRBY, ZRANGEBYSCORE

---

## Overview

Sorted Sets (ZSETs) are like Sets but each member has an associated score. Members are unique, but scores can repeat. Elements are ordered by score (ascending by default).

Perfect for: leaderboards, priority queues, time-series data, rate limiting.

---

## Part 1: Basic Sorted Set Operations

### Exercise 1: Creating and Querying

```redis
# TODO: Create a leaderboard "game:scores" with these players and scores:
# alice: 1500, bob: 2200, charlie: 1800, diana: 2200, eve: 1000


# TODO: Get all players ordered by score (lowest to highest)


# TODO: Get all players ordered by score (highest to lowest)
# Use ZREVRANGE


# TODO: Get top 3 players (highest scores)


# TODO: Get alice's rank (0-indexed, highest score = rank 0)
# Use ZREVRANK


# TODO: Get bob's score

```

### Exercise 2: Score Manipulation

```redis
# TODO: Alice wins a game, add 300 to her score
# Use ZINCRBY


# TODO: Eve loses points, subtract 200 from her score
# Hint: ZINCRBY with negative value


# TODO: Show updated leaderboard (top to bottom)


# TODO: Update charlie's score to exactly 2000 (not increment)

```

---

## Part 2: Range Queries

### Exercise 3: Query by Score Range

```redis
# TODO: Find all players with scores between 1500 and 2000 (inclusive)
# Use ZRANGEBYSCORE


# TODO: Find all players with scores above 1800
# Hint: use +inf for upper bound


# TODO: Find all players with scores below 1500


# TODO: Count how many players have scores between 1000 and 2000
# Use ZCOUNT

```

### Exercise 4: Removing Elements

```redis
# TODO: Remove the player with the lowest score
# Use ZPOPMIN


# TODO: Remove all players with scores below 1200
# Use ZREMRANGEBYSCORE


# TODO: Show remaining players

```

---

## Part 3: Real-World Applications

### Exercise 5: Priority Queue

Implement a job queue where lower score = higher priority (processed first).

```redis
# TODO: Add jobs with priorities (lower = more urgent):
# "backup_db": 100 (low priority)
# "send_email": 50 (medium priority)
# "process_payment": 10 (high priority)
# "generate_report": 75 (medium priority)


# TODO: Get the highest priority job (lowest score) WITHOUT removing
# Use ZRANGE with LIMIT


# TODO: Process (remove and return) the highest priority job
# Use ZPOPMIN


# TODO: Add a new urgent job "fix_bug" with priority 5


# TODO: Show current queue ordered by priority

```

### Exercise 6: Leaderboard with Timestamps

Use sorted sets to track time-based data (scores are Unix timestamps).

```redis
# TODO: Track user logins (score = Unix timestamp)
# User "user:1" logged in at 1700000000
# User "user:2" logged in at 1700000100
# User "user:3" logged in at 1700000050


# TODO: Get users who logged in between timestamps 1700000000 and 1700000060


# TODO: Get the most recent login
# Use ZREVRANGE with LIMIT 1


# TODO: Get logins from the last hour (assuming current time is 1700003600)
# Calculate: 1700003600 - 3600 = 1700000000

```

### Exercise 7: Rate Limiting

Implement a sliding window rate limiter: max 5 requests per 60 seconds.

```redis
# Assume current time is 1700000000

# TODO: Record a request from user "api:user:123"
# Add with score = current timestamp


# TODO: Remove old entries outside the window (older than 60 seconds ago)
# Use ZREMRANGEBYSCORE with -inf to (current_time - 60)


# TODO: Count requests in the current window
# If count >= 5, rate limit exceeded


# TODO: Simulate 6 requests at different times and check if 6th is blocked
# Times: 1700000000, 1700000010, 1700000020, 1700000030, 1700000040, 1700000050


```

---

## Part 4: Combining Operations

### Exercise 8: Multi-Key Operations

```redis
# Create two leaderboards
# TODO: game1:scores -> alice:100, bob:200, charlie:150
# TODO: game2:scores -> alice:300, bob:100, diana:250


# TODO: Create a combined leaderboard (sum of scores from both games)
# Use ZUNIONSTORE with AGGREGATE SUM


# TODO: Find players who played BOTH games
# Use ZINTERSTORE


# TODO: Show the combined leaderboard

```

---

## Key Commands Reference

| Command | Description |
|---------|-------------|
| `ZADD key score member` | Add member with score |
| `ZRANGE key start stop` | Get members by rank (ascending) |
| `ZREVRANGE key start stop` | Get members by rank (descending) |
| `ZRANGEBYSCORE key min max` | Get members by score range |
| `ZSCORE key member` | Get score of member |
| `ZRANK key member` | Get rank (ascending) |
| `ZREVRANK key member` | Get rank (descending) |
| `ZINCRBY key increment member` | Increment score |
| `ZCOUNT key min max` | Count members in score range |
| `ZPOPMIN key` | Remove and return lowest scored member |
| `ZPOPMAX key` | Remove and return highest scored member |
| `ZUNIONSTORE dest numkeys key...` | Union of sorted sets |
| `ZINTERSTORE dest numkeys key...` | Intersection of sorted sets |
