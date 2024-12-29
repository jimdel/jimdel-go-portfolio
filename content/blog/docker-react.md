---
title: "Developing with Docker & React"
blurb: "A recipe for developing React applications with Docker & docker-compose"
date: "2022-09-30"
---

## Hello, hello, hello!

My goal with this article is to provide a recipe for `dockerizing` a React application. Along the way, we'll review important Docker concepts & tools that we can use to improve our productivity & save time. By the end of this article you will:

- Write a `Dockerfile` & `docker-compose` file
- Understand how "hot reloading" works & how to implement it
- Learn about `volumes`, `bind mounts`, and `.dockerignore` files

## What is Docker?

> Docker takes away repetitive, mundane configuration tasks and is used throughout the development lifecycle for fast, easy and portable application development...Docker’s comprehensive end to end platform includes UIs, CLIs, APIs and security that are engineered to work together across the entire application delivery lifecycle.

Docker attempts to solve the ~_it runs on your machine but not on mine_~ problem that almost all software engineers will inevitably encounter when they attempt to share their creations with the world. We don't need to understand the nitty-gritty details of how Docker works\* but there are several important concepts we do need to understand. The first is understanding the difference between **images** and **containers**.

\*[the nitty-gritty details](https://docs.docker.com/get-started/overview/)

## Images & Containers

> Docker images are the basis of containers. An Image is an ordered collection of root filesystem changes and the corresponding execution parameters for use within a container runtime. An image typically contains a union of layered filesystems stacked on top of each other. An image does not have state and it never changes.
>
> A container is a runtime instance of a docker image.
> A Docker container consists of:
>
> - A Docker image
> - An execution environment
> - A standard set of instructions
>
> The concept is borrowed from shipping containers, which define a standard to ship goods globally. Docker defines a standard to ship software.

An **image** is a read-only bundle of source code, libraries, tools, dependencies, assets, and whatever else an application needs to run. Images are defined in a `Dockerfile`.

A **container** is a running **image**. Often, an application will require many containers each depending on: specific ports, environmental variables, and ways to communicate with the host system. `docker-compose` allows us to manage all of this.

We will be writing both `Dockerfile` and `docker-compose.yml` files. The example source code used can be [found here](https://github.com/JamesDeLay/docker-react-dev). It's a generic React application generated using `create-react-app`

## Dockerizing React for Development

### `Dockerfile.dev`

```Dockerfile
FROM node:alpine3.15

WORKDIR /app

COPY package*.json ./
RUN npm install
COPY . .

CMD ["npm", "start"]
```

### `docker-compose.yml`

```yml
version: "3.8"
services:
  web:
    build:
      context: .
      dockerfile: Dockerfile.dev
    environment:
      - CHOKIDAR_USEPOLLING=true
    ports:
      - "3000:3000"
    volumes:
      - .:/app
      - /app/node_modules
```

### Overview

#### `Dockerfile`

Here's what is happening, line by line, in our `Dockerfile`:

1. Define a new image from `node-alpine`
2. Create a working directory called `app` to house our application
3. Copy only the `package.json && package-lock.json` files to the working directory inside the container
4. Install `node_modules`
5. Copy the entire application to the working directory inside the container
6. Execute the command `npm start`

#### `docker-compose`

Here's what is happening, line by line, in our docker-compose file:

1. Declare which `docker-compose` version to use
2. Define a service called `web`
3. Tell the build to use the `Dockerfile.dev` in the directory set by the `context`
   - In our case, the current directory: `.`
4. Set the `CHOKIDAR_USEPOLLING` environment variable to `true`
   - This is Node specific configuration
5. Bind port `3000` of the host machine to port `3000` of the running container
6. Bind mount the current `.` directory to the `/app` directory
7. Persist `node_modules` using a volume: `/app/node_modules`

### Explanation

##### Rapid Builds

Docker knows we want quick builds so it will cache each build step to improve the speed at which future containers spin up. We can use this knowledge in combination with Docker **volumes** and `.dockerignore` to make our re-build times lightening fast.

```Dockerfile
# Dockerfile

COPY package*.json .
RUN npm install
COPY . .
```

```ini
#.dockerignore

node_modules
```

```yaml
# docker-compose.yml

environment:
  - CHOKIDAR_USEPOLLING=true
volumes:
  - /app/node_modules
```

What we've done is:

- Separate the dependency installation from the copy application step in our `Dockerfile`
- Add the `node_module` entry into our `.dockerignore` file to avoid copying dependencies into the container
- Used a [volume](https://docs.docker.com/storage/volumes/) to persist the container's `node_modules` beyond and outside of the container's lifecycle and set a Node environmental variable `CHOKIDAR_USEPOLLING=true`
  - For the curious, [chokidar](https://github.com/paulmillr/chokidar) is a file watcher for Node and we want to tell it to use [polling](<https://en.wikipedia.org/wiki/Polling_(computer_science)>)

##### Hot Reloading

When the user hits "save" their changes should be reflected immediately in the running container. We can achieve this using a `bind` mount.

> Bind mounts have limited functionality compared to volumes. When you use a bind mount, a file or directory on the host machine is mounted into a container. The file or directory is referenced by its absolute path on the host machine. By contrast, when you use a volume, a new directory is created within Docker’s storage directory on the host machine, and Docker manages that directory’s contents.

Here is the `bind` mount in our `docker-compose.yml`:

```yaml
# docker-compose.yml

volume: -.:/app
```

Now we can code in our editor, smash the save button, and see the changes because our codebase is bound to the container. Docker knows to update our container when files on our host machine change because this binding.

And thats it folks, thanks for reading!

---

## Notes, Links, & Resources

- Quotes taken directly from Docker's Documentation Site
- `docker system prune --all --force` - cleanup Docker memory issues
- [Chokidar](https://github.com/paulmillr/chokidar)
- [Polling in Programming](<https://en.wikipedia.org/wiki/Polling_(computer_science)>)
- [Docker Documentation](https://docs.docker.com/)
- [Source code](https://github.com/JamesDeLay/docker-react-dev)
