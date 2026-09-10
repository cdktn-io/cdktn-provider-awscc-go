// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package entityresolutionidnamespace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EntityresolutionIdNamespaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#id_namespace_name EntityresolutionIdNamespace#id_namespace_name}.
	IdNamespaceName *string `field:"required" json:"idNamespaceName" yaml:"idNamespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#type EntityresolutionIdNamespace#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#description EntityresolutionIdNamespace#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#id_mapping_workflow_properties EntityresolutionIdNamespace#id_mapping_workflow_properties}.
	IdMappingWorkflowProperties interface{} `field:"optional" json:"idMappingWorkflowProperties" yaml:"idMappingWorkflowProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#input_source_config EntityresolutionIdNamespace#input_source_config}.
	InputSourceConfig interface{} `field:"optional" json:"inputSourceConfig" yaml:"inputSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#role_arn EntityresolutionIdNamespace#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/entityresolution_id_namespace#tags EntityresolutionIdNamespace#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

