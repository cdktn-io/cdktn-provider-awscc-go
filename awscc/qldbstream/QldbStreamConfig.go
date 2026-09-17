// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package qldbstream

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QldbStreamConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#inclusive_start_time QldbStream#inclusive_start_time}.
	InclusiveStartTime *string `field:"required" json:"inclusiveStartTime" yaml:"inclusiveStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#kinesis_configuration QldbStream#kinesis_configuration}.
	KinesisConfiguration *QldbStreamKinesisConfiguration `field:"required" json:"kinesisConfiguration" yaml:"kinesisConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#ledger_name QldbStream#ledger_name}.
	LedgerName *string `field:"required" json:"ledgerName" yaml:"ledgerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#role_arn QldbStream#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#stream_name QldbStream#stream_name}.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#exclusive_end_time QldbStream#exclusive_end_time}.
	ExclusiveEndTime *string `field:"optional" json:"exclusiveEndTime" yaml:"exclusiveEndTime"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/qldb_stream#tags QldbStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

