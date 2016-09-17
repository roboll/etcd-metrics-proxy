# etcd-metrics-proxy [![CircleCI](https://circleci.com/gh/roboll/etcd-metrics-proxy.svg?style=svg)](https://circleci.com/gh/roboll/etcd-metrics-proxy)

[![Docker Repository on Quay](https://quay.io/repository/roboll/etcd-metrics-proxy/status "Docker Repository on Quay")](https://quay.io/repository/roboll/etcd-metrics-proxy)

Proxy metrics from secured etcd over http. This keeps credentials locally scoped to etcd, and exposes only the metrics path for scraping from prometheus without having to give it client certs to access etcd.

```
  -etcd-ca string
       	The CA file for etcd tls.
  -etcd-cert string
       	The cert file for etcd tls.
  -etcd-key string
       	The key file for etcd tls.
  -port int
       	Port to bind to. (default 2381)
  -upstream-host string
       	The upstream etcd host. (default "localhost")
  -upstream-port int
       	The upstream etcd port. (default 2379)
  -upstream-server-name string
       	The upstream tls server name. (default "localhost")
```
