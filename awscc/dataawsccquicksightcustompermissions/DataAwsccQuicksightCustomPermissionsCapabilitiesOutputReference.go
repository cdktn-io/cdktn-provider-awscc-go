// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccquicksightcustompermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccquicksightcustompermissions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference interface {
	cdktn.ComplexObject
	AccessAppsNativeDataStore() *string
	Action() *string
	AddOrRunAnomalyDetectionForAnalyses() *string
	AmazonBedrockArsAction() *string
	AmazonBedrockFsAction() *string
	AmazonBedrockKrsAction() *string
	AmazonSThreeAction() *string
	Analysis() *string
	ApproveFlowShareRequests() *string
	Apps() *string
	AsanaAction() *string
	Automate() *string
	BambooHrAction() *string
	BoxAgentAction() *string
	BuildCalculatedFieldWithQ() *string
	CanvaAgentAction() *string
	ChatAgent() *string
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
	ComprehendAction() *string
	ComprehendMedicalAction() *string
	ConfluenceAction() *string
	CreateAndUpdateAmazonBedrockArsAction() *string
	CreateAndUpdateAmazonBedrockFsAction() *string
	CreateAndUpdateAmazonBedrockKrsAction() *string
	CreateAndUpdateAmazonSThreeAction() *string
	CreateAndUpdateApps() *string
	CreateAndUpdateAsanaAction() *string
	CreateAndUpdateBambooHrAction() *string
	CreateAndUpdateBoxAgentAction() *string
	CreateAndUpdateCanvaAgentAction() *string
	CreateAndUpdateComprehendAction() *string
	CreateAndUpdateComprehendMedicalAction() *string
	CreateAndUpdateConfluenceAction() *string
	CreateAndUpdateDashboardEmailReports() *string
	CreateAndUpdateDatasets() *string
	CreateAndUpdateDataSources() *string
	CreateAndUpdateFactSetAction() *string
	CreateAndUpdateGenericHttpAction() *string
	CreateAndUpdateGithubAction() *string
	CreateAndUpdateGoogleCalendarAction() *string
	CreateAndUpdateHubspotAction() *string
	CreateAndUpdateHuggingFaceAction() *string
	CreateAndUpdateIntercomAction() *string
	CreateAndUpdateJiraAction() *string
	CreateAndUpdateKnowledgeBases() *string
	CreateAndUpdateLinearAction() *string
	CreateAndUpdateMcpAction() *string
	CreateAndUpdateMondayAction() *string
	CreateAndUpdateMsExchangeAction() *string
	CreateAndUpdateMsTeamsAction() *string
	CreateAndUpdateNewRelicAction() *string
	CreateAndUpdateNotionAction() *string
	CreateAndUpdateOneDriveAction() *string
	CreateAndUpdateOpenApiAction() *string
	CreateAndUpdatePagerDutyAction() *string
	CreateAndUpdateSalesforceAction() *string
	CreateAndUpdateSandPGlobalEnergyAction() *string
	CreateAndUpdateSandPgmiAction() *string
	CreateAndUpdateSapBillOfMaterialAction() *string
	CreateAndUpdateSapBusinessPartnerAction() *string
	CreateAndUpdateSapMaterialStockAction() *string
	CreateAndUpdateSapPhysicalInventoryAction() *string
	CreateAndUpdateSapProductMasterDataAction() *string
	CreateAndUpdateServiceNowAction() *string
	CreateAndUpdateSharePointAction() *string
	CreateAndUpdateSlackAction() *string
	CreateAndUpdateSmartsheetAction() *string
	CreateAndUpdateTextractAction() *string
	CreateAndUpdateThemes() *string
	CreateAndUpdateThresholdAlerts() *string
	CreateAndUpdateZendeskAction() *string
	CreateChatAgents() *string
	CreateDashboardExecutiveSummaryWithQ() *string
	CreateSharedFolders() *string
	CreateSpaces() *string
	CreateSpiceDataset() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dashboard() *string
	EditVisualWithQ() *string
	ExportToCsv() *string
	ExportToCsvInScheduledReports() *string
	ExportToExcel() *string
	ExportToExcelInScheduledReports() *string
	ExportToPdf() *string
	ExportToPdfInScheduledReports() *string
	Extension() *string
	FactSetAction() *string
	Flow() *string
	// Experimental.
	Fqn() *string
	GenericHttpAction() *string
	GithubAction() *string
	GoogleCalendarAction() *string
	HubspotAction() *string
	HuggingFaceAction() *string
	IncludeContentInScheduledReportsEmail() *string
	IntercomAction() *string
	InternalValue() *DataAwsccQuicksightCustomPermissionsCapabilities
	SetInternalValue(val *DataAwsccQuicksightCustomPermissionsCapabilities)
	InvokeAppsAiInference() *string
	JiraAction() *string
	KnowledgeBase() *string
	LinearAction() *string
	ManageSharedFolders() *string
	McpAction() *string
	MondayAction() *string
	MsExchangeAction() *string
	MsTeamsAction() *string
	NewRelicAction() *string
	NotionAction() *string
	OneDriveAction() *string
	OpenApiAction() *string
	PagerDutyAction() *string
	PerformFlowUiTask() *string
	PrintReports() *string
	PublishWithoutApproval() *string
	RenameSharedFolders() *string
	Research() *string
	SalesforceAction() *string
	SandPGlobalEnergyAction() *string
	SandPgmiAction() *string
	SapBillOfMaterialAction() *string
	SapBusinessPartnerAction() *string
	SapMaterialStockAction() *string
	SapPhysicalInventoryAction() *string
	SapProductMasterDataAction() *string
	ServiceNowAction() *string
	ShareAmazonBedrockArsAction() *string
	ShareAmazonBedrockFsAction() *string
	ShareAmazonBedrockKrsAction() *string
	ShareAmazonSThreeAction() *string
	ShareAnalyses() *string
	ShareApps() *string
	ShareAsanaAction() *string
	ShareBambooHrAction() *string
	ShareBoxAgentAction() *string
	ShareCanvaAgentAction() *string
	ShareChatAgents() *string
	ShareComprehendAction() *string
	ShareComprehendMedicalAction() *string
	ShareConfluenceAction() *string
	ShareDashboards() *string
	ShareDatasets() *string
	ShareDataSources() *string
	ShareFactSetAction() *string
	ShareGenericHttpAction() *string
	ShareGithubAction() *string
	ShareGoogleCalendarAction() *string
	ShareHubspotAction() *string
	ShareHuggingFaceAction() *string
	ShareIntercomAction() *string
	ShareJiraAction() *string
	ShareKnowledgeBases() *string
	ShareLinearAction() *string
	ShareMcpAction() *string
	ShareMondayAction() *string
	ShareMsExchangeAction() *string
	ShareMsTeamsAction() *string
	ShareNewRelicAction() *string
	ShareNotionAction() *string
	ShareOneDriveAction() *string
	ShareOpenApiAction() *string
	SharePagerDutyAction() *string
	SharePointAction() *string
	ShareSalesforceAction() *string
	ShareSandPGlobalEnergyAction() *string
	ShareSandPgmiAction() *string
	ShareSapBillOfMaterialAction() *string
	ShareSapBusinessPartnerAction() *string
	ShareSapMaterialStockAction() *string
	ShareSapPhysicalInventoryAction() *string
	ShareSapProductMasterDataAction() *string
	ShareServiceNowAction() *string
	ShareSharePointAction() *string
	ShareSlackAction() *string
	ShareSmartsheetAction() *string
	ShareSpaces() *string
	ShareTextractAction() *string
	ShareZendeskAction() *string
	SlackAction() *string
	SmartsheetAction() *string
	Space() *string
	SubscribeDashboardEmailReports() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextractAction() *string
	Topic() *string
	UseAgentWebSearch() *string
	UseAmazonBedrockArsAction() *string
	UseAmazonBedrockFsAction() *string
	UseAmazonBedrockKrsAction() *string
	UseAmazonSThreeAction() *string
	UseAsanaAction() *string
	UseBambooHrAction() *string
	UseBedrockModels() *string
	UseBoxAgentAction() *string
	UseCanvaAgentAction() *string
	UseComprehendAction() *string
	UseComprehendMedicalAction() *string
	UseConfluenceAction() *string
	UseFactSetAction() *string
	UseGenericHttpAction() *string
	UseGithubAction() *string
	UseGoogleCalendarAction() *string
	UseHubspotAction() *string
	UseHuggingFaceAction() *string
	UseIntercomAction() *string
	UseJiraAction() *string
	UseLinearAction() *string
	UseMcpAction() *string
	UseMondayAction() *string
	UseMsExchangeAction() *string
	UseMsTeamsAction() *string
	UseNewRelicAction() *string
	UseNotionAction() *string
	UseOneDriveAction() *string
	UseOpenApiAction() *string
	UsePagerDutyAction() *string
	UseSalesforceAction() *string
	UseSandPGlobalEnergyAction() *string
	UseSandPgmiAction() *string
	UseSapBillOfMaterialAction() *string
	UseSapBusinessPartnerAction() *string
	UseSapMaterialStockAction() *string
	UseSapPhysicalInventoryAction() *string
	UseSapProductMasterDataAction() *string
	UseServiceNowAction() *string
	UseSharePointAction() *string
	UseSlackAction() *string
	UseSmartsheetAction() *string
	UseTextractAction() *string
	UseZendeskAction() *string
	ViewAccountSpiceCapacity() *string
	ZendeskAction() *string
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

// The jsii proxy struct for DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference
type jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AccessAppsNativeDataStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessAppsNativeDataStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AddOrRunAnomalyDetectionForAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Analysis() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ApproveFlowShareRequests() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveFlowShareRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Apps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Automate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) BambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) BoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) BuildCalculatedFieldWithQ() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCalculatedFieldWithQ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ChatAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chatAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateApps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAsanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBoxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateCanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateCanvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateComprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateComprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateConfluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateFactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateFactSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGenericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGithubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGoogleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateHubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateHubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateHuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateHuggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateIntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateIntercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateJiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateJiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateKnowledgeBases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateKnowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateLinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateLinearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMcpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMsExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMsTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateNewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateNewRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateNotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateNotionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOpenApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdatePagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdatePagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSalesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateServiceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSlackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSmartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateTextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateTextractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThemes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThresholdAlerts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateZendeskAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateChatAgents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createChatAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateDashboardExecutiveSummaryWithQ() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDashboardExecutiveSummaryWithQ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpaces() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpiceDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Dashboard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) EditVisualWithQ() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editVisualWithQ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Extension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) FactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"factSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Flow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) HubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) HuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) IncludeContentInScheduledReportsEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) IntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InternalValue() *DataAwsccQuicksightCustomPermissionsCapabilities {
	var returns *DataAwsccQuicksightCustomPermissionsCapabilities
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InvokeAppsAiInference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeAppsAiInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) JiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) LinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ManageSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manageSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) McpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) MondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) MsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) MsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) NewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) NotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) OneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) OpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PerformFlowUiTask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performFlowUiTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PrintReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PublishWithoutApproval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishWithoutApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) RenameSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Research() *string {
	var returns *string
	_jsii_.Get(
		j,
		"research",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareApps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAsanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareBambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareBoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBoxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareCanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareCanvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareChatAgents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareChatAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareComprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareComprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareConfluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDashboards() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareFactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareFactSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGenericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGithubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareHubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareHubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareHuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareHuggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareIntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareIntercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareJiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareJiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareKnowledgeBases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareKnowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareLinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareLinearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMcpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMsExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMsTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareNewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNewRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareNotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNotionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareOneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareOpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOpenApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SharePagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSalesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareServiceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSlackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSmartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSpaces() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareTextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTextractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareZendeskAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Space() *string {
	var returns *string
	_jsii_.Get(
		j,
		"space",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SubscribeDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseAgentWebSearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAgentWebSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseAsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAsanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseBambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseBedrockModels() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBedrockModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseBoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBoxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseCanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useCanvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useComprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useComprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useConfluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseFactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useFactSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseGenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGenericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseGithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGithubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseGoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGoogleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseHubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useHubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseHuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useHuggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseIntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useIntercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseJiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useJiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseLinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useLinearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseMcpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseMondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseMsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMsExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseMsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMsTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseNewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useNewRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseNotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useNotionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseOneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseOpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOpenApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UsePagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSalesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useServiceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSlackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSmartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseTextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useTextractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useZendeskAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ViewAccountSpiceCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendeskAction",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightCustomPermissionsCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightCustomPermissions.DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference_Override(d DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccQuicksightCustomPermissions.DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetInternalValue(val *DataAwsccQuicksightCustomPermissionsCapabilities) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

