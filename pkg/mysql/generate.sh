#!/bin/bash

mysql_username=root
mysql_password=gloria

docker exec mysql mysql -u$mysql_user -p; << $mysql_password