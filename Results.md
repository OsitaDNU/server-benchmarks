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
| Server | Requests/sec | Latency (ms) | 
|--------|--------------|--------------|
| Go     | 24,380     | 2.15ms       |
| PHP    | 555     | 59.61ms        | 
| Node.js| 2830     | 24.29ms        |

## Conclusion

Based on the benchmark results, it appears that Go is the fastest server technology in terms of both requests per second and throughput. However, the latency of PHP and Node.js is also relatively low, making them suitable for certain use cases where low latency is important. Ultimately, the choice of server technology depends on the specific requirements of the application and the trade-offs between performance and other factors such as latency and ease of use.