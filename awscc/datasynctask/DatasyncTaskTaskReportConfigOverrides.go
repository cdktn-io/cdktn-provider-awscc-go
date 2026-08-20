// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskTaskReportConfigOverrides struct {
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to delete in your destination location.
	//
	// This only applies if you configure your task to delete data in the destination that isn't in the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#deleted DatasyncTask#deleted}
	Deleted *DatasyncTaskTaskReportConfigOverridesDeleted `field:"optional" json:"deleted" yaml:"deleted"`
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to skip during your transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#skipped DatasyncTask#skipped}
	Skipped *DatasyncTaskTaskReportConfigOverridesSkipped `field:"optional" json:"skipped" yaml:"skipped"`
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#transferred DatasyncTask#transferred}
	Transferred *DatasyncTaskTaskReportConfigOverridesTransferred `field:"optional" json:"transferred" yaml:"transferred"`
	// Specifies the level of reporting for the files, objects, and directories that Datasync attempted to verify at the end of your transfer.
	//
	// This only applies if you configure your task to verify data during and after the transfer (which Datasync does by default)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datasync_task#verified DatasyncTask#verified}
	Verified *DatasyncTaskTaskReportConfigOverridesVerified `field:"optional" json:"verified" yaml:"verified"`
}

