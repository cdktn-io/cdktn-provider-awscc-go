// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rumappmonitor


type RumAppMonitorDeobfuscationConfigurationJavaScriptSourceMaps struct {
	// The S3Uri of the bucket or folder that stores the source map files.
	//
	// It is required if status is ENABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#s3_uri RumAppMonitor#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// Specifies whether JavaScript error stack traces should be unminified for this app monitor.
	//
	// The default is for JavaScript error stack trace unminification to be DISABLED
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rum_app_monitor#status RumAppMonitor#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

