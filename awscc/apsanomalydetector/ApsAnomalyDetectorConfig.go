// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apsanomalydetector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApsAnomalyDetectorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The AnomalyDetector alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#alias ApsAnomalyDetector#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Determines the anomaly detector's algorithm and its configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#configuration ApsAnomalyDetector#configuration}
	Configuration *ApsAnomalyDetectorConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// Required to identify a specific APS Workspace associated with this Anomaly Detector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#workspace ApsAnomalyDetector#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// The AnomalyDetector period of detection and metric generation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#evaluation_interval_in_seconds ApsAnomalyDetector#evaluation_interval_in_seconds}
	EvaluationIntervalInSeconds *float64 `field:"optional" json:"evaluationIntervalInSeconds" yaml:"evaluationIntervalInSeconds"`
	// An array of key-value pairs to provide meta-data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#labels ApsAnomalyDetector#labels}
	Labels interface{} `field:"optional" json:"labels" yaml:"labels"`
	// The action to perform when running the expression returns no data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#missing_data_action ApsAnomalyDetector#missing_data_action}
	MissingDataAction *ApsAnomalyDetectorMissingDataAction `field:"optional" json:"missingDataAction" yaml:"missingDataAction"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/aps_anomaly_detector#tags ApsAnomalyDetector#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

