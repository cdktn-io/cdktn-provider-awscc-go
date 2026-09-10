// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor


type RumAppMonitorAppMonitorConfiguration struct {
	// If you set this to true, the RUM web client sets two cookies, a session cookie and a user cookie.
	//
	// The cookies allow the RUM web client to collect data relating to the number of users an application has and the behavior of the application across a sequence of events. Cookies are stored in the top-level domain of the current page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#allow_cookies RumAppMonitor#allow_cookies}
	AllowCookies interface{} `field:"optional" json:"allowCookies" yaml:"allowCookies"`
	// If you set this to true, RUM enables xray tracing for the user sessions that RUM samples.
	//
	// RUM adds an xray trace header to allowed HTTP requests. It also records an xray segment for allowed HTTP requests. You can see traces and segments from these user sessions in the xray console and the CW ServiceLens console.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#enable_x_ray RumAppMonitor#enable_x_ray}
	EnableXRay interface{} `field:"optional" json:"enableXRay" yaml:"enableXRay"`
	// A list of URLs in your website or application to exclude from RUM data collection.
	//
	// You can't include both ExcludedPages and IncludedPages in the same operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#excluded_pages RumAppMonitor#excluded_pages}
	ExcludedPages *[]*string `field:"optional" json:"excludedPages" yaml:"excludedPages"`
	// A list of pages in the RUM console that are to be displayed with a favorite icon.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#favorite_pages RumAppMonitor#favorite_pages}
	FavoritePages *[]*string `field:"optional" json:"favoritePages" yaml:"favoritePages"`
	// The ARN of the guest IAM role that is attached to the identity pool that is used to authorize the sending of data to RUM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#guest_role_arn RumAppMonitor#guest_role_arn}
	GuestRoleArn *string `field:"optional" json:"guestRoleArn" yaml:"guestRoleArn"`
	// The ID of the identity pool that is used to authorize the sending of data to RUM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#identity_pool_id RumAppMonitor#identity_pool_id}
	IdentityPoolId *string `field:"optional" json:"identityPoolId" yaml:"identityPoolId"`
	// If this app monitor is to collect data from only certain pages in your application, this structure lists those pages.
	//
	// You can't include both ExcludedPages and IncludedPages in the same operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#included_pages RumAppMonitor#included_pages}
	IncludedPages *[]*string `field:"optional" json:"includedPages" yaml:"includedPages"`
	// An array of structures which define the destinations and the metrics that you want to send.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#metric_destinations RumAppMonitor#metric_destinations}
	MetricDestinations interface{} `field:"optional" json:"metricDestinations" yaml:"metricDestinations"`
	// Specifies the percentage of user sessions to use for RUM data collection.
	//
	// Choosing a higher percentage gives you more data but also incurs more costs. The number you specify is the percentage of user sessions that will be used. If you omit this parameter, the default of 10 is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#session_sample_rate RumAppMonitor#session_sample_rate}
	SessionSampleRate *float64 `field:"optional" json:"sessionSampleRate" yaml:"sessionSampleRate"`
	// An array that lists the types of telemetry data that this app monitor is to collect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rum_app_monitor#telemetries RumAppMonitor#telemetries}
	Telemetries *[]*string `field:"optional" json:"telemetries" yaml:"telemetries"`
}

