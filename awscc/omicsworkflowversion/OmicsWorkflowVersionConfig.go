// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsworkflowversion

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsWorkflowVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#version_name OmicsWorkflowVersion#version_name}.
	VersionName *string `field:"required" json:"versionName" yaml:"versionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#workflow_id OmicsWorkflowVersion#workflow_id}.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#accelerators OmicsWorkflowVersion#accelerators}.
	Accelerators *string `field:"optional" json:"accelerators" yaml:"accelerators"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#definition_uri OmicsWorkflowVersion#definition_uri}.
	DefinitionUri *string `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#description OmicsWorkflowVersion#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#engine OmicsWorkflowVersion#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#main OmicsWorkflowVersion#main}.
	Main *string `field:"optional" json:"main" yaml:"main"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#parameter_template OmicsWorkflowVersion#parameter_template}.
	ParameterTemplate interface{} `field:"optional" json:"parameterTemplate" yaml:"parameterTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#storage_capacity OmicsWorkflowVersion#storage_capacity}.
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#storage_type OmicsWorkflowVersion#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// A map of resource tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#tags OmicsWorkflowVersion#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/omics_workflow_version#workflow_bucket_owner_id OmicsWorkflowVersion#workflow_bucket_owner_id}.
	WorkflowBucketOwnerId *string `field:"optional" json:"workflowBucketOwnerId" yaml:"workflowBucketOwnerId"`
}

