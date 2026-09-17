// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package organizationsorganization

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OrganizationsOrganizationConfig struct {
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
	// Specifies the feature set supported by the new organization. Each feature set supports different levels of functionality.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/organizations_organization#feature_set OrganizationsOrganization#feature_set}
	FeatureSet *string `field:"optional" json:"featureSet" yaml:"featureSet"`
}

