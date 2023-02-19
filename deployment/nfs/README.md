# NFS

## Why nfs?

When your pods down... All datas will disappear from there.
Just because you store you all data in this instace rather than presistence volume.

So we need a `NFS`, a file system for distributed systems... with no more explaination.

- It can provide the `volume` for backend 
- Presistent storage
- Accesss from different `namespcae` or `pods


## nfs-client-provisioner

>nfs-client-provisioner 是一个Kubernetes的简易NFS的外部provisioner，本身不提供NFS，需要现有的NFS服务器提供存储
>-    PV以 ${namespace}-${pvcName}-${pvName}的命名格式提供（在NFS服务器上）
>-    PV回收的时候以 archieved-${namespace}-${pvcName}-${pvName} 的命名格式（在NFS服务器上）
