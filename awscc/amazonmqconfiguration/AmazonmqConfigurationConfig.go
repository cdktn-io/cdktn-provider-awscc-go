// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package amazonmqconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AmazonmqConfigurationConfig struct {
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
	// The type of broker engine. Note: Currently, Amazon MQ only supports ACTIVEMQ for creating and editing broker configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amazonmq_configuration#engine_type AmazonmqConfiguration#engine_type}
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The name of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amazonmq_configuration#name AmazonmqConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The authentication strategy associated with the configuration. The default is SIMPLE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amazonmq_configuration#authentication_strategy AmazonmqConfiguration#authentication_strategy}
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// The base64-encoded XML configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amazonmq_configuration#data AmazonmqConfiguration#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amazonmq_configuration#description AmazonmqConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The version of the broker engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amazonmq_configuration#engine_version AmazonmqConfiguration#engine_version}
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Create tags when creating the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/amazonmq_configuration#tags AmazonmqConfiguration#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

