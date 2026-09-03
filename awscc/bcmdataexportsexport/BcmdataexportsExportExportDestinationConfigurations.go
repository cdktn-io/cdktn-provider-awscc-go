// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport


type BcmdataexportsExportExportDestinationConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/bcmdataexports_export#s3_destination BcmdataexportsExport#s3_destination}.
	S3Destination *BcmdataexportsExportExportDestinationConfigurationsS3Destination `field:"required" json:"s3Destination" yaml:"s3Destination"`
}

