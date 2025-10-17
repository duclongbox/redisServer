# Redis-Server

Redis-Server is a high-performance Redis-compatible server implemented in Go. It supports basic Redis commands and is designed for efficiency and speed.


## Functional Requirements
- Implement a Redis-compatible server supporting basic commands: PING, SET, GET, TTL, DEL.
- Handle multiple client connections concurrently using single-threaded I/O multiplexing (epoll on Linux, kqueue on macOS).
- Parse and respond using the RESP (Redis Serialization Protocol).
- Store data in-memory with support for key expiration and periodic cleanup of expired keys.

## Non-functional Requirements
- High throughput (e.g., 162,127 requests per second in benchmarks) and low latency.
- Scalable to handle up to 20,000 concurrent connections.
- Efficient memory usage with single-threaded design to avoid concurrency overhead (can update to multi-thread).

## Functional Complexity
- Core operations like GET and SET are O(1) time complexity using hash maps.
- Expiration handling: O(1) lazy checks on access, O(k) periodic active sampling (k=20 by default).
- Overall low complexity due to single-threaded event loop, avoiding locks and race conditions.

## Technologies
- Go programming language for the implementation.
- Low-level system calls for I/O multiplexing (epoll/kqueue).
- In-memory data structures: hash maps for key-value storage and expiration tracking.

## Algorithms
- Primarily uses hash tables for O(1) lookups and storage.
- Active key expiration: probabilistic sampling to delete expired keys periodically.
- Approximate LRU eviction: randomly samples keys (sample size 5) for memory management instead of full tracking (as described in learning notes, for efficiency).

## Data Storage and Caching
- Data is stored in an in-memory hash map acting as a key-value store.
- Functions as a cache with optional TTL (time-to-live) for automatic expiration.
- Single-threaded TCP server design ensures fast access without threading overhead.

## Benchmarks

### Throughput Test

```
1000000 requests completed in 6.17 seconds
50 parallel clients
3 bytes payload
keep alive: 1
multi-thread: no

Summary:
throughput summary: 162127.11 requests per second
latency summary (msec):
        avg       min       p50       p95       p99       max
      0.179     0.048     0.159     0.287     0.471     4.735
```

## Contributing

Feel free to open issues or pull requests for improvements.

## License

This project is licensed under the MIT License. (Note: Add LICENSE file if needed)
