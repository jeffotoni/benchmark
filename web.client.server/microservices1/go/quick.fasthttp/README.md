#go quick + client fastHttp

Abaixo o melhor resultado encontrado para o quick server e o client que dispara para nosso server.client na porta 3000 é feito em
fastHttp só para vermos a dimensão e poder desta lib.

### 90 segundos
```bash
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


running (1m30.0s), 000/100 VUs, 4441067 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 17 GB   186 MB/s
     data_sent......................: 529 MB  5.9 MB/s
     http_req_blocked...............: avg=2.7µs   min=670ns    med=1.45µs  max=101.23ms p(90)=1.9µs   p(95)=2.25µs 
     http_req_connecting............: avg=103ns   min=0s       med=0s      max=33.78ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.94ms  min=109.01µs med=1.65ms  max=109.46ms p(90)=3.6ms   p(95)=4.47ms 
       { expected_response:true }...: avg=1.94ms  min=109.01µs med=1.65ms  max=109.46ms p(90)=3.6ms   p(95)=4.47ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 4441067
     http_req_receiving.............: avg=38.93µs min=7.31µs   med=18.52µs max=103.67ms p(90)=25.72µs p(95)=33.13µs
     http_req_sending...............: avg=12.62µs min=3.38µs   med=7.1µs   max=101.02ms p(90)=9.02µs  p(95)=12.66µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.89ms  min=80.66µs  med=1.61ms  max=107.69ms p(90)=3.53ms  p(95)=4.38ms 
     http_reqs......................: 4441067 49342.575251/s
     iteration_duration.............: avg=2.01ms  min=140.6µs  med=1.71ms  max=109.51ms p(90)=3.69ms  p(95)=4.58ms 
     iterations.....................: 4441067 49342.575251/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  

```

#### wrk
```bash
http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.30ms    1.09ms  23.32ms   78.85%
    Req/Sec     6.97k   732.31    16.82k    75.68%
  7490630 requests in 1.50m, 26.29GB read
Requests/sec:  83167.71
Transfer/sec:    298.94MB

```

#### 15 segundos

```bash
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


running (15.0s), 000/100 VUs, 711678 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 2.7 GB 179 MB/s
     data_sent......................: 85 MB  5.6 MB/s
     http_req_blocked...............: avg=3.63µs  min=702ns    med=1.45µs  max=46.73ms  p(90)=2.06µs  p(95)=2.47µs 
     http_req_connecting............: avg=817ns   min=0s       med=0s      max=46.63ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.01ms  min=103.74µs med=1.65ms  max=91.42ms  p(90)=3.86ms  p(95)=4.93ms 
       { expected_response:true }...: avg=2.01ms  min=103.74µs med=1.65ms  max=91.42ms  p(90)=3.86ms  p(95)=4.93ms 
     http_req_failed................: 0.00%  ✓ 0            ✗ 711678
     http_req_receiving.............: avg=43.99µs min=8.34µs   med=18.31µs max=20.81ms  p(90)=28.53µs p(95)=88.2µs 
     http_req_sending...............: avg=14.71µs min=3.26µs   med=6.9µs   max=47ms     p(90)=11.68µs p(95)=14.79µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.95ms  min=82.65µs  med=1.6ms   max=91.25ms  p(90)=3.77ms  p(95)=4.8ms  
     http_reqs......................: 711678 47439.242982/s
     iteration_duration.............: avg=2.09ms  min=132.21µs med=1.71ms  max=108.61ms p(90)=3.97ms  p(95)=5.07ms 
     iterations.....................: 711678 47439.242982/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 


```

#### wrk
```bash
wrk -t12 -c100 -d15s http://localhost:8080/v1/user  
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.28ms    1.13ms  27.93ms   83.13%
    Req/Sec     7.12k     0.94k   26.25k    90.09%
  1280911 requests in 15.10s, 4.50GB read
Requests/sec:  84828.26
Transfer/sec:    304.91MB
```