// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramAdBreaks struct {
	// Defines a list of key/value pairs that MediaTailor generates within the EXT-X-ASSET tag for SCTE35_ENHANCED output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#ad_break_metadata MediatailorProgram#ad_break_metadata}
	AdBreakMetadata interface{} `field:"optional" json:"adBreakMetadata" yaml:"adBreakMetadata"`
	// The SCTE-35 ad insertion type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#message_type MediatailorProgram#message_type}
	MessageType *string `field:"optional" json:"messageType" yaml:"messageType"`
	// How long (in milliseconds) after the beginning of the program that an ad starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#offset_millis MediatailorProgram#offset_millis}
	OffsetMillis *float64 `field:"optional" json:"offsetMillis" yaml:"offsetMillis"`
	// Slate VOD source configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#slate MediatailorProgram#slate}
	Slate *MediatailorProgramAdBreaksSlate `field:"optional" json:"slate" yaml:"slate"`
	// Splice insert message configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#splice_insert_message MediatailorProgram#splice_insert_message}
	SpliceInsertMessage *MediatailorProgramAdBreaksSpliceInsertMessage `field:"optional" json:"spliceInsertMessage" yaml:"spliceInsertMessage"`
	// The SCTE-35 time_signal message configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/mediatailor_program#time_signal_message MediatailorProgram#time_signal_message}
	TimeSignalMessage *MediatailorProgramAdBreaksTimeSignalMessage `field:"optional" json:"timeSignalMessage" yaml:"timeSignalMessage"`
}

