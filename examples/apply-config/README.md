# Running the Apply Config Example

This is just a simple example code to apply a `WorkloadOptimizer` config.
You can instrument any other Apply Config option you would like.
We documented only the `.spec.schedule` and `.spec.workloadTargetRef`.

```
# make sure you have the SF agent installed on the cluster
go run main.go --workload=test --resource=deployments --namespace=default --schedule="@hourly"

```

The output will be something like this:

```
{
  "kind": "WorkloadOptimizer",
  "apiVersion": "optimize.stormforge.io/v1",
  "metadata": {
    "name": "test",
    "namespace": "default"
  },
  "spec": {
    "workloadTargetRef": {
      "kind": "deployments",
      "name": "test"
    },
    "schedule": "@hourly"
  }
}
Workload Optimizer 'test' created successfully in namespace 'default'
```

And the object will be created/updated like this:

```
% k get wo test -o yaml
apiVersion: optimize.stormforge.io/v1
kind: WorkloadOptimizer
metadata:
  creationTimestamp: "2025-04-03T15:22:01Z"
  generation: 1
  name: test
  namespace: default
  resourceVersion: "88561"
  uid: d35fa4df-3381-46ad-bbcf-1f9b9e11fe5f
spec:
  schedule: '@hourly'
  workloadTargetRef:
    kind: deployments
    name: test
```