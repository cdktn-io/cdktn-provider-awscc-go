// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubsecuritycontrol

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SecurityhubSecurityControlConfig struct {
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
	// An object that identifies the name of a control parameter, its current value, and whether it has been customized.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#parameters SecurityhubSecurityControl#parameters}
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The most recent reason for updating the customizable properties of a security control.
	//
	// This differs from the UpdateReason field of the BatchUpdateStandardsControlAssociations API, which tracks the reason for updating the enablement status of a control. This field accepts alphanumeric characters in addition to white spaces, dashes, and underscores.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#last_update_reason SecurityhubSecurityControl#last_update_reason}
	LastUpdateReason *string `field:"optional" json:"lastUpdateReason" yaml:"lastUpdateReason"`
	// The Amazon Resource Name (ARN) for a security control across standards, such as `arn:aws:securityhub:eu-central-1:123456789012:security-control/S3.1`. This parameter doesn't mention a specific standard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#security_control_arn SecurityhubSecurityControl#security_control_arn}
	SecurityControlArn *string `field:"optional" json:"securityControlArn" yaml:"securityControlArn"`
	// The unique identifier of a security control across standards.
	//
	// Values for this field typically consist of an AWS service name and a number, such as APIGateway.3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/securityhub_security_control#security_control_id SecurityhubSecurityControl#security_control_id}
	SecurityControlId *string `field:"optional" json:"securityControlId" yaml:"securityControlId"`
}

