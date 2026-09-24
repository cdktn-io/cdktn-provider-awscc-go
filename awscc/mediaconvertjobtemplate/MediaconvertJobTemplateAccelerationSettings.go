// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconvertjobtemplate


type MediaconvertJobTemplateAccelerationSettings struct {
	// Specify the conditions when the service will run your job with accelerated transcoding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconvert_job_template#mode MediaconvertJobTemplate#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

