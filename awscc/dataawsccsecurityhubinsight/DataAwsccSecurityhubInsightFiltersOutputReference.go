// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccsecurityhubinsight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccsecurityhubinsight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccSecurityhubInsightFiltersOutputReference interface {
	cdktn.ComplexObject
	AwsAccountId() DataAwsccSecurityhubInsightFiltersAwsAccountIdList
	AwsAccountName() DataAwsccSecurityhubInsightFiltersAwsAccountNameList
	CompanyName() DataAwsccSecurityhubInsightFiltersCompanyNameList
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ComplianceAssociatedStandardsId() DataAwsccSecurityhubInsightFiltersComplianceAssociatedStandardsIdList
	ComplianceSecurityControlId() DataAwsccSecurityhubInsightFiltersComplianceSecurityControlIdList
	ComplianceSecurityControlParametersName() DataAwsccSecurityhubInsightFiltersComplianceSecurityControlParametersNameList
	ComplianceSecurityControlParametersValue() DataAwsccSecurityhubInsightFiltersComplianceSecurityControlParametersValueList
	ComplianceStatus() DataAwsccSecurityhubInsightFiltersComplianceStatusList
	Confidence() DataAwsccSecurityhubInsightFiltersConfidenceList
	CreatedAt() DataAwsccSecurityhubInsightFiltersCreatedAtList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Criticality() DataAwsccSecurityhubInsightFiltersCriticalityList
	Description() DataAwsccSecurityhubInsightFiltersDescriptionList
	FindingProviderFieldsConfidence() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsConfidenceList
	FindingProviderFieldsCriticality() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsCriticalityList
	FindingProviderFieldsRelatedFindingsId() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsIdList
	FindingProviderFieldsRelatedFindingsProductArn() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArnList
	FindingProviderFieldsSeverityLabel() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsSeverityLabelList
	FindingProviderFieldsSeverityOriginal() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsSeverityOriginalList
	FindingProviderFieldsTypes() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsTypesList
	FirstObservedAt() DataAwsccSecurityhubInsightFiltersFirstObservedAtList
	// Experimental.
	Fqn() *string
	GeneratorId() DataAwsccSecurityhubInsightFiltersGeneratorIdList
	Id() DataAwsccSecurityhubInsightFiltersIdList
	InternalValue() *DataAwsccSecurityhubInsightFilters
	SetInternalValue(val *DataAwsccSecurityhubInsightFilters)
	Keyword() DataAwsccSecurityhubInsightFiltersKeywordList
	LastObservedAt() DataAwsccSecurityhubInsightFiltersLastObservedAtList
	MalwareName() DataAwsccSecurityhubInsightFiltersMalwareNameList
	MalwarePath() DataAwsccSecurityhubInsightFiltersMalwarePathList
	MalwareState() DataAwsccSecurityhubInsightFiltersMalwareStateList
	MalwareType() DataAwsccSecurityhubInsightFiltersMalwareTypeList
	NetworkDestinationDomain() DataAwsccSecurityhubInsightFiltersNetworkDestinationDomainList
	NetworkDestinationIpV4() DataAwsccSecurityhubInsightFiltersNetworkDestinationIpV4List
	NetworkDestinationIpV6() DataAwsccSecurityhubInsightFiltersNetworkDestinationIpV6List
	NetworkDestinationPort() DataAwsccSecurityhubInsightFiltersNetworkDestinationPortList
	NetworkDirection() DataAwsccSecurityhubInsightFiltersNetworkDirectionList
	NetworkProtocol() DataAwsccSecurityhubInsightFiltersNetworkProtocolList
	NetworkSourceDomain() DataAwsccSecurityhubInsightFiltersNetworkSourceDomainList
	NetworkSourceIpV4() DataAwsccSecurityhubInsightFiltersNetworkSourceIpV4List
	NetworkSourceIpV6() DataAwsccSecurityhubInsightFiltersNetworkSourceIpV6List
	NetworkSourceMac() DataAwsccSecurityhubInsightFiltersNetworkSourceMacList
	NetworkSourcePort() DataAwsccSecurityhubInsightFiltersNetworkSourcePortList
	NoteText() DataAwsccSecurityhubInsightFiltersNoteTextList
	NoteUpdatedAt() DataAwsccSecurityhubInsightFiltersNoteUpdatedAtList
	NoteUpdatedBy() DataAwsccSecurityhubInsightFiltersNoteUpdatedByList
	ProcessLaunchedAt() DataAwsccSecurityhubInsightFiltersProcessLaunchedAtList
	ProcessName() DataAwsccSecurityhubInsightFiltersProcessNameList
	ProcessParentPid() DataAwsccSecurityhubInsightFiltersProcessParentPidList
	ProcessPath() DataAwsccSecurityhubInsightFiltersProcessPathList
	ProcessPid() DataAwsccSecurityhubInsightFiltersProcessPidList
	ProcessTerminatedAt() DataAwsccSecurityhubInsightFiltersProcessTerminatedAtList
	ProductArn() DataAwsccSecurityhubInsightFiltersProductArnList
	ProductFields() DataAwsccSecurityhubInsightFiltersProductFieldsList
	ProductName() DataAwsccSecurityhubInsightFiltersProductNameList
	RecommendationText() DataAwsccSecurityhubInsightFiltersRecommendationTextList
	RecordState() DataAwsccSecurityhubInsightFiltersRecordStateList
	Region() DataAwsccSecurityhubInsightFiltersRegionList
	RelatedFindingsId() DataAwsccSecurityhubInsightFiltersRelatedFindingsIdList
	RelatedFindingsProductArn() DataAwsccSecurityhubInsightFiltersRelatedFindingsProductArnList
	ResourceApplicationArn() DataAwsccSecurityhubInsightFiltersResourceApplicationArnList
	ResourceApplicationName() DataAwsccSecurityhubInsightFiltersResourceApplicationNameList
	ResourceAwsEc2InstanceIamInstanceProfileArn() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArnList
	ResourceAwsEc2InstanceImageId() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceImageIdList
	ResourceAwsEc2InstanceIpV4Addresses() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIpV4AddressesList
	ResourceAwsEc2InstanceIpV6Addresses() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIpV6AddressesList
	ResourceAwsEc2InstanceKeyName() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceKeyNameList
	ResourceAwsEc2InstanceLaunchedAt() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtList
	ResourceAwsEc2InstanceSubnetId() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceSubnetIdList
	ResourceAwsEc2InstanceType() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceTypeList
	ResourceAwsEc2InstanceVpcId() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceVpcIdList
	ResourceAwsIamAccessKeyCreatedAt() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtList
	ResourceAwsIamAccessKeyPrincipalName() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalNameList
	ResourceAwsIamAccessKeyStatus() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyStatusList
	ResourceAwsIamAccessKeyUserName() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyUserNameList
	ResourceAwsIamUserUserName() DataAwsccSecurityhubInsightFiltersResourceAwsIamUserUserNameList
	ResourceAwsS3BucketOwnerId() DataAwsccSecurityhubInsightFiltersResourceAwsS3BucketOwnerIdList
	ResourceAwsS3BucketOwnerName() DataAwsccSecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList
	ResourceContainerImageId() DataAwsccSecurityhubInsightFiltersResourceContainerImageIdList
	ResourceContainerImageName() DataAwsccSecurityhubInsightFiltersResourceContainerImageNameList
	ResourceContainerLaunchedAt() DataAwsccSecurityhubInsightFiltersResourceContainerLaunchedAtList
	ResourceContainerName() DataAwsccSecurityhubInsightFiltersResourceContainerNameList
	ResourceDetailsOther() DataAwsccSecurityhubInsightFiltersResourceDetailsOtherList
	ResourceId() DataAwsccSecurityhubInsightFiltersResourceIdList
	ResourcePartition() DataAwsccSecurityhubInsightFiltersResourcePartitionList
	ResourceRegion() DataAwsccSecurityhubInsightFiltersResourceRegionList
	ResourceTags() DataAwsccSecurityhubInsightFiltersResourceTagsList
	ResourceType() DataAwsccSecurityhubInsightFiltersResourceTypeList
	Sample() DataAwsccSecurityhubInsightFiltersSampleList
	SeverityLabel() DataAwsccSecurityhubInsightFiltersSeverityLabelList
	SeverityNormalized() DataAwsccSecurityhubInsightFiltersSeverityNormalizedList
	SeverityProduct() DataAwsccSecurityhubInsightFiltersSeverityProductList
	SourceUrl() DataAwsccSecurityhubInsightFiltersSourceUrlList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	ThreatIntelIndicatorCategory() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorCategoryList
	ThreatIntelIndicatorLastObservedAt() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtList
	ThreatIntelIndicatorSource() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorSourceList
	ThreatIntelIndicatorSourceUrl() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorSourceUrlList
	ThreatIntelIndicatorType() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorTypeList
	ThreatIntelIndicatorValue() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorValueList
	Title() DataAwsccSecurityhubInsightFiltersTitleList
	Type() DataAwsccSecurityhubInsightFiltersTypeList
	UpdatedAt() DataAwsccSecurityhubInsightFiltersUpdatedAtList
	UserDefinedFields() DataAwsccSecurityhubInsightFiltersUserDefinedFieldsList
	VerificationState() DataAwsccSecurityhubInsightFiltersVerificationStateList
	VulnerabilitiesExploitAvailable() DataAwsccSecurityhubInsightFiltersVulnerabilitiesExploitAvailableList
	VulnerabilitiesFixAvailable() DataAwsccSecurityhubInsightFiltersVulnerabilitiesFixAvailableList
	WorkflowState() DataAwsccSecurityhubInsightFiltersWorkflowStateList
	WorkflowStatus() DataAwsccSecurityhubInsightFiltersWorkflowStatusList
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccSecurityhubInsightFiltersOutputReference
type jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) AwsAccountId() DataAwsccSecurityhubInsightFiltersAwsAccountIdList {
	var returns DataAwsccSecurityhubInsightFiltersAwsAccountIdList
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) AwsAccountName() DataAwsccSecurityhubInsightFiltersAwsAccountNameList {
	var returns DataAwsccSecurityhubInsightFiltersAwsAccountNameList
	_jsii_.Get(
		j,
		"awsAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) CompanyName() DataAwsccSecurityhubInsightFiltersCompanyNameList {
	var returns DataAwsccSecurityhubInsightFiltersCompanyNameList
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComplianceAssociatedStandardsId() DataAwsccSecurityhubInsightFiltersComplianceAssociatedStandardsIdList {
	var returns DataAwsccSecurityhubInsightFiltersComplianceAssociatedStandardsIdList
	_jsii_.Get(
		j,
		"complianceAssociatedStandardsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComplianceSecurityControlId() DataAwsccSecurityhubInsightFiltersComplianceSecurityControlIdList {
	var returns DataAwsccSecurityhubInsightFiltersComplianceSecurityControlIdList
	_jsii_.Get(
		j,
		"complianceSecurityControlId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComplianceSecurityControlParametersName() DataAwsccSecurityhubInsightFiltersComplianceSecurityControlParametersNameList {
	var returns DataAwsccSecurityhubInsightFiltersComplianceSecurityControlParametersNameList
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComplianceSecurityControlParametersValue() DataAwsccSecurityhubInsightFiltersComplianceSecurityControlParametersValueList {
	var returns DataAwsccSecurityhubInsightFiltersComplianceSecurityControlParametersValueList
	_jsii_.Get(
		j,
		"complianceSecurityControlParametersValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComplianceStatus() DataAwsccSecurityhubInsightFiltersComplianceStatusList {
	var returns DataAwsccSecurityhubInsightFiltersComplianceStatusList
	_jsii_.Get(
		j,
		"complianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Confidence() DataAwsccSecurityhubInsightFiltersConfidenceList {
	var returns DataAwsccSecurityhubInsightFiltersConfidenceList
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) CreatedAt() DataAwsccSecurityhubInsightFiltersCreatedAtList {
	var returns DataAwsccSecurityhubInsightFiltersCreatedAtList
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Criticality() DataAwsccSecurityhubInsightFiltersCriticalityList {
	var returns DataAwsccSecurityhubInsightFiltersCriticalityList
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Description() DataAwsccSecurityhubInsightFiltersDescriptionList {
	var returns DataAwsccSecurityhubInsightFiltersDescriptionList
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FindingProviderFieldsConfidence() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsConfidenceList {
	var returns DataAwsccSecurityhubInsightFiltersFindingProviderFieldsConfidenceList
	_jsii_.Get(
		j,
		"findingProviderFieldsConfidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FindingProviderFieldsCriticality() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsCriticalityList {
	var returns DataAwsccSecurityhubInsightFiltersFindingProviderFieldsCriticalityList
	_jsii_.Get(
		j,
		"findingProviderFieldsCriticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsId() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsIdList {
	var returns DataAwsccSecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsIdList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FindingProviderFieldsRelatedFindingsProductArn() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArnList {
	var returns DataAwsccSecurityhubInsightFiltersFindingProviderFieldsRelatedFindingsProductArnList
	_jsii_.Get(
		j,
		"findingProviderFieldsRelatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityLabel() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsSeverityLabelList {
	var returns DataAwsccSecurityhubInsightFiltersFindingProviderFieldsSeverityLabelList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FindingProviderFieldsSeverityOriginal() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsSeverityOriginalList {
	var returns DataAwsccSecurityhubInsightFiltersFindingProviderFieldsSeverityOriginalList
	_jsii_.Get(
		j,
		"findingProviderFieldsSeverityOriginal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FindingProviderFieldsTypes() DataAwsccSecurityhubInsightFiltersFindingProviderFieldsTypesList {
	var returns DataAwsccSecurityhubInsightFiltersFindingProviderFieldsTypesList
	_jsii_.Get(
		j,
		"findingProviderFieldsTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) FirstObservedAt() DataAwsccSecurityhubInsightFiltersFirstObservedAtList {
	var returns DataAwsccSecurityhubInsightFiltersFirstObservedAtList
	_jsii_.Get(
		j,
		"firstObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GeneratorId() DataAwsccSecurityhubInsightFiltersGeneratorIdList {
	var returns DataAwsccSecurityhubInsightFiltersGeneratorIdList
	_jsii_.Get(
		j,
		"generatorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Id() DataAwsccSecurityhubInsightFiltersIdList {
	var returns DataAwsccSecurityhubInsightFiltersIdList
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) InternalValue() *DataAwsccSecurityhubInsightFilters {
	var returns *DataAwsccSecurityhubInsightFilters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Keyword() DataAwsccSecurityhubInsightFiltersKeywordList {
	var returns DataAwsccSecurityhubInsightFiltersKeywordList
	_jsii_.Get(
		j,
		"keyword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) LastObservedAt() DataAwsccSecurityhubInsightFiltersLastObservedAtList {
	var returns DataAwsccSecurityhubInsightFiltersLastObservedAtList
	_jsii_.Get(
		j,
		"lastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) MalwareName() DataAwsccSecurityhubInsightFiltersMalwareNameList {
	var returns DataAwsccSecurityhubInsightFiltersMalwareNameList
	_jsii_.Get(
		j,
		"malwareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) MalwarePath() DataAwsccSecurityhubInsightFiltersMalwarePathList {
	var returns DataAwsccSecurityhubInsightFiltersMalwarePathList
	_jsii_.Get(
		j,
		"malwarePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) MalwareState() DataAwsccSecurityhubInsightFiltersMalwareStateList {
	var returns DataAwsccSecurityhubInsightFiltersMalwareStateList
	_jsii_.Get(
		j,
		"malwareState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) MalwareType() DataAwsccSecurityhubInsightFiltersMalwareTypeList {
	var returns DataAwsccSecurityhubInsightFiltersMalwareTypeList
	_jsii_.Get(
		j,
		"malwareType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkDestinationDomain() DataAwsccSecurityhubInsightFiltersNetworkDestinationDomainList {
	var returns DataAwsccSecurityhubInsightFiltersNetworkDestinationDomainList
	_jsii_.Get(
		j,
		"networkDestinationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkDestinationIpV4() DataAwsccSecurityhubInsightFiltersNetworkDestinationIpV4List {
	var returns DataAwsccSecurityhubInsightFiltersNetworkDestinationIpV4List
	_jsii_.Get(
		j,
		"networkDestinationIpV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkDestinationIpV6() DataAwsccSecurityhubInsightFiltersNetworkDestinationIpV6List {
	var returns DataAwsccSecurityhubInsightFiltersNetworkDestinationIpV6List
	_jsii_.Get(
		j,
		"networkDestinationIpV6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkDestinationPort() DataAwsccSecurityhubInsightFiltersNetworkDestinationPortList {
	var returns DataAwsccSecurityhubInsightFiltersNetworkDestinationPortList
	_jsii_.Get(
		j,
		"networkDestinationPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkDirection() DataAwsccSecurityhubInsightFiltersNetworkDirectionList {
	var returns DataAwsccSecurityhubInsightFiltersNetworkDirectionList
	_jsii_.Get(
		j,
		"networkDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkProtocol() DataAwsccSecurityhubInsightFiltersNetworkProtocolList {
	var returns DataAwsccSecurityhubInsightFiltersNetworkProtocolList
	_jsii_.Get(
		j,
		"networkProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkSourceDomain() DataAwsccSecurityhubInsightFiltersNetworkSourceDomainList {
	var returns DataAwsccSecurityhubInsightFiltersNetworkSourceDomainList
	_jsii_.Get(
		j,
		"networkSourceDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkSourceIpV4() DataAwsccSecurityhubInsightFiltersNetworkSourceIpV4List {
	var returns DataAwsccSecurityhubInsightFiltersNetworkSourceIpV4List
	_jsii_.Get(
		j,
		"networkSourceIpV4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkSourceIpV6() DataAwsccSecurityhubInsightFiltersNetworkSourceIpV6List {
	var returns DataAwsccSecurityhubInsightFiltersNetworkSourceIpV6List
	_jsii_.Get(
		j,
		"networkSourceIpV6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkSourceMac() DataAwsccSecurityhubInsightFiltersNetworkSourceMacList {
	var returns DataAwsccSecurityhubInsightFiltersNetworkSourceMacList
	_jsii_.Get(
		j,
		"networkSourceMac",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NetworkSourcePort() DataAwsccSecurityhubInsightFiltersNetworkSourcePortList {
	var returns DataAwsccSecurityhubInsightFiltersNetworkSourcePortList
	_jsii_.Get(
		j,
		"networkSourcePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NoteText() DataAwsccSecurityhubInsightFiltersNoteTextList {
	var returns DataAwsccSecurityhubInsightFiltersNoteTextList
	_jsii_.Get(
		j,
		"noteText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NoteUpdatedAt() DataAwsccSecurityhubInsightFiltersNoteUpdatedAtList {
	var returns DataAwsccSecurityhubInsightFiltersNoteUpdatedAtList
	_jsii_.Get(
		j,
		"noteUpdatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) NoteUpdatedBy() DataAwsccSecurityhubInsightFiltersNoteUpdatedByList {
	var returns DataAwsccSecurityhubInsightFiltersNoteUpdatedByList
	_jsii_.Get(
		j,
		"noteUpdatedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProcessLaunchedAt() DataAwsccSecurityhubInsightFiltersProcessLaunchedAtList {
	var returns DataAwsccSecurityhubInsightFiltersProcessLaunchedAtList
	_jsii_.Get(
		j,
		"processLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProcessName() DataAwsccSecurityhubInsightFiltersProcessNameList {
	var returns DataAwsccSecurityhubInsightFiltersProcessNameList
	_jsii_.Get(
		j,
		"processName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProcessParentPid() DataAwsccSecurityhubInsightFiltersProcessParentPidList {
	var returns DataAwsccSecurityhubInsightFiltersProcessParentPidList
	_jsii_.Get(
		j,
		"processParentPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProcessPath() DataAwsccSecurityhubInsightFiltersProcessPathList {
	var returns DataAwsccSecurityhubInsightFiltersProcessPathList
	_jsii_.Get(
		j,
		"processPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProcessPid() DataAwsccSecurityhubInsightFiltersProcessPidList {
	var returns DataAwsccSecurityhubInsightFiltersProcessPidList
	_jsii_.Get(
		j,
		"processPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProcessTerminatedAt() DataAwsccSecurityhubInsightFiltersProcessTerminatedAtList {
	var returns DataAwsccSecurityhubInsightFiltersProcessTerminatedAtList
	_jsii_.Get(
		j,
		"processTerminatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProductArn() DataAwsccSecurityhubInsightFiltersProductArnList {
	var returns DataAwsccSecurityhubInsightFiltersProductArnList
	_jsii_.Get(
		j,
		"productArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProductFields() DataAwsccSecurityhubInsightFiltersProductFieldsList {
	var returns DataAwsccSecurityhubInsightFiltersProductFieldsList
	_jsii_.Get(
		j,
		"productFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ProductName() DataAwsccSecurityhubInsightFiltersProductNameList {
	var returns DataAwsccSecurityhubInsightFiltersProductNameList
	_jsii_.Get(
		j,
		"productName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) RecommendationText() DataAwsccSecurityhubInsightFiltersRecommendationTextList {
	var returns DataAwsccSecurityhubInsightFiltersRecommendationTextList
	_jsii_.Get(
		j,
		"recommendationText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) RecordState() DataAwsccSecurityhubInsightFiltersRecordStateList {
	var returns DataAwsccSecurityhubInsightFiltersRecordStateList
	_jsii_.Get(
		j,
		"recordState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Region() DataAwsccSecurityhubInsightFiltersRegionList {
	var returns DataAwsccSecurityhubInsightFiltersRegionList
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) RelatedFindingsId() DataAwsccSecurityhubInsightFiltersRelatedFindingsIdList {
	var returns DataAwsccSecurityhubInsightFiltersRelatedFindingsIdList
	_jsii_.Get(
		j,
		"relatedFindingsId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) RelatedFindingsProductArn() DataAwsccSecurityhubInsightFiltersRelatedFindingsProductArnList {
	var returns DataAwsccSecurityhubInsightFiltersRelatedFindingsProductArnList
	_jsii_.Get(
		j,
		"relatedFindingsProductArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceApplicationArn() DataAwsccSecurityhubInsightFiltersResourceApplicationArnList {
	var returns DataAwsccSecurityhubInsightFiltersResourceApplicationArnList
	_jsii_.Get(
		j,
		"resourceApplicationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceApplicationName() DataAwsccSecurityhubInsightFiltersResourceApplicationNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceApplicationNameList
	_jsii_.Get(
		j,
		"resourceApplicationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIamInstanceProfileArn() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArnList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIamInstanceProfileArnList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceImageId() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceImageIdList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceImageIdList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpV4Addresses() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIpV4AddressesList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIpV4AddressesList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpV4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceIpV6Addresses() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIpV6AddressesList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceIpV6AddressesList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceIpV6Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceKeyName() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceKeyNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceKeyNameList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceLaunchedAt() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceLaunchedAtList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceSubnetId() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceSubnetIdList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceSubnetIdList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceType() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceTypeList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceTypeList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsEc2InstanceVpcId() DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceVpcIdList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsEc2InstanceVpcIdList
	_jsii_.Get(
		j,
		"resourceAwsEc2InstanceVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyCreatedAt() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyCreatedAtList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyCreatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyPrincipalName() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyPrincipalNameList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyPrincipalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyStatus() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyStatusList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyStatusList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsIamAccessKeyUserName() DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyUserNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsIamAccessKeyUserNameList
	_jsii_.Get(
		j,
		"resourceAwsIamAccessKeyUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsIamUserUserName() DataAwsccSecurityhubInsightFiltersResourceAwsIamUserUserNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsIamUserUserNameList
	_jsii_.Get(
		j,
		"resourceAwsIamUserUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerId() DataAwsccSecurityhubInsightFiltersResourceAwsS3BucketOwnerIdList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsS3BucketOwnerIdList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceAwsS3BucketOwnerName() DataAwsccSecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceAwsS3BucketOwnerNameList
	_jsii_.Get(
		j,
		"resourceAwsS3BucketOwnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceContainerImageId() DataAwsccSecurityhubInsightFiltersResourceContainerImageIdList {
	var returns DataAwsccSecurityhubInsightFiltersResourceContainerImageIdList
	_jsii_.Get(
		j,
		"resourceContainerImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceContainerImageName() DataAwsccSecurityhubInsightFiltersResourceContainerImageNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceContainerImageNameList
	_jsii_.Get(
		j,
		"resourceContainerImageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceContainerLaunchedAt() DataAwsccSecurityhubInsightFiltersResourceContainerLaunchedAtList {
	var returns DataAwsccSecurityhubInsightFiltersResourceContainerLaunchedAtList
	_jsii_.Get(
		j,
		"resourceContainerLaunchedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceContainerName() DataAwsccSecurityhubInsightFiltersResourceContainerNameList {
	var returns DataAwsccSecurityhubInsightFiltersResourceContainerNameList
	_jsii_.Get(
		j,
		"resourceContainerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceDetailsOther() DataAwsccSecurityhubInsightFiltersResourceDetailsOtherList {
	var returns DataAwsccSecurityhubInsightFiltersResourceDetailsOtherList
	_jsii_.Get(
		j,
		"resourceDetailsOther",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceId() DataAwsccSecurityhubInsightFiltersResourceIdList {
	var returns DataAwsccSecurityhubInsightFiltersResourceIdList
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourcePartition() DataAwsccSecurityhubInsightFiltersResourcePartitionList {
	var returns DataAwsccSecurityhubInsightFiltersResourcePartitionList
	_jsii_.Get(
		j,
		"resourcePartition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceRegion() DataAwsccSecurityhubInsightFiltersResourceRegionList {
	var returns DataAwsccSecurityhubInsightFiltersResourceRegionList
	_jsii_.Get(
		j,
		"resourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceTags() DataAwsccSecurityhubInsightFiltersResourceTagsList {
	var returns DataAwsccSecurityhubInsightFiltersResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ResourceType() DataAwsccSecurityhubInsightFiltersResourceTypeList {
	var returns DataAwsccSecurityhubInsightFiltersResourceTypeList
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Sample() DataAwsccSecurityhubInsightFiltersSampleList {
	var returns DataAwsccSecurityhubInsightFiltersSampleList
	_jsii_.Get(
		j,
		"sample",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) SeverityLabel() DataAwsccSecurityhubInsightFiltersSeverityLabelList {
	var returns DataAwsccSecurityhubInsightFiltersSeverityLabelList
	_jsii_.Get(
		j,
		"severityLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) SeverityNormalized() DataAwsccSecurityhubInsightFiltersSeverityNormalizedList {
	var returns DataAwsccSecurityhubInsightFiltersSeverityNormalizedList
	_jsii_.Get(
		j,
		"severityNormalized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) SeverityProduct() DataAwsccSecurityhubInsightFiltersSeverityProductList {
	var returns DataAwsccSecurityhubInsightFiltersSeverityProductList
	_jsii_.Get(
		j,
		"severityProduct",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) SourceUrl() DataAwsccSecurityhubInsightFiltersSourceUrlList {
	var returns DataAwsccSecurityhubInsightFiltersSourceUrlList
	_jsii_.Get(
		j,
		"sourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorCategory() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorCategoryList {
	var returns DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorCategoryList
	_jsii_.Get(
		j,
		"threatIntelIndicatorCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorLastObservedAt() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtList {
	var returns DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorLastObservedAtList
	_jsii_.Get(
		j,
		"threatIntelIndicatorLastObservedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSource() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorSourceList {
	var returns DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorSourceList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorSourceUrl() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorSourceUrlList {
	var returns DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorSourceUrlList
	_jsii_.Get(
		j,
		"threatIntelIndicatorSourceUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorType() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorTypeList {
	var returns DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorTypeList
	_jsii_.Get(
		j,
		"threatIntelIndicatorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ThreatIntelIndicatorValue() DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorValueList {
	var returns DataAwsccSecurityhubInsightFiltersThreatIntelIndicatorValueList
	_jsii_.Get(
		j,
		"threatIntelIndicatorValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Title() DataAwsccSecurityhubInsightFiltersTitleList {
	var returns DataAwsccSecurityhubInsightFiltersTitleList
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Type() DataAwsccSecurityhubInsightFiltersTypeList {
	var returns DataAwsccSecurityhubInsightFiltersTypeList
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) UpdatedAt() DataAwsccSecurityhubInsightFiltersUpdatedAtList {
	var returns DataAwsccSecurityhubInsightFiltersUpdatedAtList
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) UserDefinedFields() DataAwsccSecurityhubInsightFiltersUserDefinedFieldsList {
	var returns DataAwsccSecurityhubInsightFiltersUserDefinedFieldsList
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) VerificationState() DataAwsccSecurityhubInsightFiltersVerificationStateList {
	var returns DataAwsccSecurityhubInsightFiltersVerificationStateList
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) VulnerabilitiesExploitAvailable() DataAwsccSecurityhubInsightFiltersVulnerabilitiesExploitAvailableList {
	var returns DataAwsccSecurityhubInsightFiltersVulnerabilitiesExploitAvailableList
	_jsii_.Get(
		j,
		"vulnerabilitiesExploitAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) VulnerabilitiesFixAvailable() DataAwsccSecurityhubInsightFiltersVulnerabilitiesFixAvailableList {
	var returns DataAwsccSecurityhubInsightFiltersVulnerabilitiesFixAvailableList
	_jsii_.Get(
		j,
		"vulnerabilitiesFixAvailable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) WorkflowState() DataAwsccSecurityhubInsightFiltersWorkflowStateList {
	var returns DataAwsccSecurityhubInsightFiltersWorkflowStateList
	_jsii_.Get(
		j,
		"workflowState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) WorkflowStatus() DataAwsccSecurityhubInsightFiltersWorkflowStatusList {
	var returns DataAwsccSecurityhubInsightFiltersWorkflowStatusList
	_jsii_.Get(
		j,
		"workflowStatus",
		&returns,
	)
	return returns
}


func NewDataAwsccSecurityhubInsightFiltersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccSecurityhubInsightFiltersOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccSecurityhubInsightFiltersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSecurityhubInsight.DataAwsccSecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccSecurityhubInsightFiltersOutputReference_Override(d DataAwsccSecurityhubInsightFiltersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccSecurityhubInsight.DataAwsccSecurityhubInsightFiltersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference)SetInternalValue(val *DataAwsccSecurityhubInsightFilters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSecurityhubInsightFiltersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

