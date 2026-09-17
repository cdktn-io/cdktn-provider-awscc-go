// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ceanomalysubscription

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CeAnomalySubscriptionConfig struct {
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
	// The frequency at which anomaly reports are sent over email.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ce_anomaly_subscription#frequency CeAnomalySubscription#frequency}
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// A list of cost anomaly monitors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ce_anomaly_subscription#monitor_arn_list CeAnomalySubscription#monitor_arn_list}
	MonitorArnList *[]*string `field:"required" json:"monitorArnList" yaml:"monitorArnList"`
	// A list of subscriber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ce_anomaly_subscription#subscribers CeAnomalySubscription#subscribers}
	Subscribers interface{} `field:"required" json:"subscribers" yaml:"subscribers"`
	// The name of the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ce_anomaly_subscription#subscription_name CeAnomalySubscription#subscription_name}
	SubscriptionName *string `field:"required" json:"subscriptionName" yaml:"subscriptionName"`
	// Tags to assign to subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ce_anomaly_subscription#resource_tags CeAnomalySubscription#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// The dollar value that triggers a notification if the threshold is exceeded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ce_anomaly_subscription#threshold CeAnomalySubscription#threshold}
	Threshold *float64 `field:"optional" json:"threshold" yaml:"threshold"`
	// An Expression object in JSON String format used to specify the anomalies that you want to generate alerts for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/ce_anomaly_subscription#threshold_expression CeAnomalySubscription#threshold_expression}
	ThresholdExpression *string `field:"optional" json:"thresholdExpression" yaml:"thresholdExpression"`
}

