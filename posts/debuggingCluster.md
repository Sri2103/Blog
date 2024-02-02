# How to debug the crashed apiserver

To debug the crashed api server. first check the status of the containers using container runtime

```bash
crictl  ps -a
```

check if the kube-apiserver is running or not, you can see all pods running/exited.
If it's exited with a nonzero code then use `crictl logs` command to get more details about what went wrong.

Inspect the logs of the container fix.

Fix the manifest file  and save in `/etc/kubernetes/manifests/kube-apiserver.yaml`.

if there are any issues related to image pull policy, try changing ImagePullPolicy from Always to IfNotPresent or Never. If there are any errors in the yaml file, kubectl will report them when trying If there are any issues with the configuration then restarting the control plane will fix them. If there are any issues with the configuration, this will help you identify them. If the issue still persists, try restarting the control plane components one by one.
