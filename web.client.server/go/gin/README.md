#go gin 

Abaixo o melhor resultado encontrado para o gin utilizando keep-alive ativado.

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


running (1m30.0s), 000/100 VUs, 2024602 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 7.6 GB  85 MB/s
     data_sent......................: 253 MB  2.8 MB/s
     http_req_blocked...............: avg=2.13µs  min=679ns    med=1.57µs  max=22.13ms p(90)=2.01µs  p(95)=2.28µs 
     http_req_connecting............: avg=129ns   min=0s       med=0s      max=21.77ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=4.37ms  min=152.75µs med=3.81ms  max=52.06ms p(90)=8.44ms  p(95)=10.2ms 
       { expected_response:true }...: avg=4.37ms  min=152.75µs med=3.81ms  max=52.06ms p(90)=8.44ms  p(95)=10.2ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 2024602
     http_req_receiving.............: avg=29.98µs min=8.78µs   med=20.93µs max=11.93ms p(90)=29.53µs p(95)=35.36µs
     http_req_sending...............: avg=10.3µs  min=3.4µs    med=7.79µs  max=22.19ms p(90)=9.44µs  p(95)=10.59µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=4.33ms  min=125.38µs med=3.77ms  max=51.87ms p(90)=8.4ms   p(95)=10.15ms
     http_reqs......................: 2024602 22494.571748/s
     iteration_duration.............: avg=4.43ms  min=193.59µs med=3.87ms  max=63.87ms p(90)=8.51ms  p(95)=10.26ms
     iterations.....................: 2024602 22494.571748/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  

```