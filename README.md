# kubectlsafe
Plugin for `kubectl` for safe write operations in Kubernetes

Before every write (apply, create, etc...) operation, `kubectl` will ask you about your choice

#### Installation
It's simple as 1, 2, 3:

Download kubectl-safe to `/usr/local/bin`, add eXecutable bit and add to your aliases

```sh
sudo wget https://raw.githubusercontent.com/esin/kubectlsafe/master/kubectl-safe -O /usr/local/bin/kubectl-safe
sudo chmod +x /usr/local/bin/kubectl-safe
echo "alias kubectl='kubectl safe'" >> $HOME/.bashrc
. $HOME/.basrc
```

If you don't have root access, you can do like this:
```sh
mkdir $HOME/bin
wget https://raw.githubusercontent.com/esin/kubectlsafe/master/kubectl-safe -O $HOME/bin/kubectl-safe
chmod +x $HOME/bin/kubectl-safe
echo "alias kubectl='kubectl safe'" >> $HOME/.bashrc
. $HOME/.basrc
```

That's all

#### Example

```sh
andrey@andreypc:~$ kubectl delete pod nginx-549cfdbc34-jdos4
Current context is kube-prod. Show must go on? y #Only after that operation will begin
pod "nginx-549cfdbc34-jdos4" deleted
```


#### Contributions

Feel free to ask questions and apply new ideas

Happy kubectling! :)
