// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package translateparalleldata

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TranslateParallelDataConfig struct {
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
	// A custom name for the parallel data resource. Must be unique in the account and region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/translate_parallel_data#name TranslateParallelData#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the format and S3 location of the parallel data input file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/translate_parallel_data#parallel_data_config TranslateParallelData#parallel_data_config}
	ParallelDataConfig *TranslateParallelDataParallelDataConfig `field:"required" json:"parallelDataConfig" yaml:"parallelDataConfig"`
	// A custom description for the parallel data resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/translate_parallel_data#description TranslateParallelData#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The encryption key used to encrypt this object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/translate_parallel_data#encryption_key TranslateParallelData#encryption_key}
	EncryptionKey *TranslateParallelDataEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Tags associated with the parallel data resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/translate_parallel_data#tags TranslateParallelData#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

