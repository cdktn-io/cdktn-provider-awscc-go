// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconvertjobtemplate


type MediaconvertJobTemplateHopDestinations struct {
	// Optional. A different relative priority for the job in the destination queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#priority MediaconvertJobTemplate#priority}
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Optional. The destination queue for queue hopping.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#queue MediaconvertJobTemplate#queue}
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
	// Required for queue hopping. Minimum wait time in minutes until the job can hop to the destination queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#wait_minutes MediaconvertJobTemplate#wait_minutes}
	WaitMinutes *float64 `field:"optional" json:"waitMinutes" yaml:"waitMinutes"`
}

