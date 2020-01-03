# K8S 

## Content

- Install
- Command
- Example
- Q&A

### Install

### Command

0) Run KVM Management

`$ virt-manager`

1) Run a single cluster

`$ minikube start --vm-driver=kvm2 --registry-mirror=https://registry.docker-cn.com --image-mirror-country=cn --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers`

2) Run a pod

`$ kubectl run test-nginx --image=nginx --port=80`


3) [Forward a local port to a port on the pod](https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/#forward-a-local-port-to-a-port-on-the-pod)

`$ kubectl port-forward nginx-deployment-9f46bb5-4nqtf 8080:80`


### Example


### Service


### Dasbhoard UI

### SSH the minikube server

use the `$ kubectl cluster-info` to get the cluster ip, the output like blow:

```
Kubernetes master is running at https://192.168.39.128:8443
KubeDNS is running at https://192.168.39.128:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
```

And the `192.168.39.128` is the cluster server ip

`$ ssh docker@[ip]`

The default password: `tcuser`

Run the command: `$ kubectl get pods -o wide` 

We can use `curl 172.17.0.7` to view the api.

About the Service Forwar the port : https://kubernetes.io/zh/docs/concepts/services-networking/connect-applications-service/


### Q&A

1) How to install k8s clusters in physical machine ?

2) If we can use kubectl to connect the cluster ?

3) Can we use the webui to use the kubectl ?# learn_k8s
