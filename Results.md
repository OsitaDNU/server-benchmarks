# Benchmark Results
This document contains the benchmark results for the Go, PHP, and Node.js servers using `wrk`. See the [README.md](README.md) file for more information on the benchmark setup and the commands used to run the benchmarks.

## Benchmark Setup
- **Threads:** 12
- **Connections:** 400
- **Duration:** 30 seconds
- Intel Core i7-12800H CPU (14C, 20T)

The benchmark command used:

```sh
wrk -t12 -c400 -d30s http://localhost:8080
```

## Results
Per thread results:
| Server            | Requests/sec | Latency (ms) |
|-------------------|--------------|--------------|
| Go                | 24,380       | 2.15ms       |
| PHP               | 555          | 59.61ms      |
| Node.js           | 2,830        | 24.29ms      |
| PHP (OpenSwoole)  | 7,250        | 4.58ms       |

