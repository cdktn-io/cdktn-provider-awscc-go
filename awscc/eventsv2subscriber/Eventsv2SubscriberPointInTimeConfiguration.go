// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsv2subscriber


type Eventsv2SubscriberPointInTimeConfiguration struct {
	// An optional time to stop delivering events at, in seconds since the Unix epoch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#end_point Eventsv2Subscriber#end_point}
	EndPoint *float64 `field:"optional" json:"endPoint" yaml:"endPoint"`
	// Where to start: HORIZON starts from the earliest available event; TIMESTAMP starts from the StartingPoint timestamp.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#point_type Eventsv2Subscriber#point_type}
	PointType *string `field:"optional" json:"pointType" yaml:"pointType"`
	// The time to start delivering events from, in seconds since the Unix epoch. Required when PointType is TIMESTAMP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/eventsv2_subscriber#starting_point Eventsv2Subscriber#starting_point}
	StartingPoint *float64 `field:"optional" json:"startingPoint" yaml:"startingPoint"`
}

