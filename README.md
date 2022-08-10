# Container Runtime Compliance Checker (CRCC)

<p align="center">
<a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a>
<a href="https://www.docker.com/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/docker/docker-original-wordmark.svg" alt="docker" width="40" height="40"/> </a>
<a href="https://www.postgresql.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/postgresql/postgresql-original-wordmark.svg" alt="postgresql" width="40" height="40"/> </a>
<a href="https://nuxtjs.org/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/nuxtjs/nuxtjs-icon.svg" alt="nuxtjs" width="40" height="40"/> </a>
<a href="https://vuejs.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/vuejs/vuejs-original-wordmark.svg" alt="vuejs" width="40" height="40"/> </a>
<a href="https://www.chartjs.org" target="_blank" rel="noreferrer"> <img src="https://www.chartjs.org/media/logo-title.svg" alt="chartjs" width="40" height="40"/> </a>
<a href="https://git-scm.com/" target="_blank" rel="noreferrer"> <img src="https://www.vectorlogo.zone/logos/git-scm/git-scm-icon.svg" alt="git" width="40" height="40"/> </a>
<a href="https://www.linux.org/" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/linux/linux-original.svg" alt="linux" width="40" height="40"/> </a>
</p>

## Overview
Our project is security focused, specifically on the containers realm.

We provide security compliance for the containers that run on different operating systems. 

Our compliance is based on 4 main pillars:
1. File System Analysis
2. Process Analysis
3. Network Analysis
4. DNS Analysis

### App Architecture
<img src="./docs/app_architecture.png" width="500" height="400"/>

---
## Technologies
- [Go](https://go.dev/) for the Backend with [Standard Go Project Layout](https://github.com/golang-standards/project-layout) the project structure we follow
- NuxtJs framework with Aragon theme [live preview](https://nuxt-argon-dashboard-laravel.creative-tim.com/dashboard/).
- [Docker & Docker compose](https://www.docker.com/) for containers management and Docker APIs for interacting with containers
- [PostgreSQL](https://www.postgresql.org/) DBMS for saving our scan histories 


## How to run the project?
- start the application by running

```
sh start.sh
```
- browse the app at: http://localhost:3000
