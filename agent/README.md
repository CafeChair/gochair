agent
===
agent用于执行任务job

####原理
- agent向redis注册自己的UUID
- 监控自己的UUID并获取任务
- 执行任务并将日志记录到本地文件和存储到redis中
- 删除自己UUID的任务内容
- 循环监听

####架构