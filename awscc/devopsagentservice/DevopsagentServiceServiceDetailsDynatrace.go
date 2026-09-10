// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package devopsagentservice


type DevopsagentServiceServiceDetailsDynatrace struct {
	// Dynatrace resource account URN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#account_urn DevopsagentService#account_urn}
	AccountUrn *string `field:"optional" json:"accountUrn" yaml:"accountUrn"`
	// Dynatrace OAuth authorization configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/devopsagent_service#authorization_config DevopsagentService#authorization_config}
	AuthorizationConfig *DevopsagentServiceServiceDetailsDynatraceAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
}

