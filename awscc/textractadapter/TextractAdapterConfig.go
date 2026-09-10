// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package textractadapter

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type TextractAdapterConfig struct {
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
	// The name to be assigned to the adapter being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/textract_adapter#adapter_name TextractAdapter#adapter_name}
	AdapterName *string `field:"required" json:"adapterName" yaml:"adapterName"`
	// The type of feature that the adapter is being trained on. Currently, supported feature types are: QUERIES.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/textract_adapter#feature_types TextractAdapter#feature_types}
	FeatureTypes *[]*string `field:"required" json:"featureTypes" yaml:"featureTypes"`
	// Controls whether or not the adapter should automatically update.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/textract_adapter#auto_update TextractAdapter#auto_update}
	AutoUpdate *string `field:"optional" json:"autoUpdate" yaml:"autoUpdate"`
	// The description to be assigned to the adapter being created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/textract_adapter#description TextractAdapter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A list of tags to be added to the adapter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/textract_adapter#tags TextractAdapter#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

