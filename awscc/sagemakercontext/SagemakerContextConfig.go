// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercontext

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerContextConfig struct {
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
	// The name of the context. Must be unique to your account in an AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_context#context_name SagemakerContext#context_name}
	ContextName *string `field:"required" json:"contextName" yaml:"contextName"`
	// The context type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_context#context_type SagemakerContext#context_type}
	ContextType *string `field:"required" json:"contextType" yaml:"contextType"`
	// The source type, ID, and URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_context#source SagemakerContext#source}
	Source *SagemakerContextSource `field:"required" json:"source" yaml:"source"`
	// The description of the context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_context#description SagemakerContext#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of properties to add to the context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_context#properties SagemakerContext#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// A list of tags to apply to the context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_context#tags SagemakerContext#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

