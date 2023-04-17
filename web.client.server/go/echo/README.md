#go echo 

Abaixo o melhor resultado encontrado para o echo utilizando keep-alive ativado.

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


running (1m30.0s), 000/100 VUs, 1733048 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 6.5 GB  72 MB/s
     data_sent......................: 217 MB  2.4 MB/s
     http_req_blocked...............: avg=2.17µs  min=688ns    med=1.62µs  max=29.97ms p(90)=2.02µs  p(95)=2.24µs 
     http_req_connecting............: avg=227ns   min=0s       med=0s      max=29.52ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=5.13ms  min=162.42µs med=4.29ms  max=82.33ms p(90)=10.33ms p(95)=12.54ms
       { expected_response:true }...: avg=5.13ms  min=162.42µs med=4.29ms  max=82.33ms p(90)=10.33ms p(95)=12.54ms
     http_req_failed................: 0.00%   ✓ 0            ✗ 1733048
     http_req_receiving.............: avg=27.46µs min=9.78µs   med=21.78µs max=10.52ms p(90)=29.53µs p(95)=33.7µs 
     http_req_sending...............: avg=9.4µs   min=3.56µs   med=7.91µs  max=7.94ms  p(90)=9.37µs  p(95)=10.16µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=5.09ms  min=138.16µs med=4.26ms  max=82.23ms p(90)=10.29ms p(95)=12.5ms 
     http_reqs......................: 1733048 19254.925888/s
     iteration_duration.............: avg=5.18ms  min=203.5µs  med=4.35ms  max=82.5ms  p(90)=10.38ms p(95)=12.6ms 
     iterations.....................: 1733048 19254.925888/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  
```