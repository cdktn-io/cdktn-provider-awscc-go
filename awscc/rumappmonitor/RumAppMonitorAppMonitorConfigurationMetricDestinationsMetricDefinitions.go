// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor


type RumAppMonitorAppMonitorConfigurationMetricDestinationsMetricDefinitions struct {
	// Use this field only if you are sending the metric to CloudWatch.
	//
	// This field is a map of field paths to dimension names. It defines the dimensions to associate with this metric in CloudWatch. For extended metrics, valid values for the entries in this field are the following:
	//
	// "metadata.pageId": "PageId"
	//
	// "metadata.browserName": "BrowserName"
	//
	// "metadata.deviceType": "DeviceType"
	//
	// "metadata.osName": "OSName"
	//
	// "metadata.countryCode": "CountryCode"
	//
	// "event_details.fileType": "FileType"
	//
	// All dimensions listed in this field must also be included in EventPattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#dimension_keys RumAppMonitor#dimension_keys}
	DimensionKeys *map[string]*string `field:"optional" json:"dimensionKeys" yaml:"dimensionKeys"`
	// The pattern that defines the metric, specified as a JSON object.
	//
	// RUM checks events that happen in a user's session against the pattern, and events that match the pattern are sent to the metric destination.
	//
	// When you define extended metrics, the metric definition is not valid if EventPattern is omitted.
	//
	// Example event patterns:
	//
	// '{ "event_type": ["com.amazon.rum.js_error_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], } }'
	//
	// '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Firefox" ] }, "event_details": { "duration": [{ "numeric": [ "<", 2000 ] }] } }'
	//
	// '{ "event_type": ["com.amazon.rum.performance_navigation_event"], "metadata": { "browserName": [ "Chrome", "Safari" ], "countryCode": [ "US" ] }, "event_details": { "duration": [{ "numeric": [ ">=", 2000, "<", 8000 ] }] } }'
	//
	// If the metrics destination' is CloudWatch and the event also matches a value in DimensionKeys, then the metric is published with the specified dimensions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#event_pattern RumAppMonitor#event_pattern}
	EventPattern *string `field:"optional" json:"eventPattern" yaml:"eventPattern"`
	// The name for the metric that is defined in this structure. For extended metrics, valid values are the following:.
	//
	// PerformanceNavigationDuration
	//
	// PerformanceResourceDuration
	//
	// NavigationSatisfiedTransaction
	//
	// NavigationToleratedTransaction
	//
	// NavigationFrustratedTransaction
	//
	// WebVitalsCumulativeLayoutShift
	//
	// WebVitalsFirstInputDelay
	//
	// WebVitalsLargestContentfulPaint
	//
	// WebVitalsInteractionToNextPaint
	//
	// JsErrorCount
	//
	// HttpErrorCount
	//
	// SessionCount
	//
	// PageViewCount
	//
	// Http4xxCount
	//
	// Http5xxCount
	//
	// SessionDuration
	//
	// PageViewCountPerSession
	//
	// JsErrorCountPerSession
	//
	// Http4xxCountPerSession
	//
	// Http5xxCountPerSession
	//
	// JsErrorCountPerPageView
	//
	// Http4xxCountPerPageView
	//
	// Http5xxCountPerPageView
	//
	// TimeOnPage
	//
	// ColdLaunchTime
	//
	// WarmLaunchTime
	//
	// CrashCount
	//
	// ANRCount
	//
	// AppHangCount
	//
	// ScreenLoadCount
	//
	// ScreenLoadTime
	//
	// NetworkLatency
	//
	// SpanPayloadSize
	//
	// LogEventPayloadSize
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#name RumAppMonitor#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace used by CloudWatch Metrics for the metric that is defined in this structure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#namespace RumAppMonitor#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The CloudWatch metric unit to use for this metric.
	//
	// If you omit this field, the metric is recorded with no unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#unit_label RumAppMonitor#unit_label}
	UnitLabel *string `field:"optional" json:"unitLabel" yaml:"unitLabel"`
	// The field within the event object that the metric value is sourced from.
	//
	// If you omit this field, a hardcoded value of 1 is pushed as the metric value. This is useful if you just want to count the number of events that the filter catches.
	//
	// If this metric is sent to Evidently, this field will be passed to Evidently raw and Evidently will handle data extraction from the event. Note: Evidently has been discontinued.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#value_key RumAppMonitor#value_key}
	ValueKey *string `field:"optional" json:"valueKey" yaml:"valueKey"`
}

