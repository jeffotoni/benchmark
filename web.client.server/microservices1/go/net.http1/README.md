#go net/http (puro)

Abaixo o melhor resultado encontrado para o net/http puro sem libs ou frameworks.

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
running (1m30.0s), 000/100 VUs, 1925840 complete and 0 interrupted iterations
default ✓ [===============] 100 VUs  1m30s

     data_received..................: 7.3 GB  81 MB/s
     data_sent......................: 241 MB  2.7 MB/s
     http_req_blocked...............: avg=2.12µs  min=672ns    med=1.61µs  max=29.81ms p(90)=2.02µs  p(95)=2.24µs 
     http_req_connecting............: avg=154ns   min=0s       med=0s      max=27ms    p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=4.61ms  min=154.58µs med=3.78ms  max=44.82ms p(90)=9.26ms  p(95)=11.27ms
       { expected_response:true }...: avg=4.61ms  min=154.58µs med=3.78ms  max=44.82ms p(90)=9.26ms  p(95)=11.27ms
     http_req_failed................: 0.00%   ✓ 0            ✗ 1925840
     http_req_receiving.............: avg=28.28µs min=9.46µs   med=21.27µs max=8.08ms  p(90)=29.24µs p(95)=34.18µs
     http_req_sending...............: avg=10.42µs min=3.71µs   med=7.85µs  max=30.07ms p(90)=9.33µs  p(95)=10.25µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=4.57ms  min=130.11µs med=3.75ms  max=44.78ms p(90)=9.21ms  p(95)=11.23ms
     http_reqs......................: 1925840 21397.657295/s
     iteration_duration.............: avg=4.66ms  min=192.8µs  med=3.84ms  max=45.16ms p(90)=9.31ms  p(95)=11.33ms
     iterations.....................: 1925840 21397.657295/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  


```