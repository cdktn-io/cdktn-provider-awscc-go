// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package controltowerlandingzone

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ControltowerLandingZoneConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/controltower_landing_zone#manifest ControltowerLandingZone#manifest}.
	Manifest *string `field:"required" json:"manifest" yaml:"manifest"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/controltower_landing_zone#version ControltowerLandingZone#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/controltower_landing_zone#remediation_types ControltowerLandingZone#remediation_types}.
	RemediationTypes *[]*string `field:"optional" json:"remediationTypes" yaml:"remediationTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/controltower_landing_zone#tags ControltowerLandingZone#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

