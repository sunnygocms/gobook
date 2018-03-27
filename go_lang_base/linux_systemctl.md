#CentOS 7 上systemctl 的用法

我们对service和chkconfig两个命令都不陌生，systemctl 是管制服务的主要工具， 它整合了chkconfig 与 service功能于一体。

	systemctl is-enabled iptables.service
	systemctl is-enabled servicename.service #查询服务是否开机启动
	systemctl enable *.service #开机运行服务
	systemctl disable *.service #取消开机运行
	systemctl start *.service #启动服务
	systemctl stop *.service #停止服务
	systemctl restart *.service #重启服务
	systemctl reload *.service #重新加载服务配置文件
	systemctl status *.service #查询服务运行状态
	systemctl --failed #显示启动失败的服务

注：*代表某个服务的名字，如http的服务名为httpd

例如在CentOS 7 上安装http

	[root@CentOS7 ~]# yum -y install httpd

	启动服务（等同于service httpd start）
	systemctl start httpd.service
	
	停止服务（等同于service httpd stop）
	systemctl stop httpd.service

	重启服务（等同于service httpd restart）
	systemctl restart httpd.service

	查看服务是否运行（等同于service httpd status）
	systemctl status httpd.service

	开机自启动服务（等同于chkconfig httpd on）
	systemctl enable httpd.service

	开机时禁用服务（等同于chkconfig httpd on）
	systemctl disable httpd.service
    
    
## 链接
- [目录](https://github.com/sunnygocms/gobook/blob/master/menu.md)    