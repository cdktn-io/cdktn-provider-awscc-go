// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudwatchalarm


type CloudwatchAlarmEvaluationWindowWallClockWindow struct {
	// The time zone to use when the alarm aligns the evaluation window to clock boundaries.
	//
	// You can specify an IANA time zone name (for example, ``America/New_York``), a fixed UTC offset (for example, ``+05:30``), or an offset-prefixed identifier (for example, ``UTC+05:30``). The offset must be aligned to a multiple of 5 minutes. If you don't specify a time zone, CloudWatch uses ``UTC``.
	//  The time zone affects window alignment for all periods, including periods of one hour or shorter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudwatch_alarm#timezone CloudwatchAlarm#timezone}
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

