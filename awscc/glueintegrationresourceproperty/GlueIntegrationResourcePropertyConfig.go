// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package glueintegrationresourceproperty

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GlueIntegrationResourcePropertyConfig struct {
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
	// The connection ARN of the source, or the database ARN of the target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_integration_resource_property#resource_arn GlueIntegrationResourceProperty#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The resource properties associated with the integration source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_integration_resource_property#source_processing_properties GlueIntegrationResourceProperty#source_processing_properties}
	SourceProcessingProperties *GlueIntegrationResourcePropertySourceProcessingProperties `field:"optional" json:"sourceProcessingProperties" yaml:"sourceProcessingProperties"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_integration_resource_property#tags GlueIntegrationResourceProperty#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The resource properties associated with the integration target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_integration_resource_property#target_processing_properties GlueIntegrationResourceProperty#target_processing_properties}
	TargetProcessingProperties *GlueIntegrationResourcePropertyTargetProcessingProperties `field:"optional" json:"targetProcessingProperties" yaml:"targetProcessingProperties"`
}

