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


Usamos o k6 para fazer nossos testes de stress.

```bash
$ cd k6
$ k6 run -d 90s -u 100 script-get.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io
```

O nosso objetivo é só termos uma noção quando o assunto é memória, cpu, quantidade de requisições 
suportadas e os tempos médios de respostas, nada muito complexo só para matar a curiosidade mesmo.
