#!/usr/bin/env bash

#docker pull and save images to tar

docker_images='mysql:5.7.25 redis:4.0.14 daocloud.io/sakeven/2048:v1.0.3'
save_dir=./images.tar

echo "prepare pull images"
for IMAGE in $docker_images; do
    echo "PULLING  IMAGE: $IMAGE"
    docker pull $IMAGE
done
echo " pull images done"


echo "saving $docker_images images to $save_dir..."
    docker save -o $save_dir $docker_images
echo "saving images done..."