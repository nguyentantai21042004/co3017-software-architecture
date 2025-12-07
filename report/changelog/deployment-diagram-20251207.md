# Changelog: Enhanced Deployment Diagram

**Date:** 2025-12-07
**Author:** AI Assistant
**Type:** Addition

## Summary

Added enhanced deployment diagram with detailed infrastructure components, network topology, and high availability configuration.

## Rationale

Task 2.12 required enhancing the deployment diagram to show:

- Kubernetes nodes and pods
- Load balancers
- Database clusters
- Message broker setup
- Network topology

## Changes Made

### 1. Created PlantUML Source File

- **File:** `report/images/enhanced_deployment.puml`
- **Content:** Detailed deployment diagram showing:
  - DMZ Layer with HAProxy/NGINX Load Balancer
  - Kubernetes Cluster with Master Nodes (HA - 3 nodes)
  - Node Pool A (Go Services - Compute Optimized)
  - Node Pool B (Java Services - Memory Optimized)
  - Data Layer (PostgreSQL HA, RabbitMQ Cluster, Redis Sentinel)
  - Monitoring Layer (Prometheus, Grafana, Loki, Alertmanager, Jaeger)
  - Storage Layer (Rook-Ceph/Longhorn, SAN/NAS)

### 2. Updated LaTeX File

- **File:** `report/contents/4.3_allocation_view.tex`
- **Changes:**
  - Added reference to enhanced deployment diagram
  - Added table describing deployment layers
  - Added network topology and security zones (VLANs)
  - Added high availability configuration details

### 3. Infrastructure Details Added

- **Network Topology:**

  - VLAN 10 (DMZ): Load Balancer
  - VLAN 20 (Application): Kubernetes workers
  - VLAN 30 (Data): Databases, message broker
  - VLAN 40 (Management): Monitoring, K8s masters

- **HA Configuration:**
  - Kubernetes: 3 master nodes with etcd cluster
  - PostgreSQL: Patroni with 1 Primary + 2 Replicas
  - RabbitMQ: 3-node cluster with mirrored queues
  - Redis: Sentinel with 1 Master + 2 Replicas + 3 Sentinels

## Verification

- PlantUML source file created successfully
- LaTeX file updated with new content
- Pending: Generate PNG from PlantUML and compile LaTeX

## Related Issues

- Task 2.12: Enhance Deployment Diagram
- Addresses gap identified in `report/issues/architecture-views-gaps.md`

## Note

The PNG image needs to be generated from the PlantUML source file using:

```bash
plantuml report/images/enhanced_deployment.puml
```

Or using an online PlantUML renderer.
