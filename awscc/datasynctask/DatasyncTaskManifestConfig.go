// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskManifestConfig struct {
	// Specifies what DataSync uses the manifest for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#action DatasyncTask#action}
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Specifies the file format of your manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#format DatasyncTask#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Specifies the manifest that you want DataSync to use and where it's hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#source DatasyncTask#source}
	Source *DatasyncTaskManifestConfigSource `field:"optional" json:"source" yaml:"source"`
}

