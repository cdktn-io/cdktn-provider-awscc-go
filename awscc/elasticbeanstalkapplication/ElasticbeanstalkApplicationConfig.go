// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticbeanstalkApplicationConfig struct {
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
	// A name for the Elastic Beanstalk application.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the application name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_application#application_name ElasticbeanstalkApplication#application_name}
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Your description of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_application#description ElasticbeanstalkApplication#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/elasticbeanstalk_application#resource_lifecycle_config ElasticbeanstalkApplication#resource_lifecycle_config}
	ResourceLifecycleConfig *ElasticbeanstalkApplicationResourceLifecycleConfig `field:"optional" json:"resourceLifecycleConfig" yaml:"resourceLifecycleConfig"`
}

