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


running (15.0s), 000/100 VUs, 797082 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 3.0 GB 201 MB/s
     data_sent......................: 95 MB  6.3 MB/s
     http_req_blocked...............: avg=4.25µs  min=671ns    med=1.47µs  max=57.48ms p(90)=2.1µs   p(95)=2.58µs  
     http_req_connecting............: avg=327ns   min=0s       med=0s      max=21.11ms p(90)=0s      p(95)=0s      
     http_req_duration..............: avg=1.77ms  min=88.63µs  med=1.35ms  max=62.8ms  p(90)=3.62ms  p(95)=4.69ms  
       { expected_response:true }...: avg=1.77ms  min=88.63µs  med=1.35ms  max=62.8ms  p(90)=3.62ms  p(95)=4.69ms  
     http_req_failed................: 0.00%  ✓ 0            ✗ 797082
     http_req_receiving.............: avg=51.53µs min=7.44µs   med=18.24µs max=58.29ms p(90)=60.24µs p(95)=144.39µs
     http_req_sending...............: avg=16.72µs min=3.28µs   med=7µs     max=54.52ms p(90)=13.15µs p(95)=26.18µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=1.7ms   min=63.52µs  med=1.3ms   max=62.5ms  p(90)=3.53ms  p(95)=4.57ms  
     http_reqs......................: 797082 53134.624104/s
     iteration_duration.............: avg=1.86ms  min=123.99µs med=1.43ms  max=64.46ms p(90)=3.74ms  p(95)=4.82ms  
     iterations.....................: 797082 53134.624104/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 


```