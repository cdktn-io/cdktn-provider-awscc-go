// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsworkflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsWorkflowConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#accelerators OmicsWorkflow#accelerators}.
	Accelerators *string `field:"optional" json:"accelerators" yaml:"accelerators"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#definition_uri OmicsWorkflow#definition_uri}.
	DefinitionUri *string `field:"optional" json:"definitionUri" yaml:"definitionUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#description OmicsWorkflow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#engine OmicsWorkflow#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#main OmicsWorkflow#main}.
	Main *string `field:"optional" json:"main" yaml:"main"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#name OmicsWorkflow#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#parameter_template OmicsWorkflow#parameter_template}.
	ParameterTemplate interface{} `field:"optional" json:"parameterTemplate" yaml:"parameterTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#storage_capacity OmicsWorkflow#storage_capacity}.
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#storage_type OmicsWorkflow#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// A map of resource tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/omics_workflow#tags OmicsWorkflow#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

