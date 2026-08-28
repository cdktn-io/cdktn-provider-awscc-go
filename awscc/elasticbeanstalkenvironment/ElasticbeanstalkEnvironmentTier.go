// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticbeanstalkenvironment


type ElasticbeanstalkEnvironmentTier struct {
	// The name of this environment tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticbeanstalk_environment#name ElasticbeanstalkEnvironment#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of this environment tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticbeanstalk_environment#type ElasticbeanstalkEnvironment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The version of this environment tier.
	//
	// When you don't set a value to it, Elastic Beanstalk uses the latest compatible worker tier version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticbeanstalk_environment#version ElasticbeanstalkEnvironment#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

