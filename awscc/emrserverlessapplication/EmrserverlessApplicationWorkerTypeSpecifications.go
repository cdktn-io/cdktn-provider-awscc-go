// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package emrserverlessapplication


type EmrserverlessApplicationWorkerTypeSpecifications struct {
	// The image configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/emrserverless_application#image_configuration EmrserverlessApplication#image_configuration}
	ImageConfiguration *EmrserverlessApplicationWorkerTypeSpecificationsImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

