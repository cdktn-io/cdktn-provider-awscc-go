// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmcontactsrotation

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SsmcontactsRotationConfig struct {
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
	// Members of the rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_rotation#contact_ids SsmcontactsRotation#contact_ids}
	ContactIds *[]*string `field:"required" json:"contactIds" yaml:"contactIds"`
	// Name of the Rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_rotation#name SsmcontactsRotation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Information about when an on-call rotation is in effect and how long the rotation period lasts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_rotation#recurrence SsmcontactsRotation#recurrence}
	Recurrence *SsmcontactsRotationRecurrence `field:"required" json:"recurrence" yaml:"recurrence"`
	// Start time of the first shift of Oncall Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_rotation#start_time SsmcontactsRotation#start_time}
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// TimeZone Identifier for the Oncall Schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_rotation#time_zone_id SsmcontactsRotation#time_zone_id}
	TimeZoneId *string `field:"required" json:"timeZoneId" yaml:"timeZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ssmcontacts_rotation#tags SsmcontactsRotation#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

