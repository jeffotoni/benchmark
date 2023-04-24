#rust hyper


Instalei alguns pacotes para visualizar o flamegraph.

#### install pacotes UBUNTU para usar flamegraph para Rust
```bash
$ sudo apt-get install -y linux-tools-common linux-tools-generic linux-tools-$(uname -r)
```

```bash
$ cargo install inferno
```

```bash
$ cargo build --release
```

Coletando dados com perf da amostragem da CPU:
```bash
$ sudo perf record -F 99 -g --call-graph dwarf target/release/hyper_server
$ sudo perf script > perf.script
```

Converta em um Flame Graph:
```bash
$ cat perf.script | inferno-collapse-perf > stacks.folded
$ cat stacks.folded | inferno-flamegraph > flamegraph.svg
```

Abaixo o melhor resultado encontrado para o Rust usando framework Hyper.

#### Criar imagem do Serviço 
Foi criado o Dockerfile para que possamos subir 

```bash
$ docker build -f Dockerfile --no-cache -t jeffotoni/m1rusthyper .
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
$ docker run --name rust.server.hype --rm -p 3000:3000 jeffotoni/m1rusthyper
Listening on http://0.0.0.0:8080
```

#### cURL 

Irá apresentar uma lista de avatar com suas imagens e nomes.

```bash
$ curl -i -XGET http://localhost:3000/v1/user
HTTP/1.1 200 OK
content-type: application/json
engine: Rust/Hype
location: /v1/user
date: 2023-04-21T12:40:19.925350124+00:00
content-length: 3696

[{ "createdAt": "2023-04-21T12:40:19.925Z", "name":"BillyStoltenberg","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1164.jpg","id":"1" } , { "createdAt":"2022-11-05T09:01:38.207Z","name":"JodiKertzmann","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/502.jpg","id":"2" } , { "createdAt":"2022-11-05T15:36:31.390Z","name":"AngelKuhlmanIV","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/75.jpg","id":"3" } , { "createdAt":"2022-11-04T19:23:42.344Z","name":"KellyNolan","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/787.jpg","id":"4" } , { "createdAt":"2022-11-05T06:33:55.777Z","name":"JeremySchneider","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1153.jpg","id":"5" } , { "createdAt":"2022-11-05T07:29:47.957Z","name":"EvanDuBuque","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/99.jpg","id":"6" } , { "createdAt":"2022-12-01T10:43:35.133Z","name":"DeloresDoyle","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/331.jpg","id":"7" } , { "createdAt":"2022-12-01T22:53:22.031Z","name":"Ms.JeremyBruen","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/383.jpg","id":"8" } , { "createdAt":"2022-12-02T00:59:05.444Z","name":"JoannRitchie","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/130.jpg","id":"9" } , { "createdAt":"2022-12-01T21:37:21.680Z","name":"Mrs.StevenCummerata","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/771.jpg","id":"10" } , { "createdAt":"2022-12-02T09:01:39.388Z","name":"HattiePfeffer","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/1005.jpg","id":"11" } , { "createdAt":"2022-12-01T10:11:38.831Z","name":"EdithMacejkovic","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/882.jpg","id":"12" } , { "createdAt":"2022-12-02T08:34:12.087Z","name":"BettyBlock","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/216.jpg","id":"13" } , { "createdAt":"2022-12-02T06:56:11.901Z","name":"ArmandoBecker","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/822.jpg","id":"14" } , { "createdAt":"2022-12-01T17:22:36.489Z","name":"MargueriteWilliamson","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/205.jpg","id":"15" } , { "createdAt":"2022-12-02T02:59:48.845Z","name":"LorenePaucek","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/31.jpg","id":"16" } , { "createdAt":"2022-12-01T22:23:37.039Z","name":"TonyaHeidenreich","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/685.jpg","id":"17" } , { "createdAt":"2022-12-01T15:55:35.929Z","name":"MissAntoniaReynolds","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/208.jpg","id":"18" } , { "createdAt":"2022-12-01T14:15:27.068Z","name":"Ms.GlendaSchimmel","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/145.jpg","id":"19" } , { "createdAt":"2022-12-01T17:29:53.880Z","name":"JodySchadenDDS","avatar":"https://cloudflare-ipfs.com/ipfs/Qmd3W5DuhgHirLHGVixi6V76LhCkZUz6pnFt5AJBiyvHye/avatar/35.jpg","id":"20" }]
```

### Rust -> Go

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
CPU: 31,05
MEMORIA: 39,09
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


running (15.0s), 000/100 VUs, 715312 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 2.7 GB 180 MB/s
     data_sent......................: 85 MB  5.7 MB/s
     http_req_blocked...............: avg=2.57µs  min=687ns    med=1.46µs  max=18.81ms p(90)=1.98µs  p(95)=2.35µs 
     http_req_connecting............: avg=295ns   min=0s       med=0s      max=18.77ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=2.02ms  min=128.22µs med=1.8ms   max=59.93ms p(90)=3.34ms  p(95)=4.01ms 
       { expected_response:true }...: avg=2.02ms  min=128.22µs med=1.8ms   max=59.93ms p(90)=3.34ms  p(95)=4.01ms 
     http_req_failed................: 0.00%  ✓ 0            ✗ 715312
     http_req_receiving.............: avg=34.95µs min=7.34µs   med=18.21µs max=20.18ms p(90)=26.62µs p(95)=42.37µs
     http_req_sending...............: avg=12.38µs min=3.39µs   med=6.95µs  max=19.54ms p(90)=10.18µs p(95)=13.83µs
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.97ms  min=100.07µs med=1.77ms  max=44.05ms p(90)=3.28ms  p(95)=3.92ms 
     http_reqs......................: 715312 47680.117272/s
     iteration_duration.............: avg=2.08ms  min=168.05µs med=1.86ms  max=61.7ms  p(90)=3.42ms  p(95)=4.1ms  
     iterations.....................: 715312 47680.117272/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 

```

```bash
CPU: 49,09
MEMORIA: 40,02
wrk -t12 -c100 -d15s http://localhost:8080/v1/user
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.27ms    0.86ms  20.17ms   78.85%
    Req/Sec     6.64k     0.93k   29.52k    92.36%
  1193850 requests in 15.10s, 4.19GB read
Requests/sec:  79068.77
Transfer/sec:    284.43MB

```

#### Rust -> Rust
```bash
CPU: 30,20
MEMORIA: 37,09
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


running (15.0s), 000/100 VUs, 834807 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  15s

     data_received..................: 3.2 GB 212 MB/s
     data_sent......................: 99 MB  6.6 MB/s
     http_req_blocked...............: avg=2.38µs  min=663ns    med=1.42µs  max=34.56ms p(90)=1.86µs  p(95)=2.19µs 
     http_req_connecting............: avg=520ns   min=0s       med=0s      max=34.51ms p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.73ms  min=121.33µs med=1.56ms  max=85.91ms p(90)=2.66ms  p(95)=3.21ms 
       { expected_response:true }...: avg=1.73ms  min=121.33µs med=1.56ms  max=85.91ms p(90)=2.66ms  p(95)=3.21ms 
     http_req_failed................: 0.00%  ✓ 0            ✗ 834807
     http_req_receiving.............: avg=27.38µs min=7.17µs   med=16.68µs max=12.37ms p(90)=23.12µs p(95)=27.06µs
     http_req_sending...............: avg=10.59µs min=3.32µs   med=6.73µs  max=33.72ms p(90)=8.93µs  p(95)=12.7µs 
     http_req_tls_handshaking.......: avg=0s      min=0s       med=0s      max=0s      p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.7ms   min=93.59µs  med=1.53ms  max=85.86ms p(90)=2.61ms  p(95)=3.14ms 
     http_reqs......................: 834807 55647.592415/s
     iteration_duration.............: avg=1.78ms  min=159.3µs  med=1.6ms   max=86.23ms p(90)=2.72ms  p(95)=3.29ms 
     iterations.....................: 834807 55647.592415/s
     vus............................: 100    min=100        max=100 
     vus_max........................: 100    min=100        max=100 


CPU: 34,79
MEMORIA: 42,04
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


running (1m30.0s), 000/100 VUs, 4923007 complete and 0 interrupted iterations
default ✓ [======================================] 100 VUs  1m30s

     data_received..................: 19 GB   211 MB/s
     data_sent......................: 586 MB  6.5 MB/s
     http_req_blocked...............: avg=1.96µs min=607ns    med=1.39µs  max=135.58ms p(90)=1.73µs  p(95)=2.03µs 
     http_req_connecting............: avg=81ns   min=0s       med=0s      max=34.71ms  p(90)=0s      p(95)=0s     
     http_req_duration..............: avg=1.76ms min=102.17µs med=1.62ms  max=164.37ms p(90)=2.66ms  p(95)=3.11ms 
       { expected_response:true }...: avg=1.76ms min=102.17µs med=1.62ms  max=164.37ms p(90)=2.66ms  p(95)=3.11ms 
     http_req_failed................: 0.00%   ✓ 0            ✗ 4923007
     http_req_receiving.............: avg=27.9µs min=7.71µs   med=17.76µs max=150.89ms p(90)=22.36µs p(95)=25.53µs
     http_req_sending...............: avg=9.33µs min=2.94µs   med=6.88µs  max=135.04ms p(90)=8.19µs  p(95)=9.31µs 
     http_req_tls_handshaking.......: avg=0s     min=0s       med=0s      max=0s       p(90)=0s      p(95)=0s     
     http_req_waiting...............: avg=1.73ms min=87.22µs  med=1.59ms  max=164.15ms p(90)=2.62ms  p(95)=3.06ms 
     http_reqs......................: 4923007 54696.415882/s
     iteration_duration.............: avg=1.82ms min=125.9µs  med=1.66ms  max=164.42ms p(90)=2.72ms  p(95)=3.18ms 
     iterations.....................: 4923007 54696.415882/s
     vus............................: 100     min=100        max=100  
     vus_max........................: 100     min=100        max=100 

```

```bash
CPU: 58,83
MEMORIA: 23,09
wrk -t12 -c100 -d15s http://localhost:8080/v1/user
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.98ms  742.81us  49.06ms   95.65%
    Req/Sec     8.52k     0.91k   18.04k    86.95%
  1532236 requests in 15.10s, 5.43GB read
Requests/sec: 101480.67
Transfer/sec:    368.34MB

```

```bash
CPU: 59,31
MEMORIA: 24,09
wrk -t12 -c100 -d90s http://localhost:8080/v1/user
Running 2m test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     0.99ms  402.94us  17.45ms   80.36%
    Req/Sec     8.16k   595.25    15.95k    82.96%
  8770850 requests in 1.50m, 31.09GB read
Requests/sec:  97365.41
Transfer/sec:    353.41MB

```