# Cenário Client Server

Este cenário é um client server, onde criamos a comunicação entre dois seviços ambos utilizando rEST  feito em diversas linguagens de programação.

A comunicação entre estes dois serviços conseguiremos observar o tempo gasto pelas requisições e diversas outras variáveis que o k6 irá disponibilizar.


##### Docker-Compose

Você irá conseguir rodar os serviços no docker-compose

```bash
$ docker-compose build --no-cache
```

```bash
$ docker-compose ps -a
        Name           Command   State                    Ports                  
---------------------------------------------------------------------------------
webclientserver_m1_1   /app      Up      0.0.0.0:8080->8080/tcp,:::8080->8080/tcp
webclientserver_m2_1   /app      Up      0.0.0.0:3000->3000/tcp,:::3000->3000/tcp

$ docker-compose up
Starting webclientserver_m2_1 ... done
Starting webclientserver_m1_1 ... done
Attaching to webclientserver_m2_1, webclientserver_m1_1
m1_1  | Run Server Quick:0.0.0.0:8080
m1_1  | 2023/04/21 15:20:09 Run Server Port 0.0.0.0:8080
m1_1  | 2023/04/21 15:20:09 [GET]  /v1/user
m2_1  | Run Server Quick:0.0.0.0:3000
m2_1  | 2023/04/21 15:20:08 Run Server port 0.0.0.0:3000
m2_1  | 2023/04/21 15:20:08 [GET]  /v1/avatar

```