// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package securityhubinsight


type SecurityhubInsightFilters struct {
	// The AWS account ID in which a finding is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#aws_account_id SecurityhubInsight#aws_account_id}
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The name of the AWS account in which a finding is generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#aws_account_name SecurityhubInsight#aws_account_name}
	AwsAccountName interface{} `field:"optional" json:"awsAccountName" yaml:"awsAccountName"`
	// The name of the findings provider (company) that owns the solution (product) that generates findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#company_name SecurityhubInsight#company_name}
	CompanyName interface{} `field:"optional" json:"companyName" yaml:"companyName"`
	// The unique identifier of a standard in which a control is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#compliance_associated_standards_id SecurityhubInsight#compliance_associated_standards_id}
	ComplianceAssociatedStandardsId interface{} `field:"optional" json:"complianceAssociatedStandardsId" yaml:"complianceAssociatedStandardsId"`
	// The unique identifier of a control across standards.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#compliance_security_control_id SecurityhubInsight#compliance_security_control_id}
	ComplianceSecurityControlId interface{} `field:"optional" json:"complianceSecurityControlId" yaml:"complianceSecurityControlId"`
	// The name of a security control parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#compliance_security_control_parameters_name SecurityhubInsight#compliance_security_control_parameters_name}
	ComplianceSecurityControlParametersName interface{} `field:"optional" json:"complianceSecurityControlParametersName" yaml:"complianceSecurityControlParametersName"`
	// The current value of a security control parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#compliance_security_control_parameters_value SecurityhubInsight#compliance_security_control_parameters_value}
	ComplianceSecurityControlParametersValue interface{} `field:"optional" json:"complianceSecurityControlParametersValue" yaml:"complianceSecurityControlParametersValue"`
	// Exclusive to findings that are generated as the result of a check run against a specific rule in a supported standard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#compliance_status SecurityhubInsight#compliance_status}
	ComplianceStatus interface{} `field:"optional" json:"complianceStatus" yaml:"complianceStatus"`
	// A finding's confidence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#confidence SecurityhubInsight#confidence}
	Confidence interface{} `field:"optional" json:"confidence" yaml:"confidence"`
	// An ISO8601-formatted timestamp that indicates when the security findings provider captured the potential security issue that a finding captured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#created_at SecurityhubInsight#created_at}
	CreatedAt interface{} `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The level of importance assigned to the resources associated with the finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#criticality SecurityhubInsight#criticality}
	Criticality interface{} `field:"optional" json:"criticality" yaml:"criticality"`
	// A finding's description.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#description SecurityhubInsight#description}
	Description interface{} `field:"optional" json:"description" yaml:"description"`
	// The finding provider value for the finding confidence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#finding_provider_fields_confidence SecurityhubInsight#finding_provider_fields_confidence}
	FindingProviderFieldsConfidence interface{} `field:"optional" json:"findingProviderFieldsConfidence" yaml:"findingProviderFieldsConfidence"`
	// The finding provider value for the level of importance assigned to the resources associated with the findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#finding_provider_fields_criticality SecurityhubInsight#finding_provider_fields_criticality}
	FindingProviderFieldsCriticality interface{} `field:"optional" json:"findingProviderFieldsCriticality" yaml:"findingProviderFieldsCriticality"`
	// The finding identifier of a related finding that is identified by the finding provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#finding_provider_fields_related_findings_id SecurityhubInsight#finding_provider_fields_related_findings_id}
	FindingProviderFieldsRelatedFindingsId interface{} `field:"optional" json:"findingProviderFieldsRelatedFindingsId" yaml:"findingProviderFieldsRelatedFindingsId"`
	// The ARN of the solution that generated a related finding that is identified by the finding provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#finding_provider_fields_related_findings_product_arn SecurityhubInsight#finding_provider_fields_related_findings_product_arn}
	FindingProviderFieldsRelatedFindingsProductArn interface{} `field:"optional" json:"findingProviderFieldsRelatedFindingsProductArn" yaml:"findingProviderFieldsRelatedFindingsProductArn"`
	// The finding provider value for the severity label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#finding_provider_fields_severity_label SecurityhubInsight#finding_provider_fields_severity_label}
	FindingProviderFieldsSeverityLabel interface{} `field:"optional" json:"findingProviderFieldsSeverityLabel" yaml:"findingProviderFieldsSeverityLabel"`
	// The finding provider's original value for the severity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#finding_provider_fields_severity_original SecurityhubInsight#finding_provider_fields_severity_original}
	FindingProviderFieldsSeverityOriginal interface{} `field:"optional" json:"findingProviderFieldsSeverityOriginal" yaml:"findingProviderFieldsSeverityOriginal"`
	// One or more finding types that the finding provider assigned to the finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#finding_provider_fields_types SecurityhubInsight#finding_provider_fields_types}
	FindingProviderFieldsTypes interface{} `field:"optional" json:"findingProviderFieldsTypes" yaml:"findingProviderFieldsTypes"`
	// An ISO8601-formatted timestamp that indicates when the security findings provider first observed the potential security issue that a finding captured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#first_observed_at SecurityhubInsight#first_observed_at}
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// The identifier for the solution-specific component (a discrete unit of logic) that generated a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#generator_id SecurityhubInsight#generator_id}
	GeneratorId interface{} `field:"optional" json:"generatorId" yaml:"generatorId"`
	// The security findings provider-specific identifier for a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#id SecurityhubInsight#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id interface{} `field:"optional" json:"id" yaml:"id"`
	// A keyword for a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#keyword SecurityhubInsight#keyword}
	Keyword interface{} `field:"optional" json:"keyword" yaml:"keyword"`
	// An ISO8601-formatted timestamp that indicates when the security findings provider most recently observed the potential security issue that a finding captured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#last_observed_at SecurityhubInsight#last_observed_at}
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// The name of the malware that was observed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#malware_name SecurityhubInsight#malware_name}
	MalwareName interface{} `field:"optional" json:"malwareName" yaml:"malwareName"`
	// The filesystem path of the malware that was observed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#malware_path SecurityhubInsight#malware_path}
	MalwarePath interface{} `field:"optional" json:"malwarePath" yaml:"malwarePath"`
	// The state of the malware that was observed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#malware_state SecurityhubInsight#malware_state}
	MalwareState interface{} `field:"optional" json:"malwareState" yaml:"malwareState"`
	// The type of the malware that was observed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#malware_type SecurityhubInsight#malware_type}
	MalwareType interface{} `field:"optional" json:"malwareType" yaml:"malwareType"`
	// The destination domain of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_destination_domain SecurityhubInsight#network_destination_domain}
	NetworkDestinationDomain interface{} `field:"optional" json:"networkDestinationDomain" yaml:"networkDestinationDomain"`
	// The destination IPv4 address of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_destination_ip_v4 SecurityhubInsight#network_destination_ip_v4}
	NetworkDestinationIpV4 interface{} `field:"optional" json:"networkDestinationIpV4" yaml:"networkDestinationIpV4"`
	// The destination IPv6 address of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_destination_ip_v6 SecurityhubInsight#network_destination_ip_v6}
	NetworkDestinationIpV6 interface{} `field:"optional" json:"networkDestinationIpV6" yaml:"networkDestinationIpV6"`
	// The destination port of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_destination_port SecurityhubInsight#network_destination_port}
	NetworkDestinationPort interface{} `field:"optional" json:"networkDestinationPort" yaml:"networkDestinationPort"`
	// Indicates the direction of network traffic associated with a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_direction SecurityhubInsight#network_direction}
	NetworkDirection interface{} `field:"optional" json:"networkDirection" yaml:"networkDirection"`
	// The protocol of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_protocol SecurityhubInsight#network_protocol}
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// The source domain of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_source_domain SecurityhubInsight#network_source_domain}
	NetworkSourceDomain interface{} `field:"optional" json:"networkSourceDomain" yaml:"networkSourceDomain"`
	// The source IPv4 address of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_source_ip_v4 SecurityhubInsight#network_source_ip_v4}
	NetworkSourceIpV4 interface{} `field:"optional" json:"networkSourceIpV4" yaml:"networkSourceIpV4"`
	// The source IPv6 address of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_source_ip_v6 SecurityhubInsight#network_source_ip_v6}
	NetworkSourceIpV6 interface{} `field:"optional" json:"networkSourceIpV6" yaml:"networkSourceIpV6"`
	// The source media access control (MAC) address of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_source_mac SecurityhubInsight#network_source_mac}
	NetworkSourceMac interface{} `field:"optional" json:"networkSourceMac" yaml:"networkSourceMac"`
	// The source port of network-related information about a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#network_source_port SecurityhubInsight#network_source_port}
	NetworkSourcePort interface{} `field:"optional" json:"networkSourcePort" yaml:"networkSourcePort"`
	// The text of a note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#note_text SecurityhubInsight#note_text}
	NoteText interface{} `field:"optional" json:"noteText" yaml:"noteText"`
	// The timestamp of when the note was updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#note_updated_at SecurityhubInsight#note_updated_at}
	NoteUpdatedAt interface{} `field:"optional" json:"noteUpdatedAt" yaml:"noteUpdatedAt"`
	// The principal that created a note.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#note_updated_by SecurityhubInsight#note_updated_by}
	NoteUpdatedBy interface{} `field:"optional" json:"noteUpdatedBy" yaml:"noteUpdatedBy"`
	// A timestamp that identifies when the process was launched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#process_launched_at SecurityhubInsight#process_launched_at}
	ProcessLaunchedAt interface{} `field:"optional" json:"processLaunchedAt" yaml:"processLaunchedAt"`
	// The name of the process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#process_name SecurityhubInsight#process_name}
	ProcessName interface{} `field:"optional" json:"processName" yaml:"processName"`
	// The parent process ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#process_parent_pid SecurityhubInsight#process_parent_pid}
	ProcessParentPid interface{} `field:"optional" json:"processParentPid" yaml:"processParentPid"`
	// The path to the process executable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#process_path SecurityhubInsight#process_path}
	ProcessPath interface{} `field:"optional" json:"processPath" yaml:"processPath"`
	// The process ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#process_pid SecurityhubInsight#process_pid}
	ProcessPid interface{} `field:"optional" json:"processPid" yaml:"processPid"`
	// A timestamp that identifies when the process was terminated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#process_terminated_at SecurityhubInsight#process_terminated_at}
	ProcessTerminatedAt interface{} `field:"optional" json:"processTerminatedAt" yaml:"processTerminatedAt"`
	// The ARN generated by Security Hub that uniquely identifies a third-party company (security findings provider) after this provider's product (solution that generates findings) is registered with Security Hub.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#product_arn SecurityhubInsight#product_arn}
	ProductArn interface{} `field:"optional" json:"productArn" yaml:"productArn"`
	// A data type where security findings providers can include additional solution-specific details that aren't part of the defined AwsSecurityFinding format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#product_fields SecurityhubInsight#product_fields}
	ProductFields interface{} `field:"optional" json:"productFields" yaml:"productFields"`
	// The name of the solution (product) that generates findings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#product_name SecurityhubInsight#product_name}
	ProductName interface{} `field:"optional" json:"productName" yaml:"productName"`
	// The recommendation of what to do about the issue described in a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#recommendation_text SecurityhubInsight#recommendation_text}
	RecommendationText interface{} `field:"optional" json:"recommendationText" yaml:"recommendationText"`
	// The updated record state for the finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#record_state SecurityhubInsight#record_state}
	RecordState interface{} `field:"optional" json:"recordState" yaml:"recordState"`
	// The Region from which the finding was generated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#region SecurityhubInsight#region}
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// The solution-generated identifier for a related finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#related_findings_id SecurityhubInsight#related_findings_id}
	RelatedFindingsId interface{} `field:"optional" json:"relatedFindingsId" yaml:"relatedFindingsId"`
	// The ARN of the solution that generated a related finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#related_findings_product_arn SecurityhubInsight#related_findings_product_arn}
	RelatedFindingsProductArn interface{} `field:"optional" json:"relatedFindingsProductArn" yaml:"relatedFindingsProductArn"`
	// The ARN of the application that is related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_application_arn SecurityhubInsight#resource_application_arn}
	ResourceApplicationArn interface{} `field:"optional" json:"resourceApplicationArn" yaml:"resourceApplicationArn"`
	// The name of the application that is related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_application_name SecurityhubInsight#resource_application_name}
	ResourceApplicationName interface{} `field:"optional" json:"resourceApplicationName" yaml:"resourceApplicationName"`
	// The IAM profile ARN of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_iam_instance_profile_arn SecurityhubInsight#resource_aws_ec_2_instance_iam_instance_profile_arn}
	ResourceAwsEc2InstanceIamInstanceProfileArn interface{} `field:"optional" json:"resourceAwsEc2InstanceIamInstanceProfileArn" yaml:"resourceAwsEc2InstanceIamInstanceProfileArn"`
	// The Amazon Machine Image (AMI) ID of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_image_id SecurityhubInsight#resource_aws_ec_2_instance_image_id}
	ResourceAwsEc2InstanceImageId interface{} `field:"optional" json:"resourceAwsEc2InstanceImageId" yaml:"resourceAwsEc2InstanceImageId"`
	// The IPv4 addresses associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_ip_v4_addresses SecurityhubInsight#resource_aws_ec_2_instance_ip_v4_addresses}
	ResourceAwsEc2InstanceIpV4Addresses interface{} `field:"optional" json:"resourceAwsEc2InstanceIpV4Addresses" yaml:"resourceAwsEc2InstanceIpV4Addresses"`
	// The IPv6 addresses associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_ip_v6_addresses SecurityhubInsight#resource_aws_ec_2_instance_ip_v6_addresses}
	ResourceAwsEc2InstanceIpV6Addresses interface{} `field:"optional" json:"resourceAwsEc2InstanceIpV6Addresses" yaml:"resourceAwsEc2InstanceIpV6Addresses"`
	// The key name associated with the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_key_name SecurityhubInsight#resource_aws_ec_2_instance_key_name}
	ResourceAwsEc2InstanceKeyName interface{} `field:"optional" json:"resourceAwsEc2InstanceKeyName" yaml:"resourceAwsEc2InstanceKeyName"`
	// The date and time the instance was launched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_launched_at SecurityhubInsight#resource_aws_ec_2_instance_launched_at}
	ResourceAwsEc2InstanceLaunchedAt interface{} `field:"optional" json:"resourceAwsEc2InstanceLaunchedAt" yaml:"resourceAwsEc2InstanceLaunchedAt"`
	// The identifier of the subnet that the instance was launched in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_subnet_id SecurityhubInsight#resource_aws_ec_2_instance_subnet_id}
	ResourceAwsEc2InstanceSubnetId interface{} `field:"optional" json:"resourceAwsEc2InstanceSubnetId" yaml:"resourceAwsEc2InstanceSubnetId"`
	// The instance type of the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_type SecurityhubInsight#resource_aws_ec_2_instance_type}
	ResourceAwsEc2InstanceType interface{} `field:"optional" json:"resourceAwsEc2InstanceType" yaml:"resourceAwsEc2InstanceType"`
	// The identifier of the VPC that the instance was launched in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_ec_2_instance_vpc_id SecurityhubInsight#resource_aws_ec_2_instance_vpc_id}
	ResourceAwsEc2InstanceVpcId interface{} `field:"optional" json:"resourceAwsEc2InstanceVpcId" yaml:"resourceAwsEc2InstanceVpcId"`
	// The creation date/time of the IAM access key related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_iam_access_key_created_at SecurityhubInsight#resource_aws_iam_access_key_created_at}
	ResourceAwsIamAccessKeyCreatedAt interface{} `field:"optional" json:"resourceAwsIamAccessKeyCreatedAt" yaml:"resourceAwsIamAccessKeyCreatedAt"`
	// The name of the principal that is associated with an IAM access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_iam_access_key_principal_name SecurityhubInsight#resource_aws_iam_access_key_principal_name}
	ResourceAwsIamAccessKeyPrincipalName interface{} `field:"optional" json:"resourceAwsIamAccessKeyPrincipalName" yaml:"resourceAwsIamAccessKeyPrincipalName"`
	// The status of the IAM access key related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_iam_access_key_status SecurityhubInsight#resource_aws_iam_access_key_status}
	ResourceAwsIamAccessKeyStatus interface{} `field:"optional" json:"resourceAwsIamAccessKeyStatus" yaml:"resourceAwsIamAccessKeyStatus"`
	// The user associated with the IAM access key related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_iam_access_key_user_name SecurityhubInsight#resource_aws_iam_access_key_user_name}
	ResourceAwsIamAccessKeyUserName interface{} `field:"optional" json:"resourceAwsIamAccessKeyUserName" yaml:"resourceAwsIamAccessKeyUserName"`
	// The name of an IAM user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_iam_user_user_name SecurityhubInsight#resource_aws_iam_user_user_name}
	ResourceAwsIamUserUserName interface{} `field:"optional" json:"resourceAwsIamUserUserName" yaml:"resourceAwsIamUserUserName"`
	// The canonical user ID of the owner of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_s3_bucket_owner_id SecurityhubInsight#resource_aws_s3_bucket_owner_id}
	ResourceAwsS3BucketOwnerId interface{} `field:"optional" json:"resourceAwsS3BucketOwnerId" yaml:"resourceAwsS3BucketOwnerId"`
	// The display name of the owner of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_aws_s3_bucket_owner_name SecurityhubInsight#resource_aws_s3_bucket_owner_name}
	ResourceAwsS3BucketOwnerName interface{} `field:"optional" json:"resourceAwsS3BucketOwnerName" yaml:"resourceAwsS3BucketOwnerName"`
	// The identifier of the image related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_container_image_id SecurityhubInsight#resource_container_image_id}
	ResourceContainerImageId interface{} `field:"optional" json:"resourceContainerImageId" yaml:"resourceContainerImageId"`
	// The name of the image related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_container_image_name SecurityhubInsight#resource_container_image_name}
	ResourceContainerImageName interface{} `field:"optional" json:"resourceContainerImageName" yaml:"resourceContainerImageName"`
	// A timestamp that identifies when the container was started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_container_launched_at SecurityhubInsight#resource_container_launched_at}
	ResourceContainerLaunchedAt interface{} `field:"optional" json:"resourceContainerLaunchedAt" yaml:"resourceContainerLaunchedAt"`
	// The name of the container related to a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_container_name SecurityhubInsight#resource_container_name}
	ResourceContainerName interface{} `field:"optional" json:"resourceContainerName" yaml:"resourceContainerName"`
	// The details of a resource that doesn't have a specific subfield for the resource type defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_details_other SecurityhubInsight#resource_details_other}
	ResourceDetailsOther interface{} `field:"optional" json:"resourceDetailsOther" yaml:"resourceDetailsOther"`
	// The canonical identifier for the given resource type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_id SecurityhubInsight#resource_id}
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The canonical AWS partition name that the Region is assigned to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_partition SecurityhubInsight#resource_partition}
	ResourcePartition interface{} `field:"optional" json:"resourcePartition" yaml:"resourcePartition"`
	// The canonical AWS external Region name where this resource is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_region SecurityhubInsight#resource_region}
	ResourceRegion interface{} `field:"optional" json:"resourceRegion" yaml:"resourceRegion"`
	// A list of AWS tags associated with a resource at the time the finding was processed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_tags SecurityhubInsight#resource_tags}
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// Specifies the type of the resource that details are provided for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#resource_type SecurityhubInsight#resource_type}
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Indicates whether or not sample findings are included in the filter results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#sample SecurityhubInsight#sample}
	Sample interface{} `field:"optional" json:"sample" yaml:"sample"`
	// The label of a finding's severity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#severity_label SecurityhubInsight#severity_label}
	SeverityLabel interface{} `field:"optional" json:"severityLabel" yaml:"severityLabel"`
	// The normalized severity of a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#severity_normalized SecurityhubInsight#severity_normalized}
	SeverityNormalized interface{} `field:"optional" json:"severityNormalized" yaml:"severityNormalized"`
	// The native severity as defined by the security findings provider's solution that generated the finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#severity_product SecurityhubInsight#severity_product}
	SeverityProduct interface{} `field:"optional" json:"severityProduct" yaml:"severityProduct"`
	// A URL that links to a page about the current finding in the security findings provider's solution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#source_url SecurityhubInsight#source_url}
	SourceUrl interface{} `field:"optional" json:"sourceUrl" yaml:"sourceUrl"`
	// The category of a threat intelligence indicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#threat_intel_indicator_category SecurityhubInsight#threat_intel_indicator_category}
	ThreatIntelIndicatorCategory interface{} `field:"optional" json:"threatIntelIndicatorCategory" yaml:"threatIntelIndicatorCategory"`
	// A timestamp that identifies the last observation of a threat intelligence indicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#threat_intel_indicator_last_observed_at SecurityhubInsight#threat_intel_indicator_last_observed_at}
	ThreatIntelIndicatorLastObservedAt interface{} `field:"optional" json:"threatIntelIndicatorLastObservedAt" yaml:"threatIntelIndicatorLastObservedAt"`
	// The source of the threat intelligence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#threat_intel_indicator_source SecurityhubInsight#threat_intel_indicator_source}
	ThreatIntelIndicatorSource interface{} `field:"optional" json:"threatIntelIndicatorSource" yaml:"threatIntelIndicatorSource"`
	// The URL for more details from the source of the threat intelligence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#threat_intel_indicator_source_url SecurityhubInsight#threat_intel_indicator_source_url}
	ThreatIntelIndicatorSourceUrl interface{} `field:"optional" json:"threatIntelIndicatorSourceUrl" yaml:"threatIntelIndicatorSourceUrl"`
	// The type of a threat intelligence indicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#threat_intel_indicator_type SecurityhubInsight#threat_intel_indicator_type}
	ThreatIntelIndicatorType interface{} `field:"optional" json:"threatIntelIndicatorType" yaml:"threatIntelIndicatorType"`
	// The value of a threat intelligence indicator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#threat_intel_indicator_value SecurityhubInsight#threat_intel_indicator_value}
	ThreatIntelIndicatorValue interface{} `field:"optional" json:"threatIntelIndicatorValue" yaml:"threatIntelIndicatorValue"`
	// A finding's title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#title SecurityhubInsight#title}
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// A finding type in the format of namespace/category/classifier that classifies a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#type SecurityhubInsight#type}
	Type interface{} `field:"optional" json:"type" yaml:"type"`
	// An ISO8601-formatted timestamp that indicates when the security findings provider last updated the finding record.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#updated_at SecurityhubInsight#updated_at}
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// A list of name/value string pairs associated with the finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#user_defined_fields SecurityhubInsight#user_defined_fields}
	UserDefinedFields interface{} `field:"optional" json:"userDefinedFields" yaml:"userDefinedFields"`
	// The veracity of a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#verification_state SecurityhubInsight#verification_state}
	VerificationState interface{} `field:"optional" json:"verificationState" yaml:"verificationState"`
	// Indicates whether a software vulnerability in your environment has a known exploit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#vulnerabilities_exploit_available SecurityhubInsight#vulnerabilities_exploit_available}
	VulnerabilitiesExploitAvailable interface{} `field:"optional" json:"vulnerabilitiesExploitAvailable" yaml:"vulnerabilitiesExploitAvailable"`
	// Indicates whether a vulnerability is fixed in a newer version of the affected software packages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#vulnerabilities_fix_available SecurityhubInsight#vulnerabilities_fix_available}
	VulnerabilitiesFixAvailable interface{} `field:"optional" json:"vulnerabilitiesFixAvailable" yaml:"vulnerabilitiesFixAvailable"`
	// The workflow state of a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#workflow_state SecurityhubInsight#workflow_state}
	WorkflowState interface{} `field:"optional" json:"workflowState" yaml:"workflowState"`
	// The status of the investigation into a finding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/securityhub_insight#workflow_status SecurityhubInsight#workflow_status}
	WorkflowStatus interface{} `field:"optional" json:"workflowStatus" yaml:"workflowStatus"`
}

