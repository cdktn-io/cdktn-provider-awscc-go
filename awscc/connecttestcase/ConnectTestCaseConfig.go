// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package connecttestcase

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ConnectTestCaseConfig struct {
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
	// The content of the test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#content ConnectTestCase#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The identifier of the Amazon Connect instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#instance_arn ConnectTestCase#instance_arn}
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#name ConnectTestCase#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#description ConnectTestCase#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Entry point for Testcase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#entry_point ConnectTestCase#entry_point}
	EntryPoint *ConnectTestCaseEntryPoint `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// The initialization data of the test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#initialization_data ConnectTestCase#initialization_data}
	InitializationData *string `field:"optional" json:"initializationData" yaml:"initializationData"`
	// The status of the test case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#status ConnectTestCase#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/connect_test_case#tags ConnectTestCase#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

