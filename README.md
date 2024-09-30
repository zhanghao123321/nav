```shell
# cp /nav/api/data/data.db 文件到/data/nav/目录下
mkdir -p /data/nav/
cp nav/api/data/data.db /data/nav/

# Docker运行：
docker run -d -p 8888:8888 -v /data/nav/:/app/data/ --env DATABASE_URI=sqlite:///app/data/data.db --name nav registry.cn-shanghai.aliyuncs.com/hooz/nav

# Docker持久化数据，下载三个文件，然后使用navicat打开sqlite数据库保存即可
cd /data/nav/ 
data.db、data.db-shm、data.db-wal

# web访问：
# 前台：
127.0.0.1:8888
# 后台：
127.0.0.1:8888/admin
```
