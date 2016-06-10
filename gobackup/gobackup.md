##Gobackup
####framework
```

```

####agent config
```
[global]
server_id = 39005900000000
game_id = 39
op_id = 590
redis_host = 127.0.0.1
redis_port = 6379
redis_queue = msgqueue
my_ip = 127.0.0.1
log = /usr/local/agent/log/backup.log
error_log = /usr/local/agent/log/backup_error.log

[mysql3308]
backup_type = mysql
instance = 3308
rsync_model = backup/database_3308
back_dir = /data/backup/database_3308/
back_log = /var/log/mysql_3308.log
script = /bin/bash /usr/local/agent/script/mysql_backup.sh /etc/my.cnf

[redis6381]
backup_type = redis
instance = 6381
rsync_model = backup/redisbase_6381
back_dir = /data/backup/redisbase_6381/
back_log = /var/log/redis_6381.log
script = /bin/bash /usr/local/agent/script/redis_backup.sh /etc/redis.conf
```

####Msg
```
{
    "server_id":39005900000000,
    "game_id":39,
    "op_id":590,
    "server_ip":"127.0.0.1",
    "backup_log":"ok",
    "success":true,
    "backup_type":"mysql",
    "last_all_filename":"127.0.0.1_2016-05-17_01-01_3308.tar.gz",
    "last_all_time":1463472061,
    "filename":"127.0.0.1_2016-05-17_02-01_3308-increase.tar.gz",
    "instance":3308,
    "rsync_model":"backup/database_3308",
    "file_size":477384,
    "create_time":1463496138
}
```

####server
```
[global]
rsync_bwlimit = 40960
server_root_path = /backup/
redis_host = 127.0.0.1
redis_port = 6381
redis_queue = msgqueue
log = /usr/local/agent/log/back.log
error_log = /usr/local/agent/log/back_error.log
my_ip=127.0.0.1
retry = 3
```
