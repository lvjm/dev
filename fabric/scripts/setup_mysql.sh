docker pull mysql:5.7
docker run  -e MYSQL_ROOT_PASSWORD=Npassword123 -v "$PWD/mysql":/var/lib/mysql  --mount type=bind,source="$PWD/tfs_mysql",target=/docker-entrypoint-initdb.d  -p 3306:3306 -d mysql:5.7 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci

# Npassword123  is default password of root user. $PWD/mysql is for my location which store mysql configuration after docker starts.   $PWD/tfs_mysql folder is location of sql which need to create database tfs_api and tfs_admin, you can copy from the two codebases