# penn-roguelike

---

## Prerequisites

- Docker

- Docker Compose v3

<br>
<details>
<summary>Installing Docker and Docker Compose on Fedora </summary>

Install Docker:

```
sudo dnf -y install dnf-plugins-core
sudo dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo
sudo dnf install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

Install Docker Compose:

```console
sudo dnf -y install docker-compose
```

Start Docker:

```console
sudo systemctl start docker
```

</details>

## How to Use

1. Clone the repository.

2. Open a terminal and navigate to penn-roguelike/deploy. 
   
3. (Optional) To modify the server type, use docker-compose build --build-arg API=TYPE, replacing TYPE with http (default), grpc, or graphql. This changes the server launched.

4. Run `docker-compose up` to start the application.

5. In another terminal, run the following commands to manage the program.



### Submitting a Level

This endpoint expects a 2d array of integers from zero to four, inclusive.

- `0` for an open tile
- `1` for a wall
- `2` for a pit (player takes one damage)
- `3` for arrows (player takes two damage)
- `4` for the player

HTTP example:
```shell
curl -X POST http://localhost:8080/level/submit \
-H "Content-Type: application/json" \
-d '[[0,0,0,0,2],[0,0,4,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]'
```

Constraints:
- Map must be rectangular
- Map height and map width can not exceed 100 units
- Map spaces can only consist of zero to four, inclusive
- There is one, and only one, player

### Moving the Player

Successfully submitting a level returns a unique map `id`. 

It can move a player on a specific map in conjunction with `direction`.

The `direction` parameter is defined as follows:

- `0` for left
- `1` for right
- `2` for up
- `3` for down

HTTP example:
```shell
curl -X POST http://localhost:8080/player/move \
     -H "Content-Type: application/json" \
     -d '{"id":1,"direction":0}'
```

Alternatively, with HTTP, you may use "left", "right", "up", or "down".
```shell
curl -X POST http://localhost:8080/player/move \
     -H "Content-Type: application/json" \
     -d '{"id":1,"direction":"left"}'
```

## Metrics
You can view Prometheus metrics by visiting: <br>
http://localhost:2112/metrics

You can access Prometheus's built-in expression browser by visiting: <br>
http://localhost:9090/graph

## Grafana
To access Grafana dashboards, visit: http://localhost:3000/ and use the following credentials:
- Username: admin
- Password: grafana

## GraphQL
If running with GraphQL, visit http://localhost:9101/ to access GraphQL Playground.

## Directory Structure

https://github.com/golang-standards/project-layout
