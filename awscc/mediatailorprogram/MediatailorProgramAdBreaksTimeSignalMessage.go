// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramAdBreaksTimeSignalMessage struct {
	// The configurations for the SCTE-35 segmentation_descriptor message(s).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#segmentation_descriptors MediatailorProgram#segmentation_descriptors}
	SegmentationDescriptors interface{} `field:"optional" json:"segmentationDescriptors" yaml:"segmentationDescriptors"`
}

