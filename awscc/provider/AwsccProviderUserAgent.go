// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type AwsccProviderUserAgent struct {
	// Product name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs#product_name AwsccProvider#product_name}
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// Comment describing any additional product details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs#comment AwsccProvider#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Product version. Optional, and should only be set when `product_name` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs#product_version AwsccProvider#product_version}
	ProductVersion *string `field:"optional" json:"productVersion" yaml:"productVersion"`
}

