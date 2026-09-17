// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkconfigurationtemplate


type ElasticbeanstalkConfigurationTemplateOptionSettings struct {
	// A unique namespace that identifies the option's associated AWS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_configuration_template#namespace ElasticbeanstalkConfigurationTemplate#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the configuration option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_configuration_template#option_name ElasticbeanstalkConfigurationTemplate#option_name}
	OptionName *string `field:"optional" json:"optionName" yaml:"optionName"`
	// A unique resource name for the option setting. Use it for a time–based scaling configuration option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_configuration_template#resource_name ElasticbeanstalkConfigurationTemplate#resource_name}
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// The current value for the configuration option.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_configuration_template#value ElasticbeanstalkConfigurationTemplate#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

