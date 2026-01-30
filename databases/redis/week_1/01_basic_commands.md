# Redis Problem 1: Basic Commands

**Tests:** SET, GET, DEL, EXISTS, TTL - fundamental key-value operations

## Setup
Install Redis locally or use Docker:
```bash
docker run -d --name redis-practice -p 6379:6379 redis
docker exec -it redis-practice redis-cli
```

## Tasks

1. **SET and GET**
   - Set a key `user:1:name` with value `"Alice"`
   - Get the value back
   - What happens if you GET a key that doesn't exist?

2. **Overwriting**
   - Set `counter` to `10`
   - Set `counter` to `20`
   - What value does GET return?

3. **EXISTS and DEL**
   - Check if `user:1:name` exists
   - Delete it
   - Check again - what does EXISTS return now?

4. **TTL (Time To Live)**
   - Set `session:abc` to `"active"` with a 60-second expiration: `SET session:abc "active" EX 60`
   - Check remaining time with `TTL session:abc`
   - What does TTL return for a key with no expiration?
   - What does TTL return for a key that doesn't exist?

## Expected Outputs
Document your commands and outputs below:

```redis
# Your commands here
```

## Questions to Answer
1. What is the time complexity of GET and SET?
2. Why might you use key naming conventions like `user:1:name`?
