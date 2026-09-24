// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskManifestConfigSource struct {
	// Specifies the S3 bucket where you're hosting the manifest that you want AWS DataSync to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_task#s3 DatasyncTask#s3}
	S3 *DatasyncTaskManifestConfigSourceS3 `field:"optional" json:"s3" yaml:"s3"`
}

