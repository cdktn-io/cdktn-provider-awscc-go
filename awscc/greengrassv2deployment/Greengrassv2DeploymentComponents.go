// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment


type Greengrassv2DeploymentComponents struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#component_version Greengrassv2Deployment#component_version}.
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#configuration_update Greengrassv2Deployment#configuration_update}.
	ConfigurationUpdate *Greengrassv2DeploymentComponentsConfigurationUpdate `field:"optional" json:"configurationUpdate" yaml:"configurationUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/greengrassv2_deployment#run_with Greengrassv2Deployment#run_with}.
	RunWith *Greengrassv2DeploymentComponentsRunWith `field:"optional" json:"runWith" yaml:"runWith"`
}

