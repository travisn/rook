#!/bin/bash

kubectl -n rook-ceph get pods --selector="app==rook-ceph-ganesha" --output=custom-columns=:.metadata.labels.instance,:.status.podIP | tail -n +2 |
while read line; do
	instance=$(echo $line | awk '{print $1}')
	ip=$(echo $line | awk '{print $2}')
	echo "$instance:$ip"
	kubectl -n rook-ceph exec rook-ceph-tools -- rados -p nfs-ganesha setomapval clustermap $instance $ip
done
