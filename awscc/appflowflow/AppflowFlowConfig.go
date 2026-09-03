// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appflowflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AppflowFlowConfig struct {
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
	// List of Destination connectors of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#destination_flow_config_list AppflowFlow#destination_flow_config_list}
	DestinationFlowConfigList interface{} `field:"required" json:"destinationFlowConfigList" yaml:"destinationFlowConfigList"`
	// Name of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#flow_name AppflowFlow#flow_name}
	FlowName *string `field:"required" json:"flowName" yaml:"flowName"`
	// Configurations of Source connector of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#source_flow_config AppflowFlow#source_flow_config}
	SourceFlowConfig *AppflowFlowSourceFlowConfig `field:"required" json:"sourceFlowConfig" yaml:"sourceFlowConfig"`
	// List of tasks for the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#tasks AppflowFlow#tasks}
	Tasks interface{} `field:"required" json:"tasks" yaml:"tasks"`
	// Trigger settings of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#trigger_config AppflowFlow#trigger_config}
	TriggerConfig *AppflowFlowTriggerConfig `field:"required" json:"triggerConfig" yaml:"triggerConfig"`
	// Description of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#description AppflowFlow#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Flow activation status for Scheduled- and Event-triggered flows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#flow_status AppflowFlow#flow_status}
	FlowStatus *string `field:"optional" json:"flowStatus" yaml:"flowStatus"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables.
	//
	// If it's not provided, AWS Lambda uses a default service key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#kms_arn AppflowFlow#kms_arn}
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
	// Configurations of metadata catalog of the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#metadata_catalog_config AppflowFlow#metadata_catalog_config}
	MetadataCatalogConfig *AppflowFlowMetadataCatalogConfig `field:"optional" json:"metadataCatalogConfig" yaml:"metadataCatalogConfig"`
	// List of Tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/appflow_flow#tags AppflowFlow#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

