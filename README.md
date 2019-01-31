# kubectlsafe
Safe run operations in Kubernetes

Before every write (apply, create, etc...) operation kubectl will ask you about you choise

#### Installation
Download kubectl-safe to /usr/local/bin

```sh
echo "alias kubectl='kubectl safe'" >> $HOME/.bashrc
. $HOME/.basrc
```

That's all

#### Example

```sh
andrey@andreypc:$ kubectl delete pod nginx-549cfdbc34-jdos4
Current context is kube-prod. Show must go on? y
```
