// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsoptiongroup


type RdsOptionGroupOptionConfigurations struct {
	// A list of DB security groups used for this option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_option_group#db_security_group_memberships RdsOptionGroup#db_security_group_memberships}
	DbSecurityGroupMemberships *[]*string `field:"optional" json:"dbSecurityGroupMemberships" yaml:"dbSecurityGroupMemberships"`
	// The configuration of options to include in a group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_option_group#option_name RdsOptionGroup#option_name}
	OptionName *string `field:"optional" json:"optionName" yaml:"optionName"`
	// The option settings to include in an option group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_option_group#option_settings RdsOptionGroup#option_settings}
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The version for the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_option_group#option_version RdsOptionGroup#option_version}
	OptionVersion *string `field:"optional" json:"optionVersion" yaml:"optionVersion"`
	// The optional port for the option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_option_group#port RdsOptionGroup#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// A list of VPC security group names used for this option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_option_group#vpc_security_group_memberships RdsOptionGroup#vpc_security_group_memberships}
	VpcSecurityGroupMemberships *[]*string `field:"optional" json:"vpcSecurityGroupMemberships" yaml:"vpcSecurityGroupMemberships"`
}

