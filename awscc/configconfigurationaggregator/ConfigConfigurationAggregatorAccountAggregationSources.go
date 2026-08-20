// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package configconfigurationaggregator


type ConfigConfigurationAggregatorAccountAggregationSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_configuration_aggregator#account_ids ConfigConfigurationAggregator#account_ids}.
	AccountIds *[]*string `field:"optional" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_configuration_aggregator#all_aws_regions ConfigConfigurationAggregator#all_aws_regions}.
	AllAwsRegions interface{} `field:"optional" json:"allAwsRegions" yaml:"allAwsRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/config_configuration_aggregator#aws_regions ConfigConfigurationAggregator#aws_regions}.
	AwsRegions *[]*string `field:"optional" json:"awsRegions" yaml:"awsRegions"`
}

