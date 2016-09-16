FROM scratch

ADD etcd-metrics-proxy /
ENTRYPOINT [ "/etcd-metrics-proxy" ]
