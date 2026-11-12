// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityagentprivateconnection


type SecurityagentPrivateConnectionMode struct {
	// The configuration for a self-managed private connection, where you manage your own resource configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_private_connection#self_managed SecurityagentPrivateConnection#self_managed}
	SelfManaged *SecurityagentPrivateConnectionModeSelfManaged `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// The configuration for a service-managed private connection, where the service manages the resource gateway lifecycle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/securityagent_private_connection#service_managed SecurityagentPrivateConnection#service_managed}
	ServiceManaged *SecurityagentPrivateConnectionModeServiceManaged `field:"optional" json:"serviceManaged" yaml:"serviceManaged"`
}

