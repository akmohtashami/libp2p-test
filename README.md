To compile the chat clients go to the `chat` directory and run `go build`.

Bootstrap node can be built in the same way by going to the `bootstrap` directory.

First run the bootstrap node. You should receive an output similar to the following:
```
This node:  QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9   [/ip4/127.0.0.1/tcp/3000 /ip4/192.168.0.217/tcp/3000]
```

Then use the following command to start one or more chat clients:
```
chat/chat -peer /ip4/127.0.0.1/tcp/3000/p2p/QmVbcMycaK8ni5CeiM7JRjBRAdmwky6dQ6KcoxLesZDPk9
```
