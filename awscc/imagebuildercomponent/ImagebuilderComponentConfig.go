// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuildercomponent

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ImagebuilderComponentConfig struct {
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
	// The name of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#name ImagebuilderComponent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The platform of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#platform ImagebuilderComponent#platform}
	Platform *string `field:"required" json:"platform" yaml:"platform"`
	// The version of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#version ImagebuilderComponent#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The change description of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#change_description ImagebuilderComponent#change_description}
	ChangeDescription *string `field:"optional" json:"changeDescription" yaml:"changeDescription"`
	// The data of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#data ImagebuilderComponent#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
	// The description of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#description ImagebuilderComponent#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The KMS key identifier used to encrypt the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#kms_key_id ImagebuilderComponent#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The operating system (OS) version supported by the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#supported_os_versions ImagebuilderComponent#supported_os_versions}
	SupportedOsVersions *[]*string `field:"optional" json:"supportedOsVersions" yaml:"supportedOsVersions"`
	// The tags associated with the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#tags ImagebuilderComponent#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The uri of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/imagebuilder_component#uri ImagebuilderComponent#uri}
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}

