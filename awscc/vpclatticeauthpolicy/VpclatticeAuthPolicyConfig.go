// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vpclatticeauthpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VpclatticeAuthPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/vpclattice_auth_policy#policy VpclatticeAuthPolicy#policy}.
	Policy *string `field:"required" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/vpclattice_auth_policy#resource_identifier VpclatticeAuthPolicy#resource_identifier}.
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
}

