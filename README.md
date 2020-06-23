## gin框架的学习记录

参考自：

https://github.com/eddycjy/go-gin-example

https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md

https://eddycjy.com/tags/gin/

```shell script
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1 -d mysql
# 修改conf/app.ini中的database host之后，执行如下命令
docker build -t gin-blog-docker .
docker run --link mysql:mysql -p 8000:8000 gin-blog-docker
```