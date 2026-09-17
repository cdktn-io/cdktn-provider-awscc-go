// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package arczonalshiftzonalautoshiftconfiguration

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ArczonalshiftZonalAutoshiftConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/arczonalshift_zonal_autoshift_configuration#practice_run_configuration ArczonalshiftZonalAutoshiftConfiguration#practice_run_configuration}.
	PracticeRunConfiguration *ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfiguration `field:"optional" json:"practiceRunConfiguration" yaml:"practiceRunConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/arczonalshift_zonal_autoshift_configuration#resource_identifier ArczonalshiftZonalAutoshiftConfiguration#resource_identifier}.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/arczonalshift_zonal_autoshift_configuration#zonal_autoshift_status ArczonalshiftZonalAutoshiftConfiguration#zonal_autoshift_status}.
	ZonalAutoshiftStatus *string `field:"optional" json:"zonalAutoshiftStatus" yaml:"zonalAutoshiftStatus"`
}

