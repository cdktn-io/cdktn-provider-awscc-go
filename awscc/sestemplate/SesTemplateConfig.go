// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sestemplate

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesTemplateConfig struct {
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
	// The tags (keys and values) associated with the email template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_template#tags SesTemplate#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The content of the email, composed of a subject line, an HTML part, and a text-only part.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ses_template#template SesTemplate#template}
	Template *SesTemplateTemplate `field:"optional" json:"template" yaml:"template"`
}

