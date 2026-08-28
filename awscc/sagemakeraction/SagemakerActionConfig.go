// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraction

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerActionConfig struct {
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
	// The name of the action. Must be unique to your account in an AWS Region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#action_name SagemakerAction#action_name}
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// The action type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#action_type SagemakerAction#action_type}
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// The source type, ID, and URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#source SagemakerAction#source}
	Source *SagemakerActionSource `field:"required" json:"source" yaml:"source"`
	// The description of the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#description SagemakerAction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata properties of the tracking entity, trial, or trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#metadata_properties SagemakerAction#metadata_properties}
	MetadataProperties *SagemakerActionMetadataProperties `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// A list of properties to add to the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#properties SagemakerAction#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// The status of the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#status SagemakerAction#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// A list of tags to apply to the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_action#tags SagemakerAction#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

