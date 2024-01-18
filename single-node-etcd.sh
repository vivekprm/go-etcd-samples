#!/bin/zsh

# https://etcd.io/docs/v3.5/op-guide/container/
export NODE1=192.168.1.21

docker volume create --name etcd-data
export DATA_DIR="etcd-data"

REGISTRY=quay.io/coreos/etcd
# available from v3.2.5
# REGISTRY=gcr.io/etcd-development/etcd

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

# go get go.etcd.io/etcd/client/v3
# etcdctl --endpoints=http://${NODE1}:2379 member list
# etcdctl --endpoints=http://0.0.0.0:2379 put greeting "Hello, etcd"
# etcdctl --endpoints=http://0.0.0.0:2379 get greeting 

# docker run -d -p 8080:8080 nikfoundas/etcd-viewer
#
# cd ../../../github.com/henszey/etcd-browser
# docker build -t etcd-browser .
# docker run --rm --name etcd-browser -p 0.0.0.0:8000:8000 --env ETCD_HOST=10.10.0.1 --env AUTH_PASS=doe -t -i etcd-browser