# How to upgrade k8s cluster

## upgrade steps

upgrade first control plane

```bash
apt update
apt-cache madison kubeadm
```

 upgrade kubeadm
 clear the hold on kubeadm with *apt-mark* then upgrade again and place the hold on the version

```bash
  sudo apt-mark unhold kubeadm
  sudo apt-get install kubeadm='1.28.6-1.1'
  sudo apt-mark hold kubeadm
```

  verify the kubeadm version

```bash
  sudo kubeadm version
```

  check the version upgrade plan prior to upgrade

```bash
  sudo kubeadm upgrade plan
```

  check the output if the system is ready for upgrade, then apply the upgrade

```bash
  sudo kubeadm upgrade apply v1.28.6

```

## Upgrade kubectl and kubelet

 upgrading kubectl and  kubelet on control plane

```bash
  kubectl drain cka-control-1 --ignore-daemonsets
```

  Now upgrade kubelet and kubectl

```bash
  sudo apt-mark unhold kubelet kubectl
  sudo apt-get install kubelet='1.28.6-1.1' kubectl='1.28.6-1.1'
  sudo apt-mark hold kubelet kubectl
```

  restart system daemon and kubelet now

```bash
  sudo systemctl daemon-reload
  sudo systemctl restart kubelet
```

  Now is the time to check the cluster version.

```bash
  kubectl get nodes
```

In the results you should see the version of the node and the control plane. If everything went well, all your control plane is running Kubernetes `v1.28.6`

Now lets uncordon the control plane node and enable it to set new pods

```bash
    kubectl uncordon cka-control
    kubectl get nodes
```

## upgrading firstNode

### upgrading the remaining control plane nodes

```bash
sudo apt-mark unhold kubeadm
sudo apt-get install kubeadm='1.28.6-1.1'
sudo apt-mark hold kubeadm
```

You can use `kubeadm upgrade plan` to see what will be changed on your system. If everything looks good, proceed with:

`kubeadm upgrade plan` should show that `firstNode` does not need to be upgraded.
Let’s start by upgrading `secondNode`.

first drain the node except daemon sets

```bash
kubectl drain control-2 --ignore-daemonsets
```

### update kubelet and kubectl

```bash
sudo apt-mark unhold kubelet kubectl
sudo apt-get install kubelet='1.28.6-1.1' kubectl='1.28.6-1.1'
sudo apt-mark hold kubelet kubectl
```

Now restart the system daemon and kubelet and verify nodes.

```bash
sudo systemctl daemon-reload
sudo systemctl restart kubelet
kubectl get nodes
```

once the kubelet is restarted. Uncordon the control plane

```bash
kubectl uncordon control-2
```

## upgrade worker nodes

 Repeat same steps for all other worker nodes, one after another.

* Drain the node first except the daemonset
* update kubeadm
* update kubectl and kubelet
* restart kubelet
* Verify nodes again
* UnCordon the node

### Drain the node except the daemonset

```bash
kubectl drain worker-1 --ignore-daemonsets
```

### update  kubeadm

```bash
sudo apt-mark unhold  kubeadm
sudo apt-get install kubeadm='1.28.6-1.1'
sudo apt-mark hold kubeadm
```

### update kubelet and kubectl on worker node

```bash
sudo apt-mark unhold kubelet kubectl
sudo apt-get install kubelet='1.28.6-1.1' kubectl='1.28.6-1.1'
sudo apt-mark hold kubelet kubectl
```

### restart daemon and kubelet

```bash
sudo system daemon-reload
sudo systemctl restart kubelet
kubectl get nodes
```

### uncordon the worker node

```bash
kubectl uncordon worker-1
```

### verify the nodes

```bash
kubectl get nodes
```

You should see `worker-1` in the list of ready nodes

## summary

This is just a brief explanation
