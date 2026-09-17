// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mwaaserverlessworkflow

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type MwaaserverlessWorkflowConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#definition_s3_location MwaaserverlessWorkflow#definition_s3_location}.
	DefinitionS3Location *MwaaserverlessWorkflowDefinitionS3Location `field:"required" json:"definitionS3Location" yaml:"definitionS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#role_arn MwaaserverlessWorkflow#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The location of code artifacts in Amazon S3 for the workflow.
	//
	// Modeled as a single-member container so it stays extensible to future artifact types (e.g. OCI images).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#code MwaaserverlessWorkflow#code}
	Code *MwaaserverlessWorkflowCode `field:"optional" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#description MwaaserverlessWorkflow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#encryption_configuration MwaaserverlessWorkflow#encryption_configuration}.
	EncryptionConfiguration *MwaaserverlessWorkflowEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#logging_configuration MwaaserverlessWorkflow#logging_configuration}.
	LoggingConfiguration *MwaaserverlessWorkflowLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#name MwaaserverlessWorkflow#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#network_configuration MwaaserverlessWorkflow#network_configuration}.
	NetworkConfiguration *MwaaserverlessWorkflowNetworkConfiguration `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// A map of key-value pairs to be applied as tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#tags MwaaserverlessWorkflow#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mwaaserverless_workflow#trigger_mode MwaaserverlessWorkflow#trigger_mode}.
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}

