// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightanalysis

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightAnalysisConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#analysis_id QuicksightAnalysis#analysis_id}.
	AnalysisId *string `field:"required" json:"analysisId" yaml:"analysisId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#aws_account_id QuicksightAnalysis#aws_account_id}.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// <p>The source entity of an analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#source_entity QuicksightAnalysis#source_entity}
	SourceEntity *QuicksightAnalysisSourceEntity `field:"required" json:"sourceEntity" yaml:"sourceEntity"`
	// <p>Errors associated with the analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#errors QuicksightAnalysis#errors}
	Errors interface{} `field:"optional" json:"errors" yaml:"errors"`
	// <p>The descriptive name of the analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#name QuicksightAnalysis#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// <p>A list of QuickSight parameters and the list's override values.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#parameters QuicksightAnalysis#parameters}
	Parameters *QuicksightAnalysisParameters `field:"optional" json:"parameters" yaml:"parameters"`
	// <p>A structure that describes the principals and the resource-level permissions on an             analysis.
	//
	// You can use the <code>Permissions</code> structure to grant permissions by
	//             providing a list of AWS Identity and Access Management (IAM) action information for each
	//             principal listed by Amazon Resource Name (ARN). </p>
	//
	//         <p>To specify no permissions, omit <code>Permissions</code>.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#permissions QuicksightAnalysis#permissions}
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// <p>Contains a map of the key-value pairs for the resource tag or tags assigned to the             analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#tags QuicksightAnalysis#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// <p>The ARN of the theme of the analysis.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_analysis#theme_arn QuicksightAnalysis#theme_arn}
	ThemeArn *string `field:"optional" json:"themeArn" yaml:"themeArn"`
}

