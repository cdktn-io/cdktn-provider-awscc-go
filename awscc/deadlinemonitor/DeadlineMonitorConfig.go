// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package deadlinemonitor

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DeadlineMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_monitor#display_name DeadlineMonitor#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_monitor#identity_center_instance_arn DeadlineMonitor#identity_center_instance_arn}.
	IdentityCenterInstanceArn *string `field:"required" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_monitor#role_arn DeadlineMonitor#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_monitor#subdomain DeadlineMonitor#subdomain}.
	Subdomain *string `field:"required" json:"subdomain" yaml:"subdomain"`
	// The AWS region where IAM Identity Center is enabled.
	//
	// Required when Identity Center is in a different region than the monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_monitor#identity_center_region DeadlineMonitor#identity_center_region}
	IdentityCenterRegion *string `field:"optional" json:"identityCenterRegion" yaml:"identityCenterRegion"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/deadline_monitor#tags DeadlineMonitor#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

