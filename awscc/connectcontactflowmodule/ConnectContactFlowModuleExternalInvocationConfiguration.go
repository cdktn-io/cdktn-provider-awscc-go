// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectcontactflowmodule


type ConnectContactFlowModuleExternalInvocationConfiguration struct {
	// Specifies whether the flow module resource is enabled for external invocation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/connect_contact_flow_module#enabled ConnectContactFlowModule#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

