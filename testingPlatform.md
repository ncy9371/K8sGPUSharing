# The Testing Platform
If you are interested in the sharing GPUs between Pods, please contact ta.yeh@lsalab.cs.nthu.edu.tw.

## WARNING
It's dangerous to put your critical contents in the testing platform. Basically, having the right to manipulate Pods is to have the permission for accessing the host and everything.

## Accessing Kubernetes
A file located at ```~/.kube/config``` is configured to access Kubernetes.

## Possible Commands
Get MtgpuPods from all namespaces
```
kubectl get mtgpupods.lsalab.nthu --all-namespaces
```
Deploy a MtgpuPod, follow the template in section [Using Shareable GPU](README.md#using-shareable-gpu).
```
kubectl create -f mtgpupod.yaml
```
Verify the GPU UUID got by Pod
```
# query inside Pod
kubectl exec {POD_NAME} -- nvidia-smi -L
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

## Issues
Currently the GPU memory usage control may not work properly.
