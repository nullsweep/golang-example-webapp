#### **Docker Images**

**List Images:**
```bash
docker images
```

**Pull an Image:**
```bash
docker pull <image>:<tag>
```

**Build an Image:**
```bash
docker build -t <image>:<tag> <path_to_dockerfile>
```

**Remove an Image:**
```bash
docker rmi <image>:<tag>
```

**Tag an Image:**
```bash
docker tag <image>:<tag> <repository>/<image>:<tag>
```

**Inspect an Image:**
```bash
docker inspect <image>:<tag>
```

**Prune Unused Images:**
```bash
docker image prune
```

---

#### **Docker Containers**

**List Running Containers:**
```bash
docker ps
```

**List All Containers:**
```bash
docker ps -a
```

**Run a Container:**
```bash
docker run -d --name <container_name> <image>:<tag>
```

**Run a Container with Port Mapping:**
```bash
docker run -d -p <host_port>:<container_port> <image>:<tag>
```

**Run a Container with Volume:**
```bash
docker run -d -v <host_path>:<container_path> <image>:<tag>
```

**Run a Container with Environment Variables:**
```bash
docker run -d -e <ENV_VAR>=<value> <image>:<tag>
```

**Stop a Container:**
```bash
docker stop <container_id>
```

**Start a Container:**
```bash
docker start <container_id>
```

**Remove a Container:**
```bash
docker rm <container_id>
```

**Remove All Stopped Containers:**
```bash
docker container prune
```

**Restart a Container:**
```bash
docker restart <container_id>
```

**Attach to a Running Container:**
```bash
docker attach <container_id>
```

**View Container Logs:**
```bash
docker logs <container_id>
```

**Follow Container Logs:**
```bash
docker logs -f <container_id>
```

**Inspect a Container:**
```bash
docker inspect <container_id>
```

**Execute a Command in a Running Container:**
```bash
docker exec -it <container_id> <command>
```

---

#### **Docker Networks**

**List Networks:**
```bash
docker network ls
```

**Create a Network:**
```bash
docker network create <network_name>
```

**Remove a Network:**
```bash
docker network rm <network_name>
```

**Connect a Container to a Network:**
```bash
docker network connect <network_name> <container_id>
```

**Disconnect a Container from a Network:**
```bash
docker network disconnect <network_name> <container_id>
```

**Inspect a Network:**
```bash
docker network inspect <network_name>
```

---

#### **Docker Volumes**

**List Volumes:**
```bash
docker volume ls
```

**Create a Volume:**
```bash
docker volume create <volume_name>
```

**Remove a Volume:**
```bash
docker volume rm <volume_name>
```

**Inspect a Volume:**
```bash
docker volume inspect <volume_name>
```

**Prune Unused Volumes:**
```bash
docker volume prune
```

---

#### **Docker Compose**

**Start Services:**
```bash
docker-compose up
```

**Start Services in Detached Mode:**
```bash
docker-compose up -d
```

**Stop Services:**
```bash
docker-compose down
```

**Rebuild Services:**
```bash
docker-compose up --build
```

**List Services:**
```bash
docker-compose ps
```

**View Service Logs:**
```bash
docker-compose logs
```

**Follow Service Logs:**
```bash
docker-compose logs -f
```

**Execute Command in a Service Container:**
```bash
docker-compose exec <service_name> <command>
```

---

#### **Docker Swarm**

**Initialize Swarm:**
```bash
docker swarm init
```

**Join Swarm:**
```bash
docker swarm join --token <token> <manager_ip>:<port>
```

**List Nodes:**
```bash
docker node ls
```

**Inspect Node:**
```bash
docker node inspect <node_id>
```

**Remove Node:**
```bash
docker node rm <node_id>
```

**Create Service:**
```bash
docker service create --name <service_name> <image>:<tag>
```

**List Services:**
```bash
docker service ls
```

**Inspect Service:**
```bash
docker service inspect <service_name>
```

**Scale Service:**
```bash
docker service scale <service_name>=<replica_count>
```

**Remove Service:**
```bash
docker service rm <service_name>
```

---

#### **Docker Security**

**Run Container with Read-Only Filesystem:**
```bash
docker run --read-only <image>:<tag>
```

**Limit Container Memory:**
```bash
docker run -m <memory_limit> <image>:<tag>
```

**Limit Container CPU:**
```bash
docker run --cpus=<cpu_limit> <image>:<tag>
```

**User Namespace Remapping:**
```bash
docker run --userns-remap=<user> <image>:<tag>
```

**Use Seccomp Profile:**
```bash
docker run --security-opt seccomp=<profile.json> <image>:<tag>
```

**Use AppArmor Profile:**
```bash
docker run --security-opt apparmor=<profile> <image>:<tag>
```

**Use SELinux Labels:**
```bash
docker run --security-opt label:type:<type_label> <image>:<tag>
```

---