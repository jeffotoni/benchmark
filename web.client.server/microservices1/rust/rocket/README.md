#rust rocket


Abaixo o melhor resultado encontrado para o Rust usando framework Hyper.

### Rust -> Rust


```bash
CPU: 58,83
MEMORIA: 23,09
wrk -t12 -c100 -d15s http://localhost:8080/v1/user
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.20ms  451.22us  12.60ms   78.89%
    Req/Sec     6.77k   673.32    12.31k    87.43%
  1215866 requests in 15.06s, 4.45GB read
Requests/sec:  80722.58
Transfer/sec:    302.77MB

```

```bash
CPU: 63,44
MEMORIA: 15,9
wrk -t12 -c100 -d90s http://localhost:8080/v1/user
Running 2m test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.21ms  437.74us  20.52ms   77.96%
    Req/Sec     6.67k   450.55    14.53k    85.61%
  7173815 requests in 1.50m, 26.28GB read
Requests/sec:  79639.75
Transfer/sec:    298.71MB

```