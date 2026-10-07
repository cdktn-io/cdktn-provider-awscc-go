// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediatailorprogram


type MediatailorProgramAdBreaksSpliceInsertMessage struct {
	// This is written to splice_insert.avail_num.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#avail_num MediatailorProgram#avail_num}
	AvailNum *float64 `field:"optional" json:"availNum" yaml:"availNum"`
	// This is written to splice_insert.avails_expected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#avails_expected MediatailorProgram#avails_expected}
	AvailsExpected *float64 `field:"optional" json:"availsExpected" yaml:"availsExpected"`
	// This is written to splice_insert.splice_event_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#splice_event_id MediatailorProgram#splice_event_id}
	SpliceEventId *float64 `field:"optional" json:"spliceEventId" yaml:"spliceEventId"`
	// This is written to splice_insert.unique_program_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/mediatailor_program#unique_program_id MediatailorProgram#unique_program_id}
	UniqueProgramId *float64 `field:"optional" json:"uniqueProgramId" yaml:"uniqueProgramId"`
}

