# Cleanup

## Location

Linux

    /var/lib/docker

Mac

    $HOME/Library/Containers/com.docker.docker/Data/com.docker.driver.amd64-linux/

## General

Find images and containers

    docker images
    docker images -f dangling=true
    docker ps -a
    docker system df

Deletes

    docker rmi image_name
    docker rm -v container_name     # -v deletes the unnamed volume.

Start
    docker start <docker_name> -ai

## Volumes

    docker volume ls -f dangling=true
    docker volume prune
