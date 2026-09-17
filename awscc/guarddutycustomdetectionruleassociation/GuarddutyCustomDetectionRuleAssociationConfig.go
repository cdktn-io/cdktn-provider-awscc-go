// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package guarddutycustomdetectionruleassociation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type GuarddutyCustomDetectionRuleAssociationConfig struct {
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
	// Whether the rule runs in LIVE mode (generates findings) or DRY_RUN mode (evaluates without generating findings).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/guardduty_custom_detection_rule_association#mode GuarddutyCustomDetectionRuleAssociation#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The catalog identifier of the custom detection rule to associate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/guardduty_custom_detection_rule_association#rule_id GuarddutyCustomDetectionRuleAssociation#rule_id}
	RuleId *string `field:"required" json:"ruleId" yaml:"ruleId"`
	// The tags applied to the association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/guardduty_custom_detection_rule_association#tags GuarddutyCustomDetectionRuleAssociation#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

