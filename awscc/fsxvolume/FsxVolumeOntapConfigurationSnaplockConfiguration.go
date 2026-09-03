// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationSnaplockConfiguration struct {
	// Enables or disables the audit log volume for an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#audit_log_volume FsxVolume#audit_log_volume}
	AuditLogVolume *string `field:"optional" json:"auditLogVolume" yaml:"auditLogVolume"`
	// The configuration object for setting the autocommit period of files in an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#autocommit_period FsxVolume#autocommit_period}
	AutocommitPeriod *FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriod `field:"optional" json:"autocommitPeriod" yaml:"autocommitPeriod"`
	// Enables, disables, or permanently disables privileged delete on an FSx for ONTAP SnapLock Enterprise volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#privileged_delete FsxVolume#privileged_delete}
	PrivilegedDelete *string `field:"optional" json:"privilegedDelete" yaml:"privilegedDelete"`
	// Specifies the retention period of an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#retention_period FsxVolume#retention_period}
	RetentionPeriod *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriod `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Specifies the retention mode of an FSx for ONTAP SnapLock volume. After it is set, it can't be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#snaplock_type FsxVolume#snaplock_type}
	SnaplockType *string `field:"optional" json:"snaplockType" yaml:"snaplockType"`
	// Enables or disables volume-append mode on an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#volume_append_mode_enabled FsxVolume#volume_append_mode_enabled}
	VolumeAppendModeEnabled *string `field:"optional" json:"volumeAppendModeEnabled" yaml:"volumeAppendModeEnabled"`
}

