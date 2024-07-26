 ##vipbind的部署与介绍
 ###镜像制作
- 准备项目代码
  ```
  cd ~/vipbind/
  ```
- 构建镜像
  ```
  docker build -t vipbind:v1 .
  ```
- 保存镜像到本地
  ```
  docker save -0 vipbind.tar vipbind:v1
  ```
###配置介绍
   ```
    [k8s]
    ip       = 172.18.70.241
    port     = 6443

    [vip]
    vip      = 172.18.70.241
   ```
配置文件见configmap
ip及vip字段统一为云平台vip，port端口为k8s集群api端口

###部署vipbind
  ```
  kubectl apply -f vipbind.yaml
  ```

###功能及原理

- 保证底层服务kubevirt-controller-manager与云平台vip始终保持在同节点
- 节点vip因重启等原因发生vip迁移，触发vipbind修改集群node节点label，重新调度底层服务




  