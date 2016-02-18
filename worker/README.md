etcd:(获取任务)
    /alive/project_name/workerid(上报存活)
    1.主动上报worker的serverid，存储格式是：(/alive/project_name/workerid)
    /broker/project_name/workerid(上报worker key)
    1.主动上报worker的Serverid，存储格式是：(/broker/project_name/workerid)

redis:(worker产生的日志记录)

log:(worker产生日志文件记录)
    详细的worker产生的日志

