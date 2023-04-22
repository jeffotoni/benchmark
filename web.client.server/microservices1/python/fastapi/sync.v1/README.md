# fastAPI

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


running (1m30.1s), 000/100 VUs, 156095 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 597 MB 6.6 MB/s
     data_sent......................: 20 MB  217 kB/s
     http_req_blocked...............: avg=7.03µs  min=767ns   med=1.89µs  max=45.65ms  p(90)=2.44µs  p(95)=2.67µs 
     http_req_connecting............: avg=3.61µs  min=0s      med=0s      max=44.61ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=57.59ms min=3.15ms  med=55.93ms max=144.05ms p(90)=72.01ms p(95)=79.95ms
       { expected_response:true }...: avg=57.59ms min=3.15ms  med=55.93ms max=144.05ms p(90)=72.01ms p(95)=79.95ms
     http_req_failed................: 0.00%  ✓ 0           ✗ 156095
     http_req_receiving.............: avg=42.75ms min=16.65µs med=42.33ms max=50.37ms  p(90)=46.01ms p(95)=46.85ms
     http_req_sending...............: avg=9.75µs  min=4.11µs  med=9.19µs  max=44.92ms  p(90)=11.58µs p(95)=12.4µs 
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=14.83ms min=1.33ms  med=13.13ms max=100.97ms p(90)=30.68ms p(95)=38.01ms
     http_reqs......................: 156095 1733.187978/s
     iteration_duration.............: avg=57.66ms min=3.19ms  med=55.99ms max=144.11ms p(90)=72.07ms p(95)=80.01ms
     iterations.....................: 156095 1733.187978/s
     vus............................: 100    min=100       max=100 
     vus_max........................: 100    min=100       max=100 
```
