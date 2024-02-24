# penn-roguelike

---

## Prerequisites

 <br>

### Fedora

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

## Directory Structure

https://github.com/golang-standards/project-layout
