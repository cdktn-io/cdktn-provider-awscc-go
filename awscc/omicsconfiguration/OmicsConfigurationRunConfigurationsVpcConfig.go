// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsconfiguration


type OmicsConfigurationRunConfigurationsVpcConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/omics_configuration#security_group_ids OmicsConfiguration#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/omics_configuration#subnet_ids OmicsConfiguration#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// The Amazon Web Services account ID of the cross-account VPC owner for RAM-shared VPC configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.108.0/docs/resources/omics_configuration#vpc_owner_account_id OmicsConfiguration#vpc_owner_account_id}
	VpcOwnerAccountId *string `field:"optional" json:"vpcOwnerAccountId" yaml:"vpcOwnerAccountId"`
}

