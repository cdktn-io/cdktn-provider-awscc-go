// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentprivateconnection


type DevopsagentPrivateConnectionConnectionConfigurationSelfManaged struct {
	// The ARN of the Resource Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/devopsagent_private_connection#resource_configuration_id DevopsagentPrivateConnection#resource_configuration_id}
	ResourceConfigurationId *string `field:"optional" json:"resourceConfigurationId" yaml:"resourceConfigurationId"`
}

