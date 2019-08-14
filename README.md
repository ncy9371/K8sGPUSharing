# K8sGPUSharing
Make GPU shareable in Kubernetes

## Limitation
Only support Nvidia GPU device plugin with nvidia-docker2 in K8s.  
Not compatible with docker (version>=19) using newer GPU resource API.

## Prerequisite
* A K8s cluster with Nvidia GPU device plugin.
* kubectl with admin permissions.

## Installation
```
git clone https://github.com/ncy9371/K8sGPUSharing.git
kubectl create -f crd.yaml -f controller.yaml -f daemonset.yaml
```

## Using Shareable GPU
In order to add extra information (some environment variables, volume mounts needed by shareable GPU) to Pods created by user, we create a new CustomResourceDefinition (CRD) named MtgpuPod (Multi-tenant GPU Pod) as the basic execution unit which is originally represented by Pod.

An example for MtgpuPod spec:
```
apiVersion: lsalab.nthu/v1
kind: MtgpuPod
metadata:
  name: pod1
  annotations:
    "lsalab.nthu/gpu_request": "0.5" # GPU request, same as Pod cpu request
    "lsalab.nthu/gpu_limit": "1.0" # GPU limit, same as Pod cpu limit
    "lsalab.nthu/gpu_mem": "1073741824" # GPU memory use 1Gi, in bytes
    "lsalab.nthu/GPUID": "abc"
spec:
  containers:
  nodeName: node1
  - name: sleep
    image: nvidia/cuda:9.0-base
    command: ["sh", "-c"]
    args:
    - 'nvidia-smi -L'
    resources:
      requests:
        cpu: "1"
        memory: "500Mi"
      limits:
        cpu: "1"
        memory: "500Mi"
```

## Uninstallation
```
kubectl delete -f crd.yaml -f controller.yaml -f daemonset.yaml
```
