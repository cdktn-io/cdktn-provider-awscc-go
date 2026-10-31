// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightcustompermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/quicksightcustompermissions/internal"
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
	AdobeAnalyticsDataSource() *string
	SetAdobeAnalyticsDataSource(val *string)
	AdobeAnalyticsDataSourceInput() *string
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
	AthenaDataSource() *string
	SetAthenaDataSource(val *string)
	AthenaDataSourceInput() *string
	AuroraDataSource() *string
	SetAuroraDataSource(val *string)
	AuroraDataSourceInput() *string
	Automate() *string
	SetAutomate(val *string)
	AutomateInput() *string
	BambooHrAction() *string
	SetBambooHrAction(val *string)
	BambooHrActionInput() *string
	BedrockManagedKnowledgeBase() *string
	SetBedrockManagedKnowledgeBase(val *string)
	BedrockManagedKnowledgeBaseInput() *string
	BoxAgentAction() *string
	SetBoxAgentAction(val *string)
	BoxAgentActionInput() *string
	BoxKnowledgeBase() *string
	SetBoxKnowledgeBase(val *string)
	BoxKnowledgeBaseInput() *string
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
	ConfluenceKnowledgeBase() *string
	SetConfluenceKnowledgeBase(val *string)
	ConfluenceKnowledgeBaseInput() *string
	CreateAdobeAnalyticsDataSource() *string
	SetCreateAdobeAnalyticsDataSource(val *string)
	CreateAdobeAnalyticsDataSourceInput() *string
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
	CreateAndUpdateBedrockManagedKnowledgeBase() *string
	SetCreateAndUpdateBedrockManagedKnowledgeBase(val *string)
	CreateAndUpdateBedrockManagedKnowledgeBaseInput() *string
	CreateAndUpdateBoxAgentAction() *string
	SetCreateAndUpdateBoxAgentAction(val *string)
	CreateAndUpdateBoxAgentActionInput() *string
	CreateAndUpdateBoxKnowledgeBase() *string
	SetCreateAndUpdateBoxKnowledgeBase(val *string)
	CreateAndUpdateBoxKnowledgeBaseInput() *string
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
	CreateAndUpdateConfluenceKnowledgeBase() *string
	SetCreateAndUpdateConfluenceKnowledgeBase(val *string)
	CreateAndUpdateConfluenceKnowledgeBaseInput() *string
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
	CreateAndUpdateGoogleDriveKnowledgeBase() *string
	SetCreateAndUpdateGoogleDriveKnowledgeBase(val *string)
	CreateAndUpdateGoogleDriveKnowledgeBaseInput() *string
	CreateAndUpdateHubspotAction() *string
	SetCreateAndUpdateHubspotAction(val *string)
	CreateAndUpdateHubspotActionInput() *string
	CreateAndUpdateHuggingFaceAction() *string
	SetCreateAndUpdateHuggingFaceAction(val *string)
	CreateAndUpdateHuggingFaceActionInput() *string
	CreateAndUpdateIdcKnowledgeBase() *string
	SetCreateAndUpdateIdcKnowledgeBase(val *string)
	CreateAndUpdateIdcKnowledgeBaseInput() *string
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
	CreateAndUpdateOneDriveKnowledgeBase() *string
	SetCreateAndUpdateOneDriveKnowledgeBase(val *string)
	CreateAndUpdateOneDriveKnowledgeBaseInput() *string
	CreateAndUpdateOpenApiAction() *string
	SetCreateAndUpdateOpenApiAction(val *string)
	CreateAndUpdateOpenApiActionInput() *string
	CreateAndUpdatePagerDutyAction() *string
	SetCreateAndUpdatePagerDutyAction(val *string)
	CreateAndUpdatePagerDutyActionInput() *string
	CreateAndUpdateQBusinessKnowledgeBase() *string
	SetCreateAndUpdateQBusinessKnowledgeBase(val *string)
	CreateAndUpdateQBusinessKnowledgeBaseInput() *string
	CreateAndUpdateS3KnowledgeBase() *string
	SetCreateAndUpdateS3KnowledgeBase(val *string)
	CreateAndUpdateS3KnowledgeBaseInput() *string
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
	CreateAndUpdateSharePointKnowledgeBase() *string
	SetCreateAndUpdateSharePointKnowledgeBase(val *string)
	CreateAndUpdateSharePointKnowledgeBaseInput() *string
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
	CreateAndUpdateWebCrawlerKnowledgeBase() *string
	SetCreateAndUpdateWebCrawlerKnowledgeBase(val *string)
	CreateAndUpdateWebCrawlerKnowledgeBaseInput() *string
	CreateAndUpdateZendeskAction() *string
	SetCreateAndUpdateZendeskAction(val *string)
	CreateAndUpdateZendeskActionInput() *string
	CreateAthenaDataSource() *string
	SetCreateAthenaDataSource(val *string)
	CreateAthenaDataSourceInput() *string
	CreateAuroraDataSource() *string
	SetCreateAuroraDataSource(val *string)
	CreateAuroraDataSourceInput() *string
	CreateChatAgents() *string
	SetCreateChatAgents(val *string)
	CreateChatAgentsInput() *string
	CreateDashboardExecutiveSummaryWithQ() *string
	SetCreateDashboardExecutiveSummaryWithQ(val *string)
	CreateDashboardExecutiveSummaryWithQInput() *string
	CreateDatabricksDataSource() *string
	SetCreateDatabricksDataSource(val *string)
	CreateDatabricksDataSourceInput() *string
	CreateDb2DataSource() *string
	SetCreateDb2DataSource(val *string)
	CreateDb2DataSourceInput() *string
	CreateDenodoDataSource() *string
	SetCreateDenodoDataSource(val *string)
	CreateDenodoDataSourceInput() *string
	CreateDocumentDbDataSource() *string
	SetCreateDocumentDbDataSource(val *string)
	CreateDocumentDbDataSourceInput() *string
	CreateDremioDataSource() *string
	SetCreateDremioDataSource(val *string)
	CreateDremioDataSourceInput() *string
	CreateDynamoDbDataSource() *string
	SetCreateDynamoDbDataSource(val *string)
	CreateDynamoDbDataSourceInput() *string
	CreateExasolDataSource() *string
	SetCreateExasolDataSource(val *string)
	CreateExasolDataSourceInput() *string
	CreateFileDataSource() *string
	SetCreateFileDataSource(val *string)
	CreateFileDataSourceInput() *string
	CreateGitHubDataSource() *string
	SetCreateGitHubDataSource(val *string)
	CreateGitHubDataSourceInput() *string
	CreateGoogleAnalyticsDataSource() *string
	SetCreateGoogleAnalyticsDataSource(val *string)
	CreateGoogleAnalyticsDataSourceInput() *string
	CreateGoogleBigQueryDataSource() *string
	SetCreateGoogleBigQueryDataSource(val *string)
	CreateGoogleBigQueryDataSourceInput() *string
	CreateGoogleSheetsDataSource() *string
	SetCreateGoogleSheetsDataSource(val *string)
	CreateGoogleSheetsDataSourceInput() *string
	CreateImpalaDataSource() *string
	SetCreateImpalaDataSource(val *string)
	CreateImpalaDataSourceInput() *string
	CreateJiraDataSource() *string
	SetCreateJiraDataSource(val *string)
	CreateJiraDataSourceInput() *string
	CreateMariaDbDataSource() *string
	SetCreateMariaDbDataSource(val *string)
	CreateMariaDbDataSourceInput() *string
	CreateMongoAtlasDataSource() *string
	SetCreateMongoAtlasDataSource(val *string)
	CreateMongoAtlasDataSourceInput() *string
	CreateMongoDbDataSource() *string
	SetCreateMongoDbDataSource(val *string)
	CreateMongoDbDataSourceInput() *string
	CreateMySqlDataSource() *string
	SetCreateMySqlDataSource(val *string)
	CreateMySqlDataSourceInput() *string
	CreateOpenSearchDataSource() *string
	SetCreateOpenSearchDataSource(val *string)
	CreateOpenSearchDataSourceInput() *string
	CreateOracleDataSource() *string
	SetCreateOracleDataSource(val *string)
	CreateOracleDataSourceInput() *string
	CreatePayPalDataSource() *string
	SetCreatePayPalDataSource(val *string)
	CreatePayPalDataSourceInput() *string
	CreatePostgreSqlDataSource() *string
	SetCreatePostgreSqlDataSource(val *string)
	CreatePostgreSqlDataSourceInput() *string
	CreatePrestoDataSource() *string
	SetCreatePrestoDataSource(val *string)
	CreatePrestoDataSourceInput() *string
	CreateRadiantDataSource() *string
	SetCreateRadiantDataSource(val *string)
	CreateRadiantDataSourceInput() *string
	CreateRdsDataSource() *string
	SetCreateRdsDataSource(val *string)
	CreateRdsDataSourceInput() *string
	CreateRedshiftAutoDiscoveredDataSource() *string
	SetCreateRedshiftAutoDiscoveredDataSource(val *string)
	CreateRedshiftAutoDiscoveredDataSourceInput() *string
	CreateRedshiftManualDataSource() *string
	SetCreateRedshiftManualDataSource(val *string)
	CreateRedshiftManualDataSourceInput() *string
	CreateS3AnalyticsDataSource() *string
	SetCreateS3AnalyticsDataSource(val *string)
	CreateS3AnalyticsDataSourceInput() *string
	CreateS3DataSource() *string
	SetCreateS3DataSource(val *string)
	CreateS3DataSourceInput() *string
	CreateS3TablesDataSource() *string
	SetCreateS3TablesDataSource(val *string)
	CreateS3TablesDataSourceInput() *string
	CreateSalesforceDataSource() *string
	SetCreateSalesforceDataSource(val *string)
	CreateSalesforceDataSourceInput() *string
	CreateSapHanaDataSource() *string
	SetCreateSapHanaDataSource(val *string)
	CreateSapHanaDataSourceInput() *string
	CreateServiceNowDataSource() *string
	SetCreateServiceNowDataSource(val *string)
	CreateServiceNowDataSourceInput() *string
	CreateSharedFolders() *string
	SetCreateSharedFolders(val *string)
	CreateSharedFoldersInput() *string
	CreateSnowflakeDataSource() *string
	SetCreateSnowflakeDataSource(val *string)
	CreateSnowflakeDataSourceInput() *string
	CreateSpaces() *string
	SetCreateSpaces(val *string)
	CreateSpacesInput() *string
	CreateSparkDataSource() *string
	SetCreateSparkDataSource(val *string)
	CreateSparkDataSourceInput() *string
	CreateSpiceDataset() *string
	SetCreateSpiceDataset(val *string)
	CreateSpiceDatasetInput() *string
	CreateSqlServerDataSource() *string
	SetCreateSqlServerDataSource(val *string)
	CreateSqlServerDataSourceInput() *string
	CreateSquareDataSource() *string
	SetCreateSquareDataSource(val *string)
	CreateSquareDataSourceInput() *string
	CreateStarburstDataSource() *string
	SetCreateStarburstDataSource(val *string)
	CreateStarburstDataSourceInput() *string
	CreateTeradataDataSource() *string
	SetCreateTeradataDataSource(val *string)
	CreateTeradataDataSourceInput() *string
	CreateTimestreamDataSource() *string
	SetCreateTimestreamDataSource(val *string)
	CreateTimestreamDataSourceInput() *string
	CreateTrinoDataSource() *string
	SetCreateTrinoDataSource(val *string)
	CreateTrinoDataSourceInput() *string
	CreateTwitterDataSource() *string
	SetCreateTwitterDataSource(val *string)
	CreateTwitterDataSourceInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dashboard() *string
	SetDashboard(val *string)
	DashboardInput() *string
	DatabricksDataSource() *string
	SetDatabricksDataSource(val *string)
	DatabricksDataSourceInput() *string
	Db2DataSource() *string
	SetDb2DataSource(val *string)
	Db2DataSourceInput() *string
	DenodoDataSource() *string
	SetDenodoDataSource(val *string)
	DenodoDataSourceInput() *string
	DocumentDbDataSource() *string
	SetDocumentDbDataSource(val *string)
	DocumentDbDataSourceInput() *string
	DremioDataSource() *string
	SetDremioDataSource(val *string)
	DremioDataSourceInput() *string
	DynamoDbDataSource() *string
	SetDynamoDbDataSource(val *string)
	DynamoDbDataSourceInput() *string
	EditVisualWithQ() *string
	SetEditVisualWithQ(val *string)
	EditVisualWithQInput() *string
	ExasolDataSource() *string
	SetExasolDataSource(val *string)
	ExasolDataSourceInput() *string
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
	FileDataSource() *string
	SetFileDataSource(val *string)
	FileDataSourceInput() *string
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
	GitHubDataSource() *string
	SetGitHubDataSource(val *string)
	GitHubDataSourceInput() *string
	GoogleAnalyticsDataSource() *string
	SetGoogleAnalyticsDataSource(val *string)
	GoogleAnalyticsDataSourceInput() *string
	GoogleBigQueryDataSource() *string
	SetGoogleBigQueryDataSource(val *string)
	GoogleBigQueryDataSourceInput() *string
	GoogleCalendarAction() *string
	SetGoogleCalendarAction(val *string)
	GoogleCalendarActionInput() *string
	GoogleDriveKnowledgeBase() *string
	SetGoogleDriveKnowledgeBase(val *string)
	GoogleDriveKnowledgeBaseInput() *string
	GoogleSheetsDataSource() *string
	SetGoogleSheetsDataSource(val *string)
	GoogleSheetsDataSourceInput() *string
	HubspotAction() *string
	SetHubspotAction(val *string)
	HubspotActionInput() *string
	HuggingFaceAction() *string
	SetHuggingFaceAction(val *string)
	HuggingFaceActionInput() *string
	IdcKnowledgeBase() *string
	SetIdcKnowledgeBase(val *string)
	IdcKnowledgeBaseInput() *string
	ImpalaDataSource() *string
	SetImpalaDataSource(val *string)
	ImpalaDataSourceInput() *string
	InboundEmailTrigger() *string
	SetInboundEmailTrigger(val *string)
	InboundEmailTriggerInput() *string
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
	JiraDataSource() *string
	SetJiraDataSource(val *string)
	JiraDataSourceInput() *string
	KnowledgeBase() *string
	SetKnowledgeBase(val *string)
	KnowledgeBaseInput() *string
	LinearAction() *string
	SetLinearAction(val *string)
	LinearActionInput() *string
	ManageSharedFolders() *string
	SetManageSharedFolders(val *string)
	ManageSharedFoldersInput() *string
	MariaDbDataSource() *string
	SetMariaDbDataSource(val *string)
	MariaDbDataSourceInput() *string
	McpAction() *string
	SetMcpAction(val *string)
	McpActionInput() *string
	MondayAction() *string
	SetMondayAction(val *string)
	MondayActionInput() *string
	MongoAtlasDataSource() *string
	SetMongoAtlasDataSource(val *string)
	MongoAtlasDataSourceInput() *string
	MongoDbDataSource() *string
	SetMongoDbDataSource(val *string)
	MongoDbDataSourceInput() *string
	MsExchangeAction() *string
	SetMsExchangeAction(val *string)
	MsExchangeActionInput() *string
	MsTeamsAction() *string
	SetMsTeamsAction(val *string)
	MsTeamsActionInput() *string
	MySqlDataSource() *string
	SetMySqlDataSource(val *string)
	MySqlDataSourceInput() *string
	NewRelicAction() *string
	SetNewRelicAction(val *string)
	NewRelicActionInput() *string
	NotionAction() *string
	SetNotionAction(val *string)
	NotionActionInput() *string
	OneDriveAction() *string
	SetOneDriveAction(val *string)
	OneDriveActionInput() *string
	OneDriveKnowledgeBase() *string
	SetOneDriveKnowledgeBase(val *string)
	OneDriveKnowledgeBaseInput() *string
	OpenApiAction() *string
	SetOpenApiAction(val *string)
	OpenApiActionInput() *string
	OpenSearchDataSource() *string
	SetOpenSearchDataSource(val *string)
	OpenSearchDataSourceInput() *string
	OracleDataSource() *string
	SetOracleDataSource(val *string)
	OracleDataSourceInput() *string
	PagerDutyAction() *string
	SetPagerDutyAction(val *string)
	PagerDutyActionInput() *string
	PayPalDataSource() *string
	SetPayPalDataSource(val *string)
	PayPalDataSourceInput() *string
	PerformFlowUiTask() *string
	SetPerformFlowUiTask(val *string)
	PerformFlowUiTaskInput() *string
	PostgreSqlDataSource() *string
	SetPostgreSqlDataSource(val *string)
	PostgreSqlDataSourceInput() *string
	PrestoDataSource() *string
	SetPrestoDataSource(val *string)
	PrestoDataSourceInput() *string
	PrintReports() *string
	SetPrintReports(val *string)
	PrintReportsInput() *string
	PublishWithoutApproval() *string
	SetPublishWithoutApproval(val *string)
	PublishWithoutApprovalInput() *string
	QBusinessKnowledgeBase() *string
	SetQBusinessKnowledgeBase(val *string)
	QBusinessKnowledgeBaseInput() *string
	QuickEventTrigger() *string
	SetQuickEventTrigger(val *string)
	QuickEventTriggerInput() *string
	RadiantDataSource() *string
	SetRadiantDataSource(val *string)
	RadiantDataSourceInput() *string
	RdsDataSource() *string
	SetRdsDataSource(val *string)
	RdsDataSourceInput() *string
	RedshiftAutoDiscoveredDataSource() *string
	SetRedshiftAutoDiscoveredDataSource(val *string)
	RedshiftAutoDiscoveredDataSourceInput() *string
	RedshiftManualDataSource() *string
	SetRedshiftManualDataSource(val *string)
	RedshiftManualDataSourceInput() *string
	RenameSharedFolders() *string
	SetRenameSharedFolders(val *string)
	RenameSharedFoldersInput() *string
	Research() *string
	SetResearch(val *string)
	ResearchInput() *string
	S3AnalyticsDataSource() *string
	SetS3AnalyticsDataSource(val *string)
	S3AnalyticsDataSourceInput() *string
	S3DataSource() *string
	SetS3DataSource(val *string)
	S3DataSourceInput() *string
	S3KnowledgeBase() *string
	SetS3KnowledgeBase(val *string)
	S3KnowledgeBaseInput() *string
	S3TablesDataSource() *string
	SetS3TablesDataSource(val *string)
	S3TablesDataSourceInput() *string
	SalesforceAction() *string
	SetSalesforceAction(val *string)
	SalesforceActionInput() *string
	SalesforceDataSource() *string
	SetSalesforceDataSource(val *string)
	SalesforceDataSourceInput() *string
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
	SapHanaDataSource() *string
	SetSapHanaDataSource(val *string)
	SapHanaDataSourceInput() *string
	SapMaterialStockAction() *string
	SetSapMaterialStockAction(val *string)
	SapMaterialStockActionInput() *string
	SapPhysicalInventoryAction() *string
	SetSapPhysicalInventoryAction(val *string)
	SapPhysicalInventoryActionInput() *string
	SapProductMasterDataAction() *string
	SetSapProductMasterDataAction(val *string)
	SapProductMasterDataActionInput() *string
	ScheduleTrigger() *string
	SetScheduleTrigger(val *string)
	ScheduleTriggerInput() *string
	ServiceNowAction() *string
	SetServiceNowAction(val *string)
	ServiceNowActionInput() *string
	ServiceNowDataSource() *string
	SetServiceNowDataSource(val *string)
	ServiceNowDataSourceInput() *string
	ShareAdobeAnalyticsDataSource() *string
	SetShareAdobeAnalyticsDataSource(val *string)
	ShareAdobeAnalyticsDataSourceInput() *string
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
	ShareAthenaDataSource() *string
	SetShareAthenaDataSource(val *string)
	ShareAthenaDataSourceInput() *string
	ShareAuroraDataSource() *string
	SetShareAuroraDataSource(val *string)
	ShareAuroraDataSourceInput() *string
	ShareBambooHrAction() *string
	SetShareBambooHrAction(val *string)
	ShareBambooHrActionInput() *string
	ShareBedrockManagedKnowledgeBase() *string
	SetShareBedrockManagedKnowledgeBase(val *string)
	ShareBedrockManagedKnowledgeBaseInput() *string
	ShareBoxAgentAction() *string
	SetShareBoxAgentAction(val *string)
	ShareBoxAgentActionInput() *string
	ShareBoxKnowledgeBase() *string
	SetShareBoxKnowledgeBase(val *string)
	ShareBoxKnowledgeBaseInput() *string
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
	ShareConfluenceKnowledgeBase() *string
	SetShareConfluenceKnowledgeBase(val *string)
	ShareConfluenceKnowledgeBaseInput() *string
	ShareDashboards() *string
	SetShareDashboards(val *string)
	ShareDashboardsInput() *string
	ShareDatabricksDataSource() *string
	SetShareDatabricksDataSource(val *string)
	ShareDatabricksDataSourceInput() *string
	ShareDatasets() *string
	SetShareDatasets(val *string)
	ShareDatasetsInput() *string
	ShareDataSources() *string
	SetShareDataSources(val *string)
	ShareDataSourcesInput() *string
	ShareDb2DataSource() *string
	SetShareDb2DataSource(val *string)
	ShareDb2DataSourceInput() *string
	ShareDenodoDataSource() *string
	SetShareDenodoDataSource(val *string)
	ShareDenodoDataSourceInput() *string
	ShareDocumentDbDataSource() *string
	SetShareDocumentDbDataSource(val *string)
	ShareDocumentDbDataSourceInput() *string
	ShareDremioDataSource() *string
	SetShareDremioDataSource(val *string)
	ShareDremioDataSourceInput() *string
	ShareDynamoDbDataSource() *string
	SetShareDynamoDbDataSource(val *string)
	ShareDynamoDbDataSourceInput() *string
	ShareExasolDataSource() *string
	SetShareExasolDataSource(val *string)
	ShareExasolDataSourceInput() *string
	ShareFactSetAction() *string
	SetShareFactSetAction(val *string)
	ShareFactSetActionInput() *string
	ShareFileDataSource() *string
	SetShareFileDataSource(val *string)
	ShareFileDataSourceInput() *string
	ShareGenericHttpAction() *string
	SetShareGenericHttpAction(val *string)
	ShareGenericHttpActionInput() *string
	ShareGithubAction() *string
	SetShareGithubAction(val *string)
	ShareGithubActionInput() *string
	ShareGitHubDataSource() *string
	SetShareGitHubDataSource(val *string)
	ShareGitHubDataSourceInput() *string
	ShareGoogleAnalyticsDataSource() *string
	SetShareGoogleAnalyticsDataSource(val *string)
	ShareGoogleAnalyticsDataSourceInput() *string
	ShareGoogleBigQueryDataSource() *string
	SetShareGoogleBigQueryDataSource(val *string)
	ShareGoogleBigQueryDataSourceInput() *string
	ShareGoogleCalendarAction() *string
	SetShareGoogleCalendarAction(val *string)
	ShareGoogleCalendarActionInput() *string
	ShareGoogleDriveKnowledgeBase() *string
	SetShareGoogleDriveKnowledgeBase(val *string)
	ShareGoogleDriveKnowledgeBaseInput() *string
	ShareGoogleSheetsDataSource() *string
	SetShareGoogleSheetsDataSource(val *string)
	ShareGoogleSheetsDataSourceInput() *string
	ShareHubspotAction() *string
	SetShareHubspotAction(val *string)
	ShareHubspotActionInput() *string
	ShareHuggingFaceAction() *string
	SetShareHuggingFaceAction(val *string)
	ShareHuggingFaceActionInput() *string
	ShareIdcKnowledgeBase() *string
	SetShareIdcKnowledgeBase(val *string)
	ShareIdcKnowledgeBaseInput() *string
	ShareImpalaDataSource() *string
	SetShareImpalaDataSource(val *string)
	ShareImpalaDataSourceInput() *string
	ShareIntercomAction() *string
	SetShareIntercomAction(val *string)
	ShareIntercomActionInput() *string
	ShareJiraAction() *string
	SetShareJiraAction(val *string)
	ShareJiraActionInput() *string
	ShareJiraDataSource() *string
	SetShareJiraDataSource(val *string)
	ShareJiraDataSourceInput() *string
	ShareKnowledgeBases() *string
	SetShareKnowledgeBases(val *string)
	ShareKnowledgeBasesInput() *string
	ShareLinearAction() *string
	SetShareLinearAction(val *string)
	ShareLinearActionInput() *string
	ShareMariaDbDataSource() *string
	SetShareMariaDbDataSource(val *string)
	ShareMariaDbDataSourceInput() *string
	ShareMcpAction() *string
	SetShareMcpAction(val *string)
	ShareMcpActionInput() *string
	ShareMondayAction() *string
	SetShareMondayAction(val *string)
	ShareMondayActionInput() *string
	ShareMongoAtlasDataSource() *string
	SetShareMongoAtlasDataSource(val *string)
	ShareMongoAtlasDataSourceInput() *string
	ShareMongoDbDataSource() *string
	SetShareMongoDbDataSource(val *string)
	ShareMongoDbDataSourceInput() *string
	ShareMsExchangeAction() *string
	SetShareMsExchangeAction(val *string)
	ShareMsExchangeActionInput() *string
	ShareMsTeamsAction() *string
	SetShareMsTeamsAction(val *string)
	ShareMsTeamsActionInput() *string
	ShareMySqlDataSource() *string
	SetShareMySqlDataSource(val *string)
	ShareMySqlDataSourceInput() *string
	ShareNewRelicAction() *string
	SetShareNewRelicAction(val *string)
	ShareNewRelicActionInput() *string
	ShareNotionAction() *string
	SetShareNotionAction(val *string)
	ShareNotionActionInput() *string
	ShareOneDriveAction() *string
	SetShareOneDriveAction(val *string)
	ShareOneDriveActionInput() *string
	ShareOneDriveKnowledgeBase() *string
	SetShareOneDriveKnowledgeBase(val *string)
	ShareOneDriveKnowledgeBaseInput() *string
	ShareOpenApiAction() *string
	SetShareOpenApiAction(val *string)
	ShareOpenApiActionInput() *string
	ShareOpenSearchDataSource() *string
	SetShareOpenSearchDataSource(val *string)
	ShareOpenSearchDataSourceInput() *string
	ShareOracleDataSource() *string
	SetShareOracleDataSource(val *string)
	ShareOracleDataSourceInput() *string
	SharePagerDutyAction() *string
	SetSharePagerDutyAction(val *string)
	SharePagerDutyActionInput() *string
	SharePayPalDataSource() *string
	SetSharePayPalDataSource(val *string)
	SharePayPalDataSourceInput() *string
	SharePointAction() *string
	SetSharePointAction(val *string)
	SharePointActionInput() *string
	SharePointKnowledgeBase() *string
	SetSharePointKnowledgeBase(val *string)
	SharePointKnowledgeBaseInput() *string
	SharePostgreSqlDataSource() *string
	SetSharePostgreSqlDataSource(val *string)
	SharePostgreSqlDataSourceInput() *string
	SharePrestoDataSource() *string
	SetSharePrestoDataSource(val *string)
	SharePrestoDataSourceInput() *string
	ShareQBusinessKnowledgeBase() *string
	SetShareQBusinessKnowledgeBase(val *string)
	ShareQBusinessKnowledgeBaseInput() *string
	ShareRadiantDataSource() *string
	SetShareRadiantDataSource(val *string)
	ShareRadiantDataSourceInput() *string
	ShareRdsDataSource() *string
	SetShareRdsDataSource(val *string)
	ShareRdsDataSourceInput() *string
	ShareRedshiftAutoDiscoveredDataSource() *string
	SetShareRedshiftAutoDiscoveredDataSource(val *string)
	ShareRedshiftAutoDiscoveredDataSourceInput() *string
	ShareRedshiftManualDataSource() *string
	SetShareRedshiftManualDataSource(val *string)
	ShareRedshiftManualDataSourceInput() *string
	ShareS3AnalyticsDataSource() *string
	SetShareS3AnalyticsDataSource(val *string)
	ShareS3AnalyticsDataSourceInput() *string
	ShareS3DataSource() *string
	SetShareS3DataSource(val *string)
	ShareS3DataSourceInput() *string
	ShareS3KnowledgeBase() *string
	SetShareS3KnowledgeBase(val *string)
	ShareS3KnowledgeBaseInput() *string
	ShareS3TablesDataSource() *string
	SetShareS3TablesDataSource(val *string)
	ShareS3TablesDataSourceInput() *string
	ShareSalesforceAction() *string
	SetShareSalesforceAction(val *string)
	ShareSalesforceActionInput() *string
	ShareSalesforceDataSource() *string
	SetShareSalesforceDataSource(val *string)
	ShareSalesforceDataSourceInput() *string
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
	ShareSapHanaDataSource() *string
	SetShareSapHanaDataSource(val *string)
	ShareSapHanaDataSourceInput() *string
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
	ShareServiceNowDataSource() *string
	SetShareServiceNowDataSource(val *string)
	ShareServiceNowDataSourceInput() *string
	ShareSharePointAction() *string
	SetShareSharePointAction(val *string)
	ShareSharePointActionInput() *string
	ShareSharePointKnowledgeBase() *string
	SetShareSharePointKnowledgeBase(val *string)
	ShareSharePointKnowledgeBaseInput() *string
	ShareSlackAction() *string
	SetShareSlackAction(val *string)
	ShareSlackActionInput() *string
	ShareSmartsheetAction() *string
	SetShareSmartsheetAction(val *string)
	ShareSmartsheetActionInput() *string
	ShareSnowflakeDataSource() *string
	SetShareSnowflakeDataSource(val *string)
	ShareSnowflakeDataSourceInput() *string
	ShareSpaces() *string
	SetShareSpaces(val *string)
	ShareSpacesInput() *string
	ShareSparkDataSource() *string
	SetShareSparkDataSource(val *string)
	ShareSparkDataSourceInput() *string
	ShareSqlServerDataSource() *string
	SetShareSqlServerDataSource(val *string)
	ShareSqlServerDataSourceInput() *string
	ShareSquareDataSource() *string
	SetShareSquareDataSource(val *string)
	ShareSquareDataSourceInput() *string
	ShareStarburstDataSource() *string
	SetShareStarburstDataSource(val *string)
	ShareStarburstDataSourceInput() *string
	ShareTeradataDataSource() *string
	SetShareTeradataDataSource(val *string)
	ShareTeradataDataSourceInput() *string
	ShareTextractAction() *string
	SetShareTextractAction(val *string)
	ShareTextractActionInput() *string
	ShareTimestreamDataSource() *string
	SetShareTimestreamDataSource(val *string)
	ShareTimestreamDataSourceInput() *string
	ShareTrinoDataSource() *string
	SetShareTrinoDataSource(val *string)
	ShareTrinoDataSourceInput() *string
	ShareTwitterDataSource() *string
	SetShareTwitterDataSource(val *string)
	ShareTwitterDataSourceInput() *string
	ShareWebCrawlerKnowledgeBase() *string
	SetShareWebCrawlerKnowledgeBase(val *string)
	ShareWebCrawlerKnowledgeBaseInput() *string
	ShareZendeskAction() *string
	SetShareZendeskAction(val *string)
	ShareZendeskActionInput() *string
	SlackAction() *string
	SetSlackAction(val *string)
	SlackActionInput() *string
	SmartsheetAction() *string
	SetSmartsheetAction(val *string)
	SmartsheetActionInput() *string
	SnowflakeDataSource() *string
	SetSnowflakeDataSource(val *string)
	SnowflakeDataSourceInput() *string
	Space() *string
	SetSpace(val *string)
	SpaceInput() *string
	SparkDataSource() *string
	SetSparkDataSource(val *string)
	SparkDataSourceInput() *string
	SqlServerDataSource() *string
	SetSqlServerDataSource(val *string)
	SqlServerDataSourceInput() *string
	SquareDataSource() *string
	SetSquareDataSource(val *string)
	SquareDataSourceInput() *string
	StarburstDataSource() *string
	SetStarburstDataSource(val *string)
	StarburstDataSourceInput() *string
	SubscribeDashboardEmailReports() *string
	SetSubscribeDashboardEmailReports(val *string)
	SubscribeDashboardEmailReportsInput() *string
	TeradataDataSource() *string
	SetTeradataDataSource(val *string)
	TeradataDataSourceInput() *string
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
	TimestreamDataSource() *string
	SetTimestreamDataSource(val *string)
	TimestreamDataSourceInput() *string
	Topic() *string
	SetTopic(val *string)
	TopicInput() *string
	Trigger() *string
	SetTrigger(val *string)
	TriggerInput() *string
	TrinoDataSource() *string
	SetTrinoDataSource(val *string)
	TrinoDataSourceInput() *string
	TwitterDataSource() *string
	SetTwitterDataSource(val *string)
	TwitterDataSourceInput() *string
	UpdateAdobeAnalyticsDataSource() *string
	SetUpdateAdobeAnalyticsDataSource(val *string)
	UpdateAdobeAnalyticsDataSourceInput() *string
	UpdateAthenaDataSource() *string
	SetUpdateAthenaDataSource(val *string)
	UpdateAthenaDataSourceInput() *string
	UpdateAuroraDataSource() *string
	SetUpdateAuroraDataSource(val *string)
	UpdateAuroraDataSourceInput() *string
	UpdateDatabricksDataSource() *string
	SetUpdateDatabricksDataSource(val *string)
	UpdateDatabricksDataSourceInput() *string
	UpdateDb2DataSource() *string
	SetUpdateDb2DataSource(val *string)
	UpdateDb2DataSourceInput() *string
	UpdateDenodoDataSource() *string
	SetUpdateDenodoDataSource(val *string)
	UpdateDenodoDataSourceInput() *string
	UpdateDocumentDbDataSource() *string
	SetUpdateDocumentDbDataSource(val *string)
	UpdateDocumentDbDataSourceInput() *string
	UpdateDremioDataSource() *string
	SetUpdateDremioDataSource(val *string)
	UpdateDremioDataSourceInput() *string
	UpdateDynamoDbDataSource() *string
	SetUpdateDynamoDbDataSource(val *string)
	UpdateDynamoDbDataSourceInput() *string
	UpdateExasolDataSource() *string
	SetUpdateExasolDataSource(val *string)
	UpdateExasolDataSourceInput() *string
	UpdateFileDataSource() *string
	SetUpdateFileDataSource(val *string)
	UpdateFileDataSourceInput() *string
	UpdateGitHubDataSource() *string
	SetUpdateGitHubDataSource(val *string)
	UpdateGitHubDataSourceInput() *string
	UpdateGoogleAnalyticsDataSource() *string
	SetUpdateGoogleAnalyticsDataSource(val *string)
	UpdateGoogleAnalyticsDataSourceInput() *string
	UpdateGoogleBigQueryDataSource() *string
	SetUpdateGoogleBigQueryDataSource(val *string)
	UpdateGoogleBigQueryDataSourceInput() *string
	UpdateGoogleSheetsDataSource() *string
	SetUpdateGoogleSheetsDataSource(val *string)
	UpdateGoogleSheetsDataSourceInput() *string
	UpdateImpalaDataSource() *string
	SetUpdateImpalaDataSource(val *string)
	UpdateImpalaDataSourceInput() *string
	UpdateJiraDataSource() *string
	SetUpdateJiraDataSource(val *string)
	UpdateJiraDataSourceInput() *string
	UpdateMariaDbDataSource() *string
	SetUpdateMariaDbDataSource(val *string)
	UpdateMariaDbDataSourceInput() *string
	UpdateMongoAtlasDataSource() *string
	SetUpdateMongoAtlasDataSource(val *string)
	UpdateMongoAtlasDataSourceInput() *string
	UpdateMongoDbDataSource() *string
	SetUpdateMongoDbDataSource(val *string)
	UpdateMongoDbDataSourceInput() *string
	UpdateMySqlDataSource() *string
	SetUpdateMySqlDataSource(val *string)
	UpdateMySqlDataSourceInput() *string
	UpdateOpenSearchDataSource() *string
	SetUpdateOpenSearchDataSource(val *string)
	UpdateOpenSearchDataSourceInput() *string
	UpdateOracleDataSource() *string
	SetUpdateOracleDataSource(val *string)
	UpdateOracleDataSourceInput() *string
	UpdatePayPalDataSource() *string
	SetUpdatePayPalDataSource(val *string)
	UpdatePayPalDataSourceInput() *string
	UpdatePostgreSqlDataSource() *string
	SetUpdatePostgreSqlDataSource(val *string)
	UpdatePostgreSqlDataSourceInput() *string
	UpdatePrestoDataSource() *string
	SetUpdatePrestoDataSource(val *string)
	UpdatePrestoDataSourceInput() *string
	UpdateRadiantDataSource() *string
	SetUpdateRadiantDataSource(val *string)
	UpdateRadiantDataSourceInput() *string
	UpdateRdsDataSource() *string
	SetUpdateRdsDataSource(val *string)
	UpdateRdsDataSourceInput() *string
	UpdateRedshiftAutoDiscoveredDataSource() *string
	SetUpdateRedshiftAutoDiscoveredDataSource(val *string)
	UpdateRedshiftAutoDiscoveredDataSourceInput() *string
	UpdateRedshiftManualDataSource() *string
	SetUpdateRedshiftManualDataSource(val *string)
	UpdateRedshiftManualDataSourceInput() *string
	UpdateS3AnalyticsDataSource() *string
	SetUpdateS3AnalyticsDataSource(val *string)
	UpdateS3AnalyticsDataSourceInput() *string
	UpdateS3DataSource() *string
	SetUpdateS3DataSource(val *string)
	UpdateS3DataSourceInput() *string
	UpdateS3TablesDataSource() *string
	SetUpdateS3TablesDataSource(val *string)
	UpdateS3TablesDataSourceInput() *string
	UpdateSalesforceDataSource() *string
	SetUpdateSalesforceDataSource(val *string)
	UpdateSalesforceDataSourceInput() *string
	UpdateSapHanaDataSource() *string
	SetUpdateSapHanaDataSource(val *string)
	UpdateSapHanaDataSourceInput() *string
	UpdateServiceNowDataSource() *string
	SetUpdateServiceNowDataSource(val *string)
	UpdateServiceNowDataSourceInput() *string
	UpdateSnowflakeDataSource() *string
	SetUpdateSnowflakeDataSource(val *string)
	UpdateSnowflakeDataSourceInput() *string
	UpdateSparkDataSource() *string
	SetUpdateSparkDataSource(val *string)
	UpdateSparkDataSourceInput() *string
	UpdateSqlServerDataSource() *string
	SetUpdateSqlServerDataSource(val *string)
	UpdateSqlServerDataSourceInput() *string
	UpdateSquareDataSource() *string
	SetUpdateSquareDataSource(val *string)
	UpdateSquareDataSourceInput() *string
	UpdateStarburstDataSource() *string
	SetUpdateStarburstDataSource(val *string)
	UpdateStarburstDataSourceInput() *string
	UpdateTeradataDataSource() *string
	SetUpdateTeradataDataSource(val *string)
	UpdateTeradataDataSourceInput() *string
	UpdateTimestreamDataSource() *string
	SetUpdateTimestreamDataSource(val *string)
	UpdateTimestreamDataSourceInput() *string
	UpdateTrinoDataSource() *string
	SetUpdateTrinoDataSource(val *string)
	UpdateTrinoDataSourceInput() *string
	UpdateTwitterDataSource() *string
	SetUpdateTwitterDataSource(val *string)
	UpdateTwitterDataSourceInput() *string
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
	UseBedrockManagedKnowledgeBase() *string
	SetUseBedrockManagedKnowledgeBase(val *string)
	UseBedrockManagedKnowledgeBaseInput() *string
	UseBedrockModels() *string
	SetUseBedrockModels(val *string)
	UseBedrockModelsInput() *string
	UseBoxAgentAction() *string
	SetUseBoxAgentAction(val *string)
	UseBoxAgentActionInput() *string
	UseBoxKnowledgeBase() *string
	SetUseBoxKnowledgeBase(val *string)
	UseBoxKnowledgeBaseInput() *string
	UseBrowserExtension() *string
	SetUseBrowserExtension(val *string)
	UseBrowserExtensionInput() *string
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
	UseConfluenceKnowledgeBase() *string
	SetUseConfluenceKnowledgeBase(val *string)
	UseConfluenceKnowledgeBaseInput() *string
	UseExcelAddInExtension() *string
	SetUseExcelAddInExtension(val *string)
	UseExcelAddInExtensionInput() *string
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
	UseGoogleDriveKnowledgeBase() *string
	SetUseGoogleDriveKnowledgeBase(val *string)
	UseGoogleDriveKnowledgeBaseInput() *string
	UseHubspotAction() *string
	SetUseHubspotAction(val *string)
	UseHubspotActionInput() *string
	UseHuggingFaceAction() *string
	SetUseHuggingFaceAction(val *string)
	UseHuggingFaceActionInput() *string
	UseIdcKnowledgeBase() *string
	SetUseIdcKnowledgeBase(val *string)
	UseIdcKnowledgeBaseInput() *string
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
	UseOneDriveKnowledgeBase() *string
	SetUseOneDriveKnowledgeBase(val *string)
	UseOneDriveKnowledgeBaseInput() *string
	UseOpenApiAction() *string
	SetUseOpenApiAction(val *string)
	UseOpenApiActionInput() *string
	UseOutlookAddInExtension() *string
	SetUseOutlookAddInExtension(val *string)
	UseOutlookAddInExtensionInput() *string
	UsePagerDutyAction() *string
	SetUsePagerDutyAction(val *string)
	UsePagerDutyActionInput() *string
	UsePowerpointAddInExtension() *string
	SetUsePowerpointAddInExtension(val *string)
	UsePowerpointAddInExtensionInput() *string
	UseQBusinessKnowledgeBase() *string
	SetUseQBusinessKnowledgeBase(val *string)
	UseQBusinessKnowledgeBaseInput() *string
	UseS3KnowledgeBase() *string
	SetUseS3KnowledgeBase(val *string)
	UseS3KnowledgeBaseInput() *string
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
	UseSharePointKnowledgeBase() *string
	SetUseSharePointKnowledgeBase(val *string)
	UseSharePointKnowledgeBaseInput() *string
	UseSlackAction() *string
	SetUseSlackAction(val *string)
	UseSlackActionInput() *string
	UseSmartsheetAction() *string
	SetUseSmartsheetAction(val *string)
	UseSmartsheetActionInput() *string
	UseTextractAction() *string
	SetUseTextractAction(val *string)
	UseTextractActionInput() *string
	UseWebCrawlerKnowledgeBase() *string
	SetUseWebCrawlerKnowledgeBase(val *string)
	UseWebCrawlerKnowledgeBaseInput() *string
	UseWordAddInExtension() *string
	SetUseWordAddInExtension(val *string)
	UseWordAddInExtensionInput() *string
	UseZendeskAction() *string
	SetUseZendeskAction(val *string)
	UseZendeskActionInput() *string
	ViewAccountSpiceCapacity() *string
	SetViewAccountSpiceCapacity(val *string)
	ViewAccountSpiceCapacityInput() *string
	WebCrawlerKnowledgeBase() *string
	SetWebCrawlerKnowledgeBase(val *string)
	WebCrawlerKnowledgeBaseInput() *string
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
	ResetAdobeAnalyticsDataSource()
	ResetAmazonBedrockArsAction()
	ResetAmazonBedrockFsAction()
	ResetAmazonBedrockKrsAction()
	ResetAmazonSThreeAction()
	ResetAnalysis()
	ResetApproveFlowShareRequests()
	ResetApps()
	ResetAsanaAction()
	ResetAthenaDataSource()
	ResetAuroraDataSource()
	ResetAutomate()
	ResetBambooHrAction()
	ResetBedrockManagedKnowledgeBase()
	ResetBoxAgentAction()
	ResetBoxKnowledgeBase()
	ResetBuildCalculatedFieldWithQ()
	ResetCanvaAgentAction()
	ResetChatAgent()
	ResetComprehendAction()
	ResetComprehendMedicalAction()
	ResetConfluenceAction()
	ResetConfluenceKnowledgeBase()
	ResetCreateAdobeAnalyticsDataSource()
	ResetCreateAndUpdateAmazonBedrockArsAction()
	ResetCreateAndUpdateAmazonBedrockFsAction()
	ResetCreateAndUpdateAmazonBedrockKrsAction()
	ResetCreateAndUpdateAmazonSThreeAction()
	ResetCreateAndUpdateApps()
	ResetCreateAndUpdateAsanaAction()
	ResetCreateAndUpdateBambooHrAction()
	ResetCreateAndUpdateBedrockManagedKnowledgeBase()
	ResetCreateAndUpdateBoxAgentAction()
	ResetCreateAndUpdateBoxKnowledgeBase()
	ResetCreateAndUpdateCanvaAgentAction()
	ResetCreateAndUpdateComprehendAction()
	ResetCreateAndUpdateComprehendMedicalAction()
	ResetCreateAndUpdateConfluenceAction()
	ResetCreateAndUpdateConfluenceKnowledgeBase()
	ResetCreateAndUpdateDashboardEmailReports()
	ResetCreateAndUpdateDatasets()
	ResetCreateAndUpdateDataSources()
	ResetCreateAndUpdateFactSetAction()
	ResetCreateAndUpdateGenericHttpAction()
	ResetCreateAndUpdateGithubAction()
	ResetCreateAndUpdateGoogleCalendarAction()
	ResetCreateAndUpdateGoogleDriveKnowledgeBase()
	ResetCreateAndUpdateHubspotAction()
	ResetCreateAndUpdateHuggingFaceAction()
	ResetCreateAndUpdateIdcKnowledgeBase()
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
	ResetCreateAndUpdateOneDriveKnowledgeBase()
	ResetCreateAndUpdateOpenApiAction()
	ResetCreateAndUpdatePagerDutyAction()
	ResetCreateAndUpdateQBusinessKnowledgeBase()
	ResetCreateAndUpdateS3KnowledgeBase()
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
	ResetCreateAndUpdateSharePointKnowledgeBase()
	ResetCreateAndUpdateSlackAction()
	ResetCreateAndUpdateSmartsheetAction()
	ResetCreateAndUpdateTextractAction()
	ResetCreateAndUpdateThemes()
	ResetCreateAndUpdateThresholdAlerts()
	ResetCreateAndUpdateWebCrawlerKnowledgeBase()
	ResetCreateAndUpdateZendeskAction()
	ResetCreateAthenaDataSource()
	ResetCreateAuroraDataSource()
	ResetCreateChatAgents()
	ResetCreateDashboardExecutiveSummaryWithQ()
	ResetCreateDatabricksDataSource()
	ResetCreateDb2DataSource()
	ResetCreateDenodoDataSource()
	ResetCreateDocumentDbDataSource()
	ResetCreateDremioDataSource()
	ResetCreateDynamoDbDataSource()
	ResetCreateExasolDataSource()
	ResetCreateFileDataSource()
	ResetCreateGitHubDataSource()
	ResetCreateGoogleAnalyticsDataSource()
	ResetCreateGoogleBigQueryDataSource()
	ResetCreateGoogleSheetsDataSource()
	ResetCreateImpalaDataSource()
	ResetCreateJiraDataSource()
	ResetCreateMariaDbDataSource()
	ResetCreateMongoAtlasDataSource()
	ResetCreateMongoDbDataSource()
	ResetCreateMySqlDataSource()
	ResetCreateOpenSearchDataSource()
	ResetCreateOracleDataSource()
	ResetCreatePayPalDataSource()
	ResetCreatePostgreSqlDataSource()
	ResetCreatePrestoDataSource()
	ResetCreateRadiantDataSource()
	ResetCreateRdsDataSource()
	ResetCreateRedshiftAutoDiscoveredDataSource()
	ResetCreateRedshiftManualDataSource()
	ResetCreateS3AnalyticsDataSource()
	ResetCreateS3DataSource()
	ResetCreateS3TablesDataSource()
	ResetCreateSalesforceDataSource()
	ResetCreateSapHanaDataSource()
	ResetCreateServiceNowDataSource()
	ResetCreateSharedFolders()
	ResetCreateSnowflakeDataSource()
	ResetCreateSpaces()
	ResetCreateSparkDataSource()
	ResetCreateSpiceDataset()
	ResetCreateSqlServerDataSource()
	ResetCreateSquareDataSource()
	ResetCreateStarburstDataSource()
	ResetCreateTeradataDataSource()
	ResetCreateTimestreamDataSource()
	ResetCreateTrinoDataSource()
	ResetCreateTwitterDataSource()
	ResetDashboard()
	ResetDatabricksDataSource()
	ResetDb2DataSource()
	ResetDenodoDataSource()
	ResetDocumentDbDataSource()
	ResetDremioDataSource()
	ResetDynamoDbDataSource()
	ResetEditVisualWithQ()
	ResetExasolDataSource()
	ResetExportToCsv()
	ResetExportToCsvInScheduledReports()
	ResetExportToExcel()
	ResetExportToExcelInScheduledReports()
	ResetExportToPdf()
	ResetExportToPdfInScheduledReports()
	ResetExtension()
	ResetFactSetAction()
	ResetFileDataSource()
	ResetFlow()
	ResetGenericHttpAction()
	ResetGithubAction()
	ResetGitHubDataSource()
	ResetGoogleAnalyticsDataSource()
	ResetGoogleBigQueryDataSource()
	ResetGoogleCalendarAction()
	ResetGoogleDriveKnowledgeBase()
	ResetGoogleSheetsDataSource()
	ResetHubspotAction()
	ResetHuggingFaceAction()
	ResetIdcKnowledgeBase()
	ResetImpalaDataSource()
	ResetInboundEmailTrigger()
	ResetIncludeContentInScheduledReportsEmail()
	ResetIntercomAction()
	ResetInvokeAppsAiInference()
	ResetJiraAction()
	ResetJiraDataSource()
	ResetKnowledgeBase()
	ResetLinearAction()
	ResetManageSharedFolders()
	ResetMariaDbDataSource()
	ResetMcpAction()
	ResetMondayAction()
	ResetMongoAtlasDataSource()
	ResetMongoDbDataSource()
	ResetMsExchangeAction()
	ResetMsTeamsAction()
	ResetMySqlDataSource()
	ResetNewRelicAction()
	ResetNotionAction()
	ResetOneDriveAction()
	ResetOneDriveKnowledgeBase()
	ResetOpenApiAction()
	ResetOpenSearchDataSource()
	ResetOracleDataSource()
	ResetPagerDutyAction()
	ResetPayPalDataSource()
	ResetPerformFlowUiTask()
	ResetPostgreSqlDataSource()
	ResetPrestoDataSource()
	ResetPrintReports()
	ResetPublishWithoutApproval()
	ResetQBusinessKnowledgeBase()
	ResetQuickEventTrigger()
	ResetRadiantDataSource()
	ResetRdsDataSource()
	ResetRedshiftAutoDiscoveredDataSource()
	ResetRedshiftManualDataSource()
	ResetRenameSharedFolders()
	ResetResearch()
	ResetS3AnalyticsDataSource()
	ResetS3DataSource()
	ResetS3KnowledgeBase()
	ResetS3TablesDataSource()
	ResetSalesforceAction()
	ResetSalesforceDataSource()
	ResetSandPGlobalEnergyAction()
	ResetSandPgmiAction()
	ResetSapBillOfMaterialAction()
	ResetSapBusinessPartnerAction()
	ResetSapHanaDataSource()
	ResetSapMaterialStockAction()
	ResetSapPhysicalInventoryAction()
	ResetSapProductMasterDataAction()
	ResetScheduleTrigger()
	ResetServiceNowAction()
	ResetServiceNowDataSource()
	ResetShareAdobeAnalyticsDataSource()
	ResetShareAmazonBedrockArsAction()
	ResetShareAmazonBedrockFsAction()
	ResetShareAmazonBedrockKrsAction()
	ResetShareAmazonSThreeAction()
	ResetShareAnalyses()
	ResetShareApps()
	ResetShareAsanaAction()
	ResetShareAthenaDataSource()
	ResetShareAuroraDataSource()
	ResetShareBambooHrAction()
	ResetShareBedrockManagedKnowledgeBase()
	ResetShareBoxAgentAction()
	ResetShareBoxKnowledgeBase()
	ResetShareCanvaAgentAction()
	ResetShareChatAgents()
	ResetShareComprehendAction()
	ResetShareComprehendMedicalAction()
	ResetShareConfluenceAction()
	ResetShareConfluenceKnowledgeBase()
	ResetShareDashboards()
	ResetShareDatabricksDataSource()
	ResetShareDatasets()
	ResetShareDataSources()
	ResetShareDb2DataSource()
	ResetShareDenodoDataSource()
	ResetShareDocumentDbDataSource()
	ResetShareDremioDataSource()
	ResetShareDynamoDbDataSource()
	ResetShareExasolDataSource()
	ResetShareFactSetAction()
	ResetShareFileDataSource()
	ResetShareGenericHttpAction()
	ResetShareGithubAction()
	ResetShareGitHubDataSource()
	ResetShareGoogleAnalyticsDataSource()
	ResetShareGoogleBigQueryDataSource()
	ResetShareGoogleCalendarAction()
	ResetShareGoogleDriveKnowledgeBase()
	ResetShareGoogleSheetsDataSource()
	ResetShareHubspotAction()
	ResetShareHuggingFaceAction()
	ResetShareIdcKnowledgeBase()
	ResetShareImpalaDataSource()
	ResetShareIntercomAction()
	ResetShareJiraAction()
	ResetShareJiraDataSource()
	ResetShareKnowledgeBases()
	ResetShareLinearAction()
	ResetShareMariaDbDataSource()
	ResetShareMcpAction()
	ResetShareMondayAction()
	ResetShareMongoAtlasDataSource()
	ResetShareMongoDbDataSource()
	ResetShareMsExchangeAction()
	ResetShareMsTeamsAction()
	ResetShareMySqlDataSource()
	ResetShareNewRelicAction()
	ResetShareNotionAction()
	ResetShareOneDriveAction()
	ResetShareOneDriveKnowledgeBase()
	ResetShareOpenApiAction()
	ResetShareOpenSearchDataSource()
	ResetShareOracleDataSource()
	ResetSharePagerDutyAction()
	ResetSharePayPalDataSource()
	ResetSharePointAction()
	ResetSharePointKnowledgeBase()
	ResetSharePostgreSqlDataSource()
	ResetSharePrestoDataSource()
	ResetShareQBusinessKnowledgeBase()
	ResetShareRadiantDataSource()
	ResetShareRdsDataSource()
	ResetShareRedshiftAutoDiscoveredDataSource()
	ResetShareRedshiftManualDataSource()
	ResetShareS3AnalyticsDataSource()
	ResetShareS3DataSource()
	ResetShareS3KnowledgeBase()
	ResetShareS3TablesDataSource()
	ResetShareSalesforceAction()
	ResetShareSalesforceDataSource()
	ResetShareSandPGlobalEnergyAction()
	ResetShareSandPgmiAction()
	ResetShareSapBillOfMaterialAction()
	ResetShareSapBusinessPartnerAction()
	ResetShareSapHanaDataSource()
	ResetShareSapMaterialStockAction()
	ResetShareSapPhysicalInventoryAction()
	ResetShareSapProductMasterDataAction()
	ResetShareServiceNowAction()
	ResetShareServiceNowDataSource()
	ResetShareSharePointAction()
	ResetShareSharePointKnowledgeBase()
	ResetShareSlackAction()
	ResetShareSmartsheetAction()
	ResetShareSnowflakeDataSource()
	ResetShareSpaces()
	ResetShareSparkDataSource()
	ResetShareSqlServerDataSource()
	ResetShareSquareDataSource()
	ResetShareStarburstDataSource()
	ResetShareTeradataDataSource()
	ResetShareTextractAction()
	ResetShareTimestreamDataSource()
	ResetShareTrinoDataSource()
	ResetShareTwitterDataSource()
	ResetShareWebCrawlerKnowledgeBase()
	ResetShareZendeskAction()
	ResetSlackAction()
	ResetSmartsheetAction()
	ResetSnowflakeDataSource()
	ResetSpace()
	ResetSparkDataSource()
	ResetSqlServerDataSource()
	ResetSquareDataSource()
	ResetStarburstDataSource()
	ResetSubscribeDashboardEmailReports()
	ResetTeradataDataSource()
	ResetTextractAction()
	ResetTimestreamDataSource()
	ResetTopic()
	ResetTrigger()
	ResetTrinoDataSource()
	ResetTwitterDataSource()
	ResetUpdateAdobeAnalyticsDataSource()
	ResetUpdateAthenaDataSource()
	ResetUpdateAuroraDataSource()
	ResetUpdateDatabricksDataSource()
	ResetUpdateDb2DataSource()
	ResetUpdateDenodoDataSource()
	ResetUpdateDocumentDbDataSource()
	ResetUpdateDremioDataSource()
	ResetUpdateDynamoDbDataSource()
	ResetUpdateExasolDataSource()
	ResetUpdateFileDataSource()
	ResetUpdateGitHubDataSource()
	ResetUpdateGoogleAnalyticsDataSource()
	ResetUpdateGoogleBigQueryDataSource()
	ResetUpdateGoogleSheetsDataSource()
	ResetUpdateImpalaDataSource()
	ResetUpdateJiraDataSource()
	ResetUpdateMariaDbDataSource()
	ResetUpdateMongoAtlasDataSource()
	ResetUpdateMongoDbDataSource()
	ResetUpdateMySqlDataSource()
	ResetUpdateOpenSearchDataSource()
	ResetUpdateOracleDataSource()
	ResetUpdatePayPalDataSource()
	ResetUpdatePostgreSqlDataSource()
	ResetUpdatePrestoDataSource()
	ResetUpdateRadiantDataSource()
	ResetUpdateRdsDataSource()
	ResetUpdateRedshiftAutoDiscoveredDataSource()
	ResetUpdateRedshiftManualDataSource()
	ResetUpdateS3AnalyticsDataSource()
	ResetUpdateS3DataSource()
	ResetUpdateS3TablesDataSource()
	ResetUpdateSalesforceDataSource()
	ResetUpdateSapHanaDataSource()
	ResetUpdateServiceNowDataSource()
	ResetUpdateSnowflakeDataSource()
	ResetUpdateSparkDataSource()
	ResetUpdateSqlServerDataSource()
	ResetUpdateSquareDataSource()
	ResetUpdateStarburstDataSource()
	ResetUpdateTeradataDataSource()
	ResetUpdateTimestreamDataSource()
	ResetUpdateTrinoDataSource()
	ResetUpdateTwitterDataSource()
	ResetUseAgentWebSearch()
	ResetUseAmazonBedrockArsAction()
	ResetUseAmazonBedrockFsAction()
	ResetUseAmazonBedrockKrsAction()
	ResetUseAmazonSThreeAction()
	ResetUseAsanaAction()
	ResetUseBambooHrAction()
	ResetUseBedrockManagedKnowledgeBase()
	ResetUseBedrockModels()
	ResetUseBoxAgentAction()
	ResetUseBoxKnowledgeBase()
	ResetUseBrowserExtension()
	ResetUseCanvaAgentAction()
	ResetUseComprehendAction()
	ResetUseComprehendMedicalAction()
	ResetUseConfluenceAction()
	ResetUseConfluenceKnowledgeBase()
	ResetUseExcelAddInExtension()
	ResetUseFactSetAction()
	ResetUseGenericHttpAction()
	ResetUseGithubAction()
	ResetUseGoogleCalendarAction()
	ResetUseGoogleDriveKnowledgeBase()
	ResetUseHubspotAction()
	ResetUseHuggingFaceAction()
	ResetUseIdcKnowledgeBase()
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
	ResetUseOneDriveKnowledgeBase()
	ResetUseOpenApiAction()
	ResetUseOutlookAddInExtension()
	ResetUsePagerDutyAction()
	ResetUsePowerpointAddInExtension()
	ResetUseQBusinessKnowledgeBase()
	ResetUseS3KnowledgeBase()
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
	ResetUseSharePointKnowledgeBase()
	ResetUseSlackAction()
	ResetUseSmartsheetAction()
	ResetUseTextractAction()
	ResetUseWebCrawlerKnowledgeBase()
	ResetUseWordAddInExtension()
	ResetUseZendeskAction()
	ResetViewAccountSpiceCapacity()
	ResetWebCrawlerKnowledgeBase()
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adobeAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AdobeAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adobeAnalyticsDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"athenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AthenaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"athenaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) AuroraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bedrockManagedKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BedrockManagedKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bedrockManagedKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boxKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) BoxKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boxKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confluenceKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ConfluenceKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confluenceKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAdobeAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAdobeAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAdobeAnalyticsDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBedrockManagedKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBedrockManagedKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBedrockManagedKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBoxKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBoxKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBoxKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateConfluenceKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateConfluenceKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateConfluenceKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGoogleDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGoogleDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGoogleDriveKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateIdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateIdcKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateIdcKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateIdcKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOneDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOneDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOneDriveKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateQBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateQBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateQBusinessKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateQBusinessKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateS3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateS3KnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateS3KnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateS3KnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSharePointKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSharePointKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSharePointKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateWebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateWebCrawlerKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateWebCrawlerKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateWebCrawlerKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAthenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAthenaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAthenaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuroraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateAuroraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuroraDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDatabricksDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDatabricksDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDatabricksDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDb2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDb2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDb2DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDb2DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDenodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDenodoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDenodoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDocumentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDocumentDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDocumentDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDremioDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDremioDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDynamoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateDynamoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDynamoDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createExasolDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateExasolDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createExasolDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateFileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createFileDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateFileDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createFileDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGitHubDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGitHubDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleAnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleBigQueryDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleBigQueryDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleBigQueryDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleSheetsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleSheetsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleSheetsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createImpalaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateImpalaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createImpalaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateJiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createJiraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateJiraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createJiraDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMariaDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMariaDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMariaDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMongoAtlasDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMongoAtlasDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMongoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMongoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMongoDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMySqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateMySqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMySqlDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateOpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOpenSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateOpenSearchDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOpenSearchDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateOracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOracleDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateOracleDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOracleDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreatePayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPayPalDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreatePayPalDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPayPalDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreatePostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPostgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreatePostgreSqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPostgreSqlDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreatePrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPrestoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreatePrestoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPrestoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRadiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRadiantDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRadiantDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRdsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRdsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRedshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRedshiftAutoDiscoveredDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRedshiftAutoDiscoveredDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRedshiftManualDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateRedshiftManualDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRedshiftManualDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3AnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3AnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3TablesDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3TablesDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3TablesDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSalesforceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSalesforceDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSalesforceDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSapHanaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSapHanaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSapHanaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createServiceNowDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateServiceNowDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createServiceNowDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSnowflakeDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSnowflakeDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSnowflakeDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSparkDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSparkDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSparkDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSqlServerDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSqlServerDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSquareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateSquareDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSquareDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateStarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createStarburstDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateStarburstDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createStarburstDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTeradataDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTeradataDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTeradataDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimestreamDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTimestreamDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimestreamDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTrinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTrinoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTrinoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTwitterDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) CreateTwitterDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTwitterDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DatabricksDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Db2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"db2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Db2DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"db2DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DenodoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denodoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DocumentDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DremioDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dremioDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) DynamoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamoDbDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exasolDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ExasolDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exasolDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) FileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) FileDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GitHubDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitHubDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleBigQueryDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleBigQueryDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleBigQueryDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleDriveKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleSheetsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) GoogleSheetsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleSheetsDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) IdcKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impalaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ImpalaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impalaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InboundEmailTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundEmailTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) InboundEmailTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundEmailTriggerInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) JiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) JiraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariaDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MariaDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariaDbDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MongoAtlasDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoAtlasDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MongoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mySqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) MySqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mySqlDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OneDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneDriveKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OpenSearchDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openSearchDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) OracleDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payPalDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PayPalDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payPalDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PostgreSqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgreSqlDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prestoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) PrestoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prestoDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) QBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) QBusinessKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qBusinessKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) QuickEventTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickEventTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) QuickEventTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickEventTriggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RadiantDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiantDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RdsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RedshiftAutoDiscoveredDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftAutoDiscoveredDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftManualDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) RedshiftManualDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftManualDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3AnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3AnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3KnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3TablesDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) S3TablesDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3TablesDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SalesforceDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapHanaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SapHanaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapHanaDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ScheduleTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ScheduleTriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTriggerInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ServiceNowDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAdobeAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAdobeAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAdobeAnalyticsDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAthenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAthenaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAthenaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAuroraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareAuroraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAuroraDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBedrockManagedKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBedrockManagedKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBedrockManagedKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBoxKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareBoxKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBoxKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareConfluenceKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareConfluenceKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareConfluenceKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatabricksDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatabricksDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatabricksDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDb2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDb2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDb2DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDb2DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDenodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDenodoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDenodoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDocumentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDocumentDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDocumentDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDremioDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDremioDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDynamoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareDynamoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDynamoDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareExasolDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareExasolDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareExasolDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareFileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareFileDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareFileDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareFileDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGitHubDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGitHubDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleAnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleBigQueryDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleBigQueryDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleBigQueryDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleDriveKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleSheetsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleSheetsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleSheetsDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareIdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareIdcKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareIdcKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareIdcKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareImpalaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareImpalaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareImpalaDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareJiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareJiraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareJiraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareJiraDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMariaDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMariaDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMariaDbDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMongoAtlasDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMongoAtlasDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMongoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMongoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMongoDbDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMySqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareMySqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMySqlDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOneDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOneDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOneDriveKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOpenSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOpenSearchDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOpenSearchDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOracleDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareOracleDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOracleDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePayPalDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePayPalDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePayPalDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePointKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePostgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePostgreSqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePostgreSqlDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePrestoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SharePrestoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePrestoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareQBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareQBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareQBusinessKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareQBusinessKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRadiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRadiantDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRadiantDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRdsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRdsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRedshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRedshiftAutoDiscoveredDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRedshiftAutoDiscoveredDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRedshiftManualDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareRedshiftManualDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRedshiftManualDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3AnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3AnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3KnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3KnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3KnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3TablesDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3TablesDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3TablesDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSalesforceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSalesforceDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSalesforceDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapHanaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapHanaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapHanaDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareServiceNowDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareServiceNowDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareServiceNowDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSharePointKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSharePointKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSharePointKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSnowflakeDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSnowflakeDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSnowflakeDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSparkDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSparkDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSparkDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSqlServerDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSqlServerDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSquareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareSquareDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSquareDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareStarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareStarburstDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareStarburstDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareStarburstDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTeradataDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTeradataDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTeradataDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTimestreamDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTimestreamDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTimestreamDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTrinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTrinoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTrinoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTwitterDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareTwitterDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTwitterDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareWebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareWebCrawlerKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ShareWebCrawlerKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareWebCrawlerKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SnowflakeDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SparkDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SqlServerDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServerDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) SquareDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squareDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) StarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"starburstDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) StarburstDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"starburstDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teradataDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TeradataDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teradataDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestreamDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TimestreamDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestreamDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) Trigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TriggerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TrinoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trinoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twitterDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) TwitterDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twitterDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAdobeAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAdobeAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAdobeAnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAthenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAthenaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAthenaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAuroraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAuroraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAuroraDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDatabricksDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDatabricksDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDatabricksDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDb2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDb2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDb2DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDb2DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDenodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDenodoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDenodoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDocumentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDocumentDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDocumentDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDremioDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDremioDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDynamoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDynamoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDynamoDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateExasolDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateExasolDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateExasolDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateFileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateFileDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateFileDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateFileDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGitHubDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGitHubDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleAnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleAnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleBigQueryDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleBigQueryDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleBigQueryDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleSheetsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleSheetsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleSheetsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateImpalaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateImpalaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateImpalaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateJiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateJiraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateJiraDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateJiraDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMariaDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMariaDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMariaDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMongoAtlasDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMongoAtlasDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMongoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMongoDbDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMongoDbDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMySqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMySqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMySqlDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateOpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateOpenSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateOpenSearchDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateOpenSearchDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateOracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateOracleDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateOracleDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateOracleDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePayPalDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePayPalDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePayPalDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePostgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePostgreSqlDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePostgreSqlDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePrestoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePrestoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePrestoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRadiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRadiantDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRadiantDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRdsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRdsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRedshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRedshiftAutoDiscoveredDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRedshiftAutoDiscoveredDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRedshiftManualDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRedshiftManualDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRedshiftManualDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3AnalyticsDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3AnalyticsDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3DataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3DataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3TablesDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3TablesDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3TablesDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSalesforceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSalesforceDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSalesforceDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSapHanaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSapHanaDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSapHanaDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateServiceNowDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateServiceNowDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateServiceNowDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSnowflakeDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSnowflakeDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSnowflakeDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSparkDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSparkDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSparkDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSqlServerDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSqlServerDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSquareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSquareDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSquareDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateStarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateStarburstDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateStarburstDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateStarburstDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTeradataDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTeradataDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTeradataDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimestreamDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTimestreamDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimestreamDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTrinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTrinoDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTrinoDataSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTwitterDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTwitterDataSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTwitterDataSourceInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBedrockManagedKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBedrockManagedKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBedrockManagedKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBoxKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBoxKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBoxKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBrowserExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBrowserExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseBrowserExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBrowserExtensionInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useConfluenceKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseConfluenceKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useConfluenceKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseExcelAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useExcelAddInExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseExcelAddInExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useExcelAddInExtensionInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGoogleDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseGoogleDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGoogleDriveKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseIdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useIdcKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseIdcKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useIdcKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOneDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOneDriveKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOneDriveKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOutlookAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOutlookAddInExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseOutlookAddInExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOutlookAddInExtensionInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UsePowerpointAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePowerpointAddInExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UsePowerpointAddInExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePowerpointAddInExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseQBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useQBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseQBusinessKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useQBusinessKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseS3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useS3KnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseS3KnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useS3KnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSharePointKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseSharePointKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSharePointKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseWebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useWebCrawlerKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseWebCrawlerKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useWebCrawlerKnowledgeBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseWordAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useWordAddInExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) UseWordAddInExtensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useWordAddInExtensionInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) WebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webCrawlerKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) WebCrawlerKnowledgeBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webCrawlerKnowledgeBaseInput",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAdobeAnalyticsDataSource(val *string) {
	if err := j.validateSetAdobeAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adobeAnalyticsDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAthenaDataSource(val *string) {
	if err := j.validateSetAthenaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"athenaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetAuroraDataSource(val *string) {
	if err := j.validateSetAuroraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auroraDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetBedrockManagedKnowledgeBase(val *string) {
	if err := j.validateSetBedrockManagedKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetBoxKnowledgeBase(val *string) {
	if err := j.validateSetBoxKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boxKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetConfluenceKnowledgeBase(val *string) {
	if err := j.validateSetConfluenceKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confluenceKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAdobeAnalyticsDataSource(val *string) {
	if err := j.validateSetCreateAdobeAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAdobeAnalyticsDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateBedrockManagedKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateBedrockManagedKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateBedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateBoxKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateBoxKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateBoxKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateConfluenceKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateConfluenceKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateConfluenceKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateGoogleDriveKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateGoogleDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateGoogleDriveKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateIdcKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateIdcKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateIdcKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateOneDriveKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateOneDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateOneDriveKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateQBusinessKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateQBusinessKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateQBusinessKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateS3KnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateS3KnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateS3KnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateSharePointKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateSharePointKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateSharePointKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAndUpdateWebCrawlerKnowledgeBase(val *string) {
	if err := j.validateSetCreateAndUpdateWebCrawlerKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateWebCrawlerKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAthenaDataSource(val *string) {
	if err := j.validateSetCreateAthenaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAthenaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateAuroraDataSource(val *string) {
	if err := j.validateSetCreateAuroraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAuroraDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateDatabricksDataSource(val *string) {
	if err := j.validateSetCreateDatabricksDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDatabricksDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateDb2DataSource(val *string) {
	if err := j.validateSetCreateDb2DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDb2DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateDenodoDataSource(val *string) {
	if err := j.validateSetCreateDenodoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDenodoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateDocumentDbDataSource(val *string) {
	if err := j.validateSetCreateDocumentDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDocumentDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateDremioDataSource(val *string) {
	if err := j.validateSetCreateDremioDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDremioDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateDynamoDbDataSource(val *string) {
	if err := j.validateSetCreateDynamoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createDynamoDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateExasolDataSource(val *string) {
	if err := j.validateSetCreateExasolDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createExasolDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateFileDataSource(val *string) {
	if err := j.validateSetCreateFileDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createFileDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateGitHubDataSource(val *string) {
	if err := j.validateSetCreateGitHubDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createGitHubDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateGoogleAnalyticsDataSource(val *string) {
	if err := j.validateSetCreateGoogleAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createGoogleAnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateGoogleBigQueryDataSource(val *string) {
	if err := j.validateSetCreateGoogleBigQueryDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createGoogleBigQueryDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateGoogleSheetsDataSource(val *string) {
	if err := j.validateSetCreateGoogleSheetsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createGoogleSheetsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateImpalaDataSource(val *string) {
	if err := j.validateSetCreateImpalaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createImpalaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateJiraDataSource(val *string) {
	if err := j.validateSetCreateJiraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createJiraDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateMariaDbDataSource(val *string) {
	if err := j.validateSetCreateMariaDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMariaDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateMongoAtlasDataSource(val *string) {
	if err := j.validateSetCreateMongoAtlasDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMongoAtlasDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateMongoDbDataSource(val *string) {
	if err := j.validateSetCreateMongoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMongoDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateMySqlDataSource(val *string) {
	if err := j.validateSetCreateMySqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createMySqlDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateOpenSearchDataSource(val *string) {
	if err := j.validateSetCreateOpenSearchDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createOpenSearchDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateOracleDataSource(val *string) {
	if err := j.validateSetCreateOracleDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createOracleDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreatePayPalDataSource(val *string) {
	if err := j.validateSetCreatePayPalDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPayPalDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreatePostgreSqlDataSource(val *string) {
	if err := j.validateSetCreatePostgreSqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPostgreSqlDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreatePrestoDataSource(val *string) {
	if err := j.validateSetCreatePrestoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createPrestoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateRadiantDataSource(val *string) {
	if err := j.validateSetCreateRadiantDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createRadiantDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateRdsDataSource(val *string) {
	if err := j.validateSetCreateRdsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createRdsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateRedshiftAutoDiscoveredDataSource(val *string) {
	if err := j.validateSetCreateRedshiftAutoDiscoveredDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createRedshiftAutoDiscoveredDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateRedshiftManualDataSource(val *string) {
	if err := j.validateSetCreateRedshiftManualDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createRedshiftManualDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateS3AnalyticsDataSource(val *string) {
	if err := j.validateSetCreateS3AnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createS3AnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateS3DataSource(val *string) {
	if err := j.validateSetCreateS3DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createS3DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateS3TablesDataSource(val *string) {
	if err := j.validateSetCreateS3TablesDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createS3TablesDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSalesforceDataSource(val *string) {
	if err := j.validateSetCreateSalesforceDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSalesforceDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSapHanaDataSource(val *string) {
	if err := j.validateSetCreateSapHanaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSapHanaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateServiceNowDataSource(val *string) {
	if err := j.validateSetCreateServiceNowDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createServiceNowDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSnowflakeDataSource(val *string) {
	if err := j.validateSetCreateSnowflakeDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSnowflakeDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSparkDataSource(val *string) {
	if err := j.validateSetCreateSparkDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSparkDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSqlServerDataSource(val *string) {
	if err := j.validateSetCreateSqlServerDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSqlServerDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateSquareDataSource(val *string) {
	if err := j.validateSetCreateSquareDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSquareDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateStarburstDataSource(val *string) {
	if err := j.validateSetCreateStarburstDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createStarburstDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateTeradataDataSource(val *string) {
	if err := j.validateSetCreateTeradataDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTeradataDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateTimestreamDataSource(val *string) {
	if err := j.validateSetCreateTimestreamDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTimestreamDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateTrinoDataSource(val *string) {
	if err := j.validateSetCreateTrinoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTrinoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetCreateTwitterDataSource(val *string) {
	if err := j.validateSetCreateTwitterDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createTwitterDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetDatabricksDataSource(val *string) {
	if err := j.validateSetDatabricksDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databricksDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetDb2DataSource(val *string) {
	if err := j.validateSetDb2DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"db2DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetDenodoDataSource(val *string) {
	if err := j.validateSetDenodoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denodoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetDocumentDbDataSource(val *string) {
	if err := j.validateSetDocumentDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetDremioDataSource(val *string) {
	if err := j.validateSetDremioDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dremioDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetDynamoDbDataSource(val *string) {
	if err := j.validateSetDynamoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dynamoDbDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetExasolDataSource(val *string) {
	if err := j.validateSetExasolDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exasolDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetFileDataSource(val *string) {
	if err := j.validateSetFileDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGitHubDataSource(val *string) {
	if err := j.validateSetGitHubDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gitHubDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGoogleAnalyticsDataSource(val *string) {
	if err := j.validateSetGoogleAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleAnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGoogleBigQueryDataSource(val *string) {
	if err := j.validateSetGoogleBigQueryDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleBigQueryDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGoogleDriveKnowledgeBase(val *string) {
	if err := j.validateSetGoogleDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleDriveKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetGoogleSheetsDataSource(val *string) {
	if err := j.validateSetGoogleSheetsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"googleSheetsDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetIdcKnowledgeBase(val *string) {
	if err := j.validateSetIdcKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idcKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetImpalaDataSource(val *string) {
	if err := j.validateSetImpalaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"impalaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetInboundEmailTrigger(val *string) {
	if err := j.validateSetInboundEmailTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inboundEmailTrigger",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetJiraDataSource(val *string) {
	if err := j.validateSetJiraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jiraDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMariaDbDataSource(val *string) {
	if err := j.validateSetMariaDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mariaDbDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMongoAtlasDataSource(val *string) {
	if err := j.validateSetMongoAtlasDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongoAtlasDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMongoDbDataSource(val *string) {
	if err := j.validateSetMongoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mongoDbDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetMySqlDataSource(val *string) {
	if err := j.validateSetMySqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mySqlDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetOneDriveKnowledgeBase(val *string) {
	if err := j.validateSetOneDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oneDriveKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetOpenSearchDataSource(val *string) {
	if err := j.validateSetOpenSearchDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"openSearchDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetOracleDataSource(val *string) {
	if err := j.validateSetOracleDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oracleDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPayPalDataSource(val *string) {
	if err := j.validateSetPayPalDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payPalDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPostgreSqlDataSource(val *string) {
	if err := j.validateSetPostgreSqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postgreSqlDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetPrestoDataSource(val *string) {
	if err := j.validateSetPrestoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prestoDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetQBusinessKnowledgeBase(val *string) {
	if err := j.validateSetQBusinessKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qBusinessKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetQuickEventTrigger(val *string) {
	if err := j.validateSetQuickEventTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quickEventTrigger",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetRadiantDataSource(val *string) {
	if err := j.validateSetRadiantDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"radiantDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetRdsDataSource(val *string) {
	if err := j.validateSetRdsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetRedshiftAutoDiscoveredDataSource(val *string) {
	if err := j.validateSetRedshiftAutoDiscoveredDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redshiftAutoDiscoveredDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetRedshiftManualDataSource(val *string) {
	if err := j.validateSetRedshiftManualDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redshiftManualDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetS3AnalyticsDataSource(val *string) {
	if err := j.validateSetS3AnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3AnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetS3DataSource(val *string) {
	if err := j.validateSetS3DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetS3KnowledgeBase(val *string) {
	if err := j.validateSetS3KnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3KnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetS3TablesDataSource(val *string) {
	if err := j.validateSetS3TablesDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3TablesDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSalesforceDataSource(val *string) {
	if err := j.validateSetSalesforceDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salesforceDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSapHanaDataSource(val *string) {
	if err := j.validateSetSapHanaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sapHanaDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetScheduleTrigger(val *string) {
	if err := j.validateSetScheduleTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheduleTrigger",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetServiceNowDataSource(val *string) {
	if err := j.validateSetServiceNowDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceNowDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAdobeAnalyticsDataSource(val *string) {
	if err := j.validateSetShareAdobeAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAdobeAnalyticsDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAthenaDataSource(val *string) {
	if err := j.validateSetShareAthenaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAthenaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareAuroraDataSource(val *string) {
	if err := j.validateSetShareAuroraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAuroraDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareBedrockManagedKnowledgeBase(val *string) {
	if err := j.validateSetShareBedrockManagedKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareBedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareBoxKnowledgeBase(val *string) {
	if err := j.validateSetShareBoxKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareBoxKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareConfluenceKnowledgeBase(val *string) {
	if err := j.validateSetShareConfluenceKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareConfluenceKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDatabricksDataSource(val *string) {
	if err := j.validateSetShareDatabricksDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDatabricksDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDb2DataSource(val *string) {
	if err := j.validateSetShareDb2DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDb2DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDenodoDataSource(val *string) {
	if err := j.validateSetShareDenodoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDenodoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDocumentDbDataSource(val *string) {
	if err := j.validateSetShareDocumentDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDocumentDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDremioDataSource(val *string) {
	if err := j.validateSetShareDremioDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDremioDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareDynamoDbDataSource(val *string) {
	if err := j.validateSetShareDynamoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDynamoDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareExasolDataSource(val *string) {
	if err := j.validateSetShareExasolDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareExasolDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareFileDataSource(val *string) {
	if err := j.validateSetShareFileDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareFileDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGitHubDataSource(val *string) {
	if err := j.validateSetShareGitHubDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGitHubDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGoogleAnalyticsDataSource(val *string) {
	if err := j.validateSetShareGoogleAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGoogleAnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGoogleBigQueryDataSource(val *string) {
	if err := j.validateSetShareGoogleBigQueryDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGoogleBigQueryDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGoogleDriveKnowledgeBase(val *string) {
	if err := j.validateSetShareGoogleDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGoogleDriveKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareGoogleSheetsDataSource(val *string) {
	if err := j.validateSetShareGoogleSheetsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareGoogleSheetsDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareIdcKnowledgeBase(val *string) {
	if err := j.validateSetShareIdcKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareIdcKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareImpalaDataSource(val *string) {
	if err := j.validateSetShareImpalaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareImpalaDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareJiraDataSource(val *string) {
	if err := j.validateSetShareJiraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareJiraDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMariaDbDataSource(val *string) {
	if err := j.validateSetShareMariaDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMariaDbDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMongoAtlasDataSource(val *string) {
	if err := j.validateSetShareMongoAtlasDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMongoAtlasDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMongoDbDataSource(val *string) {
	if err := j.validateSetShareMongoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMongoDbDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareMySqlDataSource(val *string) {
	if err := j.validateSetShareMySqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareMySqlDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareOneDriveKnowledgeBase(val *string) {
	if err := j.validateSetShareOneDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareOneDriveKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareOpenSearchDataSource(val *string) {
	if err := j.validateSetShareOpenSearchDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareOpenSearchDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareOracleDataSource(val *string) {
	if err := j.validateSetShareOracleDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareOracleDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSharePayPalDataSource(val *string) {
	if err := j.validateSetSharePayPalDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePayPalDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSharePointKnowledgeBase(val *string) {
	if err := j.validateSetSharePointKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePointKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSharePostgreSqlDataSource(val *string) {
	if err := j.validateSetSharePostgreSqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePostgreSqlDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSharePrestoDataSource(val *string) {
	if err := j.validateSetSharePrestoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharePrestoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareQBusinessKnowledgeBase(val *string) {
	if err := j.validateSetShareQBusinessKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareQBusinessKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareRadiantDataSource(val *string) {
	if err := j.validateSetShareRadiantDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareRadiantDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareRdsDataSource(val *string) {
	if err := j.validateSetShareRdsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareRdsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareRedshiftAutoDiscoveredDataSource(val *string) {
	if err := j.validateSetShareRedshiftAutoDiscoveredDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareRedshiftAutoDiscoveredDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareRedshiftManualDataSource(val *string) {
	if err := j.validateSetShareRedshiftManualDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareRedshiftManualDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareS3AnalyticsDataSource(val *string) {
	if err := j.validateSetShareS3AnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareS3AnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareS3DataSource(val *string) {
	if err := j.validateSetShareS3DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareS3DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareS3KnowledgeBase(val *string) {
	if err := j.validateSetShareS3KnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareS3KnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareS3TablesDataSource(val *string) {
	if err := j.validateSetShareS3TablesDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareS3TablesDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSalesforceDataSource(val *string) {
	if err := j.validateSetShareSalesforceDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSalesforceDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSapHanaDataSource(val *string) {
	if err := j.validateSetShareSapHanaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSapHanaDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareServiceNowDataSource(val *string) {
	if err := j.validateSetShareServiceNowDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareServiceNowDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSharePointKnowledgeBase(val *string) {
	if err := j.validateSetShareSharePointKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSharePointKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSnowflakeDataSource(val *string) {
	if err := j.validateSetShareSnowflakeDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSnowflakeDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSparkDataSource(val *string) {
	if err := j.validateSetShareSparkDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSparkDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSqlServerDataSource(val *string) {
	if err := j.validateSetShareSqlServerDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSqlServerDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareSquareDataSource(val *string) {
	if err := j.validateSetShareSquareDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareSquareDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareStarburstDataSource(val *string) {
	if err := j.validateSetShareStarburstDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareStarburstDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareTeradataDataSource(val *string) {
	if err := j.validateSetShareTeradataDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareTeradataDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareTimestreamDataSource(val *string) {
	if err := j.validateSetShareTimestreamDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareTimestreamDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareTrinoDataSource(val *string) {
	if err := j.validateSetShareTrinoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareTrinoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareTwitterDataSource(val *string) {
	if err := j.validateSetShareTwitterDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareTwitterDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetShareWebCrawlerKnowledgeBase(val *string) {
	if err := j.validateSetShareWebCrawlerKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareWebCrawlerKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSnowflakeDataSource(val *string) {
	if err := j.validateSetSnowflakeDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snowflakeDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSparkDataSource(val *string) {
	if err := j.validateSetSparkDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sparkDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSqlServerDataSource(val *string) {
	if err := j.validateSetSqlServerDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlServerDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetSquareDataSource(val *string) {
	if err := j.validateSetSquareDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squareDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetStarburstDataSource(val *string) {
	if err := j.validateSetStarburstDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"starburstDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTeradataDataSource(val *string) {
	if err := j.validateSetTeradataDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teradataDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTimestreamDataSource(val *string) {
	if err := j.validateSetTimestreamDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timestreamDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTrigger(val *string) {
	if err := j.validateSetTriggerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trigger",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTrinoDataSource(val *string) {
	if err := j.validateSetTrinoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trinoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetTwitterDataSource(val *string) {
	if err := j.validateSetTwitterDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twitterDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateAdobeAnalyticsDataSource(val *string) {
	if err := j.validateSetUpdateAdobeAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateAdobeAnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateAthenaDataSource(val *string) {
	if err := j.validateSetUpdateAthenaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateAthenaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateAuroraDataSource(val *string) {
	if err := j.validateSetUpdateAuroraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateAuroraDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateDatabricksDataSource(val *string) {
	if err := j.validateSetUpdateDatabricksDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDatabricksDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateDb2DataSource(val *string) {
	if err := j.validateSetUpdateDb2DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDb2DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateDenodoDataSource(val *string) {
	if err := j.validateSetUpdateDenodoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDenodoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateDocumentDbDataSource(val *string) {
	if err := j.validateSetUpdateDocumentDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDocumentDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateDremioDataSource(val *string) {
	if err := j.validateSetUpdateDremioDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDremioDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateDynamoDbDataSource(val *string) {
	if err := j.validateSetUpdateDynamoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateDynamoDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateExasolDataSource(val *string) {
	if err := j.validateSetUpdateExasolDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateExasolDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateFileDataSource(val *string) {
	if err := j.validateSetUpdateFileDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateFileDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateGitHubDataSource(val *string) {
	if err := j.validateSetUpdateGitHubDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateGitHubDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateGoogleAnalyticsDataSource(val *string) {
	if err := j.validateSetUpdateGoogleAnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateGoogleAnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateGoogleBigQueryDataSource(val *string) {
	if err := j.validateSetUpdateGoogleBigQueryDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateGoogleBigQueryDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateGoogleSheetsDataSource(val *string) {
	if err := j.validateSetUpdateGoogleSheetsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateGoogleSheetsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateImpalaDataSource(val *string) {
	if err := j.validateSetUpdateImpalaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateImpalaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateJiraDataSource(val *string) {
	if err := j.validateSetUpdateJiraDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateJiraDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateMariaDbDataSource(val *string) {
	if err := j.validateSetUpdateMariaDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateMariaDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateMongoAtlasDataSource(val *string) {
	if err := j.validateSetUpdateMongoAtlasDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateMongoAtlasDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateMongoDbDataSource(val *string) {
	if err := j.validateSetUpdateMongoDbDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateMongoDbDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateMySqlDataSource(val *string) {
	if err := j.validateSetUpdateMySqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateMySqlDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateOpenSearchDataSource(val *string) {
	if err := j.validateSetUpdateOpenSearchDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateOpenSearchDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateOracleDataSource(val *string) {
	if err := j.validateSetUpdateOracleDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateOracleDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdatePayPalDataSource(val *string) {
	if err := j.validateSetUpdatePayPalDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatePayPalDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdatePostgreSqlDataSource(val *string) {
	if err := j.validateSetUpdatePostgreSqlDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatePostgreSqlDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdatePrestoDataSource(val *string) {
	if err := j.validateSetUpdatePrestoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatePrestoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateRadiantDataSource(val *string) {
	if err := j.validateSetUpdateRadiantDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateRadiantDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateRdsDataSource(val *string) {
	if err := j.validateSetUpdateRdsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateRdsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateRedshiftAutoDiscoveredDataSource(val *string) {
	if err := j.validateSetUpdateRedshiftAutoDiscoveredDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateRedshiftAutoDiscoveredDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateRedshiftManualDataSource(val *string) {
	if err := j.validateSetUpdateRedshiftManualDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateRedshiftManualDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateS3AnalyticsDataSource(val *string) {
	if err := j.validateSetUpdateS3AnalyticsDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateS3AnalyticsDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateS3DataSource(val *string) {
	if err := j.validateSetUpdateS3DataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateS3DataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateS3TablesDataSource(val *string) {
	if err := j.validateSetUpdateS3TablesDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateS3TablesDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateSalesforceDataSource(val *string) {
	if err := j.validateSetUpdateSalesforceDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSalesforceDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateSapHanaDataSource(val *string) {
	if err := j.validateSetUpdateSapHanaDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSapHanaDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateServiceNowDataSource(val *string) {
	if err := j.validateSetUpdateServiceNowDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateServiceNowDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateSnowflakeDataSource(val *string) {
	if err := j.validateSetUpdateSnowflakeDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSnowflakeDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateSparkDataSource(val *string) {
	if err := j.validateSetUpdateSparkDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSparkDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateSqlServerDataSource(val *string) {
	if err := j.validateSetUpdateSqlServerDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSqlServerDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateSquareDataSource(val *string) {
	if err := j.validateSetUpdateSquareDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateSquareDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateStarburstDataSource(val *string) {
	if err := j.validateSetUpdateStarburstDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateStarburstDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateTeradataDataSource(val *string) {
	if err := j.validateSetUpdateTeradataDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTeradataDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateTimestreamDataSource(val *string) {
	if err := j.validateSetUpdateTimestreamDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTimestreamDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateTrinoDataSource(val *string) {
	if err := j.validateSetUpdateTrinoDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTrinoDataSource",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUpdateTwitterDataSource(val *string) {
	if err := j.validateSetUpdateTwitterDataSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateTwitterDataSource",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseBedrockManagedKnowledgeBase(val *string) {
	if err := j.validateSetUseBedrockManagedKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseBoxKnowledgeBase(val *string) {
	if err := j.validateSetUseBoxKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBoxKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseBrowserExtension(val *string) {
	if err := j.validateSetUseBrowserExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBrowserExtension",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseConfluenceKnowledgeBase(val *string) {
	if err := j.validateSetUseConfluenceKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useConfluenceKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseExcelAddInExtension(val *string) {
	if err := j.validateSetUseExcelAddInExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useExcelAddInExtension",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseGoogleDriveKnowledgeBase(val *string) {
	if err := j.validateSetUseGoogleDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useGoogleDriveKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseIdcKnowledgeBase(val *string) {
	if err := j.validateSetUseIdcKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useIdcKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseOneDriveKnowledgeBase(val *string) {
	if err := j.validateSetUseOneDriveKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOneDriveKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseOutlookAddInExtension(val *string) {
	if err := j.validateSetUseOutlookAddInExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOutlookAddInExtension",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUsePowerpointAddInExtension(val *string) {
	if err := j.validateSetUsePowerpointAddInExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePowerpointAddInExtension",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseQBusinessKnowledgeBase(val *string) {
	if err := j.validateSetUseQBusinessKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useQBusinessKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseS3KnowledgeBase(val *string) {
	if err := j.validateSetUseS3KnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useS3KnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseSharePointKnowledgeBase(val *string) {
	if err := j.validateSetUseSharePointKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useSharePointKnowledgeBase",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseWebCrawlerKnowledgeBase(val *string) {
	if err := j.validateSetUseWebCrawlerKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useWebCrawlerKnowledgeBase",
		val,
	)
}

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetUseWordAddInExtension(val *string) {
	if err := j.validateSetUseWordAddInExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useWordAddInExtension",
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

func (j *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference)SetWebCrawlerKnowledgeBase(val *string) {
	if err := j.validateSetWebCrawlerKnowledgeBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webCrawlerKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAdobeAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetAdobeAnalyticsDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAthenaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetAthenaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetAuroraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetAuroraDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetBedrockManagedKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetBedrockManagedKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetBoxKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetBoxKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetConfluenceKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetConfluenceKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAdobeAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAdobeAnalyticsDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateBedrockManagedKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateBedrockManagedKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateBoxKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateBoxKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateConfluenceKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateConfluenceKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateGoogleDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateGoogleDriveKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateIdcKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateIdcKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateOneDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateOneDriveKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateQBusinessKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateQBusinessKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateS3KnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateS3KnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateSharePointKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateSharePointKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAndUpdateWebCrawlerKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAndUpdateWebCrawlerKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAthenaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAthenaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateAuroraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateAuroraDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateDatabricksDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateDatabricksDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateDb2DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateDb2DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateDenodoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateDenodoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateDocumentDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateDocumentDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateDremioDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateDremioDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateDynamoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateDynamoDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateExasolDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateExasolDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateFileDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateFileDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateGitHubDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateGitHubDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateGoogleAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateGoogleAnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateGoogleBigQueryDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateGoogleBigQueryDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateGoogleSheetsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateGoogleSheetsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateImpalaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateImpalaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateJiraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateJiraDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateMariaDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateMariaDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateMongoAtlasDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateMongoAtlasDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateMongoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateMongoDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateMySqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateMySqlDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateOpenSearchDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateOpenSearchDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateOracleDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateOracleDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreatePayPalDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreatePayPalDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreatePostgreSqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreatePostgreSqlDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreatePrestoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreatePrestoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateRadiantDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateRadiantDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateRdsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateRdsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateRedshiftAutoDiscoveredDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateRedshiftAutoDiscoveredDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateRedshiftManualDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateRedshiftManualDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateS3AnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateS3AnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateS3DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateS3DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateS3TablesDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateS3TablesDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSalesforceDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSalesforceDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSapHanaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSapHanaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateServiceNowDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateServiceNowDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSnowflakeDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSnowflakeDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSparkDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSparkDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSqlServerDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSqlServerDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateSquareDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateSquareDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateStarburstDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateStarburstDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateTeradataDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateTeradataDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateTimestreamDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateTimestreamDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateTrinoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateTrinoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetCreateTwitterDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetCreateTwitterDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetDatabricksDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetDatabricksDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetDb2DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetDb2DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetDenodoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetDenodoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetDocumentDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetDocumentDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetDremioDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetDremioDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetDynamoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetDynamoDbDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetExasolDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetExasolDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetFileDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetFileDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGitHubDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetGitHubDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGoogleAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetGoogleAnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGoogleBigQueryDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetGoogleBigQueryDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGoogleDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetGoogleDriveKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetGoogleSheetsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetGoogleSheetsDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetIdcKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetIdcKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetImpalaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetImpalaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetInboundEmailTrigger() {
	_jsii_.InvokeVoid(
		q,
		"resetInboundEmailTrigger",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetJiraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetJiraDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMariaDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetMariaDbDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMongoAtlasDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetMongoAtlasDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMongoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetMongoDbDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetMySqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetMySqlDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetOneDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetOneDriveKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetOpenSearchDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetOpenSearchDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetOracleDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetOracleDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPayPalDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetPayPalDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPostgreSqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetPostgreSqlDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetPrestoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetPrestoDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetQBusinessKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetQBusinessKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetQuickEventTrigger() {
	_jsii_.InvokeVoid(
		q,
		"resetQuickEventTrigger",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetRadiantDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetRadiantDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetRdsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetRdsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetRedshiftAutoDiscoveredDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetRedshiftAutoDiscoveredDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetRedshiftManualDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetRedshiftManualDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetS3AnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetS3AnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetS3DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetS3DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetS3KnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetS3KnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetS3TablesDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetS3TablesDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSalesforceDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSalesforceDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSapHanaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSapHanaDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetScheduleTrigger() {
	_jsii_.InvokeVoid(
		q,
		"resetScheduleTrigger",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetServiceNowDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetServiceNowDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAdobeAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAdobeAnalyticsDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAthenaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAthenaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareAuroraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareAuroraDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareBedrockManagedKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareBedrockManagedKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareBoxKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareBoxKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareConfluenceKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareConfluenceKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDatabricksDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDatabricksDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDb2DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDb2DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDenodoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDenodoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDocumentDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDocumentDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDremioDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDremioDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareDynamoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareDynamoDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareExasolDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareExasolDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareFileDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareFileDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGitHubDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGitHubDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGoogleAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGoogleAnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGoogleBigQueryDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGoogleBigQueryDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGoogleDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGoogleDriveKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareGoogleSheetsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareGoogleSheetsDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareIdcKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareIdcKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareImpalaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareImpalaDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareJiraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareJiraDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMariaDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMariaDbDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMongoAtlasDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMongoAtlasDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMongoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMongoDbDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareMySqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareMySqlDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareOneDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareOneDriveKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareOpenSearchDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareOpenSearchDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareOracleDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareOracleDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSharePayPalDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSharePayPalDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSharePointKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetSharePointKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSharePostgreSqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSharePostgreSqlDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSharePrestoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSharePrestoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareQBusinessKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareQBusinessKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareRadiantDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareRadiantDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareRdsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareRdsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareRedshiftAutoDiscoveredDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareRedshiftAutoDiscoveredDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareRedshiftManualDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareRedshiftManualDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareS3AnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareS3AnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareS3DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareS3DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareS3KnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareS3KnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareS3TablesDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareS3TablesDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSalesforceDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSalesforceDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSapHanaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSapHanaDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareServiceNowDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareServiceNowDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSharePointKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSharePointKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSnowflakeDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSnowflakeDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSparkDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSparkDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSqlServerDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSqlServerDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareSquareDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareSquareDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareStarburstDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareStarburstDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareTeradataDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareTeradataDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareTimestreamDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareTimestreamDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareTrinoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareTrinoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareTwitterDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetShareTwitterDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetShareWebCrawlerKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetShareWebCrawlerKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSnowflakeDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSnowflakeDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSparkDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSparkDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSqlServerDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSqlServerDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetSquareDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetSquareDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetStarburstDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetStarburstDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetTeradataDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetTeradataDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetTimestreamDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetTimestreamDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetTrigger() {
	_jsii_.InvokeVoid(
		q,
		"resetTrigger",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetTrinoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetTrinoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetTwitterDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetTwitterDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateAdobeAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateAdobeAnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateAthenaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateAthenaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateAuroraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateAuroraDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateDatabricksDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateDatabricksDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateDb2DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateDb2DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateDenodoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateDenodoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateDocumentDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateDocumentDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateDremioDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateDremioDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateDynamoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateDynamoDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateExasolDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateExasolDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateFileDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateFileDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateGitHubDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateGitHubDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateGoogleAnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateGoogleAnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateGoogleBigQueryDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateGoogleBigQueryDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateGoogleSheetsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateGoogleSheetsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateImpalaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateImpalaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateJiraDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateJiraDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateMariaDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateMariaDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateMongoAtlasDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateMongoAtlasDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateMongoDbDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateMongoDbDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateMySqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateMySqlDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateOpenSearchDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateOpenSearchDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateOracleDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateOracleDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdatePayPalDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdatePayPalDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdatePostgreSqlDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdatePostgreSqlDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdatePrestoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdatePrestoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateRadiantDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateRadiantDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateRdsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateRdsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateRedshiftAutoDiscoveredDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateRedshiftAutoDiscoveredDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateRedshiftManualDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateRedshiftManualDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateS3AnalyticsDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateS3AnalyticsDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateS3DataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateS3DataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateS3TablesDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateS3TablesDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateSalesforceDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateSalesforceDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateSapHanaDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateSapHanaDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateServiceNowDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateServiceNowDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateSnowflakeDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateSnowflakeDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateSparkDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateSparkDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateSqlServerDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateSqlServerDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateSquareDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateSquareDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateStarburstDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateStarburstDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateTeradataDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateTeradataDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateTimestreamDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateTimestreamDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateTrinoDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateTrinoDataSource",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUpdateTwitterDataSource() {
	_jsii_.InvokeVoid(
		q,
		"resetUpdateTwitterDataSource",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseBedrockManagedKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseBedrockManagedKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseBoxKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseBoxKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseBrowserExtension() {
	_jsii_.InvokeVoid(
		q,
		"resetUseBrowserExtension",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseConfluenceKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseConfluenceKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseExcelAddInExtension() {
	_jsii_.InvokeVoid(
		q,
		"resetUseExcelAddInExtension",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseGoogleDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseGoogleDriveKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseIdcKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseIdcKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseOneDriveKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseOneDriveKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseOutlookAddInExtension() {
	_jsii_.InvokeVoid(
		q,
		"resetUseOutlookAddInExtension",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUsePowerpointAddInExtension() {
	_jsii_.InvokeVoid(
		q,
		"resetUsePowerpointAddInExtension",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseQBusinessKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseQBusinessKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseS3KnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseS3KnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseSharePointKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseSharePointKnowledgeBase",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseWebCrawlerKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetUseWebCrawlerKnowledgeBase",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetUseWordAddInExtension() {
	_jsii_.InvokeVoid(
		q,
		"resetUseWordAddInExtension",
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

func (q *jsiiProxy_QuicksightCustomPermissionsCapabilitiesOutputReference) ResetWebCrawlerKnowledgeBase() {
	_jsii_.InvokeVoid(
		q,
		"resetWebCrawlerKnowledgeBase",
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

