#go quick + client fastHttp

Abaixo o melhor resultado encontrado para o quick server e o client que dispara para nosso server.client na porta 3000 é feito em
fastHttp só para vermos a dimensão e poder desta lib.

### 90 segundos
```bash
k6 run -d 90s -u 100 k6/script-get.js
          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/script-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 2m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 1m30s (gracefulStop: 30s)

running (1m30.0s), 000/100 VUs, 4257800 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 16 GB   178 MB/s
     data_sent......................: 532 MB  5.9 MB/s
     http_req_blocked...............: avg=2.64µs  min=641ns    med=1.43µs  max=28.08ms p(90)=1.89µs  p(95)=2.23µs 
     http_req_connecting............: avg=125ns   min=0s       med=0s      max=27.2ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.03ms  min=109.71µs med=1.72ms  max=49.52ms p(90)=3.85ms  p(95)=4.81ms 
       { expected_response:true }...: avg=2.03ms  min=109.71µs med=1.72ms  max=49.52ms p(90)=3.85ms  p(95)=4.81ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 4257800
     http_req_receiving.............: avg=36.94µs min=6.44µs   med=18.31µs max=20.27ms p(90)=25.48µs p(95)=33.93µs
     http_req_sending...............: avg=12.84µs min=3.26µs   med=7.07µs  max=26.59ms p(90)=8.96µs  p(95)=12.35µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.98ms  min=86.73µs  med=1.68ms  max=30.27ms p(90)=3.78ms  p(95)=4.73ms 
     http_reqs......................: 4257800 47307.889286/s
     iteration_duration.............: avg=2.1ms   min=150.69µs med=1.77ms  max=50.06ms p(90)=3.95ms  p(95)=4.93ms 
     iterations.....................: 4257800 47307.889286/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100 
```

#### 15 segundos

```bash
k6 run -d 15s -u 100 k6/script-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/script-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 45s max duration (incl. graceful stop):
           * default: 100 looping VUs for 15s (gracefulStop: 30s)


running (15.0s), 000/100 VUs, 546191 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 2.1 GB 137 MB/s
     data_sent......................: 68 MB  4.5 MB/s
     http_req_blocked...............: avg=3.18µs  min=722ns    med=1.48µs  max=32.32ms p(90)=2.09µs  p(95)=2.48µs 
     http_req_connecting............: avg=713ns   min=0s       med=0s      max=32.24ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.66ms  min=124.23µs med=2.3ms   max=68.22ms p(90)=4.9ms   p(95)=6.07ms 
       { expected_response:true }...: avg=2.66ms  min=124.23µs med=2.3ms   max=68.22ms p(90)=4.9ms   p(95)=6.07ms 
     http_req_failed................: 0.00%  ✓ 0            ✗ 546191
     http_req_receiving.............: avg=40.05µs min=8.5µs    med=19.33µs max=22.73ms p(90)=29.44µs p(95)=56.09µs
     http_req_sending...............: avg=13.56µs min=3.53µs   med=7.17µs  max=16.63ms p(90)=10µs    p(95)=14.34µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=2.61ms  min=99.77µs  med=2.25ms  max=67.82ms p(90)=4.83ms  p(95)=5.97ms 
     http_reqs......................: 546191 36393.423126/s
     iteration_duration.............: avg=2.73ms  min=157.01µs med=2.36ms  max=69ms    p(90)=4.99ms  p(95)=6.17ms 
     iterations.....................: 546191 36393.423126/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 

```