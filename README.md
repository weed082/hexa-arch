Hexagonal architecture microservice

1. Set Docker Container
  (1) install docker setup
    - install docker
    - install docker compose
    - manage docker as non-root user (https://docs.docker.com/engine/install/linux-postinstall/)

  (2) start container
    - development
      config : docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file .env.dev config
      build  : docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file .env.dev build
      start  : docker-compose -f docker-compose.yml -f docker-compose.dev.yml --env-file .env.dev up
      
