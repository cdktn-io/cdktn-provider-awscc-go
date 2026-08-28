// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasyncstoragesystem

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DatasyncStorageSystemConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_storage_system#agent_arns DatasyncStorageSystem#agent_arns}.
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_storage_system#server_configuration DatasyncStorageSystem#server_configuration}.
	ServerConfiguration *DatasyncStorageSystemServerConfiguration `field:"required" json:"serverConfiguration" yaml:"serverConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_storage_system#system_type DatasyncStorageSystem#system_type}.
	SystemType *string `field:"required" json:"systemType" yaml:"systemType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_storage_system#cloudwatch_log_group_arn DatasyncStorageSystem#cloudwatch_log_group_arn}.
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_storage_system#name DatasyncStorageSystem#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_storage_system#server_credentials DatasyncStorageSystem#server_credentials}.
	ServerCredentials *DatasyncStorageSystemServerCredentials `field:"optional" json:"serverCredentials" yaml:"serverCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/datasync_storage_system#tags DatasyncStorageSystem#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

