// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisemodelmanifest

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotfleetwiseModelManifestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_model_manifest#name IotfleetwiseModelManifest#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_model_manifest#signal_catalog_arn IotfleetwiseModelManifest#signal_catalog_arn}.
	SignalCatalogArn *string `field:"required" json:"signalCatalogArn" yaml:"signalCatalogArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_model_manifest#description IotfleetwiseModelManifest#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_model_manifest#nodes IotfleetwiseModelManifest#nodes}.
	Nodes *[]*string `field:"optional" json:"nodes" yaml:"nodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_model_manifest#status IotfleetwiseModelManifest#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_model_manifest#tags IotfleetwiseModelManifest#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

