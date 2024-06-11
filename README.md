# Server Benchmark Comparison: Go, PHP, and Node.js

This repository contains simple HTTP servers written in Go, PHP, and Node.js. The goal is to compare the performance of these servers by benchmarking them using `wrk`.

Use the instructions below to run and benchmark each server.

## Setup

### Prerequisites

- Install Go (ideally Go 1.20 or later)
- Install PHP (ideally PHP 8.2 or later)
- Install Node.js (ideally Node.js 20 or later)
- Install `wrk` for benchmarking

### Installing `wrk`

On Linux, you can install `wrk` with the following commands:

```sh
sudo apt-get update
sudo apt-get install build-essential libssl-dev git -y
git clone https://github.com/wg/wrk.git
cd wrk
make
sudo cp wrk /usr/local/bin
```

## Running the Servers

### Go Server
Navigate to the `go` directory and run the following command:

```sh
go run main.go
```
The server will start on `http://localhost:8080`.

### PHP Server
Navigate to the `php` directory and run the following command:
```sh
php -S localhost:8080
```
The server will start on `http://localhost:8080`.

### Node.js Server
Navigate to the `node` directory and run the following command:
```sh
node server.js
```
The server will start on `http://localhost:8080`.

## Benchmarking the Servers

Use `wrk` to benchmark the servers. Open a new terminal window and run the following command:

```sh
wrk -t12 -c400 -d30s http://localhost:8080
```

This command will run `wrk` with:
- `-t12`: 12 threads
- `-c400`: 400 connections
- `-d30s`: for a duration of 30 seconds

Adjust the parameters as needed to better understand the server’s capabilities under different loads.

### Example Output

The output will show statistics like requests per second, latency, and more. Example:

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.09ms    4.95ms  45.79ms   65.22%
    Req/Sec     3.33k     1.33k    8.18k    69.19%
  1194183 requests in 30.10s, 152.23MB read
Requests/sec:  39694.77
Transfer/sec:      5.06MB
```

## Notes

- Ensure that only one server is running at a time on port 8080 to avoid conflicts.
- Adjust the parameters of the `wrk` command as needed to better understand the server’s capabilities under different loads.

## Contributing

Feel free to fork the project and submit pull requests with any enhancements or bug fixes.

## License

[MIT License](LICENSE.md)