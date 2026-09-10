// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bcmdataexportsexport


type BcmdataexportsExportExport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bcmdataexports_export#data_query BcmdataexportsExport#data_query}.
	DataQuery *BcmdataexportsExportExportDataQuery `field:"required" json:"dataQuery" yaml:"dataQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bcmdataexports_export#destination_configurations BcmdataexportsExport#destination_configurations}.
	DestinationConfigurations *BcmdataexportsExportExportDestinationConfigurations `field:"required" json:"destinationConfigurations" yaml:"destinationConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bcmdataexports_export#name BcmdataexportsExport#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bcmdataexports_export#refresh_cadence BcmdataexportsExport#refresh_cadence}.
	RefreshCadence *BcmdataexportsExportExportRefreshCadence `field:"required" json:"refreshCadence" yaml:"refreshCadence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bcmdataexports_export#description BcmdataexportsExport#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

