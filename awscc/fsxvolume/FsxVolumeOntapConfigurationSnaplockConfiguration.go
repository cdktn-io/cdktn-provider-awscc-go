// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationSnaplockConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#audit_log_volume FsxVolume#audit_log_volume}.
	AuditLogVolume *string `field:"optional" json:"auditLogVolume" yaml:"auditLogVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#autocommit_period FsxVolume#autocommit_period}.
	AutocommitPeriod *FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriod `field:"optional" json:"autocommitPeriod" yaml:"autocommitPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#privileged_delete FsxVolume#privileged_delete}.
	PrivilegedDelete *string `field:"optional" json:"privilegedDelete" yaml:"privilegedDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#retention_period FsxVolume#retention_period}.
	RetentionPeriod *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriod `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#snaplock_type FsxVolume#snaplock_type}.
	SnaplockType *string `field:"optional" json:"snaplockType" yaml:"snaplockType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#volume_append_mode_enabled FsxVolume#volume_append_mode_enabled}.
	VolumeAppendModeEnabled *string `field:"optional" json:"volumeAppendModeEnabled" yaml:"volumeAppendModeEnabled"`
}

