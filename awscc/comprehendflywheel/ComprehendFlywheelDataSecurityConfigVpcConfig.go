// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package comprehendflywheel


type ComprehendFlywheelDataSecurityConfigVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/comprehend_flywheel#security_group_ids ComprehendFlywheel#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/comprehend_flywheel#subnets ComprehendFlywheel#subnets}.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

