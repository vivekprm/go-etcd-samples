# Introduction
etcd is a distributed reliable key-value store for the most critical data of a distributed system, with a focus on being:
- Simple: well-defined, user-facing API (gRPC)
- Secure: automatic TLS with optional client cert authentication
- Fast: benchmarked 10,000 writes/sec
- Reliable: properly distributed using Raft

etcd is written in Go and uses the [Raft](https://raft.github.io/) consensus algorithm to manage a highly-available replicated log.

etcd is used [in production by many companies](https://github.com/etcd-io/etcd/blob/main/ADOPTERS.md), and the development team stands behind it in critical deployment scenarios, where etcd is frequently teamed with applications such as [Kubernetes](http://kubernetes.io/), [locksmith](https://github.com/coreos/locksmith), [vulcand](https://github.com/vulcand/vulcand), [Doorman](https://github.com/youtube/doorman), and many others. Reliability is further ensured by rigorous [robustness testing](https://github.com/etcd-io/etcd/tree/main/tests/robustness).

See [etcdctl](https://github.com/etcd-io/etcd/tree/main/etcdctl) for a simple command line client.

# Run Single Node ETCD
```sh
export NODE1=192.168.1.21

docker volume create --name etcd-data
export DATA_DIR="etcd-data"

REGISTRY=quay.io/coreos/etcd

docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --volume=${DATA_DIR}:/etcd-data \
  --name etcd ${REGISTRY}:latest \
  /usr/local/bin/etcd \
  --data-dir=/etcd-data --name node1 \
  --initial-advertise-peer-urls http://${NODE1}:2380 --listen-peer-urls http://0.0.0.0:2380 \
  --advertise-client-urls http://${NODE1}:2379 --listen-client-urls http://0.0.0.0:2379 \
  --initial-cluster node1=http://${NODE1}:2380
```

More detailed instructions can be find [here](https://etcd.io/docs/v3.5/op-guide/container/).

listen-client-urls and listen-peer-urls specify the local addresses etcd server binds to for accepting incoming connections. To listen on a port for all interfaces, specify 0.0.0.0 as the listen IP address.

advertise-client-urls and initial-advertise-peer-urls specify the addresses etcd clients or other etcd members should use to contact the etcd server. The advertise addresses must be reachable from the remote machines. Do not advertise addresses like localhost or 0.0.0.0 for a production setup since these addresses are unreachable from remote machines.

# Using ETCDCTL to communicate
```sh
etcdctl --endpoints=http://${NODE1}:2379 member list
etcdctl --endpoints=http://0.0.0.0:2379 put greeting "Hello, etcd"
etcdctl --endpoints=http://0.0.0.0:2379 get greeting
```
