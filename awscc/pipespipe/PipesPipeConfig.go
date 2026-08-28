// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PipesPipeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#role_arn PipesPipe#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#source PipesPipe#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#target PipesPipe#target}.
	Target *string `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#description PipesPipe#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#desired_state PipesPipe#desired_state}.
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#enrichment PipesPipe#enrichment}.
	Enrichment *string `field:"optional" json:"enrichment" yaml:"enrichment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#enrichment_parameters PipesPipe#enrichment_parameters}.
	EnrichmentParameters *PipesPipeEnrichmentParameters `field:"optional" json:"enrichmentParameters" yaml:"enrichmentParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#kms_key_identifier PipesPipe#kms_key_identifier}.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#log_configuration PipesPipe#log_configuration}.
	LogConfiguration *PipesPipeLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#name PipesPipe#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#source_parameters PipesPipe#source_parameters}.
	SourceParameters *PipesPipeSourceParameters `field:"optional" json:"sourceParameters" yaml:"sourceParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#tags PipesPipe#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#target_parameters PipesPipe#target_parameters}.
	TargetParameters *PipesPipeTargetParameters `field:"optional" json:"targetParameters" yaml:"targetParameters"`
}

