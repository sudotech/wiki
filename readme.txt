windows安装
使用 nssm.exe 进行安装

网站发布
1、拷贝 logs static tools httpserver.exe 到自定义文件夹下
2、根据自己的windows系统拷贝 tools中的nssm.exe 到自定义文件夹下
3、使用cmd 进行安装
4、命令：
nssm.exe install httpserver.exe //安装
nssm.exe start httpserver.exe	//启动
nssm.exe stop httpserver.exe	//停止
nssm.exe remove httpserver.exe	//移除