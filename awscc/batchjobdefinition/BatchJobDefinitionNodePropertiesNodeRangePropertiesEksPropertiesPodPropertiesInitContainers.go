// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#args BatchJobDefinition#args}.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#command BatchJobDefinition#command}.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#env BatchJobDefinition#env}.
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#image BatchJobDefinition#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#image_pull_policy BatchJobDefinition#image_pull_policy}.
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#resources BatchJobDefinition#resources}.
	Resources *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainersResources `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#security_context BatchJobDefinition#security_context}.
	SecurityContext *BatchJobDefinitionNodePropertiesNodeRangePropertiesEksPropertiesPodPropertiesInitContainersSecurityContext `field:"optional" json:"securityContext" yaml:"securityContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/batch_job_definition#volume_mounts BatchJobDefinition#volume_mounts}.
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

