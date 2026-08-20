// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderdistributionconfiguration


type ImagebuilderDistributionConfigurationDistributionsSsmParameterConfigurations struct {
	// The account ID for the AMI to update the parameter with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_distribution_configuration#ami_account_id ImagebuilderDistributionConfiguration#ami_account_id}
	AmiAccountId *string `field:"optional" json:"amiAccountId" yaml:"amiAccountId"`
	// The data type of the SSM parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_distribution_configuration#data_type ImagebuilderDistributionConfiguration#data_type}
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The name of the SSM parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/imagebuilder_distribution_configuration#parameter_name ImagebuilderDistributionConfiguration#parameter_name}
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
}

