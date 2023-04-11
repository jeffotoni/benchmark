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

Poderemos propor diversos cenários e linguagens de programação e implementar podem ficar a vontade em enviar seu PR.
Tudo é para matar a curiosidade mesmo, nada muito complexo quando tratamos de desenvolvimento em backend para web.

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
