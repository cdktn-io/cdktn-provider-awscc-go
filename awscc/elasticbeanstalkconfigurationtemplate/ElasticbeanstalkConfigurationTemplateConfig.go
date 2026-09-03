// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkconfigurationtemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticbeanstalkConfigurationTemplateConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Elastic Beanstalk application to associate with this configuration template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticbeanstalk_configuration_template#application_name ElasticbeanstalkConfigurationTemplate#application_name}
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// An optional description for this configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticbeanstalk_configuration_template#description ElasticbeanstalkConfigurationTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of an environment whose settings you want to use to create the configuration template.
	//
	// You must specify EnvironmentId if you don't specify PlatformArn, SolutionStackName, or SourceConfiguration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticbeanstalk_configuration_template#environment_id ElasticbeanstalkConfigurationTemplate#environment_id}
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// Option values for the Elastic Beanstalk configuration, such as the instance type.
	//
	// If specified, these values override the values obtained from the solution stack or the source configuration template. For a complete list of Elastic Beanstalk configuration options, see [Option Values](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/command-options.html) in the AWS Elastic Beanstalk Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticbeanstalk_configuration_template#option_settings ElasticbeanstalkConfigurationTemplate#option_settings}
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The Amazon Resource Name (ARN) of the custom platform.
	//
	// For more information, see [Custom Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/custom-platforms.html) in the AWS Elastic Beanstalk Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticbeanstalk_configuration_template#platform_arn ElasticbeanstalkConfigurationTemplate#platform_arn}
	PlatformArn *string `field:"optional" json:"platformArn" yaml:"platformArn"`
	// The name of an Elastic Beanstalk solution stack (platform version) that this configuration uses.
	//
	// For example, 64bit Amazon Linux 2013.09 running Tomcat 7 Java 7. A solution stack specifies the operating system, runtime, and application server for a configuration template. It also determines the set of configuration options as well as the possible and default values. For more information, see [Supported Platforms](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/concepts.platforms.html) in the AWS Elastic Beanstalk Developer Guide.
	//
	//  You must specify SolutionStackName if you don't specify PlatformArn, EnvironmentId, or SourceConfiguration.
	//
	//  Use the ListAvailableSolutionStacks API to obtain a list of available solution stacks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticbeanstalk_configuration_template#solution_stack_name ElasticbeanstalkConfigurationTemplate#solution_stack_name}
	SolutionStackName *string `field:"optional" json:"solutionStackName" yaml:"solutionStackName"`
	// An Elastic Beanstalk configuration template to base this one on.
	//
	// If specified, Elastic Beanstalk uses the configuration values from the specified configuration template to create a new configuration.
	//
	// Values specified in OptionSettings override any values obtained from the SourceConfiguration.
	//
	// You must specify SourceConfiguration if you don't specify PlatformArn, EnvironmentId, or SolutionStackName.
	//
	// Constraint: If both solution stack name and source configuration are specified, the solution stack of the source configuration template must match the specified solution stack name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/elasticbeanstalk_configuration_template#source_configuration ElasticbeanstalkConfigurationTemplate#source_configuration}
	SourceConfiguration *ElasticbeanstalkConfigurationTemplateSourceConfiguration `field:"optional" json:"sourceConfiguration" yaml:"sourceConfiguration"`
}

