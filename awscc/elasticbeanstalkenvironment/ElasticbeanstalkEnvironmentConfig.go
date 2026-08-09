// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkenvironment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticbeanstalkEnvironmentConfig struct {
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
	// The name of the application that is associated with this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#application_name ElasticbeanstalkEnvironment#application_name}
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// If specified, the environment attempts to use this value as the prefix for the CNAME in your Elastic Beanstalk environment URL.
	//
	// If not specified, the CNAME is generated automatically by appending a random alphanumeric string to the environment name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#cname_prefix ElasticbeanstalkEnvironment#cname_prefix}
	CnamePrefix *string `field:"optional" json:"cnamePrefix" yaml:"cnamePrefix"`
	// Your description for this environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#description ElasticbeanstalkEnvironment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A unique name for the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#environment_name ElasticbeanstalkEnvironment#environment_name}
	EnvironmentName *string `field:"optional" json:"environmentName" yaml:"environmentName"`
	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the environment's operations role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#operations_role ElasticbeanstalkEnvironment#operations_role}
	OperationsRole *string `field:"optional" json:"operationsRole" yaml:"operationsRole"`
	// Key-value pairs defining configuration options for this environment, such as the instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#option_settings ElasticbeanstalkEnvironment#option_settings}
	OptionSettings interface{} `field:"optional" json:"optionSettings" yaml:"optionSettings"`
	// The Amazon Resource Name (ARN) of the custom platform to use with the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#platform_arn ElasticbeanstalkEnvironment#platform_arn}
	PlatformArn *string `field:"optional" json:"platformArn" yaml:"platformArn"`
	// The name of an Elastic Beanstalk solution stack (platform version) to use with the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#solution_stack_name ElasticbeanstalkEnvironment#solution_stack_name}
	SolutionStackName *string `field:"optional" json:"solutionStackName" yaml:"solutionStackName"`
	// Specifies the tags applied to resources in the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#tags ElasticbeanstalkEnvironment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the Elastic Beanstalk configuration template to use with the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#template_name ElasticbeanstalkEnvironment#template_name}
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
	// Specifies the tier to use in creating this environment.
	//
	// The environment tier that you choose determines whether Elastic Beanstalk provisions resources to support a web application that handles HTTP(S) requests or a web application that handles background-processing tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#tier ElasticbeanstalkEnvironment#tier}
	Tier *ElasticbeanstalkEnvironmentTier `field:"optional" json:"tier" yaml:"tier"`
	// The name of the application version to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/elasticbeanstalk_environment#version_label ElasticbeanstalkEnvironment#version_label}
	VersionLabel *string `field:"optional" json:"versionLabel" yaml:"versionLabel"`
}

