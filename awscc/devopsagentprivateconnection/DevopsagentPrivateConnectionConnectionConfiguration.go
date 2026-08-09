// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentprivateconnection


type DevopsagentPrivateConnectionConnectionConfiguration struct {
	// Configuration for a self-managed Private Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_private_connection#self_managed DevopsagentPrivateConnection#self_managed}
	SelfManaged *DevopsagentPrivateConnectionConnectionConfigurationSelfManaged `field:"optional" json:"selfManaged" yaml:"selfManaged"`
	// Configuration for a service-managed Private Connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/devopsagent_private_connection#service_managed DevopsagentPrivateConnection#service_managed}
	ServiceManaged *DevopsagentPrivateConnectionConnectionConfigurationServiceManaged `field:"optional" json:"serviceManaged" yaml:"serviceManaged"`
}

