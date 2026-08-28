// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2fpgaimage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Ec2FpgaImageConfig struct {
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
	// A description for the AFI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_fpga_image#description Ec2FpgaImage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The location of the encrypted design checkpoint in Amazon S3. The input must be a tarball.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_fpga_image#input_storage_location Ec2FpgaImage#input_storage_location}
	InputStorageLocation *Ec2FpgaImageInputStorageLocation `field:"optional" json:"inputStorageLocation" yaml:"inputStorageLocation"`
	// The location in Amazon S3 for the output logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_fpga_image#logs_storage_location Ec2FpgaImage#logs_storage_location}
	LogsStorageLocation *Ec2FpgaImageLogsStorageLocation `field:"optional" json:"logsStorageLocation" yaml:"logsStorageLocation"`
	// A name for the AFI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_fpga_image#name Ec2FpgaImage#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The tags assigned to the FPGA image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/ec2_fpga_image#tags Ec2FpgaImage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

