// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowNdiConfig struct {
	// A prefix for the names of the NDI sources that the flow creates.
	//
	// If a custom name isn't specified, MediaConnect generates a unique 12-character ID as the prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#machine_name MediaconnectFlow#machine_name}
	MachineName *string `field:"optional" json:"machineName" yaml:"machineName"`
	// A list of up to three NDI discovery server configurations.
	//
	// While not required by the API, this configuration is necessary for NDI functionality to work properly.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#ndi_discovery_servers MediaconnectFlow#ndi_discovery_servers}
	NdiDiscoveryServers interface{} `field:"optional" json:"ndiDiscoveryServers" yaml:"ndiDiscoveryServers"`
	// A setting that controls whether NDI sources or outputs can be used in the flow.
	//
	// The default value is DISABLED. This value must be set as ENABLED for your flow to support NDI sources or outputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_flow#ndi_state MediaconnectFlow#ndi_state}
	NdiState *string `field:"optional" json:"ndiState" yaml:"ndiState"`
}

