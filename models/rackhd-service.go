package models

// AMQP constants for voyager-secret-service
const (
  RackHDExchange     string = "voyager-rackhd-service"
  RackHDExchangeType string = "topic"
  RackHDQueueName    string = "voyager-rackhd-service-queue"
  RackHDBindingKey   string = "requests"
  RackHDConsumerTag  string = "consumer-tag"
)

// Action types for template
const (
  UploadTemplateAction string = "upload-template"
  UpdateTemplateAction string = "update-template"
  DeleteTemplateAction string = "delete-template"
)

// Action types for workflows
const (
  UploadWorkflowAction string = "upload-workflow"
  UpdateWorkflowAction string = "update-workflow"
  DeleteWorkflowAction string = "delete-workflow"
  RunWorkflowAction    string = "run-workflow"
  ListenWorkflowAction string = "listen-workflow"
)

const (
  OnEventsExchange     string = "on.events"
  OnEventsExchangeType string = "topic"
)

const (
  GraphFinishedRoutingKey string = "graph.finished"
)

// RackHD URI
const (
  RackHDTemplateURI string = "/api/common/templates/library/"
  RackHDGraphURI    string = "/api/common/workflows/graphs"
  RackHDRootURI     string = "/api/common/"
)

// RackHD files and variables for templates and workflows
const (
  RackHDCiscoNexusDeployConfigTemplate string = "ciscoNexus-deploy-config-and-images.py"
  RackHDDeployConfigWorkflow           string = "DeployStartupConfigAndBootImages.json"
  DefaultCiscoBootImage                string = "nxos.7.0.3.I2.2d.bin"
  DefaultCiscoInjectableName           string = "Graph.Switch.CiscoNexus3000.Deploy"
)

// RackHDConfigReq is used to put the config template in RackHD
type RackHDConfigReq struct {
  RackHDReq
  Name   string `json:"name"`
  Config string `json:"config"`
}

// RackHDWorkflowReq is used to put the workflows in RackHD
type RackHDWorkflowReq struct {
  RackHDReq
  Workflow string `json:"Workflow"`
}

// RackHDRunWorkflowReq is use to run workflows in RackHD
type RackHDRunWorkflowReq struct {
  RackHDReq
  RackHDWorkflowConfig
  NodeID string `json:"nodeID"`
}

// RackHDListenReq struct tells voyager-rackhd-service which exchange and routing key that it should use
type RackHDListenReq struct {
  RackHDReq
  Exchange     string `json:"exchange"`
  ExchangeType string `json:"exchangeType"`
  RoutingKey   string `json:"routingKey"`
}

// RackHDWorkflowConfig is used to define the config for the workflow to be executed
type RackHDWorkflowConfig struct {
  Name    string                `json:"name"`
  Options WorkflowConfigOptions `json:"options"`
}

// WorkflowConfigOptions is used to define the options for the workflow to be executed
type WorkflowConfigOptions struct {
  DeployConfigAndImages WorkflowConfigAndImages `json:"deploy-config-and-images"`
}

// WorkflowConfigAndImages is used to define the startup config and boot image for a switch
type WorkflowConfigAndImages struct {
  StartupConfig string `json:"startupConfig"`
  BootImage     string `json:"bootImage"`
}

// RackHDReq is used as general request for RackHD service
type RackHDReq struct {
  Action string `json:"action"`
}

// RackHDResp is used as general response for RackHD service
type RackHDResp struct {
  AmqpResp
  ServerResponse string `json:"response"`
}
