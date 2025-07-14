podman build -t basic_reader .
podman run -it localhost/basic_reader 

# -e MYSQL_USER=$MYSQL_USER -e MYSQL_DATABASE=$MYSQL_DATABASE -e MYSQL_PASSWORD=$MYSQL_PASSWORD -e MYSQL_HOSTNAME=$MYSQL_HOSTNAME
