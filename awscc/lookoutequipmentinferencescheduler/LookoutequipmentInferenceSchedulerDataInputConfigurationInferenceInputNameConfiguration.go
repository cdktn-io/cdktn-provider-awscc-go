// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lookoutequipmentinferencescheduler


type LookoutequipmentInferenceSchedulerDataInputConfigurationInferenceInputNameConfiguration struct {
	// Indicates the delimiter character used between items in the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutequipment_inference_scheduler#component_timestamp_delimiter LookoutequipmentInferenceScheduler#component_timestamp_delimiter}
	ComponentTimestampDelimiter *string `field:"optional" json:"componentTimestampDelimiter" yaml:"componentTimestampDelimiter"`
	// The format of the timestamp, whether Epoch time, or standard, with or without hyphens (-).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lookoutequipment_inference_scheduler#timestamp_format LookoutequipmentInferenceScheduler#timestamp_format}
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

