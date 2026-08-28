// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoverycontrolcontrolpanel

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53RecoverycontrolControlPanelConfig struct {
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
	// The name of the control panel. You can use any non-white space character in the name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoverycontrol_control_panel#name Route53RecoverycontrolControlPanel#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Cluster to associate with the Control Panel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoverycontrol_control_panel#cluster_arn Route53RecoverycontrolControlPanel#cluster_arn}
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoverycontrol_control_panel#tags Route53RecoverycontrolControlPanel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

