// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxdatarepositoryassociation


type FsxDataRepositoryAssociationS3 struct {
	// Specifies the type of updated objects (new, changed, deleted) that will be automatically exported from your file system to the linked S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fsx_data_repository_association#auto_export_policy FsxDataRepositoryAssociation#auto_export_policy}
	AutoExportPolicy *FsxDataRepositoryAssociationS3AutoExportPolicy `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// Specifies the type of updated objects (new, changed, deleted) that will be automatically imported from the linked S3 bucket to your file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/fsx_data_repository_association#auto_import_policy FsxDataRepositoryAssociation#auto_import_policy}
	AutoImportPolicy *FsxDataRepositoryAssociationS3AutoImportPolicy `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}

