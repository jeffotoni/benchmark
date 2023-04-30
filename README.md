# Benchmark

O objetivo deste repo é criar cenários para observar o desempenho de linguagens utilizadas na web 
para desenvolvimento de backend linguagens como:
 - Go
 - C
 - C++
 - Java
 - Python
 - Ruby
 - PHP
 - C#
 - Rust
 - Javascript Nodejs
 - Javascript Deno
 - Javascript Bun
 - Dart
 - Lua
 - Elixir

Poderemos propor diversos cenários e linguagens de programação e implementar podem ficar a vontade em enviar seu PR.
Tudo é para matar a curiosidade mesmo, nada muito complexo quando tratamos de desenvolvimento em backend para web.

A máquina que está sendo usada para fazer os testes é:

| Coluna 1 | Coluna 2 |
|----------|----------|
| Vendor | Genuine Intel |
| Model | Intel(R) Core(TM) i7-10750H CPU @ 2.60GHz |
| Cache | 12288kb |
| siblings/núcleos | 12 |
| bogomips | 5199.98 |
| address sizes | 39 bits physical, 48 bits virtual |


Usamos o k6 e wrk para fazer nossos testes de stress.

```bash
$ cd k6
$ k6 run -d 90s -u 100 script-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io
```

```bash
$ wrk -t12 -c100 -d15s http://localhost:8080/v1/user
Running 15s test @ http://localhost:8080/v1/user
  12 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.09ms    1.75ms  59.07ms   98.98%
    Req/Sec     8.06k     1.07k   26.57k    86.95%
  1449705 requests in 15.10s, 5.20GB read
Requests/sec:  96009.25
Transfer/sec:    352.69MB

```

O nosso objetivo é só termos uma noção quando o assunto é memória, cpu, quantidade de requisições 
suportadas e os tempos médios de respostas, nada muito complexo só para matar a curiosidade mesmo.
