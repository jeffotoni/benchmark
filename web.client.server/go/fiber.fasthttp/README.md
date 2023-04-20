#go fiber

Abaixo o melhor resultado encontrado para o fiber.

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


running (1m30.0s), 000/100 VUs, 3796577 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 14 GB   159 MB/s
     data_sent......................: 475 MB  5.3 MB/s
     http_req_blocked...............: avg=2.47µs  min=694ns    med=1.48µs  max=36.97ms p(90)=1.9µs   p(95)=2.23µs 
     http_req_connecting............: avg=121ns   min=0s       med=0s      max=36.4ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.29ms  min=105.45µs med=1.91ms  max=64.77ms p(90)=4.74ms  p(95)=5.62ms 
       { expected_response:true }...: avg=2.29ms  min=105.45µs med=1.91ms  max=64.77ms p(90)=4.74ms  p(95)=5.62ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 3796577
     http_req_receiving.............: avg=34.84µs min=8.33µs   med=19.22µs max=19ms    p(90)=28.18µs p(95)=79.95µs
     http_req_sending...............: avg=11.93µs min=3.42µs   med=7.33µs  max=36.27ms p(90)=9.1µs   p(95)=12.58µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=2.24ms  min=82.93µs  med=1.86ms  max=64.24ms p(90)=4.68ms  p(95)=5.56ms 
     http_reqs......................: 3796577 42183.444861/s
     iteration_duration.............: avg=2.35ms  min=143.75µs med=1.98ms  max=67.19ms p(90)=4.81ms  p(95)=5.7ms  
     iterations.....................: 3796577 42183.444861/s
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


running (15.0s), 000/100 VUs, 617701 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 2.3 GB 155 MB/s
     data_sent......................: 77 MB  5.1 MB/s
     http_req_blocked...............: avg=5.32µs  min=675ns    med=1.52µs max=36.96ms p(90)=2.09µs  p(95)=2.5µs   
     http_req_connecting............: avg=691ns   min=0s       med=0s     max=34.76ms p(90)=0s      p(95)=0s      
     http_req_duration..............: avg=2.33ms  min=105.58µs med=2.02ms max=39.55ms p(90)=4.41ms  p(95)=5.52ms  
       { expected_response:true }...: avg=2.33ms  min=105.58µs med=2.02ms max=39.55ms p(90)=4.41ms  p(95)=5.52ms  
     http_req_failed................: 0.00%  ✓ 0            ✗ 617701
     http_req_receiving.............: avg=44.84µs min=8.16µs   med=19.5µs max=23.01ms p(90)=34.3µs  p(95)=128.33µs
     http_req_sending...............: avg=15.22µs min=3.38µs   med=7.38µs max=14.68ms p(90)=10.88µs p(95)=17.28µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s     max=0s      p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=2.27ms  min=81.35µs  med=1.96ms max=39.46ms p(90)=4.33ms  p(95)=5.41ms  
     http_reqs......................: 617701 41172.921422/s
     iteration_duration.............: avg=2.41ms  min=143.51µs med=2.09ms max=76.18ms p(90)=4.5ms   p(95)=5.63ms  
     iterations.....................: 617701 41172.921422/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 

```