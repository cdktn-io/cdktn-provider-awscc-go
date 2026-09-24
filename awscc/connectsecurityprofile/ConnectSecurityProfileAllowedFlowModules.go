// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connectsecurityprofile


type ConnectSecurityProfileAllowedFlowModules struct {
	// The identifier of the application that you want to give access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_security_profile#flow_module_id ConnectSecurityProfile#flow_module_id}
	FlowModuleId *string `field:"optional" json:"flowModuleId" yaml:"flowModuleId"`
	// The type of the first-party application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/connect_security_profile#type ConnectSecurityProfile#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

