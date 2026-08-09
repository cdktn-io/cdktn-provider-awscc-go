// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightcustompermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightcustompermissions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type QuicksightCustomPermissionsCapabilitiesOutputReference interface {
	cdktn.ComplexObject
	AccessAppsNativeDataStore() *string
	SetAccessAppsNativeDataStore(val *string)
	AccessAppsNativeDataStoreInput() *string
	Action() *string
	SetAction(val *string)
	ActionInput() *string
	AddOrRunAnomalyDetectionForAnalyses() *string
	SetAddOrRunAnomalyDetectionForAnalyses(val *string)
	AddOrRunAnomalyDetectionForAnalysesInput() *string
	AmazonBedrockArsAction() *string
	SetAmazonBedrockArsAction(val *string)
	AmazonBedrockArsActionInput() *string
	AmazonBedrockFsAction() *string
	SetAmazonBedrockFsAction(val *string)
	AmazonBedrockFsActionInput() *string
	AmazonBedrockKrsAction() *string
	SetAmazonBedrockKrsAction(val *string)
	AmazonBedrockKrsActionInput() *string
	AmazonSThreeAction() *string
	SetAmazonSThreeAction(val *string)
	AmazonSThreeActionInput() *string
	Analysis() *string
	SetAnalysis(val *string)
	AnalysisInput() *string
	ApproveFlowShareRequests() *string
	SetApproveFlowShareRequests(val *string)
	ApproveFlowShareRequestsInput() *string
	Apps() *string
	SetApps(val *string)
	AppsInput() *string
	AsanaAction() *string
	SetAsanaAction(val *string)
	AsanaActionInput() *string
	Automate() *string
	SetAutomate(val *string)
	AutomateInput() *string
	BambooHrAction() *string
	SetBambooHrAction(val *string)
	BambooHrActionInput() *string
	BoxAgentAction() *string
	SetBoxAgentAction(val *string)
	BoxAgentActionInput() *string
	BuildCalculatedFieldWithQ() *string
	SetBuildCalculatedFieldWithQ(val *string)
	BuildCalculatedFieldWithQInput() *string
	CanvaAgentAction() *string
	SetCanvaAgentAction(val *string)
	CanvaAgentActionInput() *string
	ChatAgent() *string
	SetChatAgent(val *string)
	ChatAgentInput() *string
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
	SetComprehendAction(val *string)
	ComprehendActionInput() *string
	ComprehendMedicalAction() *string
	SetComprehendMedicalAction(val *string)
	ComprehendMedicalActionInput() *string
	ConfluenceAction() *string
	SetConfluenceAction(val *string)
	ConfluenceActionInput() *string
	CreateAndUpdateAmazonBedrockArsAction() *string
	SetCreateAndUpdateAmazonBedrockArsAction(val *string)
	CreateAndUpdateAmazonBedrockArsActionInput() *string
	CreateAndUpdateAmazonBedrockFsAction() *string
	SetCreateAndUpdateAmazonBedrockFsAction(val *string)
	CreateAndUpdateAmazonBedrockFsActionInput() *string
	CreateAndUpdateAmazonBedrockKrsAction() *string
	SetCreateAndUpdateAmazonBedrockKrsAction(val *string)
	CreateAndUpdateAmazonBedrockKrsActionInput() *string
	CreateAndUpdateAmazonSThreeAction() *string
	SetCreateAndUpdateAmazonSThreeAction(val *string)
	CreateAndUpdateAmazonSThreeActionInput() *string
	CreateAndUpdateApps() *string
	SetCreateAndUpdateApps(val *string)
	CreateAndUpdateAppsInput() *string
	CreateAndUpdateAsanaAction() *string
	SetCreateAndUpdateAsanaAction(val *string)
	CreateAndUpdateAsanaActionInput() *string
	CreateAndUpdateBambooHrAction() *string
	SetCreateAndUpdateBambooHrAction(val *string)
	CreateAndUpdateBambooHrActionInput() *string
	CreateAndUpdateBoxAgentAction() *string
	SetCreateAndUpdateBoxAgentAction(val *string)
	CreateAndUpdateBoxAgentActionInput() *string
	CreateAndUpdateCanvaAgentAction() *string
	SetCreateAndUpdateCanvaAgentAction(val *string)
	CreateAndUpdateCanvaAgentActionInput() *string
	CreateAndUpdateComprehendAction() *string
	SetCreateAndUpdateComprehendAction(val *string)
	CreateAndUpdateComprehendActionInput() *string
	CreateAndUpdateComprehendMedicalAction() *string
	SetCreateAndUpdateComprehendMedicalAction(val *string)
	CreateAndUpdateComprehendMedicalActionInput() *string
	CreateAndUpdateConfluenceAction() *string
	SetCreateAndUpdateConfluenceAction(val *string)
	CreateAndUpdateConfluenceActionInput() *string
	CreateAndUpdateDashboardEmailReports() *string
	SetCreateAndUpdateDashboardEmailReports(val *string)
	CreateAndUpdateDashboardEmailReportsInput() *string
	CreateAndUpdateDatasets() *string
	SetCreateAndUpdateDatasets(val *string)
	CreateAndUpdateDatasetsInput() *string
	CreateAndUpdateDataSources() *string
	SetCreateAndUpdateDataSources(val *string)
	CreateAndUpdateDataSourcesInput() *string
	CreateAndUpdateFactSetAction() *string
	SetCreateAndUpdateFactSetAction(val *string)
	CreateAndUpdateFactSetActionInput() *string
	CreateAndUpdateGenericHttpAction() *string
	SetCreateAndUpdateGenericHttpAction(val *string)
	CreateAndUpdateGenericHttpActionInput() *string
	CreateAndUpdateGithubAction() *string
	SetCreateAndUpdateGithubAction(val *string)
	CreateAndUpdateGithubActionInput() *string
	CreateAndUpdateGoogleCalendarAction() *string
	SetCreateAndUpdateGoogleCalendarAction(val *string)
	CreateAndUpdateGoogleCalendarActionInput() *string
	CreateAndUpdateHubspotAction() *string
	SetCreateAndUpdateHubspotAction(val *string)
	CreateAndUpdateHubspotActionInput() *string
	CreateAndUpdateHuggingFaceAction() *string
	SetCreateAndUpdateHuggingFaceAction(val *string)
	CreateAndUpdateHuggingFaceActionInput() *string
	CreateAndUpdateIntercomAction() *string
	SetCreateAndUpdateIntercomAction(val *string)
	CreateAndUpdateIntercomActionInput() *string
	CreateAndUpdateJiraAction() *string
	SetCreateAndUpdateJiraAction(val *string)
	CreateAndUpdateJiraActionInput() *string
	CreateAndUpdateKnowledgeBases() *string
	SetCreateAndUpdateKnowledgeBases(val *string)
	CreateAndUpdateKnowledgeBasesInput() *string
	CreateAndUpdateLinearAction() *string
	SetCreateAndUpdateLinearAction(val *string)
	CreateAndUpdateLinearActionInput() *string
	CreateAndUpdateMcpAction() *string
	SetCreateAndUpdateMcpAction(val *string)
	CreateAndUpdateMcpActionInput() *string
	CreateAndUpdateMondayAction() *string
	SetCreateAndUpdateMondayAction(val *string)
	CreateAndUpdateMondayActionInput() *string
	CreateAndUpdateMsExchangeAction() *string
	SetCreateAndUpdateMsExchangeAction(val *string)
	CreateAndUpdateMsExchangeActionInput() *string
	CreateAndUpdateMsTeamsAction() *string
	SetCreateAndUpdateMsTeamsAction(val *string)
	CreateAndUpdateMsTeamsActionInput() *string
	CreateAndUpdateNewRelicAction() *string
	SetCreateAndUpdateNewRelicAction(val *string)
	CreateAndUpdateNewRelicActionInput() *string
	CreateAndUpdateNotionAction() *string
	SetCreateAndUpdateNotionAction(val *string)
	CreateAndUpdateNotionActionInput() *string
	CreateAndUpdateOneDriveAction() *string
	SetCreateAndUpdateOneDriveAction(val *string)
	CreateAndUpdateOneDriveActionInput() *string
	CreateAndUpdateOpenApiAction() *string
	SetCreateAndUpdateOpenApiAction(val *string)
	CreateAndUpdateOpenApiActionInput() *string
	CreateAndUpdatePagerDutyAction() *string
	SetCreateAndUpdatePagerDutyAction(val *string)
	CreateAndUpdatePagerDutyActionInput() *string
	CreateAndUpdateSalesforceAction() *string
	SetCreateAndUpdateSalesforceAction(val *string)
	CreateAndUpdateSalesforceActionInput() *string
	CreateAndUpdateSandPGlobalEnergyAction() *string
	SetCreateAndUpdateSandPGlobalEnergyAction(val *string)
	CreateAndUpdateSandPGlobalEnergyActionInput() *string
	CreateAndUpdateSandPgmiAction() *string
	SetCreateAndUpdateSandPgmiAction(val *string)
	CreateAndUpdateSandPgmiActionInput() *string
	CreateAndUpdateSapBillOfMaterialAction() *string
	SetCreateAndUpdateSapBillOfMaterialAction(val *string)
	CreateAndUpdateSapBillOfMaterialActionInput() *string
	CreateAndUpdateSapBusinessPartnerAction() *string
	SetCreateAndUpdateSapBusinessPartnerAction(val *string)
	CreateAndUpdateSapBusinessPartnerActionInput() *string
	CreateAndUpdateSapMaterialStockAction() *string
	SetCreateAndUpdateSapMaterialStockAction(val *string)
	CreateAndUpdateSapMaterialStockActionInput() *string
	CreateAndUpdateSapPhysicalInventoryAction() *string
	SetCreateAndUpdateSapPhysicalInventoryAction(val *string)
	CreateAndUpdateSapPhysicalInventoryActionInput() *string
	CreateAndUpdateSapProductMasterDataAction() *string
	SetCreateAndUpdateSapProductMasterDataAction(val *string)
	CreateAndUpdateSapProductMasterDataActionInput() *string
	CreateAndUpdateServiceNowAction() *string
	SetCreateAndUpdateServiceNowAction(val *string)
	CreateAndUpdateServiceNowActionInput() *string
	CreateAndUpdateSharePointAction() *string
	SetCreateAndUpdateSharePointAction(val *string)
	CreateAndUpdateSharePointActionInput() *string
	CreateAndUpdateSlackAction() *string
	SetCreateAndUpdateSlackAction(val *string)
	CreateAndUpdateSlackActionInput() *string
	CreateAndUpdateSmartsheetAction() *string
	SetCreateAndUpdateSmartsheetAction(val *string)
	CreateAndUpdateSmartsheetActionInput() *string
	CreateAndUpdateTextractAction() *string
	SetCreateAndUpdateTextractAction(val *string)
	CreateAndUpdateTextractActionInput() *string
	CreateAndUpdateThemes() *string
	SetCreateAndUpdateThemes(val *string)
	CreateAndUpdateThemesInput() *string
	CreateAndUpdateThresholdAlerts() *string
	SetCreateAndUpdateThresholdAlerts(val *string)
	CreateAndUpdateThresholdAlertsInput() *string
	CreateAndUpdateZendeskAction() *string
	SetCreateAndUpdateZendeskAction(val *string)
	CreateAndUpdateZendeskActionInput() *string
	CreateChatAgents() *string
	SetCreateChatAgents(val *string)
	CreateChatAgentsInput() *string
	CreateDashboardExecutiveSummaryWithQ() *string
	SetCreateDashboardExecutiveSummaryWithQ(val *string)
	CreateDashboardExecutiveSummaryWithQInput() *string
	CreateSharedFolders() *string
	SetCreateSharedFolders(val *string)
	CreateSharedFoldersInput() *string
	CreateSpaces() *string
	SetCreateSpaces(val *string)
	CreateSpacesInput() *string
	CreateSpiceDataset() *string
	SetCreateSpiceDataset(val *string)
	CreateSpiceDatasetInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dashboard() *string
	SetDashboard(val *string)
	DashboardInput() *string
	EditVisualWithQ() *string
	SetEditVisualWithQ(val *string)
	EditVisualWithQInput() *string
	ExportToCsv() *string
	SetExportToCsv(val *string)
	ExportToCsvInput() *string
	ExportToCsvInScheduledReports() *string
	SetExportToCsvInScheduledReports(val *string)
	ExportToCsvInScheduledReportsInput() *string
	ExportToExcel() *string
	SetExportToExcel(val *string)
	ExportToExcelInput() *string
	ExportToExcelInScheduledReports() *string
	SetExportToExcelInScheduledReports(val *string)
	ExportToExcelInScheduledReportsInput() *string
	ExportToPdf() *string
	SetExportToPdf(val *string)
	ExportToPdfInput() *string
	ExportToPdfInScheduledReports() *string
	SetExportToPdfInScheduledReports(val *string)
	ExportToPdfInScheduledReportsInput() *string
	Extension() *string
	SetExtension(val *string)
	ExtensionInput() *string
	FactSetAction() *string
	SetFactSetAction(val *string)
	FactSetActionInput() *string
	Flow() *string
	SetFlow(val *string)
	FlowInput() *string
	// Experimental.
	Fqn() *string
	GenericHttpAction() *string
	SetGenericHttpAction(val *string)
	GenericHttpActionInput() *string
	GithubAction() *string
	SetGithubAction(val *string)
	GithubActionInput() *string
	GoogleCalendarAction() *string
	SetGoogleCalendarAction(val *string)
	GoogleCalendarActionInput() *string
	HubspotAction() *string
	SetHubspotAction(val *string)
	HubspotActionInput() *string
	HuggingFaceAction() *string
	SetHuggingFaceAction(val *string)
	HuggingFaceActionInput() *string
	IncludeContentInScheduledReportsEmail() *string
	SetIncludeContentInScheduledReportsEmail(val *string)
	IncludeContentInScheduledReportsEmailInput() *string
	IntercomAction() *string
	SetIntercomAction(val *string)
	IntercomActionInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	InvokeAppsAiInference() *string
	SetInvokeAppsAiInference(val *string)
	InvokeAppsAiInferenceInput() *string
	JiraAction() *string
	SetJiraAction(val *string)
	JiraActionInput() *string
	KnowledgeBase() *string
	SetKnowledgeBase(val *string)
	KnowledgeBaseInput() *string
	LinearAction() *string
	SetLinearAction(val *string)
	LinearActionInput() *string
	ManageSharedFolders() *string
	SetManageSharedFolders(val *string)
	ManageSharedFoldersInput() *string
	McpAction() *string
	SetMcpAction(val *string)
	McpActionInput() *string
	MondayAction() *string
	SetMondayAction(val *string)
	MondayActionInput() *string
	MsExchangeAction() *string
	SetMsExchangeAction(val *string)
	MsExchangeActionInput() *string
	MsTeamsAction() *string
	SetMsTeamsAction(val *string)
	MsTeamsActionInput() *string
	NewRelicAction() *string
	SetNewRelicAction(val *string)
	NewRelicActionInput() *string
	NotionAction() *string
	SetNotionAction(val *string)
	NotionActionInput() *string
	OneDriveAction() *string
	SetOneDriveAction(val *string)
	OneDriveActionInput() *string
	OpenApiAction() *string
	SetOpenApiAction(val *string)
	OpenApiActionInput() *string
	PagerDutyAction() *string
	SetPagerDutyAction(val *string)
	PagerDutyActionInput() *string
	PerformFlowUiTask() *string
	SetPerformFlowUiTask(val *string)
	PerformFlowUiTaskInput() *string
	PrintReports() *string
	SetPrintReports(val *string)
	PrintReportsInput() *string
	PublishWithoutApproval() *string
	SetPublishWithoutApproval(val *string)
	PublishWithoutApprovalInput() *string
	RenameSharedFolders() *string
	SetRenameSharedFolders(val *string)
	RenameSharedFoldersInput() *string
	Research() *string
	SetResearch(val *string)
	ResearchInput() *string
	SalesforceAction() *string
	SetSalesforceAction(val *string)
	SalesforceActionInput() *string
	SandPGlobalEnergyAction() *string
	SetSandPGlobalEnergyAction(val *string)
	SandPGlobalEnergyActionInput() *string
	SandPgmiAction() *string
	SetSandPgmiAction(val *string)
	SandPgmiActionInput() *string
	SapBillOfMaterialAction() *string
	SetSapBillOfMaterialAction(val *string)
	SapBillOfMaterialActionInput() *string
	SapBusinessPartnerAction() *string
	SetSapBusinessPartnerAction(val *string)
	SapBusinessPartnerActionInput() *string
	SapMaterialStockAction() *string
	SetSapMaterialStockAction(val *string)
	SapMaterialStockActionInput() *string
	SapPhysicalInventoryAction() *string
	SetSapPhysicalInventoryAction(val *string)
	SapPhysicalInventoryActionInput() *string
	SapProductMasterDataAction() *string
	SetSapProductMasterDataAction(val *string)
	SapProductMasterDataActionInput() *string
	ServiceNowAction() *string
	SetServiceNowAction(val *string)
	ServiceNowActionInput() *string
	ShareAmazonBedrockArsAction() *string
	SetShareAmazonBedrockArsAction(val *string)
	ShareAmazonBedrockArsActionInput() *string
	ShareAmazonBedrockFsAction() *string
	SetShareAmazonBedrockFsAction(val *string)
	ShareAmazonBedrockFsActionInput() *string
	ShareAmazonBedrockKrsAction() *string
	SetShareAmazonBedrockKrsAction(val *string)
	ShareAmazonBedrockKrsActionInput() *string
	ShareAmazonSThreeAction() *string
	SetShareAmazonSThreeAction(val *string)
	ShareAmazonSThreeActionInput() *string
	ShareAnalyses() *string
	SetShareAnalyses(val *string)
	ShareAnalysesInput() *string
	ShareApps() *string
	SetShareApps(val *string)
	ShareAppsInput() *string
	ShareAsanaAction() *string
	SetShareAsanaAction(val *string)
	ShareAsanaActionInput() *string
	ShareBambooHrAction() *string
	SetShareBambooHrAction(val *string)
	ShareBambooHrActionInput() *string
	ShareBoxAgentAction() *string
	SetShareBoxAgentAction(val *string)
	ShareBoxAgentActionInput() *string
	ShareCanvaAgentAction() *string
	SetShareCanvaAgentAction(val *string)
	ShareCanvaAgentActionInput() *string
	ShareChatAgents() *string
	SetShareChatAgents(val *string)
	ShareChatAgentsInput() *string
	ShareComprehendAction() *string
	SetShareComprehendAction(val *string)
	ShareComprehendActionInput() *string
	ShareComprehendMedicalAction() *string
	SetShareComprehendMedicalAction(val *string)
	ShareComprehendMedicalActionInput() *string
	ShareConfluenceAction() *string
	SetShareConfluenceAction(val *string)
	ShareConfluenceActionInput() *string
	ShareDashboards() *string
	SetShareDashboards(val *string)
	ShareDashboardsInput() *string
	ShareDatasets() *string
	SetShareDatasets(val *string)
	ShareDatasetsInput() *string
	ShareDataSources() *string
	SetShareDataSources(val *string)
	ShareDataSourcesInput() *string
	ShareFactSetAction() *string
	SetShareFactSetAction(val *string)
	ShareFactSetActionInput() *string
	ShareGenericHttpAction() *string
	SetShareGenericHttpAction(val *string)
	ShareGenericHttpActionInput() *string
	ShareGithubAction() *string
	SetShareGithubAction(val *string)
	ShareGithubActionInput() *string
	ShareGoogleCalendarAction() *string
	SetShareGoogleCalendarAction(val *string)
	ShareGoogleCalendarActionInput() *string
	ShareHubspotAction() *string
	SetShareHubspotAction(val *string)
	ShareHubspotActionInput() *string
	ShareHuggingFaceAction() *string
	SetShareHuggingFaceAction(val *string)
	ShareHuggingFaceActionInput() *string
	ShareIntercomAction() *string
	SetShareIntercomAction(val *string)
	ShareIntercomActionInput() *string
	ShareJiraAction() *string
	SetShareJiraAction(val *string)
	ShareJiraActionInput() *string
	ShareKnowledgeBases() *string
	SetShareKnowledgeBases(val *string)
	ShareKnowledgeBasesInput() *string
	ShareLinearAction() *string
	SetShareLinearAction(val *string)
	ShareLinearActionInput() *string
	ShareMcpAction() *string
	SetShareMcpAction(val *string)
	ShareMcpActionInput() *string
	ShareMondayAction() *string
	SetShareMondayAction(val *string)
	ShareMondayActionInput() *string
	ShareMsExchangeAction() *string
	SetShareMsExchangeAction(val *string)
	ShareMsExchangeActionInput() *string
	ShareMsTeamsAction() *string
	SetShareMsTeamsAction(val *string)
	ShareMsTeamsActionInput() *string
	ShareNewRelicAction() *string
	SetShareNewRelicAction(val *string)
	ShareNewRelicActionInput() *string
	ShareNotionAction() *string
	SetShareNotionAction(val *string)
	ShareNotionActionInput() *string
	ShareOneDriveAction() *string
	SetShareOneDriveAction(val *string)
	ShareOneDriveActionInput() *string
	ShareOpenApiAction() *string
	SetShareOpenApiAction(val *string)
	ShareOpenApiActionInput() *string
	SharePagerDutyAction() *string
	SetSharePagerDutyAction(val *string)
	SharePagerDutyActionInput() *string
	SharePointAction() *string
	SetSharePointAction(val *string)
	SharePointActionInput() *string
	ShareSalesforceAction() *string
	SetShareSalesforceAction(val *string)
	ShareSalesforceActionInput() *string
	ShareSandPGlobalEnergyAction() *string
	SetShareSandPGlobalEnergyAction(val *string)
	ShareSandPGlobalEnergyActionInput() *string
	ShareSandPgmiAction() *string
	SetShareSandPgmiAction(val *string)
	ShareSandPgmiActionInput() *string
	ShareSapBillOfMaterialAction() *string
	SetShareSapBillOfMaterialAction(val *string)
	ShareSapBillOfMaterialActionInput() *string
	ShareSapBusinessPartnerAction() *string
	SetShareSapBusinessPartnerAction(val *string)
	ShareSapBusinessPartnerActionInput() *string
	ShareSapMaterialStockAction() *string
	SetShareSapMaterialStockAction(val *string)
	ShareSapMaterialStockActionInput() *string
	ShareSapPhysicalInventoryAction() *string
	SetShareSapPhysicalInventoryAction(val *string)
	ShareSapPhysicalInventoryActionInput() *string
	ShareSapProductMasterDataAction() *string
	SetShareSapProductMasterDataAction(val *string)
	ShareSapProductMasterDataActionInput() *string
	ShareServiceNowAction() *string
	SetShareServiceNowAction(val *string)
	ShareServiceNowActionInput() *string
	ShareSharePointAction() *string
	SetShareSharePointAction(val *string)
	ShareSharePointActionInput() *string
	ShareSlackAction() *string
	SetShareSlackAction(val *string)
	ShareSlackActionInput() *string
	ShareSmartsheetAction() *string
	SetShareSmartsheetAction(val *string)
	ShareSmartsheetActionInput() *string
	ShareSpaces() *string
	SetShareSpaces(val *string)
	ShareSpacesInput() *string
	ShareTextractAction() *string
	SetShareTextractAction(val *string)
	ShareTextractActionInput() *string
	ShareZendeskAction() *string
	SetShareZendeskAction(val *string)
	ShareZendeskActionInput() *string
	SlackAction() *string
	SetSlackAction(val *string)
	SlackActionInput() *string
	SmartsheetAction() *string
	SetSmartsheetAction(val *string)
	SmartsheetActionInput() *string
	Space() *string
	SetSpace(val *string)
	SpaceInput() *string
	SubscribeDashboardEmailReports() *string
	SetSubscribeDashboardEmailReports(val *string)
	SubscribeDashboardEmailReportsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextractAction() *string
	SetTextractAction(val *string)
	TextractActionInput() *string
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
	UseAgentWebSearch() *string
	SetUseAgentWebSearch(val *string)
	UseAgentWebSearchInput() *string
	UseAmazonBedrockArsAction() *string
	SetUseAmazonBedrockArsAction(val *string)
	UseAmazonBedrockArsActionInput() *string
	UseAmazonBedrockFsAction() *string
	SetUseAmazonBedrockFsAction(val *string)
	UseAmazonBedrockFsActionInput() *string
	UseAmazonBedrockKrsAction() *string
	SetUseAmazonBedrockKrsAction(val *string)
	UseAmazonBedrockKrsActionInput() *string
	UseAmazonSThreeAction() *string
	SetUseAmazonSThreeAction(val *string)
	UseAmazonSThreeActionInput() *string
	UseAsanaAction() *string
	SetUseAsanaAction(val *string)
	UseAsanaActionInput() *string
	UseBambooHrAction() *string
	SetUseBambooHrAction(val *string)
	UseBambooHrActionInput() *string
	UseBedrockModels() *string
	SetUseBedrockModels(val *string)
	UseBedrockModelsInput() *string
	UseBoxAgentAction() *string
	SetUseBoxAgentAction(val *string)
	UseBoxAgentActionInput() *string
	UseCanvaAgentAction() *string
	SetUseCanvaAgentAction(val *string)
	UseCanvaAgentActionInput() *string
	UseComprehendAction() *string
	SetUseComprehendAction(val *string)
	UseComprehendActionInput() *string
	UseComprehendMedicalAction() *string
	SetUseComprehendMedicalAction(val *string)
	UseComprehendMedicalActionInput() *string
	UseConfluenceAction() *string
	SetUseConfluenceAction(val *string)
	UseConfluenceActionInput() *string
	UseFactSetAction() *string
	SetUseFactSetAction(val *string)
	UseFactSetActionInput() *string
	UseGenericHttpAction() *string
	SetUseGenericHttpAction(val *string)
	UseGenericHttpActionInput() *string
	UseGithubAction() *string
	SetUseGithubAction(val *string)
	UseGithubActionInput() *string
	UseGoogleCalendarAction() *string
	SetUseGoogleCalendarAction(val *string)
	UseGoogleCalendarActionInput() *string
	UseHubspotAction() *string
	SetUseHubspotAction(val *string)
	UseHubspotActionInput() *string
	UseHuggingFaceAction() *string
	SetUseHuggingFaceAction(val *string)
	UseHuggingFaceActionInput() *string
	UseIntercomAction() *string
	SetUseIntercomAction(val *string)
	UseIntercomActionInput() *string
	UseJiraAction() *string
	SetUseJiraAction(val *string)
	UseJiraActionInput() *string
	UseLinearAction() *string
	SetUseLinearAction(val *string)
	UseLinearActionInput() *string
	UseMcpAction() *string
	SetUseMcpAction(val *string)
	UseMcpActionInput() *string
	UseMondayAction() *string
	SetUseMondayAction(val *string)
	UseMondayActionInput() *string
	UseMsExchangeAction() *string
	SetUseMsExchangeAction(val *string)
	UseMsExchangeActionInput() *string
	UseMsTeamsAction() *string
	SetUseMsTeamsAction(val *string)
	UseMsTeamsActionInput() *string
	UseNewRelicAction() *string
	SetUseNewRelicAction(val *string)
	UseNewRelicActionInput() *string
	UseNotionAction() *string
	SetUseNotionAction(val *string)
	UseNotionActionInput() *string
	UseOneDriveAction() *string
	SetUseOneDriveAction(val *string)
	UseOneDriveActionInput() *string
	UseOpenApiAction() *string
	SetUseOpenApiAction(val *string)
	UseOpenApiActionInput() *string
	UsePagerDutyAction() *string
	SetUsePagerDutyAction(val *string)
	UsePagerDutyActionInput() *string
	UseSalesforceAction() *string
	SetUseSalesforceAction(val *string)
	UseSalesforceActionInput() *string
	UseSandPGlobalEnergyAction() *string
	SetUseSandPGlobalEnergyAction(val *string)
	UseSandPGlobalEnergyActionInput() *string
	UseSandPgmiAction() *string
	SetUseSandPgmiAction(val *string)
	UseSandPgmiActionInput() *string
	UseSapBillOfMaterialAction() *string
	SetUseSapBillOfMaterialAction(val *string)
	UseSapBillOfMaterialActionInput() *string
	UseSapBusinessPartnerAction() *string
	SetUseSapBusinessPartnerAction(val *string)
	UseSapBusinessPartnerActionInput() *string
	UseSapMaterialStockAction() *string
	SetUseSapMaterialStockAction(val *string)
	UseSapMaterialStockActionInput() *string
	UseSapPhysicalInventoryAction() *string
	SetUseSapPhysicalInventoryAction(val *string)
	UseSapPhysicalInventoryActionInput() *string
	UseSapProductMasterDataAction() *string
	SetUseSapProductMasterDataAction(val *string)
	UseSapProductMasterDataActionInput() *string
	UseServiceNowAction() *string
	SetUseServiceNowAction(val *string)
	UseServiceNowActionInput() *string
	UseSharePointAction() *string
	SetUseSharePointAction(val *string)
	UseSharePointActionInput() *string
	UseSlackAction() *string
	SetUseSlackAction(val *string)
	UseSlackActionInput() *string
	UseSmartsheetAction() *string
	SetUseSmartsheetAction(val *string)
	UseSmartsheetActionInput() *string
	UseTextractAction() *string
	SetUseTextractAction(val *string)
	UseTextractActionInput() *string
	UseZendeskAction() *string
	SetUseZendeskAction(val *string)
	UseZendeskActionInput() *string
	ViewAccountSpiceCapacity() *string
	SetViewAccountSpiceCapacity(val *string)
	ViewAccountSpiceCapacityInput() *string
	ZendeskAction() *string
	SetZendeskAction(val *string)
	ZendeskActionInput() *string
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
	ResetAccessAppsNativeDataStore()
	ResetAction()
	ResetAddOrRunAnomalyDetectionForAnalyses()
	ResetAmazonBedrockArsAction()
	ResetAmazonBedrockFsAction()
	ResetAmazonBedrockKrsAction()
	ResetAmazonSThreeAction()
	ResetAnalysis()
	ResetApproveFlowShareRequests()
	ResetApps()
	ResetAsanaAction()
	ResetAutomate()
	ResetBambooHrAction()
	ResetBoxAgentAction()
	ResetBuildCalculatedFieldWithQ()
	ResetCanvaAgentAction()
	ResetChatAgent()
	ResetComprehendAction()
	ResetComprehendMedicalAction()
	ResetConfluenceAction()
	ResetCreateAndUpdateAmazonBedrockArsAction()
	ResetCreateAndUpdateAmazonBedrockFsAction()
	ResetCreateAndUpdateAmazonBedrockKrsAction()
	ResetCreateAndUpdateAmazonSThreeAction()
	ResetCreateAndUpdateApps()
	ResetCreateAndUpdateAsanaAction()
	ResetCreateAndUpdateBambooHrAction()
	ResetCreateAndUpdateBoxAgentAction()
	ResetCreateAndUpdateCanvaAgentAction()
	ResetCreateAndUpdateComprehendAction()
	ResetCreateAndUpdateComprehendMedicalAction()
	ResetCreateAndUpdateConfluenceAction()
	ResetCreateAndUpdateDashboardEmailReports()
	ResetCreateAndUpdateDatasets()
	ResetCreateAndUpdateDataSources()
	ResetCreateAndUpdateFactSetAction()
	ResetCreateAndUpdateGenericHttpAction()
	ResetCreateAndUpdateGithubAction()
	ResetCreateAndUpdateGoogleCalendarAction()
	ResetCreateAndUpdateHubspotAction()
	ResetCreateAndUpdateHuggingFaceAction()
	ResetCreateAndUpdateIntercomAction()
	ResetCreateAndUpdateJiraAction()
	ResetCreateAndUpdateKnowledgeBases()
	ResetCreateAndUpdateLinearAction()
	ResetCreateAndUpdateMcpAction()
	ResetCreateAndUpdateMondayAction()
	ResetCreateAndUpdateMsExchangeAction()
	ResetCreateAndUpdateMsTeamsAction()
	ResetCreateAndUpdateNewRelicAction()
	ResetCreateAndUpdateNotionAction()
	ResetCreateAndUpdateOneDriveAction()
	ResetCreateAndUpdateOpenApiAction()
	ResetCreateAndUpdatePagerDutyAction()
	ResetCreateAndUpdateSalesforceAction()
	ResetCreateAndUpdateSandPGlobalEnergyAction()
	ResetCreateAndUpdateSandPgmiAction()
	ResetCreateAndUpdateSapBillOfMaterialAction()
	ResetCreateAndUpdateSapBusinessPartnerAction()
	ResetCreateAndUpdateSapMaterialStockAction()
	ResetCreateAndUpdateSapPhysicalInventoryAction()
	ResetCreateAndUpdateSapProductMasterDataAction()
	ResetCreateAndUpdateServiceNowAction()
	ResetCreateAndUpdateSharePointAction()
	ResetCreateAndUpdateSlackAction()
	ResetCreateAndUpdateSmartsheetAction()
	ResetCreateAndUpdateTextractAction()
	ResetCreateAndUpdateThemes()
	ResetCreateAndUpdateThresholdAlerts()
	ResetCreateAndUpdateZendeskAction()
	ResetCreateChatAgents()
	ResetCreateDashboardExecutiveSummaryWithQ()
	ResetCreateSharedFolders()
	ResetCreateSpaces()
	ResetCreateSpiceDataset()
	ResetDashboard()
	ResetEditVisualWithQ()
	ResetExportToCsv()
	ResetExportToCsvInScheduledReports()
	ResetExportToExcel()
	ResetExportToExcelInScheduledReports()
	ResetExportToPdf()
	ResetExportToPdfInScheduledReports()
	ResetExtension()
	ResetFactSetAction()
	ResetFlow()
	ResetGenericHttpAction()
	ResetGithubAction()
	ResetGoogleCalendarAction()
	ResetHubspotAction()
	ResetHuggingFaceAction()
	ResetIncludeContentInScheduledReportsEmail()
	ResetIntercomAction()
	ResetInvokeAppsAiInference()
	ResetJiraAction()
	ResetKnowledgeBase()
	ResetLinearAction()
	ResetManageSharedFolders()
	ResetMcpAction()
	ResetMondayAction()
	ResetMsExchangeAction()
	ResetMsTeamsAction()
	ResetNewRelicAction()
	ResetNotionAction()
	ResetOneDriveAction()
	ResetOpenApiAction()
	ResetPagerDutyAction()
	ResetPerformFlowUiTask()
	ResetPrintReports()
	ResetPublishWithoutApproval()
	ResetRenameSharedFolders()
	ResetResearch()
	ResetSalesforceAction()
	ResetSandPGlobalEnergyAction()
	ResetSandPgmiAction()
	ResetSapBillOfMaterialAction()
	ResetSapBusinessPartnerAction()
	ResetSapMaterialStockAction()
	ResetSapPhysicalInventoryAction()
	ResetSapProductMasterDataAction()
	ResetServiceNowAction()
	ResetShareAmazonBedrockArsAction()
	ResetShareAmazonBedrockFsAction()
	ResetShareAmazonBedrockKrsAction()
	ResetShareAmazonSThreeAction()
	ResetShareAnalyses()
	ResetShareApps()
	ResetShareAsanaAction()
	ResetShareBambooHrAction()
	ResetShareBoxAgentAction()
	ResetShareCanvaAgentAction()
	ResetShareChatAgents()
	ResetShareComprehendAction()
	ResetShareComprehendMedicalAction()
	ResetShareConfluenceAction()
	ResetShareDashboards()
	ResetShareDatasets()
	ResetShareDataSources()
	ResetShareFactSetAction()
	ResetShareGenericHttpAction()
	ResetShareGithubAction()
	ResetShareGoogleCalendarAction()
	ResetShareHubspotAction()
	ResetShareHuggingFaceAction()
	ResetShareIntercomAction()
	ResetShareJiraAction()
	ResetShareKnowledgeBases()
	ResetShareLinearAction()
	ResetShareMcpAction()
	ResetShareMondayAction()
	ResetShareMsExchangeAction()
	ResetShareMsTeamsAction()
	ResetShareNewRelicAction()
	ResetShareNotionAction()
	ResetShareOneDriveAction()
	ResetShareOpenApiAction()
	ResetSharePagerDutyAction()
	ResetSharePointAction()
	ResetShareSalesforceAction()
	ResetShareSandPGlobalEnergyAction()
	ResetShareSandPgmiAction()
	ResetShareSapBillOfMaterialAction()
	ResetShareSapBusinessPartnerAction()
	ResetShareSapMaterialStockAction()
	ResetShareSapPhysicalInventoryAction()
	ResetShareSapProductMasterDataAction()
	ResetShareServiceNowAction()
	ResetShareSharePointAction()
	ResetShareSlackAction()
	ResetShareSmartsheetAction()
	ResetShareSpaces()
	ResetShareTextractAction()
	ResetShareZendeskAction()
	ResetSlackAction()
	ResetSmartsheetAction()
	ResetSpace()
	ResetSubscribeDashboardEmailReports()
	ResetTextractAction()
	ResetTopic()
	ResetUseAgentWebSearch()
	ResetUseAmazonBedrockArsAction()
	ResetUseAmazonBedrockFsAction()
	ResetUseAmazonBedrockKrsAction()
	ResetUseAmazonSThreeAction()
	ResetUseAsanaAction()
	ResetUseBambooHrAction()
	ResetUseBedrockModels()
	ResetUseBoxAgentAction()
	ResetUseCanvaAgentAction()
	ResetUseComprehendAction()
	ResetUseComprehendMedicalAction()
	ResetUseConfluenceAction()
	ResetUseFactSetAction()
	ResetUseGenericHttpAction()
	ResetUseGithubAction()
	ResetUseGoogleCalendarAction()
	ResetUseHubspotAction()
	ResetUseHuggingFaceAction()
	ResetUseIntercomAction()
	ResetUseJiraAction()
	ResetUseLinearAction()
	ResetUseMcpAction()
	ResetUseMondayAction()
	ResetUseMsExchangeAction()
	ResetUseMsTeamsAction()
	ResetUseNewRelicAction()
	ResetUseNotionAction()
	ResetUseOneDriveAction()
	ResetUseOpenApiAction()
	ResetUsePagerDutyAction()
	ResetUseSalesforceAction()
	ResetUseSandPGlobalEnergyAction()
	ResetUseSandPgmiAction()
	ResetUseSapBillOfMaterialAction()
	ResetUseSapBusinessPartnerAction()
	ResetUseSapMaterialStockAction()
	ResetUseSapPhysicalInventoryAction()
	ResetUseSapProductMasterDataAction()
	ResetUseServiceNowAction()
	ResetUseSharePointAction()
	ResetUseSlackAction()
	ResetUseSmartsheetAction()
	ResetUseTextractAction()
	ResetUseZendeskAction()
	ResetViewAccountSpiceCapacity()
	ResetZendeskAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightCustomPermissionsCapabilitiesOutputReference
type jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AccessAppsNativeDataStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessAppsNativeDataStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AccessAppsNativeDataStoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessAppsNativeDataStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AddOrRunAnomalyDetectionForAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AddOrRunAnomalyDetectionForAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockArsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockArsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockFsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockFsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonBedrockKrsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonBedrockKrsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AmazonSThreeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amazonSThreeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Analysis() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AnalysisInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ApproveFlowShareRequests() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveFlowShareRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ApproveFlowShareRequestsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveFlowShareRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Apps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AppsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AsanaActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"asanaActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Automate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AutomateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BambooHrActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bambooHrActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BoxAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boxAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BuildCalculatedFieldWithQ() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCalculatedFieldWithQ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BuildCalculatedFieldWithQInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCalculatedFieldWithQInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CanvaAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"canvaAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ChatAgent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chatAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ChatAgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chatAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComprehendActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comprehendActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComprehendMedicalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comprehendMedicalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ConfluenceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confluenceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockArsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockArsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockFsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockFsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonBedrockKrsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonBedrockKrsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAmazonSThreeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAmazonSThreeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateApps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAppsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAsanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateAsanaActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateAsanaActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBambooHrActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBambooHrActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBoxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBoxAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBoxAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateCanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateCanvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateCanvaAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateCanvaAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateComprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateComprehendActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateComprehendActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateComprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateComprehendMedicalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateComprehendMedicalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateConfluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateConfluenceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateConfluenceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDashboardEmailReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDatasetsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDataSourcesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateFactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateFactSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateFactSetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateFactSetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGenericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGenericHttpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGenericHttpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGithubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGithubActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGithubActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGoogleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGoogleCalendarActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGoogleCalendarActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateHubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateHubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateHubspotActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateHubspotActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateHuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateHuggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateHuggingFaceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateHuggingFaceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateIntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateIntercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateIntercomActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateIntercomActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateJiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateJiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateJiraActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateJiraActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateKnowledgeBases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateKnowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateKnowledgeBasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateKnowledgeBasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateLinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateLinearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateLinearActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateLinearActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMcpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMcpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMcpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMondayActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMondayActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMsExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMsExchangeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMsExchangeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMsTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateMsTeamsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateMsTeamsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateNewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateNewRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateNewRelicActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateNewRelicActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateNotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateNotionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateNotionActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateNotionActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOneDriveActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOneDriveActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOpenApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOpenApiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOpenApiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdatePagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdatePagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdatePagerDutyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdatePagerDutyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSalesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSalesforceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSalesforceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSandPGlobalEnergyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSandPGlobalEnergyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSandPgmiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSandPgmiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapBillOfMaterialActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapBillOfMaterialActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapBusinessPartnerActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapBusinessPartnerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapMaterialStockActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapMaterialStockActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapPhysicalInventoryActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapPhysicalInventoryActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSapProductMasterDataActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSapProductMasterDataActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateServiceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateServiceNowActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateServiceNowActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSharePointActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSharePointActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSlackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSlackActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSlackActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSmartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSmartsheetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSmartsheetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateTextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateTextractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateTextractActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateTextractActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThemes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThemesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThresholdAlerts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThresholdAlertsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateZendeskAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateZendeskActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateZendeskActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateChatAgents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createChatAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateChatAgentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createChatAgentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDashboardExecutiveSummaryWithQ() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDashboardExecutiveSummaryWithQ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDashboardExecutiveSummaryWithQInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDashboardExecutiveSummaryWithQInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSharedFoldersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpaces() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpacesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpiceDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpiceDatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Dashboard() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DashboardInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) EditVisualWithQ() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editVisualWithQ",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) EditVisualWithQInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editVisualWithQInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Extension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) FactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"factSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) FactSetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"factSetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Flow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) FlowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GenericHttpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genericHttpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GithubActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"githubActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleCalendarActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleCalendarActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) HubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) HubspotActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hubspotActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) HuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) HuggingFaceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"huggingFaceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IncludeContentInScheduledReportsEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IncludeContentInScheduledReportsEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IntercomActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intercomActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InvokeAppsAiInference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeAppsAiInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InvokeAppsAiInferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeAppsAiInferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) JiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) JiraActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) KnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) LinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) LinearActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linearActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ManageSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manageSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ManageSharedFoldersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manageSharedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) McpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) McpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mcpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MondayActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mondayActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MsExchangeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msExchangeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MsTeamsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msTeamsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) NewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) NewRelicActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newRelicActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) NotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) NotionActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notionActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OneDriveActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneDriveActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OpenApiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openApiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PagerDutyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pagerDutyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PerformFlowUiTask() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performFlowUiTask",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PerformFlowUiTaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performFlowUiTaskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PrintReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PrintReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PublishWithoutApproval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishWithoutApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PublishWithoutApprovalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishWithoutApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RenameSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RenameSharedFoldersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Research() *string {
	var returns *string
	_jsii_.Get(
		j,
		"research",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"researchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SalesforceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SandPGlobalEnergyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sandPGlobalEnergyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SandPgmiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sandPgmiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapBillOfMaterialActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapBillOfMaterialActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapBusinessPartnerActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapBusinessPartnerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapMaterialStockActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapMaterialStockActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapPhysicalInventoryActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapPhysicalInventoryActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapProductMasterDataActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapProductMasterDataActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ServiceNowActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockArsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockArsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockFsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockFsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonBedrockKrsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonBedrockKrsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAmazonSThreeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAmazonSThreeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareApps() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareApps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAppsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAppsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAsanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAsanaActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAsanaActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBambooHrActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBambooHrActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBoxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBoxAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBoxAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareCanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareCanvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareCanvaAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareCanvaAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareChatAgents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareChatAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareChatAgentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareChatAgentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareComprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareComprehendActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareComprehendActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareComprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareComprehendMedicalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareComprehendMedicalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareConfluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareConfluenceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareConfluenceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDashboards() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDashboardsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatasetsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDataSourcesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareFactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareFactSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareFactSetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareFactSetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGenericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGenericHttpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGenericHttpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGithubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGithubActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGithubActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleCalendarActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleCalendarActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareHubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareHubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareHubspotActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareHubspotActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareHuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareHuggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareHuggingFaceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareHuggingFaceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareIntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareIntercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareIntercomActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareIntercomActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareJiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareJiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareJiraActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareJiraActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareKnowledgeBases() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareKnowledgeBases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareKnowledgeBasesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareKnowledgeBasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareLinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareLinearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareLinearActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareLinearActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMcpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMcpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMcpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMondayActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMondayActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMsExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMsExchangeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMsExchangeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMsTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMsTeamsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMsTeamsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareNewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNewRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareNewRelicActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNewRelicActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareNotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNotionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareNotionActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNotionActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOneDriveActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOneDriveActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOpenApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOpenApiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOpenApiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePagerDutyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePagerDutyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePointActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSalesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSalesforceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSalesforceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSandPGlobalEnergyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSandPGlobalEnergyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSandPgmiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSandPgmiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapBillOfMaterialActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapBillOfMaterialActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapBusinessPartnerActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapBusinessPartnerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapMaterialStockActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapMaterialStockActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapPhysicalInventoryActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapPhysicalInventoryActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapProductMasterDataActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapProductMasterDataActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareServiceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareServiceNowActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareServiceNowActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSharePointActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSharePointActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSlackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSlackActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSlackActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSmartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSmartsheetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSmartsheetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSpaces() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSpaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSpacesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSpacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTextractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTextractActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTextractActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareZendeskAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareZendeskActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareZendeskActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SlackActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slackActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SmartsheetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smartsheetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Space() *string {
	var returns *string
	_jsii_.Get(
		j,
		"space",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SpaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SubscribeDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SubscribeDashboardEmailReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TextractActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textractActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Topic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAgentWebSearch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAgentWebSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAgentWebSearchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAgentWebSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockArsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockArsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockArsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockArsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockFsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockFsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockFsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockFsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockKrsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockKrsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonBedrockKrsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonBedrockKrsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonSThreeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonSThreeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAmazonSThreeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAmazonSThreeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAsanaAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAsanaAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseAsanaActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useAsanaActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBambooHrAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBambooHrAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBambooHrActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBambooHrActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBedrockModels() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBedrockModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBedrockModelsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBedrockModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBoxAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBoxAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBoxAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBoxAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseCanvaAgentAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useCanvaAgentAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseCanvaAgentActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useCanvaAgentActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseComprehendAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useComprehendAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseComprehendActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useComprehendActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseComprehendMedicalAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useComprehendMedicalAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseComprehendMedicalActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useComprehendMedicalActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseConfluenceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useConfluenceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseConfluenceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useConfluenceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseFactSetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useFactSetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseFactSetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useFactSetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGenericHttpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGenericHttpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGenericHttpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGenericHttpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGithubAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGithubAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGithubActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGithubActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGoogleCalendarAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGoogleCalendarAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGoogleCalendarActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGoogleCalendarActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseHubspotAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useHubspotAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseHubspotActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useHubspotActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseHuggingFaceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useHuggingFaceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseHuggingFaceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useHuggingFaceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseIntercomAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useIntercomAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseIntercomActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useIntercomActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseJiraAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useJiraAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseJiraActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useJiraActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseLinearAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useLinearAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseLinearActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useLinearActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMcpAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMcpAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMcpActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMcpActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMondayAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMondayAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMondayActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMondayActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMsExchangeAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMsExchangeAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMsExchangeActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMsExchangeActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMsTeamsAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMsTeamsAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseMsTeamsActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useMsTeamsActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseNewRelicAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useNewRelicAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseNewRelicActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useNewRelicActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseNotionAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useNotionAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseNotionActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useNotionActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOneDriveAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOneDriveAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOneDriveActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOneDriveActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOpenApiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOpenApiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOpenApiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOpenApiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UsePagerDutyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePagerDutyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UsePagerDutyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePagerDutyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSalesforceAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSalesforceAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSalesforceActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSalesforceActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSandPGlobalEnergyAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSandPGlobalEnergyAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSandPGlobalEnergyActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSandPGlobalEnergyActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSandPgmiAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSandPgmiAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSandPgmiActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSandPgmiActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapBillOfMaterialAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapBillOfMaterialAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapBillOfMaterialActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapBillOfMaterialActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapBusinessPartnerAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapBusinessPartnerAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapBusinessPartnerActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapBusinessPartnerActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapMaterialStockAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapMaterialStockAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapMaterialStockActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapMaterialStockActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapPhysicalInventoryAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapPhysicalInventoryAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapPhysicalInventoryActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapPhysicalInventoryActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapProductMasterDataAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapProductMasterDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSapProductMasterDataActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSapProductMasterDataActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseServiceNowAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useServiceNowAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseServiceNowActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useServiceNowActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSharePointAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSharePointAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSharePointActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSharePointActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSlackAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSlackAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSlackActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSlackActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSmartsheetAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSmartsheetAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSmartsheetActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSmartsheetActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseTextractAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useTextractAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseTextractActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useTextractActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useZendeskAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseZendeskActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useZendeskActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ViewAccountSpiceCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ViewAccountSpiceCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ZendeskAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendeskAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ZendeskActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zendeskActionInput",
		&returns,
	)
	return returns
}


func NewQuicksightCustomPermissionsCapabilitiesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) QuicksightCustomPermissionsCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightCustomPermissionsCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightCustomPermissions.QuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightCustomPermissionsCapabilitiesOutputReference_Override(q QuicksightCustomPermissionsCapabilitiesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightCustomPermissions.QuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAccessAppsNativeDataStore(val *string) {
	if err := j.validateSetAccessAppsNativeDataStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessAppsNativeDataStore",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAddOrRunAnomalyDetectionForAnalyses(val *string) {
	if err := j.validateSetAddOrRunAnomalyDetectionForAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAmazonBedrockArsAction(val *string) {
	if err := j.validateSetAmazonBedrockArsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonBedrockArsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAmazonBedrockFsAction(val *string) {
	if err := j.validateSetAmazonBedrockFsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonBedrockFsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAmazonBedrockKrsAction(val *string) {
	if err := j.validateSetAmazonBedrockKrsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonBedrockKrsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAmazonSThreeAction(val *string) {
	if err := j.validateSetAmazonSThreeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonSThreeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAnalysis(val *string) {
	if err := j.validateSetAnalysisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analysis",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetApproveFlowShareRequests(val *string) {
	if err := j.validateSetApproveFlowShareRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approveFlowShareRequests",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetApps(val *string) {
	if err := j.validateSetAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apps",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAsanaAction(val *string) {
	if err := j.validateSetAsanaActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asanaAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAutomate(val *string) {
	if err := j.validateSetAutomateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automate",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetBambooHrAction(val *string) {
	if err := j.validateSetBambooHrActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bambooHrAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetBoxAgentAction(val *string) {
	if err := j.validateSetBoxAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boxAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetBuildCalculatedFieldWithQ(val *string) {
	if err := j.validateSetBuildCalculatedFieldWithQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCalculatedFieldWithQ",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCanvaAgentAction(val *string) {
	if err := j.validateSetCanvaAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"canvaAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetChatAgent(val *string) {
	if err := j.validateSetChatAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chatAgent",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetComprehendAction(val *string) {
	if err := j.validateSetComprehendActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comprehendAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetComprehendMedicalAction(val *string) {
	if err := j.validateSetComprehendMedicalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comprehendMedicalAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetConfluenceAction(val *string) {
	if err := j.validateSetConfluenceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confluenceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateAmazonBedrockArsAction(val *string) {
	if err := j.validateSetCreateAndUpdateAmazonBedrockArsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateAmazonBedrockArsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateAmazonBedrockFsAction(val *string) {
	if err := j.validateSetCreateAndUpdateAmazonBedrockFsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateAmazonBedrockFsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateAmazonBedrockKrsAction(val *string) {
	if err := j.validateSetCreateAndUpdateAmazonBedrockKrsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateAmazonBedrockKrsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateAmazonSThreeAction(val *string) {
	if err := j.validateSetCreateAndUpdateAmazonSThreeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateAmazonSThreeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateApps(val *string) {
	if err := j.validateSetCreateAndUpdateAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateApps",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateAsanaAction(val *string) {
	if err := j.validateSetCreateAndUpdateAsanaActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateAsanaAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateBambooHrAction(val *string) {
	if err := j.validateSetCreateAndUpdateBambooHrActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateBambooHrAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateBoxAgentAction(val *string) {
	if err := j.validateSetCreateAndUpdateBoxAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateBoxAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateCanvaAgentAction(val *string) {
	if err := j.validateSetCreateAndUpdateCanvaAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateCanvaAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateComprehendAction(val *string) {
	if err := j.validateSetCreateAndUpdateComprehendActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateComprehendAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateComprehendMedicalAction(val *string) {
	if err := j.validateSetCreateAndUpdateComprehendMedicalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateComprehendMedicalAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateConfluenceAction(val *string) {
	if err := j.validateSetCreateAndUpdateConfluenceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateConfluenceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateDashboardEmailReports(val *string) {
	if err := j.validateSetCreateAndUpdateDashboardEmailReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDashboardEmailReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateDatasets(val *string) {
	if err := j.validateSetCreateAndUpdateDatasetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDatasets",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateDataSources(val *string) {
	if err := j.validateSetCreateAndUpdateDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDataSources",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateFactSetAction(val *string) {
	if err := j.validateSetCreateAndUpdateFactSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateFactSetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateGenericHttpAction(val *string) {
	if err := j.validateSetCreateAndUpdateGenericHttpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateGenericHttpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateGithubAction(val *string) {
	if err := j.validateSetCreateAndUpdateGithubActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateGithubAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateGoogleCalendarAction(val *string) {
	if err := j.validateSetCreateAndUpdateGoogleCalendarActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateGoogleCalendarAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateHubspotAction(val *string) {
	if err := j.validateSetCreateAndUpdateHubspotActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateHubspotAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateHuggingFaceAction(val *string) {
	if err := j.validateSetCreateAndUpdateHuggingFaceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateHuggingFaceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateIntercomAction(val *string) {
	if err := j.validateSetCreateAndUpdateIntercomActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateIntercomAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateJiraAction(val *string) {
	if err := j.validateSetCreateAndUpdateJiraActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateJiraAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateKnowledgeBases(val *string) {
	if err := j.validateSetCreateAndUpdateKnowledgeBasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateKnowledgeBases",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateLinearAction(val *string) {
	if err := j.validateSetCreateAndUpdateLinearActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateLinearAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateMcpAction(val *string) {
	if err := j.validateSetCreateAndUpdateMcpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateMcpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateMondayAction(val *string) {
	if err := j.validateSetCreateAndUpdateMondayActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateMondayAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateMsExchangeAction(val *string) {
	if err := j.validateSetCreateAndUpdateMsExchangeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateMsExchangeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateMsTeamsAction(val *string) {
	if err := j.validateSetCreateAndUpdateMsTeamsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateMsTeamsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateNewRelicAction(val *string) {
	if err := j.validateSetCreateAndUpdateNewRelicActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateNewRelicAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateNotionAction(val *string) {
	if err := j.validateSetCreateAndUpdateNotionActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateNotionAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateOneDriveAction(val *string) {
	if err := j.validateSetCreateAndUpdateOneDriveActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateOneDriveAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateOpenApiAction(val *string) {
	if err := j.validateSetCreateAndUpdateOpenApiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateOpenApiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdatePagerDutyAction(val *string) {
	if err := j.validateSetCreateAndUpdatePagerDutyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdatePagerDutyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSalesforceAction(val *string) {
	if err := j.validateSetCreateAndUpdateSalesforceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSalesforceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSandPGlobalEnergyAction(val *string) {
	if err := j.validateSetCreateAndUpdateSandPGlobalEnergyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSandPGlobalEnergyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSandPgmiAction(val *string) {
	if err := j.validateSetCreateAndUpdateSandPgmiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSandPgmiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSapBillOfMaterialAction(val *string) {
	if err := j.validateSetCreateAndUpdateSapBillOfMaterialActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSapBillOfMaterialAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSapBusinessPartnerAction(val *string) {
	if err := j.validateSetCreateAndUpdateSapBusinessPartnerActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSapBusinessPartnerAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSapMaterialStockAction(val *string) {
	if err := j.validateSetCreateAndUpdateSapMaterialStockActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSapMaterialStockAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSapPhysicalInventoryAction(val *string) {
	if err := j.validateSetCreateAndUpdateSapPhysicalInventoryActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSapPhysicalInventoryAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSapProductMasterDataAction(val *string) {
	if err := j.validateSetCreateAndUpdateSapProductMasterDataActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSapProductMasterDataAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateServiceNowAction(val *string) {
	if err := j.validateSetCreateAndUpdateServiceNowActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateServiceNowAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSharePointAction(val *string) {
	if err := j.validateSetCreateAndUpdateSharePointActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSharePointAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSlackAction(val *string) {
	if err := j.validateSetCreateAndUpdateSlackActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSlackAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSmartsheetAction(val *string) {
	if err := j.validateSetCreateAndUpdateSmartsheetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSmartsheetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateTextractAction(val *string) {
	if err := j.validateSetCreateAndUpdateTextractActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateTextractAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateThemes(val *string) {
	if err := j.validateSetCreateAndUpdateThemesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateThemes",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateThresholdAlerts(val *string) {
	if err := j.validateSetCreateAndUpdateThresholdAlertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateThresholdAlerts",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateZendeskAction(val *string) {
	if err := j.validateSetCreateAndUpdateZendeskActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateZendeskAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateChatAgents(val *string) {
	if err := j.validateSetCreateChatAgentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createChatAgents",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateDashboardExecutiveSummaryWithQ(val *string) {
	if err := j.validateSetCreateDashboardExecutiveSummaryWithQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDashboardExecutiveSummaryWithQ",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSharedFolders(val *string) {
	if err := j.validateSetCreateSharedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSharedFolders",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSpaces(val *string) {
	if err := j.validateSetCreateSpacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSpaces",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSpiceDataset(val *string) {
	if err := j.validateSetCreateSpiceDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSpiceDataset",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetDashboard(val *string) {
	if err := j.validateSetDashboardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dashboard",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetEditVisualWithQ(val *string) {
	if err := j.validateSetEditVisualWithQParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"editVisualWithQ",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToCsv(val *string) {
	if err := j.validateSetExportToCsvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToCsv",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToCsvInScheduledReports(val *string) {
	if err := j.validateSetExportToCsvInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToCsvInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToExcel(val *string) {
	if err := j.validateSetExportToExcelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToExcel",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToExcelInScheduledReports(val *string) {
	if err := j.validateSetExportToExcelInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToExcelInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToPdf(val *string) {
	if err := j.validateSetExportToPdfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToPdf",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExportToPdfInScheduledReports(val *string) {
	if err := j.validateSetExportToPdfInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToPdfInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExtension(val *string) {
	if err := j.validateSetExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extension",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetFactSetAction(val *string) {
	if err := j.validateSetFactSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"factSetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetFlow(val *string) {
	if err := j.validateSetFlowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flow",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGenericHttpAction(val *string) {
	if err := j.validateSetGenericHttpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"genericHttpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGithubAction(val *string) {
	if err := j.validateSetGithubActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"githubAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGoogleCalendarAction(val *string) {
	if err := j.validateSetGoogleCalendarActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleCalendarAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetHubspotAction(val *string) {
	if err := j.validateSetHubspotActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hubspotAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetHuggingFaceAction(val *string) {
	if err := j.validateSetHuggingFaceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"huggingFaceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetIncludeContentInScheduledReportsEmail(val *string) {
	if err := j.validateSetIncludeContentInScheduledReportsEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeContentInScheduledReportsEmail",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetIntercomAction(val *string) {
	if err := j.validateSetIntercomActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intercomAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetInvokeAppsAiInference(val *string) {
	if err := j.validateSetInvokeAppsAiInferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invokeAppsAiInference",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetJiraAction(val *string) {
	if err := j.validateSetJiraActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetKnowledgeBase(val *string) {
	if err := j.validateSetKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"knowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetLinearAction(val *string) {
	if err := j.validateSetLinearActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linearAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetManageSharedFolders(val *string) {
	if err := j.validateSetManageSharedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageSharedFolders",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMcpAction(val *string) {
	if err := j.validateSetMcpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mcpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMondayAction(val *string) {
	if err := j.validateSetMondayActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mondayAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMsExchangeAction(val *string) {
	if err := j.validateSetMsExchangeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"msExchangeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMsTeamsAction(val *string) {
	if err := j.validateSetMsTeamsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"msTeamsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetNewRelicAction(val *string) {
	if err := j.validateSetNewRelicActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newRelicAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetNotionAction(val *string) {
	if err := j.validateSetNotionActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notionAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetOneDriveAction(val *string) {
	if err := j.validateSetOneDriveActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oneDriveAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetOpenApiAction(val *string) {
	if err := j.validateSetOpenApiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openApiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPagerDutyAction(val *string) {
	if err := j.validateSetPagerDutyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pagerDutyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPerformFlowUiTask(val *string) {
	if err := j.validateSetPerformFlowUiTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performFlowUiTask",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPrintReports(val *string) {
	if err := j.validateSetPrintReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPublishWithoutApproval(val *string) {
	if err := j.validateSetPublishWithoutApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishWithoutApproval",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetRenameSharedFolders(val *string) {
	if err := j.validateSetRenameSharedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renameSharedFolders",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetResearch(val *string) {
	if err := j.validateSetResearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"research",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSalesforceAction(val *string) {
	if err := j.validateSetSalesforceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salesforceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSandPGlobalEnergyAction(val *string) {
	if err := j.validateSetSandPGlobalEnergyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sandPGlobalEnergyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSandPgmiAction(val *string) {
	if err := j.validateSetSandPgmiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sandPgmiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSapBillOfMaterialAction(val *string) {
	if err := j.validateSetSapBillOfMaterialActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapBillOfMaterialAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSapBusinessPartnerAction(val *string) {
	if err := j.validateSetSapBusinessPartnerActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapBusinessPartnerAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSapMaterialStockAction(val *string) {
	if err := j.validateSetSapMaterialStockActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapMaterialStockAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSapPhysicalInventoryAction(val *string) {
	if err := j.validateSetSapPhysicalInventoryActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapPhysicalInventoryAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSapProductMasterDataAction(val *string) {
	if err := j.validateSetSapProductMasterDataActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapProductMasterDataAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetServiceNowAction(val *string) {
	if err := j.validateSetServiceNowActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNowAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAmazonBedrockArsAction(val *string) {
	if err := j.validateSetShareAmazonBedrockArsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAmazonBedrockArsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAmazonBedrockFsAction(val *string) {
	if err := j.validateSetShareAmazonBedrockFsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAmazonBedrockFsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAmazonBedrockKrsAction(val *string) {
	if err := j.validateSetShareAmazonBedrockKrsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAmazonBedrockKrsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAmazonSThreeAction(val *string) {
	if err := j.validateSetShareAmazonSThreeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAmazonSThreeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAnalyses(val *string) {
	if err := j.validateSetShareAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAnalyses",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareApps(val *string) {
	if err := j.validateSetShareAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareApps",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAsanaAction(val *string) {
	if err := j.validateSetShareAsanaActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAsanaAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareBambooHrAction(val *string) {
	if err := j.validateSetShareBambooHrActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareBambooHrAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareBoxAgentAction(val *string) {
	if err := j.validateSetShareBoxAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareBoxAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareCanvaAgentAction(val *string) {
	if err := j.validateSetShareCanvaAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareCanvaAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareChatAgents(val *string) {
	if err := j.validateSetShareChatAgentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareChatAgents",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareComprehendAction(val *string) {
	if err := j.validateSetShareComprehendActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareComprehendAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareComprehendMedicalAction(val *string) {
	if err := j.validateSetShareComprehendMedicalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareComprehendMedicalAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareConfluenceAction(val *string) {
	if err := j.validateSetShareConfluenceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareConfluenceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDashboards(val *string) {
	if err := j.validateSetShareDashboardsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDashboards",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDatasets(val *string) {
	if err := j.validateSetShareDatasetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDatasets",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDataSources(val *string) {
	if err := j.validateSetShareDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDataSources",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareFactSetAction(val *string) {
	if err := j.validateSetShareFactSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareFactSetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGenericHttpAction(val *string) {
	if err := j.validateSetShareGenericHttpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGenericHttpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGithubAction(val *string) {
	if err := j.validateSetShareGithubActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGithubAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGoogleCalendarAction(val *string) {
	if err := j.validateSetShareGoogleCalendarActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGoogleCalendarAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareHubspotAction(val *string) {
	if err := j.validateSetShareHubspotActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareHubspotAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareHuggingFaceAction(val *string) {
	if err := j.validateSetShareHuggingFaceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareHuggingFaceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareIntercomAction(val *string) {
	if err := j.validateSetShareIntercomActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareIntercomAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareJiraAction(val *string) {
	if err := j.validateSetShareJiraActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareJiraAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareKnowledgeBases(val *string) {
	if err := j.validateSetShareKnowledgeBasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareKnowledgeBases",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareLinearAction(val *string) {
	if err := j.validateSetShareLinearActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareLinearAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMcpAction(val *string) {
	if err := j.validateSetShareMcpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMcpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMondayAction(val *string) {
	if err := j.validateSetShareMondayActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMondayAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMsExchangeAction(val *string) {
	if err := j.validateSetShareMsExchangeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMsExchangeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMsTeamsAction(val *string) {
	if err := j.validateSetShareMsTeamsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMsTeamsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareNewRelicAction(val *string) {
	if err := j.validateSetShareNewRelicActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareNewRelicAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareNotionAction(val *string) {
	if err := j.validateSetShareNotionActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareNotionAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareOneDriveAction(val *string) {
	if err := j.validateSetShareOneDriveActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareOneDriveAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareOpenApiAction(val *string) {
	if err := j.validateSetShareOpenApiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareOpenApiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSharePagerDutyAction(val *string) {
	if err := j.validateSetSharePagerDutyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePagerDutyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSharePointAction(val *string) {
	if err := j.validateSetSharePointActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePointAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSalesforceAction(val *string) {
	if err := j.validateSetShareSalesforceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSalesforceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSandPGlobalEnergyAction(val *string) {
	if err := j.validateSetShareSandPGlobalEnergyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSandPGlobalEnergyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSandPgmiAction(val *string) {
	if err := j.validateSetShareSandPgmiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSandPgmiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSapBillOfMaterialAction(val *string) {
	if err := j.validateSetShareSapBillOfMaterialActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSapBillOfMaterialAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSapBusinessPartnerAction(val *string) {
	if err := j.validateSetShareSapBusinessPartnerActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSapBusinessPartnerAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSapMaterialStockAction(val *string) {
	if err := j.validateSetShareSapMaterialStockActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSapMaterialStockAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSapPhysicalInventoryAction(val *string) {
	if err := j.validateSetShareSapPhysicalInventoryActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSapPhysicalInventoryAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSapProductMasterDataAction(val *string) {
	if err := j.validateSetShareSapProductMasterDataActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSapProductMasterDataAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareServiceNowAction(val *string) {
	if err := j.validateSetShareServiceNowActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareServiceNowAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSharePointAction(val *string) {
	if err := j.validateSetShareSharePointActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSharePointAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSlackAction(val *string) {
	if err := j.validateSetShareSlackActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSlackAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSmartsheetAction(val *string) {
	if err := j.validateSetShareSmartsheetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSmartsheetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSpaces(val *string) {
	if err := j.validateSetShareSpacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSpaces",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareTextractAction(val *string) {
	if err := j.validateSetShareTextractActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareTextractAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareZendeskAction(val *string) {
	if err := j.validateSetShareZendeskActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareZendeskAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSlackAction(val *string) {
	if err := j.validateSetSlackActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slackAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSmartsheetAction(val *string) {
	if err := j.validateSetSmartsheetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smartsheetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSpace(val *string) {
	if err := j.validateSetSpaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"space",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSubscribeDashboardEmailReports(val *string) {
	if err := j.validateSetSubscribeDashboardEmailReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscribeDashboardEmailReports",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTextractAction(val *string) {
	if err := j.validateSetTextractActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textractAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTopic(val *string) {
	if err := j.validateSetTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topic",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseAgentWebSearch(val *string) {
	if err := j.validateSetUseAgentWebSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAgentWebSearch",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseAmazonBedrockArsAction(val *string) {
	if err := j.validateSetUseAmazonBedrockArsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAmazonBedrockArsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseAmazonBedrockFsAction(val *string) {
	if err := j.validateSetUseAmazonBedrockFsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAmazonBedrockFsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseAmazonBedrockKrsAction(val *string) {
	if err := j.validateSetUseAmazonBedrockKrsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAmazonBedrockKrsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseAmazonSThreeAction(val *string) {
	if err := j.validateSetUseAmazonSThreeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAmazonSThreeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseAsanaAction(val *string) {
	if err := j.validateSetUseAsanaActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAsanaAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseBambooHrAction(val *string) {
	if err := j.validateSetUseBambooHrActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBambooHrAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseBedrockModels(val *string) {
	if err := j.validateSetUseBedrockModelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBedrockModels",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseBoxAgentAction(val *string) {
	if err := j.validateSetUseBoxAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBoxAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseCanvaAgentAction(val *string) {
	if err := j.validateSetUseCanvaAgentActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCanvaAgentAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseComprehendAction(val *string) {
	if err := j.validateSetUseComprehendActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useComprehendAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseComprehendMedicalAction(val *string) {
	if err := j.validateSetUseComprehendMedicalActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useComprehendMedicalAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseConfluenceAction(val *string) {
	if err := j.validateSetUseConfluenceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useConfluenceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseFactSetAction(val *string) {
	if err := j.validateSetUseFactSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFactSetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseGenericHttpAction(val *string) {
	if err := j.validateSetUseGenericHttpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGenericHttpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseGithubAction(val *string) {
	if err := j.validateSetUseGithubActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGithubAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseGoogleCalendarAction(val *string) {
	if err := j.validateSetUseGoogleCalendarActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGoogleCalendarAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseHubspotAction(val *string) {
	if err := j.validateSetUseHubspotActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useHubspotAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseHuggingFaceAction(val *string) {
	if err := j.validateSetUseHuggingFaceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useHuggingFaceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseIntercomAction(val *string) {
	if err := j.validateSetUseIntercomActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useIntercomAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseJiraAction(val *string) {
	if err := j.validateSetUseJiraActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useJiraAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseLinearAction(val *string) {
	if err := j.validateSetUseLinearActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLinearAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseMcpAction(val *string) {
	if err := j.validateSetUseMcpActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMcpAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseMondayAction(val *string) {
	if err := j.validateSetUseMondayActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMondayAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseMsExchangeAction(val *string) {
	if err := j.validateSetUseMsExchangeActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMsExchangeAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseMsTeamsAction(val *string) {
	if err := j.validateSetUseMsTeamsActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMsTeamsAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseNewRelicAction(val *string) {
	if err := j.validateSetUseNewRelicActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNewRelicAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseNotionAction(val *string) {
	if err := j.validateSetUseNotionActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useNotionAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseOneDriveAction(val *string) {
	if err := j.validateSetUseOneDriveActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOneDriveAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseOpenApiAction(val *string) {
	if err := j.validateSetUseOpenApiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOpenApiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUsePagerDutyAction(val *string) {
	if err := j.validateSetUsePagerDutyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePagerDutyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSalesforceAction(val *string) {
	if err := j.validateSetUseSalesforceActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSalesforceAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSandPGlobalEnergyAction(val *string) {
	if err := j.validateSetUseSandPGlobalEnergyActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSandPGlobalEnergyAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSandPgmiAction(val *string) {
	if err := j.validateSetUseSandPgmiActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSandPgmiAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSapBillOfMaterialAction(val *string) {
	if err := j.validateSetUseSapBillOfMaterialActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSapBillOfMaterialAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSapBusinessPartnerAction(val *string) {
	if err := j.validateSetUseSapBusinessPartnerActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSapBusinessPartnerAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSapMaterialStockAction(val *string) {
	if err := j.validateSetUseSapMaterialStockActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSapMaterialStockAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSapPhysicalInventoryAction(val *string) {
	if err := j.validateSetUseSapPhysicalInventoryActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSapPhysicalInventoryAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSapProductMasterDataAction(val *string) {
	if err := j.validateSetUseSapProductMasterDataActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSapProductMasterDataAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseServiceNowAction(val *string) {
	if err := j.validateSetUseServiceNowActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useServiceNowAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSharePointAction(val *string) {
	if err := j.validateSetUseSharePointActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSharePointAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSlackAction(val *string) {
	if err := j.validateSetUseSlackActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSlackAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSmartsheetAction(val *string) {
	if err := j.validateSetUseSmartsheetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSmartsheetAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseTextractAction(val *string) {
	if err := j.validateSetUseTextractActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTextractAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseZendeskAction(val *string) {
	if err := j.validateSetUseZendeskActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useZendeskAction",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetViewAccountSpiceCapacity(val *string) {
	if err := j.validateSetViewAccountSpiceCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewAccountSpiceCapacity",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetZendeskAction(val *string) {
	if err := j.validateSetZendeskActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zendeskAction",
		val,
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAccessAppsNativeDataStore() {
	_jsii_.InvokeVoid(
		q,
		"resetAccessAppsNativeDataStore",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAddOrRunAnomalyDetectionForAnalyses() {
	_jsii_.InvokeVoid(
		q,
		"resetAddOrRunAnomalyDetectionForAnalyses",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAmazonBedrockArsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonBedrockArsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAmazonBedrockFsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonBedrockFsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAmazonBedrockKrsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonBedrockKrsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAmazonSThreeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetAmazonSThreeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAnalysis() {
	_jsii_.InvokeVoid(
		q,
		"resetAnalysis",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetApproveFlowShareRequests() {
	_jsii_.InvokeVoid(
		q,
		"resetApproveFlowShareRequests",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetApps() {
	_jsii_.InvokeVoid(
		q,
		"resetApps",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAsanaAction() {
	_jsii_.InvokeVoid(
		q,
		"resetAsanaAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAutomate() {
	_jsii_.InvokeVoid(
		q,
		"resetAutomate",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetBambooHrAction() {
	_jsii_.InvokeVoid(
		q,
		"resetBambooHrAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetBoxAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetBoxAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetBuildCalculatedFieldWithQ() {
	_jsii_.InvokeVoid(
		q,
		"resetBuildCalculatedFieldWithQ",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCanvaAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCanvaAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetChatAgent() {
	_jsii_.InvokeVoid(
		q,
		"resetChatAgent",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetComprehendAction() {
	_jsii_.InvokeVoid(
		q,
		"resetComprehendAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetComprehendMedicalAction() {
	_jsii_.InvokeVoid(
		q,
		"resetComprehendMedicalAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetConfluenceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetConfluenceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateAmazonBedrockArsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateAmazonBedrockArsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateAmazonBedrockFsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateAmazonBedrockFsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateAmazonBedrockKrsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateAmazonBedrockKrsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateAmazonSThreeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateAmazonSThreeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateApps() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateApps",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateAsanaAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateAsanaAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateBambooHrAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateBambooHrAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateBoxAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateBoxAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateCanvaAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateCanvaAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateComprehendAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateComprehendAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateComprehendMedicalAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateComprehendMedicalAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateConfluenceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateConfluenceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateDashboardEmailReports() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateDashboardEmailReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateDatasets() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateDatasets",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateDataSources() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateDataSources",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateFactSetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateFactSetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateGenericHttpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateGenericHttpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateGithubAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateGithubAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateGoogleCalendarAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateGoogleCalendarAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateHubspotAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateHubspotAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateHuggingFaceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateHuggingFaceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateIntercomAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateIntercomAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateJiraAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateJiraAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateKnowledgeBases() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateKnowledgeBases",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateLinearAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateLinearAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateMcpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateMcpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateMondayAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateMondayAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateMsExchangeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateMsExchangeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateMsTeamsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateMsTeamsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateNewRelicAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateNewRelicAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateNotionAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateNotionAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateOneDriveAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateOneDriveAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateOpenApiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateOpenApiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdatePagerDutyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdatePagerDutyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSalesforceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSalesforceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSandPGlobalEnergyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSandPGlobalEnergyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSandPgmiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSandPgmiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSapBillOfMaterialAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSapBillOfMaterialAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSapBusinessPartnerAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSapBusinessPartnerAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSapMaterialStockAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSapMaterialStockAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSapPhysicalInventoryAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSapPhysicalInventoryAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSapProductMasterDataAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSapProductMasterDataAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateServiceNowAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateServiceNowAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSharePointAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSharePointAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSlackAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSlackAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSmartsheetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSmartsheetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateTextractAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateTextractAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateThemes() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateThemes",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateThresholdAlerts() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateThresholdAlerts",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateZendeskAction() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateZendeskAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateChatAgents() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateChatAgents",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateDashboardExecutiveSummaryWithQ() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateDashboardExecutiveSummaryWithQ",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSharedFolders() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSharedFolders",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSpaces() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSpaces",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSpiceDataset() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSpiceDataset",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetDashboard() {
	_jsii_.InvokeVoid(
		q,
		"resetDashboard",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetEditVisualWithQ() {
	_jsii_.InvokeVoid(
		q,
		"resetEditVisualWithQ",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToCsv() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToCsv",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToCsvInScheduledReports() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToCsvInScheduledReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToExcel() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToExcel",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToExcelInScheduledReports() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToExcelInScheduledReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToPdf() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToPdf",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExportToPdfInScheduledReports() {
	_jsii_.InvokeVoid(
		q,
		"resetExportToPdfInScheduledReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExtension() {
	_jsii_.InvokeVoid(
		q,
		"resetExtension",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetFactSetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetFactSetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetFlow() {
	_jsii_.InvokeVoid(
		q,
		"resetFlow",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGenericHttpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetGenericHttpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGithubAction() {
	_jsii_.InvokeVoid(
		q,
		"resetGithubAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGoogleCalendarAction() {
	_jsii_.InvokeVoid(
		q,
		"resetGoogleCalendarAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetHubspotAction() {
	_jsii_.InvokeVoid(
		q,
		"resetHubspotAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetHuggingFaceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetHuggingFaceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetIncludeContentInScheduledReportsEmail() {
	_jsii_.InvokeVoid(
		q,
		"resetIncludeContentInScheduledReportsEmail",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetIntercomAction() {
	_jsii_.InvokeVoid(
		q,
		"resetIntercomAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetInvokeAppsAiInference() {
	_jsii_.InvokeVoid(
		q,
		"resetInvokeAppsAiInference",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetJiraAction() {
	_jsii_.InvokeVoid(
		q,
		"resetJiraAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetLinearAction() {
	_jsii_.InvokeVoid(
		q,
		"resetLinearAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetManageSharedFolders() {
	_jsii_.InvokeVoid(
		q,
		"resetManageSharedFolders",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMcpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetMcpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMondayAction() {
	_jsii_.InvokeVoid(
		q,
		"resetMondayAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMsExchangeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetMsExchangeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMsTeamsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetMsTeamsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetNewRelicAction() {
	_jsii_.InvokeVoid(
		q,
		"resetNewRelicAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetNotionAction() {
	_jsii_.InvokeVoid(
		q,
		"resetNotionAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetOneDriveAction() {
	_jsii_.InvokeVoid(
		q,
		"resetOneDriveAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetOpenApiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetOpenApiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPagerDutyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetPagerDutyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPerformFlowUiTask() {
	_jsii_.InvokeVoid(
		q,
		"resetPerformFlowUiTask",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPrintReports() {
	_jsii_.InvokeVoid(
		q,
		"resetPrintReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPublishWithoutApproval() {
	_jsii_.InvokeVoid(
		q,
		"resetPublishWithoutApproval",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetRenameSharedFolders() {
	_jsii_.InvokeVoid(
		q,
		"resetRenameSharedFolders",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetResearch() {
	_jsii_.InvokeVoid(
		q,
		"resetResearch",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSalesforceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSalesforceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSandPGlobalEnergyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSandPGlobalEnergyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSandPgmiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSandPgmiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSapBillOfMaterialAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSapBillOfMaterialAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSapBusinessPartnerAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSapBusinessPartnerAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSapMaterialStockAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSapMaterialStockAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSapPhysicalInventoryAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSapPhysicalInventoryAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSapProductMasterDataAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSapProductMasterDataAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetServiceNowAction() {
	_jsii_.InvokeVoid(
		q,
		"resetServiceNowAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAmazonBedrockArsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAmazonBedrockArsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAmazonBedrockFsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAmazonBedrockFsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAmazonBedrockKrsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAmazonBedrockKrsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAmazonSThreeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAmazonSThreeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAnalyses() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAnalyses",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareApps() {
	_jsii_.InvokeVoid(
		q,
		"resetShareApps",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAsanaAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAsanaAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareBambooHrAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareBambooHrAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareBoxAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareBoxAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareCanvaAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareCanvaAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareChatAgents() {
	_jsii_.InvokeVoid(
		q,
		"resetShareChatAgents",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareComprehendAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareComprehendAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareComprehendMedicalAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareComprehendMedicalAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareConfluenceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareConfluenceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDashboards() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDashboards",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDatasets() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDatasets",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDataSources() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDataSources",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareFactSetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareFactSetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGenericHttpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGenericHttpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGithubAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGithubAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGoogleCalendarAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGoogleCalendarAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareHubspotAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareHubspotAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareHuggingFaceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareHuggingFaceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareIntercomAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareIntercomAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareJiraAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareJiraAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareKnowledgeBases() {
	_jsii_.InvokeVoid(
		q,
		"resetShareKnowledgeBases",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareLinearAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareLinearAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMcpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMcpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMondayAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMondayAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMsExchangeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMsExchangeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMsTeamsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMsTeamsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareNewRelicAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareNewRelicAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareNotionAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareNotionAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareOneDriveAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareOneDriveAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareOpenApiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareOpenApiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSharePagerDutyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSharePagerDutyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSharePointAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSharePointAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSalesforceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSalesforceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSandPGlobalEnergyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSandPGlobalEnergyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSandPgmiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSandPgmiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSapBillOfMaterialAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSapBillOfMaterialAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSapBusinessPartnerAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSapBusinessPartnerAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSapMaterialStockAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSapMaterialStockAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSapPhysicalInventoryAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSapPhysicalInventoryAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSapProductMasterDataAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSapProductMasterDataAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareServiceNowAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareServiceNowAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSharePointAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSharePointAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSlackAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSlackAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSmartsheetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSmartsheetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSpaces() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSpaces",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareTextractAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareTextractAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareZendeskAction() {
	_jsii_.InvokeVoid(
		q,
		"resetShareZendeskAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSlackAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSlackAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSmartsheetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetSmartsheetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSpace() {
	_jsii_.InvokeVoid(
		q,
		"resetSpace",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSubscribeDashboardEmailReports() {
	_jsii_.InvokeVoid(
		q,
		"resetSubscribeDashboardEmailReports",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetTextractAction() {
	_jsii_.InvokeVoid(
		q,
		"resetTextractAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetTopic() {
	_jsii_.InvokeVoid(
		q,
		"resetTopic",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseAgentWebSearch() {
	_jsii_.InvokeVoid(
		q,
		"resetUseAgentWebSearch",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseAmazonBedrockArsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseAmazonBedrockArsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseAmazonBedrockFsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseAmazonBedrockFsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseAmazonBedrockKrsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseAmazonBedrockKrsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseAmazonSThreeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseAmazonSThreeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseAsanaAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseAsanaAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseBambooHrAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseBambooHrAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseBedrockModels() {
	_jsii_.InvokeVoid(
		q,
		"resetUseBedrockModels",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseBoxAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseBoxAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseCanvaAgentAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseCanvaAgentAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseComprehendAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseComprehendAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseComprehendMedicalAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseComprehendMedicalAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseConfluenceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseConfluenceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseFactSetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseFactSetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseGenericHttpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseGenericHttpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseGithubAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseGithubAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseGoogleCalendarAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseGoogleCalendarAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseHubspotAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseHubspotAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseHuggingFaceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseHuggingFaceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseIntercomAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseIntercomAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseJiraAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseJiraAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseLinearAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseLinearAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseMcpAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseMcpAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseMondayAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseMondayAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseMsExchangeAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseMsExchangeAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseMsTeamsAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseMsTeamsAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseNewRelicAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseNewRelicAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseNotionAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseNotionAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseOneDriveAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseOneDriveAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseOpenApiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseOpenApiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUsePagerDutyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUsePagerDutyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSalesforceAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSalesforceAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSandPGlobalEnergyAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSandPGlobalEnergyAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSandPgmiAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSandPgmiAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSapBillOfMaterialAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSapBillOfMaterialAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSapBusinessPartnerAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSapBusinessPartnerAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSapMaterialStockAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSapMaterialStockAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSapPhysicalInventoryAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSapPhysicalInventoryAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSapProductMasterDataAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSapProductMasterDataAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseServiceNowAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseServiceNowAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSharePointAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSharePointAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSlackAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSlackAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSmartsheetAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSmartsheetAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseTextractAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseTextractAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseZendeskAction() {
	_jsii_.InvokeVoid(
		q,
		"resetUseZendeskAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetViewAccountSpiceCapacity() {
	_jsii_.InvokeVoid(
		q,
		"resetViewAccountSpiceCapacity",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetZendeskAction() {
	_jsii_.InvokeVoid(
		q,
		"resetZendeskAction",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := q.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

