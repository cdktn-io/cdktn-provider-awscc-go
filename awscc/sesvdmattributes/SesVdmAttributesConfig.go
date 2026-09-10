// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesvdmattributes

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesVdmAttributesConfig struct {
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
	// Preferences regarding the Dashboard feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_vdm_attributes#dashboard_attributes SesVdmAttributes#dashboard_attributes}
	DashboardAttributes *SesVdmAttributesDashboardAttributes `field:"optional" json:"dashboardAttributes" yaml:"dashboardAttributes"`
	// Preferences regarding the Guardian feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_vdm_attributes#guardian_attributes SesVdmAttributes#guardian_attributes}
	GuardianAttributes *SesVdmAttributesGuardianAttributes `field:"optional" json:"guardianAttributes" yaml:"guardianAttributes"`
}

