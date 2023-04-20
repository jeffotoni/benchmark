#rust hyper

Abaixo o melhor resultado encontrado para o Rust usando framework Hyper.

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


running (1m30.0s), 000/100 VUs, 3661233 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 14 GB   154 MB/s
     data_sent......................: 458 MB  5.1 MB/s
     http_req_blocked...............: avg=2.17µs  min=665ns    med=1.46µs  max=38.99ms p(90)=1.85µs  p(95)=2.13µs 
     http_req_connecting............: avg=119ns   min=0s       med=0s      max=37.77ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.39ms  min=133.48µs med=2.2ms   max=40.98ms p(90)=4.02ms  p(95)=4.81ms 
       { expected_response:true }...: avg=2.39ms  min=133.48µs med=2.2ms   max=40.98ms p(90)=4.02ms  p(95)=4.81ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 3661233
     http_req_receiving.............: avg=29.23µs min=8.17µs   med=19.05µs max=12.71ms p(90)=25.41µs p(95)=31.01µs
     http_req_sending...............: avg=10.45µs min=3.48µs   med=7.25µs  max=37.84ms p(90)=8.82µs  p(95)=10.14µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=2.35ms  min=113.75µs med=2.17ms  max=38.77ms p(90)=3.98ms  p(95)=4.76ms 
     http_reqs......................: 3661233 40679.059122/s
     iteration_duration.............: avg=2.44ms  min=165.59µs med=2.26ms  max=43.13ms p(90)=4.09ms  p(95)=4.88ms 
     iterations.....................: 3661233 40679.059122/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  


```

### 15 segundos

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


running (15.0s), 000/100 VUs, 559178 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 2.1 GB 141 MB/s
     data_sent......................: 70 MB  4.7 MB/s
     http_req_blocked...............: avg=3.72µs  min=696ns    med=1.49µs  max=40.01ms p(90)=2.06µs  p(95)=2.45µs 
     http_req_connecting............: avg=939ns   min=0s       med=0s      max=39.76ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.6ms   min=156.82µs med=2.32ms  max=51.57ms p(90)=4.4ms   p(95)=5.21ms 
       { expected_response:true }...: avg=2.6ms   min=156.82µs med=2.32ms  max=51.57ms p(90)=4.4ms   p(95)=5.21ms 
     http_req_failed................: 0.00%  ✓ 0            ✗ 559178
     http_req_receiving.............: avg=38.12µs min=8.61µs   med=19.39µs max=11.45ms p(90)=29.12µs p(95)=69.22µs
     http_req_sending...............: avg=13.47µs min=3.4µs    med=7.27µs  max=36.02ms p(90)=10.02µs p(95)=14.9µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=2.55ms  min=133.14µs med=2.28ms  max=38.55ms p(90)=4.34ms  p(95)=5.12ms 
     http_reqs......................: 559178 37273.671338/s
     iteration_duration.............: avg=2.66ms  min=198.02µs med=2.39ms  max=60.3ms  p(90)=4.48ms  p(95)=5.29ms 
     iterations.....................: 559178 37273.671338/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 

```