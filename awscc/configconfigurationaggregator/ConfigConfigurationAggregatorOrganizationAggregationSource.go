// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigurationaggregator


type ConfigConfigurationAggregatorOrganizationAggregationSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/config_configuration_aggregator#all_aws_regions ConfigConfigurationAggregator#all_aws_regions}.
	AllAwsRegions interface{} `field:"optional" json:"allAwsRegions" yaml:"allAwsRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/config_configuration_aggregator#aws_regions ConfigConfigurationAggregator#aws_regions}.
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/config_configuration_aggregator#role_arn ConfigConfigurationAggregator#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

