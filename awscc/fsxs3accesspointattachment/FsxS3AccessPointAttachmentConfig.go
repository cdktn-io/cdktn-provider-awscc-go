// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxS3AccessPointAttachmentConfig struct {
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
	// The name of the S3 access point attachment; also used for the name of the S3 access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#name FsxS3AccessPointAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of Amazon FSx volume that the S3 access point is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#type FsxS3AccessPointAttachment#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The OntapConfiguration of the S3 access point attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#ontap_configuration FsxS3AccessPointAttachment#ontap_configuration}
	OntapConfiguration *FsxS3AccessPointAttachmentOntapConfiguration `field:"optional" json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The OpenZFSConfiguration of the S3 access point attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#open_zfs_configuration FsxS3AccessPointAttachment#open_zfs_configuration}
	OpenZfsConfiguration *FsxS3AccessPointAttachmentOpenZfsConfiguration `field:"optional" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// The S3 access point configuration of the S3 access point attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#s3_access_point FsxS3AccessPointAttachment#s3_access_point}
	S3AccessPoint *FsxS3AccessPointAttachmentS3AccessPoint `field:"optional" json:"s3AccessPoint" yaml:"s3AccessPoint"`
}

