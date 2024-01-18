# Introduction
etcd is a distributed reliable key-value store for the most critical data of a distributed system, with a focus on being:
- Simple: well-defined, user-facing API (gRPC)
- Secure: automatic TLS with optional client cert authentication
- Fast: benchmarked 10,000 writes/sec
- Reliable: properly distributed using Raft

etcd is written in Go and uses the [Raft](https://raft.github.io/) consensus algorithm to manage a highly-available replicated log.

etcd is used [in production by many companies](https://github.com/etcd-io/etcd/blob/main/ADOPTERS.md), and the development team stands behind it in critical deployment scenarios, where etcd is frequently teamed with applications such as Kubernetes, locksmith, vulcand, Doorman, and many others. Reliability is further ensured by rigorous robustness testing.

See [etcdctl](https://github.com/etcd-io/etcd/tree/main/etcdctl) for a simple command line client.