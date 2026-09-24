// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsailinstance


type LightsailInstanceHardwareDisks struct {
	// Instance attached to the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lightsail_instance#attached_to LightsailInstance#attached_to}
	AttachedTo *string `field:"optional" json:"attachedTo" yaml:"attachedTo"`
	// Attachment state of the disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lightsail_instance#attachment_state LightsailInstance#attachment_state}
	AttachmentState *string `field:"optional" json:"attachmentState" yaml:"attachmentState"`
	// The names to use for your new Lightsail disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lightsail_instance#disk_name LightsailInstance#disk_name}
	DiskName *string `field:"optional" json:"diskName" yaml:"diskName"`
	// IOPS of disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lightsail_instance#iops LightsailInstance#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Is the Attached disk is the system disk of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lightsail_instance#is_system_disk LightsailInstance#is_system_disk}
	IsSystemDisk interface{} `field:"optional" json:"isSystemDisk" yaml:"isSystemDisk"`
	// Path of the disk attached to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lightsail_instance#path LightsailInstance#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Size of the disk attached to the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/lightsail_instance#size_in_gb LightsailInstance#size_in_gb}
	SizeInGb *string `field:"optional" json:"sizeInGb" yaml:"sizeInGb"`
}

