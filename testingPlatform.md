# The Testing Platform
If you are interested in the sharing GPUs between Pods, please contact ta.yeh@lsalab.cs.nthu.edu.tw.

## WARNING
It's dangerous to put your critical contents in the testing platform. Basically, having the right to manipulate Pods is to have the permission for accessing the host and everything.

## Terms of Use
* Cannot launch any Pod with privileged permission. Also, cannot modify any setting of the host or any files not your own.
* Each user has a quota of 10G.

## Accessing Kubernetes
A file located at ```~/.kube/config``` is configured to access Kubernetes.

## Available Resources
* MtgpuPod (multi-tenant GPU Pod)
* TFJobs (distributed tensorflow model training with MtgpuPod)

### MtgpuPod Possible Commands
Get MtgpuPods from all namespaces
```
kubectl get mtgpupods.lsalab.nthu --all-namespaces
```
Deploy a MtgpuPod, follow the template in section [Using Shareable GPU](README.md#using-shareable-gpu). And there are some examples at ```~/examples/mtgpupod/m*.yaml```
```
kubectl create -f mtgpupod.yaml
```
Verify the GPU UUID got by Pod
```
# query inside Pod
kubectl -n {POD_NS} exec {POD_NAME} -- nvidia-smi -L
# query logs of Pod
kubectl -n {POD_NS} logs {POD_NAME}
# query from PodSpec
kubectl -n {POD_NS} get {POD_NAME} -o yaml | grep -A 1 NVIDIA_VISIBLE_DEVICES
```
Delete a MtgpuPod
```
# delete by ns/name
kubectl -n {POD_NS} delete mtgpupod {POD_NAME}
# delete by file
kubectl delete -f mtgpupod.yaml
```

### Deploy a Jupyter Notebook using MtgpuPod
An example yaml file located at ```~/examples/mtgpupod/jupyter.yaml```.
```
kubectl create -f ~/examples/mtgpupod/jupyter.yaml
```
Get the NodePort and access the notebook from ```140.114.78.243:32753``` in this example.
```
kubectl get svc jupyter1-svc
NAME           TYPE       CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
jupyter1-svc   NodePort   10.100.174.228   <none>        8888:32753/TCP   2m
```
Delete the Pod and the Service
```
kubectl delete -f ~/examples/mtgpupod/jupyter.yaml
```

### Deploy a TFJob
#### With dashboard
The dashboard located at ```140.114.78.243:8080/tfjobs/ui/```.
#### With commands
An example yaml file located at ```~/examples/tfjob/tfjob.yaml```.
```
kubectl create -f ~/examples/tfjob/tfjob.yaml
# list the running Pods
kubectl get pod
```
Delete the tfjob
```
kubectl delete -f ~/examples/tfjob/tfjob.yaml
```

## Issues
Currently the GPU memory usage control may not work properly.
