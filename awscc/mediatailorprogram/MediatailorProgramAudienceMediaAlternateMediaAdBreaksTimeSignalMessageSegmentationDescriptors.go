// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramAudienceMediaAlternateMediaAdBreaksTimeSignalMessageSegmentationDescriptors struct {
	// The Event Identifier to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#segmentation_event_id MediatailorProgram#segmentation_event_id}
	SegmentationEventId *float64 `field:"optional" json:"segmentationEventId" yaml:"segmentationEventId"`
	// The Type Identifier to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#segmentation_type_id MediatailorProgram#segmentation_type_id}
	SegmentationTypeId *float64 `field:"optional" json:"segmentationTypeId" yaml:"segmentationTypeId"`
	// The Upid to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#segmentation_upid MediatailorProgram#segmentation_upid}
	SegmentationUpid *string `field:"optional" json:"segmentationUpid" yaml:"segmentationUpid"`
	// The Upid Type to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#segmentation_upid_type MediatailorProgram#segmentation_upid_type}
	SegmentationUpidType *float64 `field:"optional" json:"segmentationUpidType" yaml:"segmentationUpidType"`
	// The segment number to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#segment_num MediatailorProgram#segment_num}
	SegmentNum *float64 `field:"optional" json:"segmentNum" yaml:"segmentNum"`
	// The number of segments expected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#segments_expected MediatailorProgram#segments_expected}
	SegmentsExpected *float64 `field:"optional" json:"segmentsExpected" yaml:"segmentsExpected"`
	// The sub-segment number to assign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#sub_segment_num MediatailorProgram#sub_segment_num}
	SubSegmentNum *float64 `field:"optional" json:"subSegmentNum" yaml:"subSegmentNum"`
	// The number of sub-segments expected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#sub_segments_expected MediatailorProgram#sub_segments_expected}
	SubSegmentsExpected *float64 `field:"optional" json:"subSegmentsExpected" yaml:"subSegmentsExpected"`
}

