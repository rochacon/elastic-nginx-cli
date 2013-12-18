Elastic NGINX cli
=================

`elastic-nginx-cli` is a simple cli for [Elastic NGINX](https://github.com/rochacon/elastic-nginx).

This allows you to add or remove a EC2 instance from a Elastic NGINX managed upstream.


Setup
-----

```
go get github.com/rochacon/elastic-nginx-cli
```


Usage
-----

```bash
% elastic-nginx-cli -h
Usage: elastic-nginx-cli [InstanceId]...
  -asg-arn="": Auto Scaling Group ARN
  -event="launch": Auto Scaling Event (launch/terminate)
  -host="127.0.0.1:5000": Elastic NGINX host (may include port, e.g. 127.0.0.1:5000)
  -topic-arn="": Topic ARN
```
