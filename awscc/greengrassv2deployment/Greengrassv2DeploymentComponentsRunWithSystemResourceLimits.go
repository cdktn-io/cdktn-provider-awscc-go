// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package greengrassv2deployment


type Greengrassv2DeploymentComponentsRunWithSystemResourceLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/greengrassv2_deployment#cpus Greengrassv2Deployment#cpus}.
	Cpus *float64 `field:"optional" json:"cpus" yaml:"cpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/greengrassv2_deployment#memory Greengrassv2Deployment#memory}.
	Memory *float64 `field:"optional" json:"memory" yaml:"memory"`
}

