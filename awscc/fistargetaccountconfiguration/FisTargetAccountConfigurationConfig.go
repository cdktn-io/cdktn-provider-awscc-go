// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fistargetaccountconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FisTargetAccountConfigurationConfig struct {
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
	// The AWS account ID of the target account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_target_account_configuration#account_id FisTargetAccountConfiguration#account_id}
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// The ID of the experiment template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_target_account_configuration#experiment_template_id FisTargetAccountConfiguration#experiment_template_id}
	ExperimentTemplateId *string `field:"required" json:"experimentTemplateId" yaml:"experimentTemplateId"`
	// The Amazon Resource Name (ARN) of an IAM role for the target account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_target_account_configuration#role_arn FisTargetAccountConfiguration#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The description of the target account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fis_target_account_configuration#description FisTargetAccountConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

