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


running (1m30.0s), 000/100 VUs, 655599 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 2.5 GB 28 MB/s
     data_sent......................: 82 MB  910 kB/s
     http_req_blocked...............: avg=3.2µs   min=676ns  med=1.59µs  max=39.36ms  p(90)=1.96µs  p(95)=2.18µs 
     http_req_connecting............: avg=1.31µs  min=0s     med=0s      max=39.33ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=13.67ms min=1.78ms med=13.36ms max=143.54ms p(90)=14.44ms p(95)=15.73ms
       { expected_response:true }...: avg=13.67ms min=1.78ms med=13.36ms max=143.54ms p(90)=14.44ms p(95)=15.73ms
     http_req_failed................: 0.00%  ✓ 0           ✗ 655599
     http_req_receiving.............: avg=24.87µs min=9.53µs med=24.23µs max=6.15ms   p(90)=30.57µs p(95)=32.72µs
     http_req_sending...............: avg=7.99µs  min=3.6µs  med=7.92µs  max=4.44ms   p(90)=9.97µs  p(95)=10.92µs
     http_req_tls_handshaking.......: avg=0s      min=0s     med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=13.64ms min=1.75ms med=13.32ms max=143.46ms p(90)=14.41ms p(95)=15.7ms 
     http_reqs......................: 655599 7283.412126/s
     iteration_duration.............: avg=13.72ms min=1.81ms med=13.4ms  max=143.61ms p(90)=14.49ms p(95)=15.79ms
     iterations.....................: 655599 7283.412126/s
     vus............................: 100    min=100       max=100 
     vus_max........................: 100    min=100       max=100 

```