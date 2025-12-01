# Architecture Diagrams Directory

High-level architecture diagrams showing system structure and deployment.

## Purpose

Store architecture diagram source files for system decomposition, component views, and deployment.

## Diagrams

1. **System Decomposition** (`system-decomposition.drawio`)
   - Shows: High-level system modules and their relationships
   - Exported to: `report/images/system_decomposition.png`
   - Referenced in: `4.1_module_view.tex` Section 3.1.1

2. **Clean Architecture Layers** (`clean-architecture-layers.drawio`)
   - Shows: Layering within a single service (Domain, Application, Infrastructure)
   - Exported to: `report/images/clean-architecture-layers.png`
   - Referenced in: `4.1_module_view.tex` Section 3.1.2

3. **Service Architecture / C4 Container** (`service-architecture.drawio`)
   - Shows: All microservices, databases, message queues, and connections
   - Exported to: `report/images/service_architecture.png`
   - Referenced in: `4.2_component_connector_view.tex` Section 3.2.1

4. **Complete Component Diagram** (`complete-component-diagram.drawio`)
   - Enhanced version showing all interfaces and dependencies
   - Exported to: `report/images/complete_component_diagram.png`
   - Referenced in: `4.2_component_connector_view.tex` Section 3.2.1

5. **Integration Patterns** (`integration-patterns.drawio`)
   - Shows: Synchronous (REST) and Asynchronous (RabbitMQ) communication
   - Exported to: `report/images/synchronous_communication.png`, `asynchronous_communication.png`
   - Referenced in: `4.2_component_connector_view.tex` Section 3.2.2

6. **Deployment Architecture** (`deployment-architecture.drawio`)
   - Shows: Kubernetes nodes, pods, load balancers, network topology
   - Exported to: `report/images/deployment_architecture_onprem.png`
   - Referenced in: `4.3_allocation_view.tex` Section 3.3.1

7. **Enhanced Deployment** (`enhanced-deployment.drawio`)
   - More detailed version with infrastructure specifics
   - Exported to: `report/images/enhanced_deployment.png`
   - Referenced in: `4.3_allocation_view.tex` Section 3.3.1

8. **AI Pipeline Data Flow** (`ai-pipeline-dataflow.drawio`)
   - Shows: Data transformations through scoring and adaptive engine
   - Exported to: `report/images/ai_pipeline_dataflow.png`
   - Referenced in: `4.2_component_connector_view.tex` Section 3.2.1
