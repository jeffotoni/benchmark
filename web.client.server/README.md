# Benchmark

Este cenário é de client server, onde criamos a comunicação entre dois seviços ambos utilizando rEST e o método GET 
e temos um RESPONSE em JSON vindo do nosso server.client.

A comunicação entre os serviços mais a saída conseguimos observar o tempo gasto em diversas linguagens de programação.

O nosso objetivo é ter uma noção quando o assunto é memória, cpu, quantidade de requisições suportadas e os 
tempos médios de respostas, nada muito complexo só para matar a curiosidade mesmo.

### Exemplo de saída do teste
```bash
stress.api.client git:(master) ✗ k6 run -d 90s -u 100 k6/script-get.js

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


running (1m30.0s), 000/100 VUs, 312283 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 1.2 GB 13 MB/s
     data_sent......................: 39 MB  434 kB/s
     http_req_blocked...............: avg=2.53µs  min=638ns   med=1.01µs  max=32.82ms  p(90)=1.51µs  p(95)=1.66µs 
     http_req_connecting............: avg=1.14µs  min=0s      med=0s      max=29.98ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=28.78ms min=17.74ms med=28.49ms max=137.97ms p(90)=31.29ms p(95)=32.11ms
       { expected_response:true }...: avg=28.78ms min=17.74ms med=28.49ms max=137.97ms p(90)=31.29ms p(95)=32.11ms
     http_req_failed................: 0.00%  ✓ 0           ✗ 312283
     http_req_receiving.............: avg=28.26µs min=14.03µs med=26.18µs max=4.81ms   p(90)=34.64µs p(95)=37.72µs
     http_req_sending...............: avg=5.75µs  min=3.54µs  med=5.05µs  max=30.52ms  p(90)=7.22µs  p(95)=7.8µs  
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=28.74ms min=17.7ms  med=28.46ms max=137.91ms p(90)=31.26ms p(95)=32.08ms
     http_reqs......................: 312283 3469.002208/s
     iteration_duration.............: avg=28.81ms min=17.77ms med=28.53ms max=165.66ms p(90)=31.33ms p(95)=32.15ms
     iterations.....................: 312283 3469.002208/s
     vus............................: 100    min=100       max=100 
     vus_max........................: 100    min=100       max=100

```