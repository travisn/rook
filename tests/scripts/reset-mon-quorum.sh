#!/bin/bash -e

: "${MON_TO_RESET:=${1}}"

if [ "$MON_TO_RESET" == "" ]; then
  echo "missing healthy mon name to reset the quorum"
  echo "reset-mon-quorum.sh <mon>"
  exit 1
fi

export mon_to_reset_deployment=rook-ceph-mon-$MON_TO_RESET
export rook_namespace=rook-ceph

# stop the operator deployment by scaling it down
echo "scaling down the operator"
kubectl -n $rook_namespace scale deployment rook-ceph-operator --replicas=0

# ensure the mon exists that we want to
echo "ensuring $mon_to_reset_deployment exists"
kubectl -n $rook_namespace get deploy $mon_to_reset_deployment

echo "get the node where the mon is assigned"
mon_node=$(kubectl -n $rook_namespace get deploy $mon_to_reset_deployment -o yaml | grep " kubernetes.io/hostname:" | awk '{print $2}')

echo "scaling mon $MON_TO_RESET to 0"
kubectl -n $rook_namespace scale deployment $mon_to_reset_deployment --replicas=0
echo "scaling other mons to 0"
#kubectl -n $rook_namespace scale deployment rook-ceph-mon-b --replicas=0
#kubectl -n $rook_namespace scale deployment rook-ceph-mon-c --replicas=0

echo "TODO: backing up configmap and secrets before updating them"


mon_public_ip=$(kubectl -n $rook_namespace get deploy $mon_to_reset_deployment -o yaml | grep " --public-addr=" | head -1 | sed 's/=/ /' | awk '{print $3}')
echo "mon_public_ip=$mon_public_ip"

echo "updating the rook-ceph-config secret with the new mon info"
kubectl -n $rook_namespace patch secret rook-ceph-config -p '{"stringData": {"mon_host": "[v2:'"$mon_public_ip"':3300,v1:'"$mon_public_ip"':6789]", "mon_initial_members": "'"${MON_TO_RESET}"'"}}'

echo "update the rook-ceph-mon-endpoints configmap with the new mon info:"
echo "  kubectl -n $rook_namespace edit cm rook-ceph-mon-endpoints"
echo "after updating the configmap,"

echo "run the toolbox job on NODE $mon_node to reset the monmap, then click any key to restart the mon and operator"
read -n 1

echo "scaling up the mon $MON_TO_RESET"
kubectl -n $rook_namespace scale deployment $mon_to_reset_deployment --replicas=1

echo "deleting obsolete mons and services"
kubectl -n $rook_namespace delete service rook-ceph-mon-b
kubectl -n $rook_namespace delete service rook-ceph-mon-c
kubectl -n $rook_namespace delete deployment rook-ceph-mon-b
kubectl -n $rook_namespace delete deployment rook-ceph-mon-c

#echo "scaling up the operator"
#kubectl -n $rook_namespace scale deployment rook-ceph-operator --replicas=1
