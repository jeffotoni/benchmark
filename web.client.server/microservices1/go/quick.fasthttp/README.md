#go quick + client fastHttp

Abaixo o melhor resultado encontrado para o quick server e o client que dispara para nosso server.client na porta 3000 é feito em
fastHttp só para vermos a dimensão e poder desta lib.

### 90 segundos
```bash
CPU: 32,22
RAM: 21,04
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


running (1m30.0s), 000/100 VUs, 4874100 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 19 GB   208 MB/s
     data_sent......................: 580 MB  6.4 MB/s
     http_req_blocked...............: avg=2.1µs   min=659ns    med=1.4µs   max=125.25ms p(90)=1.8µs   p(95)=2.12µs 
     http_req_connecting............: avg=65ns    min=0s       med=0s      max=24.54ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.77ms  min=103.61µs med=1.57ms  max=134.5ms  p(90)=2.99ms  p(95)=3.62ms 
       { expected_response:true }...: avg=1.77ms  min=103.61µs med=1.57ms  max=134.5ms  p(90)=2.99ms  p(95)=3.62ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 4874100
     http_req_receiving.............: avg=32.78µs min=7.77µs   med=18.25µs max=133.97ms p(90)=24.09µs p(95)=27.94µs
     http_req_sending...............: avg=10.68µs min=3.56µs   med=6.98µs  max=130.59ms p(90)=8.54µs  p(95)=10.64µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.73ms  min=73.23µs  med=1.54ms  max=127.88ms p(90)=2.94ms  p(95)=3.55ms 
     http_reqs......................: 4874100 54146.314828/s
     iteration_duration.............: avg=1.83ms  min=142.63µs med=1.62ms  max=142.02ms p(90)=3.06ms  p(95)=3.71ms 
     iterations.....................: 4874100 54146.314828/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  


```

#### wrk
```bash
CPU: 57,04
RAM: 21,04
 wrk -t12 -c100 -d90s http://localhost:8080/v1/user
Running 2m test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.95ms  747.83us  60.65ms   89.01%
    Req/Sec     8.99k   724.00    20.28k    77.23%
  9663324 requests in 1.50m, 34.63GB read
Requests/sec: 107317.99
Transfer/sec:    393.83MB

```

#### 15 segundos

```bash
CPU: 30,04
RAM: 20,01
k6 run -d 15s -u 100 k6/microservice1-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/microservice1-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 45s max duration (incl. graceful stop):
           * default: 100 looping VUs for 15s (gracefulStop: 30s)


running (15.0s), 000/100 VUs, 806863 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 3.1 GB 207 MB/s
     data_sent......................: 96 MB  6.4 MB/s
     http_req_blocked...............: avg=3.51µs  min=674ns    med=1.41µs  max=65.36ms p(90)=1.92µs  p(95)=2.28µs 
     http_req_connecting............: avg=1.27µs  min=0s       med=0s      max=65.32ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.78ms  min=98.48µs  med=1.53ms  max=57.63ms p(90)=3.08ms  p(95)=3.86ms 
       { expected_response:true }...: avg=1.78ms  min=98.48µs  med=1.53ms  max=57.63ms p(90)=3.08ms  p(95)=3.86ms 
     http_req_failed................: 0.00%  ✓ 0           ✗ 806863
     http_req_receiving.............: avg=35.05µs min=8.4µs    med=17.82µs max=18.82ms p(90)=25.32µs p(95)=31.64µs
     http_req_sending...............: avg=12.32µs min=3.29µs   med=6.85µs  max=38.19ms p(90)=9.38µs  p(95)=13.14µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.74ms  min=78.57µs  med=1.5ms   max=57.14ms p(90)=3.01ms  p(95)=3.76ms 
     http_reqs......................: 806863 53783.38268/s
     iteration_duration.............: avg=1.84ms  min=133.97µs med=1.58ms  max=70.41ms p(90)=3.15ms  p(95)=3.96ms 
     iterations.....................: 806863 53783.38268/s
     vus............................: 100    min=100       max=100 
     vus_max........................: 100    min=100       max=100 
```

#### wrk
```bash
CPU: 55,04
RAM: 21,01
wrk -t12 -c100 -d15s http://localhost:8080/v1/user
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.98ms  696.20us  16.02ms   84.86%
    Req/Sec     8.72k     1.02k   17.16k    86.50%
  1562269 requests in 15.06s, 5.60GB read
Requests/sec: 103763.08
Transfer/sec:    380.78MB

```