// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qbusinessdatasource


type QbusinessDataSourceVpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#security_group_ids QbusinessDataSource#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/qbusiness_data_source#subnet_ids QbusinessDataSource#subnet_ids}.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
}

