# bun 



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


running (1m30.0s), 000/100 VUs, 315317 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 1.2 GB 13 MB/s
     data_sent......................: 38 MB  417 kB/s
     http_req_blocked...............: avg=2.59µs  min=601ns  med=1.11µs  max=29.8ms  p(90)=1.62µs  p(95)=1.91µs 
     http_req_connecting............: avg=1.25µs  min=0s     med=0s      max=29.76ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=28.5ms  min=1.47ms med=28.19ms max=89.93ms p(90)=30.67ms p(95)=32.14ms
       { expected_response:true }...: avg=28.5ms  min=1.47ms med=28.19ms max=89.93ms p(90)=30.67ms p(95)=32.14ms
     http_req_failed................: 0.00%  ✓ 0          ✗ 315317
     http_req_receiving.............: avg=19.19µs min=6.75µs med=15.13µs max=6.9ms   p(90)=22.37µs p(95)=29.74µs
     http_req_sending...............: avg=7.02µs  min=3.3µs  med=5.53µs  max=3.97ms  p(90)=7.56µs  p(95)=9.42µs 
     http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=28.47ms min=1.45ms med=28.17ms max=89.91ms p(90)=30.64ms p(95)=32.1ms 
     http_reqs......................: 315317 3502.63205/s
     iteration_duration.............: avg=28.54ms min=1.5ms  med=28.23ms max=96.28ms p(90)=30.71ms p(95)=32.19ms
     iterations.....................: 315317 3502.63205/s
     vus............................: 100    min=100      max=100 
     vus_max........................: 100    min=100      max=100 

```

```bash
wrk -t12 -c100 -d90s http://localhost:8080/v1/user
Running 2m test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.21ms    7.78ms  63.70ms   64.72%
    Req/Sec   568.46    343.42     3.30k    85.19%
  611395 requests in 1.50m, 2.17GB read
Requests/sec:   6789.36
Transfer/sec:     24.64MB
```