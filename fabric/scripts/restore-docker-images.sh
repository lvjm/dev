#!/usr/bin/env bash
#restore docker images from tar

save_dir=./images.tar

echo "loading images from $save_dir.."
docker load -i $save_dir
echo "load images done"