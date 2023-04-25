#go fiber fasthttp

Abaixo o melhor resultado encontrado para o Go usando framework Fiber.

#### Criar imagem do Serviço 
Foi criado o Dockerfile para que possamos subir 

```bash
$ docker build -f Dockerfile --no-cache -t jeffotoni/m1fiberfasthttp .

```

#### Executar nossa imagem

Agora vamos abrir a porta e executar nosso microservice

```bash
$ docker run --name go.server.fiber --rm -p 3000:3000 jeffotoni/m1fiberfasthttp
Listening on http://0.0.0.0:8080
```

#### cURL 

Irá apresentar uma lista de avatar com suas imagens e nomes.

```bash
$ curl -i -XGET http://localhost:8080/v1/user
HTTP/1.1 200 OK
content-type: application/json
engine: Go/Fiber
location: /v1/user
date: 2023-04-21T12:40:19.925350124+00:00
content-length: 3696

[{ "createdAt": "2023-04-21T12:40:19.925Z", "name":"BillyStoltenberg","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1164.jpg","id":"1" } , { "createdAt":"2022-11-05T09:01:38.207Z","name":"JodiKertzmann","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg","id":"2" } , { "createdAt":"2022-11-05T15:36:31.390Z","name":"AngelKuhlmanIV","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/75.jpg","id":"3" } , { "createdAt":"2022-11-04T19:23:42.344Z","name":"KellyNolan","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/787.jpg","id":"4" } , { "createdAt":"2022-11-05T06:33:55.777Z","name":"JeremySchneider","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1153.jpg","id":"5" } , { "createdAt":"2022-11-05T07:29:47.957Z","name":"EvanDuBuque","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/99.jpg","id":"6" } , { "createdAt":"2022-12-01T10:43:35.133Z","name":"DeloresDoyle","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/331.jpg","id":"7" } , { "createdAt":"2022-12-01T22:53:22.031Z","name":"Ms.JeremyBruen","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/383.jpg","id":"8" } , { "createdAt":"2022-12-02T00:59:05.444Z","name":"JoannRitchie","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/130.jpg","id":"9" } , { "createdAt":"2022-12-01T21:37:21.680Z","name":"Mrs.StevenCummerata","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/771.jpg","id":"10" } , { "createdAt":"2022-12-02T09:01:39.388Z","name":"HattiePfeffer","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1005.jpg","id":"11" } , { "createdAt":"2022-12-01T10:11:38.831Z","name":"EdithMacejkovic","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/882.jpg","id":"12" } , { "createdAt":"2022-12-02T08:34:12.087Z","name":"BettyBlock","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/216.jpg","id":"13" } , { "createdAt":"2022-12-02T06:56:11.901Z","name":"ArmandoBecker","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/822.jpg","id":"14" } , { "createdAt":"2022-12-01T17:22:36.489Z","name":"MargueriteWilliamson","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/205.jpg","id":"15" } , { "createdAt":"2022-12-02T02:59:48.845Z","name":"LorenePaucek","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/31.jpg","id":"16" } , { "createdAt":"2022-12-01T22:23:37.039Z","name":"TonyaHeidenreich","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/685.jpg","id":"17" } , { "createdAt":"2022-12-01T15:55:35.929Z","name":"MissAntoniaReynolds","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/208.jpg","id":"18" } , { "createdAt":"2022-12-01T14:15:27.068Z","name":"Ms.GlendaSchimmel","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/145.jpg","id":"19" } , { "createdAt":"2022-12-01T17:29:53.880Z","name":"JodySchadenDDS","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/35.jpg","id":"20" }]
```

Abaixo o melhor resultado encontrado para o fiber.

#### k6 microservice1 Go -> Go microservice2

```bash
CPU: 27,14
MEMORIA: 10,03
k6 run -d 90s -u 100 k6/microservice1-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/microservice1-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 2m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 1m30s (gracefulStop: 30s)


running (1m30.0s), 000/100 VUs, 4987865 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 19 GB   209 MB/s
     data_sent......................: 594 MB  6.6 MB/s
     http_req_blocked...............: avg=3.28µs  min=695ns    med=1.46µs max=78.41ms  p(90)=1.93µs  p(95)=2.34µs 
     http_req_connecting............: avg=154ns   min=0s       med=0s     max=53.43ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.69ms  min=79.96µs  med=1.34ms max=224.19ms p(90)=3.42ms  p(95)=4.36ms 
       { expected_response:true }...: avg=1.69ms  min=79.96µs  med=1.34ms max=224.19ms p(90)=3.42ms  p(95)=4.36ms 
     http_req_failed................: 0.00%   ✓ 0           ✗ 4987865
     http_req_receiving.............: avg=44.43µs min=7.42µs   med=18.2µs max=219.66ms p(90)=46.47µs p(95)=128.7µs
     http_req_sending...............: avg=15.44µs min=3.39µs   med=7.11µs max=218.59ms p(90)=9.95µs  p(95)=17.44µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s     max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.63ms  min=57.9µs   med=1.28ms max=223.7ms  p(90)=3.34ms  p(95)=4.27ms 
     http_reqs......................: 4987865 55419.62106/s
     iteration_duration.............: avg=1.78ms  min=115.08µs med=1.42ms max=224.23ms p(90)=3.52ms  p(95)=4.48ms 
     iterations.....................: 4987865 55419.62106/s
     vus............................: 100     min=100       max=100  
     vus_max........................: 100     min=100       max=100  

```

```bash
CPU: 24,61
MEMORIA: 10,02
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


running (15.0s), 000/100 VUs, 799267 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 3.0 GB 201 MB/s
     data_sent......................: 95 MB  6.3 MB/s
     http_req_blocked...............: avg=3.3µs   min=683ns    med=1.45µs  max=18.74ms p(90)=2.06µs  p(95)=2.51µs  
     http_req_connecting............: avg=177ns   min=0s       med=0s      max=18.71ms p(90)=0s      p(95)=0s      
     http_req_duration..............: avg=1.77ms  min=81.2µs   med=1.36ms  max=69.41ms p(90)=3.61ms  p(95)=4.66ms  
       { expected_response:true }...: avg=1.77ms  min=81.2µs   med=1.36ms  max=69.41ms p(90)=3.61ms  p(95)=4.66ms  
     http_req_failed................: 0.00%  ✓ 0            ✗ 799267
     http_req_receiving.............: avg=48.43µs min=7.72µs   med=17.97µs max=26.77ms p(90)=52.03µs p(95)=137.72µs
     http_req_sending...............: avg=16.97µs min=3.23µs   med=6.87µs  max=25.85ms p(90)=12.71µs p(95)=21.22µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=1.7ms   min=57.94µs  med=1.3ms   max=68.08ms p(90)=3.51ms  p(95)=4.54ms  
     http_reqs......................: 799267 53276.732292/s
     iteration_duration.............: avg=1.85ms  min=119.41µs med=1.44ms  max=69.57ms p(90)=3.72ms  p(95)=4.79ms  
     iterations.....................: 799267 53276.732292/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 

```

#### wrk go -> go

```bash
CPU: 38,91
MEMORIA: 9,8
wrk -t12 -c100 -d15s http://localhost:8080/v1/user
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.99ms    0.88ms  24.89ms   84.00%
    Req/Sec     9.18k     1.05k   19.62k    90.49%
  1651631 requests in 15.10s, 5.81GB read
Requests/sec: 109382.11
Transfer/sec:    393.68MB
```

#### Go -> Rust

```bash
CPU: 27,04
MEMORIA: 9,8
 k6 run -d 90s -u 100 k6/microservice1-get.js 

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: k6/microservice1-get.js
     output: -

  scenarios: (100.00%) 1 scenario, 100 max VUs, 2m0s max duration (incl. graceful stop):
           * default: 100 looping VUs for 1m30s (gracefulStop: 30s)


running (1m30.0s), 000/100 VUs, 5750021 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 22 GB   246 MB/s
     data_sent......................: 685 MB  7.6 MB/s
     http_req_blocked...............: avg=2.22µs  min=622ns    med=1.39µs  max=21.94ms p(90)=1.76µs  p(95)=2.09µs 
     http_req_connecting............: avg=51ns    min=0s       med=0s      max=19.11ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.49ms  min=88.71µs  med=1.3ms   max=70.13ms p(90)=2.6ms   p(95)=3.15ms 
       { expected_response:true }...: avg=1.49ms  min=88.71µs  med=1.3ms   max=70.13ms p(90)=2.6ms   p(95)=3.15ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 5752508
     http_req_receiving.............: avg=31.61µs min=7.55µs   med=17.46µs max=52.16ms p(90)=23.44µs p(95)=31.57µs
     http_req_sending...............: avg=10.89µs min=3.25µs   med=6.81µs  max=19.77ms p(90)=8.36µs  p(95)=11.96µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.45ms  min=64.81µs  med=1.26ms  max=68.49ms p(90)=2.55ms  p(95)=3.08ms 
     http_reqs......................: 5752508 63915.647841/s
     iteration_duration.............: avg=1.55ms  min=117.56µs med=1.35ms  max=71.16ms p(90)=2.67ms  p(95)=3.24ms 
     iterations.....................: 5752508 63915.647841/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100 

CPU: 25,01
MEMORIA: 9,06
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


running (15.0s), 000/100 VUs, 828346 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 3.2 GB 213 MB/s
     data_sent......................: 99 MB  6.6 MB/s
     http_req_blocked...............: avg=6.2µs   min=684ns    med=1.46µs  max=39.56ms p(90)=2.08µs  p(95)=2.57µs  
     http_req_connecting............: avg=3.04µs  min=0s       med=0s      max=38.7ms  p(90)=0s      p(95)=0s      
     http_req_duration..............: avg=1.7ms   min=94.27µs  med=1.37ms  max=57.11ms p(90)=3.11ms  p(95)=3.97ms  
       { expected_response:true }...: avg=1.7ms   min=94.27µs  med=1.37ms  max=57.11ms p(90)=3.11ms  p(95)=3.97ms  
     http_req_failed................: 0.00%  ✓ 0            ✗ 828346
     http_req_receiving.............: avg=46.01µs min=7.91µs   med=18.27µs max=28.6ms  p(90)=31.24µs p(95)=110.68µs
     http_req_sending...............: avg=15.44µs min=3.29µs   med=6.94µs  max=35.87ms p(90)=12.54µs p(95)=15.27µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s      
     http_req_waiting...............: avg=1.64ms  min=71.46µs  med=1.32ms  max=50.95ms p(90)=3.03ms  p(95)=3.85ms  
     http_reqs......................: 828346 55216.187271/s
     iteration_duration.............: avg=1.79ms  min=128.75µs med=1.44ms  max=61.23ms p(90)=3.22ms  p(95)=4.12ms  
     iterations.....................: 828346 55216.187271/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100
```

```bash
CPU: 49,46
MEMORIA: 9,8
wrk -t12 -c100 -d15s http://localhost:8080/v1/user
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.85ms    2.39ms  96.55ms   99.30%
    Req/Sec    11.56k     1.34k   29.84k    86.77%
  2078174 requests in 15.10s, 7.46GB read
Requests/sec: 137635.74
Transfer/sec:    505.74MB
```

```bash
CPU: 49,20
MEMORIA: 9,7
wrk -t12 -c100 -d90s http://localhost:8080/v1/user
Running 2m test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   742.42us  575.76us  55.27ms   89.78%
    Req/Sec    11.29k   844.82    23.51k    77.68%
  12137200 requests in 1.50m, 43.55GB read
Requests/sec: 134728.99
Transfer/sec:    495.06MB

```