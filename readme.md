## 下载kafak
* 下载后直接解压到即可用
> 在官网下载kafka时，要选择带有bin后缀的安装包下载，这样方便我们直接运行
> 例如我解压到`D:/WorkSpace/kafka`

## 配置zookeeper
* 打开config文件夹下的`zookeeper.properties`文件，修改默认的dir目录
> 例如我改成`dataDir=D:/WorkSpace/kafka/tmp/zookeeper`

## 启动zookeeper
* 使用【管理员身份】运行命令 `windows\zookeeper-server-start.bat [zookeeper.properties path]`进行启动
> 例如我启动命令 `D:\WorkSpace\kafka\bin>windows\zookeeper-server-start.bat ../config/zookeeper.properties`
* 注意事项
> 如果下载的安装包不是带有bin后缀的话，可能会报错`无法加载主类`
> 启动命令时，config对应的路径不要写错，不然zookeeper无法通过合适的路径找到配置文件
> 在window与linux中，zookeeper的启用命令不同，注意区分，我示例为window的

## 检查启动
* 如果是默认端口启动（port:2181）,启动成功后，会打印相关端口及启动信息

## 配置kafka
* 开config文件夹下的`server.properties`文件，修改默认的log.dirs目录
> 例如我改成`log.dirs=D:/WorkSpace/kafka/tmp/kafka-logs`

## 启动kafka
* 使用【管理员身份】运行命令 `windows\windows\kafka-server-start.bat [server.properties path]`进行启动
> 例如我的启动命令`D:\WorkSpace\kafka\bin>windows\kafka-server-start.bat ../config/server.properties`

## 启动kafka消费者
* 使用【管理员身份】运行命令 `kafka-server-start.bat --bootstrap-server=[ip:port] --topic=[name] --from-beginning`进行启动
> 例如我的启动命令`D:\WorkSpace\kafka\bin>windows\kafka-console-consumer.bat --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning`