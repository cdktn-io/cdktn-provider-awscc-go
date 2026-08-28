// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessrecoverygroup

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53RecoveryreadinessRecoveryGroupConfig struct {
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
	// A list of the cell Amazon Resource Names (ARNs) in the recovery group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_recovery_group#cells Route53RecoveryreadinessRecoveryGroup#cells}
	Cells *[]*string `field:"optional" json:"cells" yaml:"cells"`
	// The name of the recovery group to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_recovery_group#recovery_group_name Route53RecoveryreadinessRecoveryGroup#recovery_group_name}
	RecoveryGroupName *string `field:"optional" json:"recoveryGroupName" yaml:"recoveryGroupName"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_recovery_group#tags Route53RecoveryreadinessRecoveryGroup#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

