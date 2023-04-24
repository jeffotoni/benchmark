#rust rocket


Abaixo o melhor resultado encontrado para o Rust usando framework Hyper.

### Rust -> Rust

```bash
CPU: 40,88
MEMORIA: 32,2
k6 run -d 90s -u 100 k6/microservice1-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/microservice1-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 2m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 1m30s (gracefulStop: 30s)


running (1m30.0s), 000/100 VUs, 4451743 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 18 GB   195 MB/s
     data_sent......................: 530 MB  5.9 MB/s
     http_req_blocked...............: avg=1.84µs  min=639ns    med=1.44µs  max=33.96ms  p(90)=1.79µs  p(95)=2.1µs  
     http_req_connecting............: avg=88ns    min=0s       med=0s      max=33.33ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.96ms  min=158.24µs med=1.82ms  max=120.24ms p(90)=2.89ms  p(95)=3.39ms 
       { expected_response:true }...: avg=1.96ms  min=158.24µs med=1.82ms  max=120.24ms p(90)=2.89ms  p(95)=3.39ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 4451743
     http_req_receiving.............: avg=25.92µs min=8.42µs   med=19.13µs max=21.53ms  p(90)=24.22µs p(95)=27.29µs
     http_req_sending...............: avg=9.13µs  min=3.39µs   med=7.08µs  max=32.56ms  p(90)=8.44µs  p(95)=9.54µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.93ms  min=133.22µs med=1.79ms  max=120.19ms p(90)=2.85ms  p(95)=3.34ms 
     http_reqs......................: 4451743 49462.527174/s
     iteration_duration.............: avg=2.01ms  min=191.38µs med=1.86ms  max=151.96ms p(90)=2.95ms  p(95)=3.46ms 
     iterations.....................: 4451743 49462.527174/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  

``

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