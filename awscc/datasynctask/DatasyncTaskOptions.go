// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynctask


type DatasyncTaskOptions struct {
	// A file metadata value that shows the last time a file was accessed (that is, when the file was read or written to).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#atime DatasyncTask#atime}
	Atime *string `field:"optional" json:"atime" yaml:"atime"`
	// A value that limits the bandwidth used by AWS DataSync.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#bytes_per_second DatasyncTask#bytes_per_second}
	BytesPerSecond *float64 `field:"optional" json:"bytesPerSecond" yaml:"bytesPerSecond"`
	// The group ID (GID) of the file's owners.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#gid DatasyncTask#gid}
	Gid *string `field:"optional" json:"gid" yaml:"gid"`
	// A value that determines the types of logs that DataSync publishes to a log stream in the Amazon CloudWatch log group that you provide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#log_level DatasyncTask#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// A value that indicates the last time that a file was modified (that is, a file was written to) before the PREPARING phase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#mtime DatasyncTask#mtime}
	Mtime *string `field:"optional" json:"mtime" yaml:"mtime"`
	// A value that determines whether object tags should be read from the source object store and written to the destination object store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#object_tags DatasyncTask#object_tags}
	ObjectTags *string `field:"optional" json:"objectTags" yaml:"objectTags"`
	// A value that determines whether files at the destination should be overwritten or preserved when copying files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#overwrite_mode DatasyncTask#overwrite_mode}
	OverwriteMode *string `field:"optional" json:"overwriteMode" yaml:"overwriteMode"`
	// A value that determines which users or groups can access a file for a specific purpose such as reading, writing, or execution of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#posix_permissions DatasyncTask#posix_permissions}
	PosixPermissions *string `field:"optional" json:"posixPermissions" yaml:"posixPermissions"`
	// A value that specifies whether files in the destination that don't exist in the source file system should be preserved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#preserve_deleted_files DatasyncTask#preserve_deleted_files}
	PreserveDeletedFiles *string `field:"optional" json:"preserveDeletedFiles" yaml:"preserveDeletedFiles"`
	// A value that determines whether AWS DataSync should preserve the metadata of block and character devices in the source file system, and recreate the files with that device name and metadata on the destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#preserve_devices DatasyncTask#preserve_devices}
	PreserveDevices *string `field:"optional" json:"preserveDevices" yaml:"preserveDevices"`
	// A value that determines which components of the SMB security descriptor are copied during transfer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#security_descriptor_copy_flags DatasyncTask#security_descriptor_copy_flags}
	SecurityDescriptorCopyFlags *string `field:"optional" json:"securityDescriptorCopyFlags" yaml:"securityDescriptorCopyFlags"`
	// A value that determines whether tasks should be queued before executing the tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#task_queueing DatasyncTask#task_queueing}
	TaskQueueing *string `field:"optional" json:"taskQueueing" yaml:"taskQueueing"`
	// A value that determines whether DataSync transfers only the data and metadata that differ between the source and the destination location, or whether DataSync transfers all the content from the source, without comparing to the destination location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#transfer_mode DatasyncTask#transfer_mode}
	TransferMode *string `field:"optional" json:"transferMode" yaml:"transferMode"`
	// The user ID (UID) of the file's owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#uid DatasyncTask#uid}
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
	// A value that determines whether a data integrity verification should be performed at the end of a task execution after all data and metadata have been transferred.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_task#verify_mode DatasyncTask#verify_mode}
	VerifyMode *string `field:"optional" json:"verifyMode" yaml:"verifyMode"`
}

