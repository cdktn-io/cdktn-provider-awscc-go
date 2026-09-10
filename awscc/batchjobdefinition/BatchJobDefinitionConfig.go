// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BatchJobDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#type BatchJobDefinition#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#consumable_resource_properties BatchJobDefinition#consumable_resource_properties}.
	ConsumableResourceProperties *BatchJobDefinitionConsumableResourceProperties `field:"optional" json:"consumableResourceProperties" yaml:"consumableResourceProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#container_properties BatchJobDefinition#container_properties}.
	ContainerProperties *BatchJobDefinitionContainerProperties `field:"optional" json:"containerProperties" yaml:"containerProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#ecs_properties BatchJobDefinition#ecs_properties}.
	EcsProperties *BatchJobDefinitionEcsProperties `field:"optional" json:"ecsProperties" yaml:"ecsProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#eks_properties BatchJobDefinition#eks_properties}.
	EksProperties *BatchJobDefinitionEksProperties `field:"optional" json:"eksProperties" yaml:"eksProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#job_definition_name BatchJobDefinition#job_definition_name}.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#node_properties BatchJobDefinition#node_properties}.
	NodeProperties *BatchJobDefinitionNodeProperties `field:"optional" json:"nodeProperties" yaml:"nodeProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#parameters BatchJobDefinition#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#platform_capabilities BatchJobDefinition#platform_capabilities}.
	PlatformCapabilities *[]*string `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#propagate_tags BatchJobDefinition#propagate_tags}.
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#resource_retention_policy BatchJobDefinition#resource_retention_policy}.
	ResourceRetentionPolicy *BatchJobDefinitionResourceRetentionPolicy `field:"optional" json:"resourceRetentionPolicy" yaml:"resourceRetentionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#retry_strategy BatchJobDefinition#retry_strategy}.
	RetryStrategy *BatchJobDefinitionRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#scheduling_priority BatchJobDefinition#scheduling_priority}.
	SchedulingPriority *float64 `field:"optional" json:"schedulingPriority" yaml:"schedulingPriority"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#tags BatchJobDefinition#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/batch_job_definition#timeout BatchJobDefinition#timeout}.
	Timeout *BatchJobDefinitionTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

