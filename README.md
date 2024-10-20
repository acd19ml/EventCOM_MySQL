# EventCOM_MySQL
```
WSL登录MySQL
mysql -u root -p

WSL启动MySQL服务
sudo service mysql start

端口映射
sudo socat TCP-LISTEN:3307,fork TCP:127.0.0.1:3306

SHOW DATABASE;
USE EventCOM;
```