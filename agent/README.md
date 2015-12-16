agent
===
agent用于执行任务job

####原理
- agent向redis注册自己的UUID
- 监听自己的UUID并获取任务
- 执行任务并将日志记录到本地文件和存储到redis中
- 删除自己UUID的任务内容
- 循环监听

####架构

                dashboard
                    |
                    |
               redis cluster
                    |
                    |
            -----------------
            |       |       |
            |       |       |
          agent   agent   agent

####配置文件

    {
        "uuid":"1001001001",
        "tags":"mobile_project",
        "redis":{
            "addr":"127.0.0.1",
            "port":6379
        },
        "task":{
            "path":"/usr/local/agent/task",
            "timeout":1000
        },
        "script":{
            "path":"/usr/local/agent/plugin",
            "timeout":1000
        },
        "log"{
            "path":"/usr/local/agent/log",
            "file":"agent.log"
        }
    }

####API(agent)

    GET /v1/agent/version
    {
        "version":"1.0"
    }

    GET /v1/agent/health
    {
        "health":"healthy|death"
    }

    GET /v1/agent/auuid/log
    {
        "uuid":"1001001001",
        "auuid":"log info"
    }

####API(dashboard)

    POST /v1/producer/uuid/task
    POST {
        "uuid":"1001001001",
        "taskname":"like linux command or command file"
    }
    return {
        "auuid":"1001001001taskname"
    }

    POST /v1/producer/project/task
    POST {
        "uuid":"1001001001",
        "taskname":"like linux command or command file"
    }
    return {
        "auuid":"1001001001taskname"
    }

    GET /v1/producer/project/auuid/log
    {
        "uuid":"1001001001",
        "taskname":"producer set taskname",
        "log":"agent run task loginfo"
    }

####redis struct(procuder)

    smembers "project_name"
    hset "1001001001" "taskname" "linux command or command file"
    hget "1001001001taskname" "uuid"
    hget "1001001001taskname" "log"
####redis struct(agent)

    sadd "mobile_project" "1001001001"
    向redis注册自己的UUID

    hget "1001001001" "taskname" "linux command or command file"
    监听自己UUID的任务内容并执行

    hset "1001001001taskname" "uuid" "1001001001"
    hset "1001001001taskname" "log" "exec task and log info"
    执行任务内容并将日志保存到redis中

    del "1001001001" 
    不管任务执行结果删除自己的UUID

