## Docker
docker是一种容器技术，使用docker能够帮助企业快速水平扩展服务，从而达到弹性部署的能力，对微服务的部署带来了极大的方便。

容器除了运行其中应用外，基本不消耗额外的系统资源，使得应用的性能很高，同时系统的开销尽量小。
Docker 在如下几个方面具有较大的优势。
1、更快速的交付和部署
2、更高效的虚拟化
	Docker 容器的运行不需要额外的 hypervisor 支持，它是**内核级的虚拟化**，因此可以实现更高的性能和效率。
3、更轻松的迁移和扩展
4、更简单的管理

**Docker daemon**: 运行在宿主机上，Docker守护进程，用户通过Docker client(Docker命令)与Docker daemon交互
**Docker client**: Docker 命令行工具，是用户使用Docker的主要方式，Docker client与Docker daemon通信并将结果返回给用户，Docker client也可以通过socket或者RESTful api访问远程的Docker daemon

**Docker image**：镜像是只读的，镜像中包含有需要运行的文件。镜像用来创建container，一个镜像可以运行多个container；镜像可以通过Dockerfile创建，也可以从Docker hub/registry上下载。
**Docker container**：容器是Docker的运行组件，启动一个镜像就是一个容器，容器是一个隔离环境，多个容器之间不会相互影响，保证容器中的程序运行在一个相对安全的环境中。
**Docker hub/registry**: 共享和管理Docker镜像，用户可以上传或者下载上面的镜像，官方地址为https://registry.hub.docker.com/，也可以搭建自己私有的Docker registry。

阿里镜像加速器：https://e9ob1yf2.mirror.aliyuncs.com

**centos7 安装docker**
```bash
# 查看内核版本
uname -r #3.10.0-862.14.4.el7.x86_64 
# 卸载docker
sudo yum remove docker docker-common  docker-selinux  docker-engine docker-ce 
# 更新yum元缓存
sudo yum makecache fast
# 配置yum源
sudo yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
# 启动docker daemon
sudo systemctl start docker
# 检验安装结果
docker run hello-world
```
**自动安装**
```
curl -fsSL get.docker.com -o get-docker.sh
sudo sh get-docker.sh --mirror Aliyun
```

