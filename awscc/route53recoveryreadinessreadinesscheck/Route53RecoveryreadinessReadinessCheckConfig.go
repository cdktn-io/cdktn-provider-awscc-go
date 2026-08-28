// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package route53recoveryreadinessreadinesscheck

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Route53RecoveryreadinessReadinessCheckConfig struct {
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
	// Name of the ReadinessCheck to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_readiness_check#readiness_check_name Route53RecoveryreadinessReadinessCheck#readiness_check_name}
	ReadinessCheckName *string `field:"optional" json:"readinessCheckName" yaml:"readinessCheckName"`
	// The name of the resource set to check.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_readiness_check#resource_set_name Route53RecoveryreadinessReadinessCheck#resource_set_name}
	ResourceSetName *string `field:"optional" json:"resourceSetName" yaml:"resourceSetName"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/route53recoveryreadiness_readiness_check#tags Route53RecoveryreadinessReadinessCheck#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

