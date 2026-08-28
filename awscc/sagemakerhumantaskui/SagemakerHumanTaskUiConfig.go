// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerhumantaskui

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SagemakerHumanTaskUiConfig struct {
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
	// The name of the human task user interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_human_task_ui#human_task_ui_name SagemakerHumanTaskUi#human_task_ui_name}
	HumanTaskUiName *string `field:"required" json:"humanTaskUiName" yaml:"humanTaskUiName"`
	// An array of key-value pairs that contain metadata to help you categorize and organize a human review workflow user interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_human_task_ui#tags SagemakerHumanTaskUi#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Liquid template for the worker user interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_human_task_ui#ui_template SagemakerHumanTaskUi#ui_template}
	UiTemplate *SagemakerHumanTaskUiUiTemplate `field:"optional" json:"uiTemplate" yaml:"uiTemplate"`
}

