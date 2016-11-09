package models

// CiscoTemplateData is used to render the cisco template configuration
type CiscoTemplateData struct {
	IpamLeaseResp
	Credentials
}

// WorkflowInstanceID is the instance ID of a RackHD workflow
type WorkflowInstanceID struct {
	InstanceID string `json:"instanceId"`
}
