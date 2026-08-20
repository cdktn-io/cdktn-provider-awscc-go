// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package inspectorv2cisscanconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Inspectorv2CisScanConfigurationConfig struct {
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
	// Name of the scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_cis_scan_configuration#scan_name Inspectorv2CisScanConfiguration#scan_name}
	ScanName *string `field:"required" json:"scanName" yaml:"scanName"`
	// Choose a Schedule cadence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_cis_scan_configuration#schedule Inspectorv2CisScanConfiguration#schedule}
	Schedule *Inspectorv2CisScanConfigurationSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_cis_scan_configuration#security_level Inspectorv2CisScanConfiguration#security_level}.
	SecurityLevel *string `field:"required" json:"securityLevel" yaml:"securityLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_cis_scan_configuration#targets Inspectorv2CisScanConfiguration#targets}.
	Targets *Inspectorv2CisScanConfigurationTargets `field:"required" json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/inspectorv2_cis_scan_configuration#tags Inspectorv2CisScanConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

