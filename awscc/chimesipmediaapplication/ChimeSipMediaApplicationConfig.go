// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimesipmediaapplication

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ChimeSipMediaApplicationConfig struct {
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
	// The AWS Region in which the SIP media application is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_sip_media_application#aws_region ChimeSipMediaApplication#aws_region}
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// List of endpoints (Lambda ARNs) specified for the SIP media application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_sip_media_application#endpoints ChimeSipMediaApplication#endpoints}
	Endpoints interface{} `field:"required" json:"endpoints" yaml:"endpoints"`
	// The name of the SIP media application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_sip_media_application#name ChimeSipMediaApplication#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Tags assigned to the SIP media application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/chime_sip_media_application#tags ChimeSipMediaApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

