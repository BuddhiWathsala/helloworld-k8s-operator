kubectl delete -f deploy/crds/helloworld.io_helloworlds_crd.yaml
kubectl delete -f deploy/service_account.yaml
kubectl delete -f deploy/role.yaml
kubectl delete -f deploy/role_binding.yaml
kubectl delete -f deploy/operator.yaml
