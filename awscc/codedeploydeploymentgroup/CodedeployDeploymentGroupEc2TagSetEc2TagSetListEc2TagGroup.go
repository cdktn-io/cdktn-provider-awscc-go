// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupEc2TagSetEc2TagSetListEc2TagGroup struct {
	// The tag filter key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codedeploy_deployment_group#key CodedeployDeploymentGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag filter type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codedeploy_deployment_group#type CodedeployDeploymentGroup#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The tag filter value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/codedeploy_deployment_group#value CodedeployDeploymentGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

