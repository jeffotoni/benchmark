#go fiber

Abaixo o melhor resultado encontrado para o fiber.

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


running (1m30.0s), 000/100 VUs, 2125655 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 8.0 GB  89 MB/s
     data_sent......................: 266 MB  3.0 MB/s
     http_req_blocked...............: avg=3.23µs  min=693ns    med=1.57µs  max=58.79ms p(90)=2.04µs  p(95)=2.3µs  
     http_req_connecting............: avg=1.17µs  min=0s       med=0s      max=58.74ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=4.16ms  min=156.53µs med=3.8ms   max=70.42ms p(90)=7.35ms  p(95)=8.81ms 
       { expected_response:true }...: avg=4.16ms  min=156.53µs med=3.8ms   max=70.42ms p(90)=7.35ms  p(95)=8.81ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 2125655
     http_req_receiving.............: avg=30.53µs min=9.29µs   med=20.76µs max=18.56ms p(90)=29.48µs p(95)=35.01µs
     http_req_sending...............: avg=10.51µs min=3.73µs   med=7.76µs  max=56.86ms p(90)=9.49µs  p(95)=10.6µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=4.12ms  min=129.99µs med=3.76ms  max=62.64ms p(90)=7.3ms   p(95)=8.75ms 
     http_reqs......................: 2125655 23617.592624/s
     iteration_duration.............: avg=4.22ms  min=197.96µs med=3.85ms  max=92.69ms p(90)=7.41ms  p(95)=8.87ms 
     iterations.....................: 2125655 23617.592624/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  

```