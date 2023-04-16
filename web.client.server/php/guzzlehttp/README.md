#php 8.1 usando nginx

Abaixo o melhor resultado encontrado para o php8.1 usando JIT e nginx.

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


running (1m30.0s), 000/100 VUs, 649014 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 2.5 GB 27 MB/s
     data_sent......................: 81 MB  901 kB/s
     http_req_blocked...............: avg=3.19µs  min=691ns  med=1.62µs  max=36.53ms p(90)=1.98µs  p(95)=2.21µs 
     http_req_connecting............: avg=1.27µs  min=0s     med=0s      max=36.49ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=13.81ms min=1.75ms med=13.55ms max=67.91ms p(90)=14.56ms p(95)=15.64ms
       { expected_response:true }...: avg=13.81ms min=1.75ms med=13.55ms max=67.91ms p(90)=14.56ms p(95)=15.64ms
     http_req_failed................: 0.00%  ✓ 0           ✗ 649014
     http_req_receiving.............: avg=25.33µs min=9.8µs  med=24.65µs max=5.64ms  p(90)=30.99µs p(95)=33.16µs
     http_req_sending...............: avg=8.14µs  min=3.62µs med=8.05µs  max=5.16ms  p(90)=10.03µs p(95)=10.98µs
     http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=13.78ms min=1.57ms med=13.52ms max=67.73ms p(90)=14.53ms p(95)=15.6ms 
     http_reqs......................: 649014 7210.198645/s
     iteration_duration.............: avg=13.86ms min=1.97ms med=13.59ms max=68.08ms p(90)=14.61ms p(95)=15.69ms
     iterations.....................: 649014 7210.198645/s
     vus............................: 100    min=100       max=100 
     vus_max........................: 100    min=100       max=100 
```