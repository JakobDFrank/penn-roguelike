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

2. Open a terminal and navigate to the root directory.

3. Run `docker-compose up --build` to start the application.

4. In another terminal, run the following commands to manage the program.



### Submitting a Level

This endpoint expects a 2d array of integers from zero to four, inclusive.

- `0` for an open tile
- `1` for a wall
- `2` for a pit (player takes one damage)
- `3` for arrows (player takes two damage)
- `4` for the player

```shell
curl -X POST http://localhost:8080/level/submit \
-H "Content-Type: application/json" \
-d '[[0,0,0,0,2],[0,4,0,0,2],[0,1,2,0,0],[0,1,1,3,0],[0,0,0,0,0]]'
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

```shell
curl -X POST http://localhost:8080/player/move \
     -H "Content-Type: application/json" \
     -d '{"id":1,"direction":0}'
```

## Directory Structure

https://github.com/golang-standards/project-layout
