// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2fpgaimage


type Ec2FpgaImageLogsStorageLocation struct {
	// The name of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_fpga_image#bucket Ec2FpgaImage#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_fpga_image#key Ec2FpgaImage#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
}

