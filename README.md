# learning-go

## Go Rest API

* Simple CRUD to MySQL
* JWT Cookie
* MUX Middleware 
* Store File to AWS S3

## Requirement
* Ubuntu 22.04
* MariaDB
* Docker
* Fully Registered Domain Name
* Fully Access to S3

## Getting Started
### Step 1 - Installing Docker
---------------------

First, update your existing list of packages:
   ```sh
   $ sudo apt update
   ```
Next, install a few prerequisite packages which let apt use packages over HTTPS:
   ```sh
   $ sudo apt install apt-transport-https ca-certificates curl software-properties-common
   ```
Then add the GPG key for the official Docker repository to your system:
   ```sh
   $ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
   ```
Add the Docker repository to APT sources:
   ```js
   $ echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
   ```
Update your existing list of packages again for the addition to be recognized:
   ```sh
   $ sudo apt update
   ```
Make sure you are about to install from the Docker repo instead of the default Ubuntu repo:
   ```sh
   $ apt-cache policy docker-ce
   ```
You’ll see output like this, although the version number for Docker may be different:
   ```sh
   docker-ce:
   Installed: (none)
   Candidate: 5:20.10.14~3-0~ubuntu-jammy
   Version table:
      5:20.10.14~3-0~ubuntu-jammy 500
         500 https://download.docker.com/linux/ubuntu jammy/stable amd64 Packages
      5:20.10.13~3-0~ubuntu-jammy 500
         500 https://download.docker.com/linux/ubuntu jammy/stable amd64 Packages
   ```
Finally, install Docker:
   ```sh
   $ sudo apt install docker-ce
   ```
Docker should now be installed, the daemon started, and the process enabled to start on boot. Check that it’s running:
   ```sh
   $ sudo systemctl status docker
   ```
The output should be similar to the following, showing that the service is active and running:
   ``` sh
   Output
   ● docker.service - Docker Application Container Engine
      Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
      Active: active (running) since Fri 2022-04-01 21:30:25 UTC; 22s ago
   TriggeredBy: ● docker.socket
         Docs: https://docs.docker.com
      Main PID: 7854 (dockerd)
         Tasks: 7
      Memory: 38.3M
         CPU: 340ms
      CGroup: /system.slice/docker.service
               └─7854 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
   ```

<br></br>
### Step 2 - Executing the Docker Command Without Sudo (Optional)
---------------------

If you want to avoid typing sudo whenever you run the docker command, add your username to the docker group:
   ```sh
   $ sudo usermod -aG docker ${USER}
   ```
To apply the new group membership, log out of the server and back in, or type the following:
   ```sh
   $ su - ${USER}
   ```
Confirm that your user is now added to the docker group by typing:
   ```sh
   $ groups
   ```
   ```sh
   Output
   sammy sudo docker
   ```
If you need to add a user to the docker group that you’re not logged in as, declare that username explicitly using:
   ```sh
   $ sudo usermod -aG docker username
   ```

<br></br>
### Step 3 - Setup Domain DNS
---------------------

A DNS “A” record with `your_domain` pointing to your server’s public IP address.

<br></br>
### Step 4 - Deploying nginx-proxy with Let’s Encrypt
---------------------

You’ll be storing the Docker Compose configuration for nginx-proxy in a file named nginx-proxy-compose.yaml. Create it by running:
   ```sh
   $ docker compose -f nginx-compose.yaml up -d
   ```
   ```sh
   Output
   [+] Running 21/21
   ⠿ letsencrypt-nginx-proxy-companion Pulled                            6.8s
      ⠿ df9b9388f04a Pull complete                                        3.1s
      ⠿ 6c6cfd4eaf5b Pull complete                                        3.9s
      ⠿ 870307501973 Pull complete                                        4.3s
      ⠿ e8ff3435d14f Pull complete                                        4.5s
      ⠿ 5b78ba945919 Pull complete                                        4.8s
      ⠿ 973b2ca26006 Pull complete                                        5.0s
   ⠿ nginx-proxy Pulled                                                  8.1s
      ⠿ 42c077c10790 Pull complete                                        3.9s
      ⠿ 62c70f376f6a Pull complete                                        5.5s
      ⠿ 915cc9bd79c2 Pull complete                                        5.6s
      ⠿ 75a963e94de0 Pull complete                                        5.7s
      ⠿ 7b1fab684d70 Pull complete                                        5.7s
      ⠿ db24d06d5af4 Pull complete                                        5.8s
      ⠿ e917373dbecf Pull complete                                        5.9s
      ⠿ 11e2be9775e9 Pull complete                                        5.9s
      ⠿ 9996fa75bc02 Pull complete                                        6.1s
      ⠿ d37674efdf77 Pull complete                                        6.3s
      ⠿ a45d84576e75 Pull complete                                        6.3s
      ⠿ a13c1f42faf7 Pull complete                                        6.4s
      ⠿ 4f4fb700ef54 Pull complete                                        6.5s
   [+] Running 3/3
   ⠿ Network go-docker_default                                Created    0.1s
   ⠿ Container go-docker-nginx-proxy-1                        Started    0.5s
   ⠿ Container go-docker-letsencrypt-nginx-proxy-companion-1  Started    0.8s
   ```

<br></br>
### Step 5 - Creating and Running the Docker Compose File
---------------------

You will store the Docker Compose configuration for the Go web app in a file named go-app-compose.yaml. Create it by running:
   ```sh
   $ nano go-app-compose.yaml
   ```
Edit `your_domain` in this file:
   ```sh
   version: '3'
   services:
   go-web-app:
      restart: always
      build:
         dockerfile: Dockerfile
         context: .
      environment:
         - VIRTUAL_HOST=your_domain
         - LETSENCRYPT_HOST=your_domain
   ```

<br></br>
### Step 6 - Running the Docker Compose File
---------------------

You will store the Docker Compose configuration for the Go web app in a file named go-app-compose.yaml. Create it by running:
   ```sh
   $ docker compose -f app-compose.yaml up -d
   ```
   ```sh
   Output
   Creating network "go-docker_default" with the default driver
   Building go-web-app
   Step 1/13 : FROM golang:alpine AS build
   ---> b97a72b8e97d
   ...
   Successfully built 71e4b1ef2e25
   Successfully tagged go-docker_go-web-app:latest
   ...
   [+] Running 1/1
   ⠿ Container go-docker-go-web-app-1  Started 
 ```

<br></br>
### Step 6 - Running the Docker Compose File
---------------------
You can now navigate to https://`your_domain`/ to access your homepage.