# microservice2 Quick

Este microservice retorna um JSON em seu response, ele recebe pelo método GET no path /v1/avatar.
O objetivo do serviço é somente retorno o JSON para simularmos os comportamentos de chamadas entre serviços utilizando http1.1.

#### Criar imagem do Serviço 
Foi criado o Dockerfile para que possamos subir 

```bash
$ docker build -f Dockerfile --no-cache -t jeffotoni/m2goquick .

```

#### Executar nossa imagem

Agora vamos abrir a porta e executar nosso microservice

```bash
$ docker run --name server.fiber --rm -p 3000:3000 jeffotoni/m2goquick
Listening on http://0.0.0.0:3000
```

#### cURL 

Irá apresentar uma lista de avatar com suas imagens e nomes.

```bash
$ curl -i -XGET http://localhost:3000/v1/avatar
HTTP/1.1 200 OK
content-type: application/json
engine: Rust/Hype
location: /v1/avatar
date: 2023-04-21T12:40:19.925350124+00:00
content-length: 3696

[{ "createdAt": "2023-04-21T12:40:19.925Z", "name":"BillyStoltenberg","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1164.jpg","id":"1" } , { "createdAt":"2022-11-05T09:01:38.207Z","name":"JodiKertzmann","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg","id":"2" } , { "createdAt":"2022-11-05T15:36:31.390Z","name":"AngelKuhlmanIV","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/75.jpg","id":"3" } , { "createdAt":"2022-11-04T19:23:42.344Z","name":"KellyNolan","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/787.jpg","id":"4" } , { "createdAt":"2022-11-05T06:33:55.777Z","name":"JeremySchneider","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1153.jpg","id":"5" } , { "createdAt":"2022-11-05T07:29:47.957Z","name":"EvanDuBuque","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/99.jpg","id":"6" } , { "createdAt":"2022-12-01T10:43:35.133Z","name":"DeloresDoyle","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/331.jpg","id":"7" } , { "createdAt":"2022-12-01T22:53:22.031Z","name":"Ms.JeremyBruen","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/383.jpg","id":"8" } , { "createdAt":"2022-12-02T00:59:05.444Z","name":"JoannRitchie","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/130.jpg","id":"9" } , { "createdAt":"2022-12-01T21:37:21.680Z","name":"Mrs.StevenCummerata","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/771.jpg","id":"10" } , { "createdAt":"2022-12-02T09:01:39.388Z","name":"HattiePfeffer","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1005.jpg","id":"11" } , { "createdAt":"2022-12-01T10:11:38.831Z","name":"EdithMacejkovic","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/882.jpg","id":"12" } , { "createdAt":"2022-12-02T08:34:12.087Z","name":"BettyBlock","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/216.jpg","id":"13" } , { "createdAt":"2022-12-02T06:56:11.901Z","name":"ArmandoBecker","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/822.jpg","id":"14" } , { "createdAt":"2022-12-01T17:22:36.489Z","name":"MargueriteWilliamson","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/205.jpg","id":"15" } , { "createdAt":"2022-12-02T02:59:48.845Z","name":"LorenePaucek","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/31.jpg","id":"16" } , { "createdAt":"2022-12-01T22:23:37.039Z","name":"TonyaHeidenreich","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/685.jpg","id":"17" } , { "createdAt":"2022-12-01T15:55:35.929Z","name":"MissAntoniaReynolds","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/208.jpg","id":"18" } , { "createdAt":"2022-12-01T14:15:27.068Z","name":"Ms.GlendaSchimmel","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/145.jpg","id":"19" } , { "createdAt":"2022-12-01T17:29:53.880Z","name":"JodySchadenDDS","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/35.jpg","id":"20" }]
```

#### Test de Stress neste serviço

Vamos observar agora como ele responde, vamos fazer um teste de carga usando o k6 e o wrk para analisarmos.
O objetivo e saber alguns detalhes sobre o comportamento do nosso serviço quando ele tomar muita porrada, muitas requisições.

Vamos fazer o teste porém rodando direto na máquina sem docker.

```bash
CPU: 	32,53
MEMORY: 17,2
k6 run -d 15s -u 100 k6/microservice2-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/microservice2-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 45s max duration (incl. graceful stop):
           * default: 100 looping VUs for 15s (gracefulStop: 30s)


running (15.0s), 000/100 VUs, 1043647 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 3.9 GB  262 MB/s
     data_sent......................: 126 MB  8.4 MB/s
     http_req_blocked...............: avg=3.12µs  min=659ns   med=1.44µs   max=40.79ms p(90)=1.99µs  p(95)=2.45µs 
     http_req_connecting............: avg=434ns   min=0s      med=0s       max=40.75ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.34ms  min=55.12µs med=905.57µs max=41.09ms p(90)=2.96ms  p(95)=3.99ms 
       { expected_response:true }...: avg=1.34ms  min=55.12µs med=905.57µs max=41.09ms p(90)=2.96ms  p(95)=3.99ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 1043647
     http_req_receiving.............: avg=39.3µs  min=6.67µs  med=17.06µs  max=39.88ms p(90)=33.66µs p(95)=97.98µs
     http_req_sending...............: avg=13.92µs min=3.21µs  med=6.72µs   max=15.77ms p(90)=12.41µs p(95)=18.29µs
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s       max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.28ms  min=35.61µs med=862.69µs max=40.57ms p(90)=2.88ms  p(95)=3.89ms 
     http_reqs......................: 1043647 69569.127977/s
     iteration_duration.............: avg=1.42ms  min=84.78µs med=972.19µs max=44.55ms p(90)=3.07ms  p(95)=4.13ms 
     iterations.....................: 1043647 69569.127977/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  

CPU:  34,72
MEMORY: 16,8
k6 run -d 90s -u 100 k6/microservice2-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/microservice2-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 2m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 1m30s (gracefulStop: 30s)


running (1m30.0s), 000/100 VUs, 6584919 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 25 GB   275 MB/s
     data_sent......................: 797 MB  8.9 MB/s
     http_req_blocked...............: avg=2.3µs   min=641ns   med=1.38µs   max=49.11ms  p(90)=1.78µs  p(95)=2.18µs 
     http_req_connecting............: avg=45ns    min=0s      med=0s       max=30.09ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.28ms  min=48.87µs med=918.96µs max=119.41ms p(90)=2.74ms  p(95)=3.63ms 
       { expected_response:true }...: avg=1.28ms  min=48.87µs med=918.96µs max=119.41ms p(90)=2.74ms  p(95)=3.63ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 6584919
     http_req_receiving.............: avg=33.84µs min=6.85µs  med=17.07µs  max=103.43ms p(90)=27.35µs p(95)=76.89µs
     http_req_sending...............: avg=11.88µs min=3.22µs  med=6.81µs   max=103.34ms p(90)=8.8µs   p(95)=13.92µs
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s       max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.24ms  min=32.3µs  med=881.71µs max=119.23ms p(90)=2.68ms  p(95)=3.55ms 
     http_reqs......................: 6584919 73164.716257/s
     iteration_duration.............: avg=1.35ms  min=72.31µs med=977.89µs max=119.69ms p(90)=2.82ms  p(95)=3.73ms 
     iterations.....................: 6584919 73164.716257/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  


```

```bash
CPU: 	68,64
MEMORY: 19,89
wrk -t12 -c100 -d15s http://localhost:3000/v1/avatar
Running 15s test @ http://localhost:3000/v1/avatar
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.92ms    1.39ms  42.11ms   88.58%
    Req/Sec    16.01k     2.70k   34.77k    81.23%
  2878225 requests in 15.07s, 10.08GB read
Requests/sec: 190992.22
Transfer/sec:    684.86MB

CPU:  70,01
MEMORY: 19,07
wrk -t12 -c100 -d90s http://localhost:3000/v1/avatar
Running 2m test @ http://localhost:3000/v1/avatar
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.93ms    1.31ms  40.87ms   87.42%
    Req/Sec    15.27k     1.72k   34.35k    73.76%
  16412788 requests in 1.50m, 57.47GB read
Requests/sec: 182204.01
Transfer/sec:    653.35MB

```