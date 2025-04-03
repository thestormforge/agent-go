Agent Custom Resource Client
============================

This repository contains the Kubernetes client for the agent's custom resources (if installed).

## Examples

We provided some examples in how to use this module.
It assumes you have installed the StormForge agent with the CRD.
This module does not install the CRD `workloadoptimizers.optimize.stormforge.io` on your cluster for you.

### Using ApplyConfig Method (preferred)

[Apply Config](./examples/apply-config/main.go)

### Using Controller Runtime Client

[Controller Runtime Client](./examples/create-object/main.go)

