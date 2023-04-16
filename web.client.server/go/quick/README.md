#go quick

Abaixo o melhor resultado encontrado para o quick.

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


running (1m30.0s), 000/100 VUs, 2081624 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 7.8 GB  87 MB/s
     data_sent......................: 260 MB  2.9 MB/s
     http_req_blocked...............: avg=2.34µs  min=693ns    med=1.54µs  max=48.24ms  p(90)=2.04µs  p(95)=2.32µs 
     http_req_connecting............: avg=303ns   min=0s       med=0s      max=48.17ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=4.25ms  min=155.9µs  med=3.6ms   max=134.36ms p(90)=7.27ms  p(95)=8.99ms 
       { expected_response:true }...: avg=4.25ms  min=155.9µs  med=3.6ms   max=134.36ms p(90)=7.27ms  p(95)=8.99ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 2081624
     http_req_receiving.............: avg=31.51µs min=8.63µs   med=20.42µs max=18.32ms  p(90)=28.86µs p(95)=34.39µs
     http_req_sending...............: avg=10.81µs min=3.49µs   med=7.64µs  max=42.42ms  p(90)=9.39µs  p(95)=10.7µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=4.21ms  min=131.34µs med=3.56ms  max=134.2ms  p(90)=7.22ms  p(95)=8.93ms 
     http_reqs......................: 2081624 23128.025952/s
     iteration_duration.............: avg=4.31ms  min=193.86µs med=3.66ms  max=134.4ms  p(90)=7.33ms  p(95)=9.06ms 
     iterations.....................: 2081624 23128.025952/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  

```