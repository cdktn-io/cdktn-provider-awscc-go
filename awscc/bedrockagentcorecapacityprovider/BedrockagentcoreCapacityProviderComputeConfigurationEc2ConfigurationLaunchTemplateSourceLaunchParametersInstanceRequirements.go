// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorecapacityprovider


type BedrockagentcoreCapacityProviderComputeConfigurationEc2ConfigurationLaunchTemplateSourceLaunchParametersInstanceRequirements struct {
	// List of allowed instance types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_capacity_provider#allowed_instance_types BedrockagentcoreCapacityProvider#allowed_instance_types}
	AllowedInstanceTypes *[]*string `field:"required" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
}

