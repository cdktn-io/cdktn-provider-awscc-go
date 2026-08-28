// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedlens

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WellarchitectedLensConfig struct {
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
	// The JSON representation of a lens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wellarchitected_lens#json_string WellarchitectedLens#json_string}
	JsonString *string `field:"optional" json:"jsonString" yaml:"jsonString"`
	// The version of the lens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wellarchitected_lens#lens_version WellarchitectedLens#lens_version}
	LensVersion *string `field:"optional" json:"lensVersion" yaml:"lensVersion"`
	// The tags assigned to the lens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/wellarchitected_lens#tags WellarchitectedLens#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

