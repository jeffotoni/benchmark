# microservice2 Rust Hyper

Este microservice retorna um JSON em seu response, ele recebe pelo método GET no path /v1/avatar.
O objetivo do serviço é somente retorno o JSON para simularmos os comportamentos de chamadas entre serviços utilizando http1.1.

#### Criar imagem do Serviço 
Foi criado o Dockerfile para que possamos subir 

```bash
$ docker build -f Dockerfile --no-cache -t jeffotoni/m2rusthyper .
Sending build context to Docker daemon  296.9MB
Step 1/12 : FROM rust:1.67.0 as builder
1.67.0: Pulling from library/rust
699c8a97647f: Pull complete 
86cd158b89fd: Pull complete 
a226e961cfaa: Pull complete 
4cec535da207: Pull complete 
225fdd30e1a3: Pull complete 
203d59dc0d36: Pull complete 

```

#### Executar nossa imagem

Agora vamos abrir a porta e executar nosso microservice

```bash
$ docker run --name rust.server.hype --rm -p 3000:3000 jeffotoni/m2rusthyper
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
CPU: 	19,22
MEMORY: 22,5

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


running (15.0s), 000/100 VUs, 1346252 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 5.2 GB  346 MB/s
     data_sent......................: 163 MB  11 MB/s
     http_req_blocked...............: avg=2.19µs min=642ns   med=1.36µs   max=35.92ms p(90)=1.82µs  p(95)=2.19µs 
     http_req_connecting............: avg=189ns  min=0s      med=0s       max=21.96ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.05ms min=40.38µs med=837.42µs max=53.96ms p(90)=1.88ms  p(95)=2.44ms 
       { expected_response:true }...: avg=1.05ms min=40.38µs med=837.42µs max=53.96ms p(90)=1.88ms  p(95)=2.44ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 1346252
     http_req_receiving.............: avg=29.3µs min=6.98µs  med=16.24µs  max=36.84ms p(90)=23.33µs p(95)=26.19µs
     http_req_sending...............: avg=10.2µs min=3.04µs  med=6.44µs   max=25.59ms p(90)=11.19µs p(95)=12.33µs
     http_req_tls_handshaking.......: avg=0s     min=0s      med=0s       max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.01ms min=23.62µs med=809.95µs max=53.81ms p(90)=1.83ms  p(95)=2.37ms 
     http_reqs......................: 1346252 89743.877581/s
     iteration_duration.............: avg=1.1ms  min=62.38µs med=878.63µs max=54.17ms p(90)=1.94ms  p(95)=2.54ms 
     iterations.....................: 1346252 89743.877581/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100

CPU:  21,18
MEMORY: 14,9
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


running (1m30.0s), 000/100 VUs, 8593139 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 33 GB   368 MB/s
     data_sent......................: 1.0 GB  12 MB/s
     http_req_blocked...............: avg=1.88µs   min=629ns   med=1.35µs   max=154.07ms p(90)=1.62µs  p(95)=1.95µs 
     http_req_connecting............: avg=47ns     min=0s      med=0s       max=34.88ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=990.66µs min=43.36µs med=826.37µs max=156.7ms  p(90)=1.76ms  p(95)=2.2ms  
       { expected_response:true }...: avg=990.66µs min=43.36µs med=826.37µs max=156.7ms  p(90)=1.76ms  p(95)=2.2ms  
     http_req_failed................: 0.00%   ✓ 0            ✗ 8593139
     http_req_receiving.............: avg=26.12µs  min=6.86µs  med=16.43µs  max=154.89ms p(90)=20.17µs p(95)=23.77µs
     http_req_sending...............: avg=8.66µs   min=3.13µs  med=6.64µs   max=151.7ms  p(90)=7.87µs  p(95)=10.67µs
     http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=955.86µs min=23.85µs med=800.24µs max=120.12ms p(90)=1.73ms  p(95)=2.16ms 
     http_reqs......................: 8593139 95478.254805/s
     iteration_duration.............: avg=1.04ms   min=72.25µs med=866.34µs max=156.74ms p(90)=1.82ms  p(95)=2.27ms 
     iterations.....................: 8593139 95478.254805/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100  

```

```bash
CPU: 	58,12
MEMORY: 12,4

Running 15s test @ http://localhost:3000/v1/avatar
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   325.86us  258.49us  10.25ms   85.67%
    Req/Sec    25.86k     2.75k   50.76k    86.24%
  4654278 requests in 15.10s, 16.72GB read
Requests/sec: 308214.88
Transfer/sec:   1.11GB

CPU:  60,63
MEMORY: 10,4
wrk -t12 -c100 -d90s http://localhost:3000/v1/avatar
Running 2m test @ http://localhost:3000/v1/avatar
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   318.04us  321.94us  32.50ms   92.29%
    Req/Sec    26.27k     1.94k   51.51k    81.62%
  28241509 requests in 1.50m, 101.34GB read
Requests/sec: 313554.11
Transfer/sec:      1.13GB

```