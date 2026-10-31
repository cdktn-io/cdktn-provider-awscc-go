// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccquicksightcustompermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/v2/dataawsccquicksightcustompermissions/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference interface {
	cdktn.ComplexObject
	AccessAppsNativeDataStore() *string
	Action() *string
	AddOrRunAnomalyDetectionForAnalyses() *string
	AdobeAnalyticsDataSource() *string
	AmazonBedrockArsAction() *string
	AmazonBedrockFsAction() *string
	AmazonBedrockKrsAction() *string
	AmazonSThreeAction() *string
	Analysis() *string
	ApproveFlowShareRequests() *string
	Apps() *string
	AsanaAction() *string
	AthenaDataSource() *string
	AuroraDataSource() *string
	Automate() *string
	BambooHrAction() *string
	BedrockManagedKnowledgeBase() *string
	BoxAgentAction() *string
	BoxKnowledgeBase() *string
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
	ConfluenceKnowledgeBase() *string
	CreateAdobeAnalyticsDataSource() *string
	CreateAndUpdateAmazonBedrockArsAction() *string
	CreateAndUpdateAmazonBedrockFsAction() *string
	CreateAndUpdateAmazonBedrockKrsAction() *string
	CreateAndUpdateAmazonSThreeAction() *string
	CreateAndUpdateApps() *string
	CreateAndUpdateAsanaAction() *string
	CreateAndUpdateBambooHrAction() *string
	CreateAndUpdateBedrockManagedKnowledgeBase() *string
	CreateAndUpdateBoxAgentAction() *string
	CreateAndUpdateBoxKnowledgeBase() *string
	CreateAndUpdateCanvaAgentAction() *string
	CreateAndUpdateComprehendAction() *string
	CreateAndUpdateComprehendMedicalAction() *string
	CreateAndUpdateConfluenceAction() *string
	CreateAndUpdateConfluenceKnowledgeBase() *string
	CreateAndUpdateDashboardEmailReports() *string
	CreateAndUpdateDatasets() *string
	CreateAndUpdateDataSources() *string
	CreateAndUpdateFactSetAction() *string
	CreateAndUpdateGenericHttpAction() *string
	CreateAndUpdateGithubAction() *string
	CreateAndUpdateGoogleCalendarAction() *string
	CreateAndUpdateGoogleDriveKnowledgeBase() *string
	CreateAndUpdateHubspotAction() *string
	CreateAndUpdateHuggingFaceAction() *string
	CreateAndUpdateIdcKnowledgeBase() *string
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
	CreateAndUpdateOneDriveKnowledgeBase() *string
	CreateAndUpdateOpenApiAction() *string
	CreateAndUpdatePagerDutyAction() *string
	CreateAndUpdateQBusinessKnowledgeBase() *string
	CreateAndUpdateS3KnowledgeBase() *string
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
	CreateAndUpdateSharePointKnowledgeBase() *string
	CreateAndUpdateSlackAction() *string
	CreateAndUpdateSmartsheetAction() *string
	CreateAndUpdateTextractAction() *string
	CreateAndUpdateThemes() *string
	CreateAndUpdateThresholdAlerts() *string
	CreateAndUpdateWebCrawlerKnowledgeBase() *string
	CreateAndUpdateZendeskAction() *string
	CreateAthenaDataSource() *string
	CreateAuroraDataSource() *string
	CreateChatAgents() *string
	CreateDashboardExecutiveSummaryWithQ() *string
	CreateDatabricksDataSource() *string
	CreateDb2DataSource() *string
	CreateDenodoDataSource() *string
	CreateDocumentDbDataSource() *string
	CreateDremioDataSource() *string
	CreateDynamoDbDataSource() *string
	CreateExasolDataSource() *string
	CreateFileDataSource() *string
	CreateGitHubDataSource() *string
	CreateGoogleAnalyticsDataSource() *string
	CreateGoogleBigQueryDataSource() *string
	CreateGoogleSheetsDataSource() *string
	CreateImpalaDataSource() *string
	CreateJiraDataSource() *string
	CreateMariaDbDataSource() *string
	CreateMongoAtlasDataSource() *string
	CreateMongoDbDataSource() *string
	CreateMySqlDataSource() *string
	CreateOpenSearchDataSource() *string
	CreateOracleDataSource() *string
	CreatePayPalDataSource() *string
	CreatePostgreSqlDataSource() *string
	CreatePrestoDataSource() *string
	CreateRadiantDataSource() *string
	CreateRdsDataSource() *string
	CreateRedshiftAutoDiscoveredDataSource() *string
	CreateRedshiftManualDataSource() *string
	CreateS3AnalyticsDataSource() *string
	CreateS3DataSource() *string
	CreateS3TablesDataSource() *string
	CreateSalesforceDataSource() *string
	CreateSapHanaDataSource() *string
	CreateServiceNowDataSource() *string
	CreateSharedFolders() *string
	CreateSnowflakeDataSource() *string
	CreateSpaces() *string
	CreateSparkDataSource() *string
	CreateSpiceDataset() *string
	CreateSqlServerDataSource() *string
	CreateSquareDataSource() *string
	CreateStarburstDataSource() *string
	CreateTeradataDataSource() *string
	CreateTimestreamDataSource() *string
	CreateTrinoDataSource() *string
	CreateTwitterDataSource() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Dashboard() *string
	DatabricksDataSource() *string
	Db2DataSource() *string
	DenodoDataSource() *string
	DocumentDbDataSource() *string
	DremioDataSource() *string
	DynamoDbDataSource() *string
	EditVisualWithQ() *string
	ExasolDataSource() *string
	ExportToCsv() *string
	ExportToCsvInScheduledReports() *string
	ExportToExcel() *string
	ExportToExcelInScheduledReports() *string
	ExportToPdf() *string
	ExportToPdfInScheduledReports() *string
	Extension() *string
	FactSetAction() *string
	FileDataSource() *string
	Flow() *string
	// Experimental.
	Fqn() *string
	GenericHttpAction() *string
	GithubAction() *string
	GitHubDataSource() *string
	GoogleAnalyticsDataSource() *string
	GoogleBigQueryDataSource() *string
	GoogleCalendarAction() *string
	GoogleDriveKnowledgeBase() *string
	GoogleSheetsDataSource() *string
	HubspotAction() *string
	HuggingFaceAction() *string
	IdcKnowledgeBase() *string
	ImpalaDataSource() *string
	InboundEmailTrigger() *string
	IncludeContentInScheduledReportsEmail() *string
	IntercomAction() *string
	InternalValue() *DataAwsccQuicksightCustomPermissionsCapabilities
	SetInternalValue(val *DataAwsccQuicksightCustomPermissionsCapabilities)
	InvokeAppsAiInference() *string
	JiraAction() *string
	JiraDataSource() *string
	KnowledgeBase() *string
	LinearAction() *string
	ManageSharedFolders() *string
	MariaDbDataSource() *string
	McpAction() *string
	MondayAction() *string
	MongoAtlasDataSource() *string
	MongoDbDataSource() *string
	MsExchangeAction() *string
	MsTeamsAction() *string
	MySqlDataSource() *string
	NewRelicAction() *string
	NotionAction() *string
	OneDriveAction() *string
	OneDriveKnowledgeBase() *string
	OpenApiAction() *string
	OpenSearchDataSource() *string
	OracleDataSource() *string
	PagerDutyAction() *string
	PayPalDataSource() *string
	PerformFlowUiTask() *string
	PostgreSqlDataSource() *string
	PrestoDataSource() *string
	PrintReports() *string
	PublishWithoutApproval() *string
	QBusinessKnowledgeBase() *string
	QuickEventTrigger() *string
	RadiantDataSource() *string
	RdsDataSource() *string
	RedshiftAutoDiscoveredDataSource() *string
	RedshiftManualDataSource() *string
	RenameSharedFolders() *string
	Research() *string
	S3AnalyticsDataSource() *string
	S3DataSource() *string
	S3KnowledgeBase() *string
	S3TablesDataSource() *string
	SalesforceAction() *string
	SalesforceDataSource() *string
	SandPGlobalEnergyAction() *string
	SandPgmiAction() *string
	SapBillOfMaterialAction() *string
	SapBusinessPartnerAction() *string
	SapHanaDataSource() *string
	SapMaterialStockAction() *string
	SapPhysicalInventoryAction() *string
	SapProductMasterDataAction() *string
	ScheduleTrigger() *string
	ServiceNowAction() *string
	ServiceNowDataSource() *string
	ShareAdobeAnalyticsDataSource() *string
	ShareAmazonBedrockArsAction() *string
	ShareAmazonBedrockFsAction() *string
	ShareAmazonBedrockKrsAction() *string
	ShareAmazonSThreeAction() *string
	ShareAnalyses() *string
	ShareApps() *string
	ShareAsanaAction() *string
	ShareAthenaDataSource() *string
	ShareAuroraDataSource() *string
	ShareBambooHrAction() *string
	ShareBedrockManagedKnowledgeBase() *string
	ShareBoxAgentAction() *string
	ShareBoxKnowledgeBase() *string
	ShareCanvaAgentAction() *string
	ShareChatAgents() *string
	ShareComprehendAction() *string
	ShareComprehendMedicalAction() *string
	ShareConfluenceAction() *string
	ShareConfluenceKnowledgeBase() *string
	ShareDashboards() *string
	ShareDatabricksDataSource() *string
	ShareDatasets() *string
	ShareDataSources() *string
	ShareDb2DataSource() *string
	ShareDenodoDataSource() *string
	ShareDocumentDbDataSource() *string
	ShareDremioDataSource() *string
	ShareDynamoDbDataSource() *string
	ShareExasolDataSource() *string
	ShareFactSetAction() *string
	ShareFileDataSource() *string
	ShareGenericHttpAction() *string
	ShareGithubAction() *string
	ShareGitHubDataSource() *string
	ShareGoogleAnalyticsDataSource() *string
	ShareGoogleBigQueryDataSource() *string
	ShareGoogleCalendarAction() *string
	ShareGoogleDriveKnowledgeBase() *string
	ShareGoogleSheetsDataSource() *string
	ShareHubspotAction() *string
	ShareHuggingFaceAction() *string
	ShareIdcKnowledgeBase() *string
	ShareImpalaDataSource() *string
	ShareIntercomAction() *string
	ShareJiraAction() *string
	ShareJiraDataSource() *string
	ShareKnowledgeBases() *string
	ShareLinearAction() *string
	ShareMariaDbDataSource() *string
	ShareMcpAction() *string
	ShareMondayAction() *string
	ShareMongoAtlasDataSource() *string
	ShareMongoDbDataSource() *string
	ShareMsExchangeAction() *string
	ShareMsTeamsAction() *string
	ShareMySqlDataSource() *string
	ShareNewRelicAction() *string
	ShareNotionAction() *string
	ShareOneDriveAction() *string
	ShareOneDriveKnowledgeBase() *string
	ShareOpenApiAction() *string
	ShareOpenSearchDataSource() *string
	ShareOracleDataSource() *string
	SharePagerDutyAction() *string
	SharePayPalDataSource() *string
	SharePointAction() *string
	SharePointKnowledgeBase() *string
	SharePostgreSqlDataSource() *string
	SharePrestoDataSource() *string
	ShareQBusinessKnowledgeBase() *string
	ShareRadiantDataSource() *string
	ShareRdsDataSource() *string
	ShareRedshiftAutoDiscoveredDataSource() *string
	ShareRedshiftManualDataSource() *string
	ShareS3AnalyticsDataSource() *string
	ShareS3DataSource() *string
	ShareS3KnowledgeBase() *string
	ShareS3TablesDataSource() *string
	ShareSalesforceAction() *string
	ShareSalesforceDataSource() *string
	ShareSandPGlobalEnergyAction() *string
	ShareSandPgmiAction() *string
	ShareSapBillOfMaterialAction() *string
	ShareSapBusinessPartnerAction() *string
	ShareSapHanaDataSource() *string
	ShareSapMaterialStockAction() *string
	ShareSapPhysicalInventoryAction() *string
	ShareSapProductMasterDataAction() *string
	ShareServiceNowAction() *string
	ShareServiceNowDataSource() *string
	ShareSharePointAction() *string
	ShareSharePointKnowledgeBase() *string
	ShareSlackAction() *string
	ShareSmartsheetAction() *string
	ShareSnowflakeDataSource() *string
	ShareSpaces() *string
	ShareSparkDataSource() *string
	ShareSqlServerDataSource() *string
	ShareSquareDataSource() *string
	ShareStarburstDataSource() *string
	ShareTeradataDataSource() *string
	ShareTextractAction() *string
	ShareTimestreamDataSource() *string
	ShareTrinoDataSource() *string
	ShareTwitterDataSource() *string
	ShareWebCrawlerKnowledgeBase() *string
	ShareZendeskAction() *string
	SlackAction() *string
	SmartsheetAction() *string
	SnowflakeDataSource() *string
	Space() *string
	SparkDataSource() *string
	SqlServerDataSource() *string
	SquareDataSource() *string
	StarburstDataSource() *string
	SubscribeDashboardEmailReports() *string
	TeradataDataSource() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TextractAction() *string
	TimestreamDataSource() *string
	Topic() *string
	Trigger() *string
	TrinoDataSource() *string
	TwitterDataSource() *string
	UpdateAdobeAnalyticsDataSource() *string
	UpdateAthenaDataSource() *string
	UpdateAuroraDataSource() *string
	UpdateDatabricksDataSource() *string
	UpdateDb2DataSource() *string
	UpdateDenodoDataSource() *string
	UpdateDocumentDbDataSource() *string
	UpdateDremioDataSource() *string
	UpdateDynamoDbDataSource() *string
	UpdateExasolDataSource() *string
	UpdateFileDataSource() *string
	UpdateGitHubDataSource() *string
	UpdateGoogleAnalyticsDataSource() *string
	UpdateGoogleBigQueryDataSource() *string
	UpdateGoogleSheetsDataSource() *string
	UpdateImpalaDataSource() *string
	UpdateJiraDataSource() *string
	UpdateMariaDbDataSource() *string
	UpdateMongoAtlasDataSource() *string
	UpdateMongoDbDataSource() *string
	UpdateMySqlDataSource() *string
	UpdateOpenSearchDataSource() *string
	UpdateOracleDataSource() *string
	UpdatePayPalDataSource() *string
	UpdatePostgreSqlDataSource() *string
	UpdatePrestoDataSource() *string
	UpdateRadiantDataSource() *string
	UpdateRdsDataSource() *string
	UpdateRedshiftAutoDiscoveredDataSource() *string
	UpdateRedshiftManualDataSource() *string
	UpdateS3AnalyticsDataSource() *string
	UpdateS3DataSource() *string
	UpdateS3TablesDataSource() *string
	UpdateSalesforceDataSource() *string
	UpdateSapHanaDataSource() *string
	UpdateServiceNowDataSource() *string
	UpdateSnowflakeDataSource() *string
	UpdateSparkDataSource() *string
	UpdateSqlServerDataSource() *string
	UpdateSquareDataSource() *string
	UpdateStarburstDataSource() *string
	UpdateTeradataDataSource() *string
	UpdateTimestreamDataSource() *string
	UpdateTrinoDataSource() *string
	UpdateTwitterDataSource() *string
	UseAgentWebSearch() *string
	UseAmazonBedrockArsAction() *string
	UseAmazonBedrockFsAction() *string
	UseAmazonBedrockKrsAction() *string
	UseAmazonSThreeAction() *string
	UseAsanaAction() *string
	UseBambooHrAction() *string
	UseBedrockManagedKnowledgeBase() *string
	UseBedrockModels() *string
	UseBoxAgentAction() *string
	UseBoxKnowledgeBase() *string
	UseBrowserExtension() *string
	UseCanvaAgentAction() *string
	UseComprehendAction() *string
	UseComprehendMedicalAction() *string
	UseConfluenceAction() *string
	UseConfluenceKnowledgeBase() *string
	UseExcelAddInExtension() *string
	UseFactSetAction() *string
	UseGenericHttpAction() *string
	UseGithubAction() *string
	UseGoogleCalendarAction() *string
	UseGoogleDriveKnowledgeBase() *string
	UseHubspotAction() *string
	UseHuggingFaceAction() *string
	UseIdcKnowledgeBase() *string
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
	UseOneDriveKnowledgeBase() *string
	UseOpenApiAction() *string
	UseOutlookAddInExtension() *string
	UsePagerDutyAction() *string
	UsePowerpointAddInExtension() *string
	UseQBusinessKnowledgeBase() *string
	UseS3KnowledgeBase() *string
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
	UseSharePointKnowledgeBase() *string
	UseSlackAction() *string
	UseSmartsheetAction() *string
	UseTextractAction() *string
	UseWebCrawlerKnowledgeBase() *string
	UseWordAddInExtension() *string
	UseZendeskAction() *string
	ViewAccountSpiceCapacity() *string
	WebCrawlerKnowledgeBase() *string
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adobeAnalyticsDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"athenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auroraDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) BedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) BoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boxKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confluenceKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAdobeAnalyticsDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateBoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateBoxKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateConfluenceKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateGoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateGoogleDriveKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateIdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateIdcKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateOneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateOneDriveKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateQBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateQBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateS3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateS3KnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateSharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateSharePointKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateWebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateWebCrawlerKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAthenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAuroraDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateDatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDatabricksDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateDb2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDb2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateDenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDenodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateDocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDocumentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateDremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateDynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createDynamoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createExasolDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateFileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createFileDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateGitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleBigQueryDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateGoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createGoogleSheetsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createImpalaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateJiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createJiraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateMariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMariaDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateMongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateMongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMongoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateMySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createMySqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateOpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOpenSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateOracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOracleDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreatePayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPayPalDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreatePostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPostgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreatePrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createPrestoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateRadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRadiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateRdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateRedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRedshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateRedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createRedshiftManualDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateS3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createS3TablesDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSalesforceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSapHanaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createServiceNowDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSnowflakeDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSparkDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSquareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateStarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createStarburstDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateTeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTeradataDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateTimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTimestreamDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateTrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTrinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateTwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTwitterDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) DatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databricksDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Db2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"db2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) DenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"denodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) DocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) DremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) DynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dynamoDbDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exasolDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) FileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleBigQueryDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"googleSheetsDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) IdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idcKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"impalaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InboundEmailTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inboundEmailTrigger",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) JiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jiraDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) MariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mariaDbDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) MongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) MongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mongoDbDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) MySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mySqlDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) OneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oneDriveKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) OpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) OracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oracleDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payPalDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prestoDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) QBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) QuickEventTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quickEventTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) RadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"radiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) RdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) RedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) RedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redshiftManualDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) S3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) S3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) S3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) S3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3TablesDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salesforceDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sapHanaDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ScheduleTrigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleTrigger",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceNowDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAdobeAnalyticsDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAthenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAuroraDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareBedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareBoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareBoxKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareConfluenceKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatabricksDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDb2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDb2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDenodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDocumentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDynamoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareExasolDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareFileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareFileDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleBigQueryDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleDriveKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareGoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareGoogleSheetsDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareIdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareIdcKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareImpalaDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareJiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareJiraDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMariaDbDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMongoDbDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareMySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareMySqlDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareOneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOneDriveKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareOpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOpenSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareOracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareOracleDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SharePayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePayPalDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePointKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SharePostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePostgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SharePrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sharePrestoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareQBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareQBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareRadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRadiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareRdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareRedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRedshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareRedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareRedshiftManualDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3KnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareS3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareS3TablesDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSalesforceDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSapHanaDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareServiceNowDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSharePointKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSnowflakeDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSparkDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareSquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareSquareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareStarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareStarburstDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareTeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTeradataDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareTimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTimestreamDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareTrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTrinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareTwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareTwitterDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareWebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareWebCrawlerKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snowflakeDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sparkDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) StarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"starburstDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"teradataDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestreamDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Trigger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twitterDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAdobeAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAdobeAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAthenaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAthenaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateAuroraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateAuroraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDatabricksDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDatabricksDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDb2DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDb2DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDenodoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDenodoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDocumentDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDocumentDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDremioDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDremioDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateDynamoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateDynamoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateExasolDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateExasolDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateFileDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateFileDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGitHubDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGitHubDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleAnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleAnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleBigQueryDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleBigQueryDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateGoogleSheetsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateGoogleSheetsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateImpalaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateImpalaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateJiraDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateJiraDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMariaDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMariaDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMongoAtlasDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMongoAtlasDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMongoDbDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMongoDbDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateMySqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateMySqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateOpenSearchDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateOpenSearchDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateOracleDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateOracleDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePayPalDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePayPalDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePostgreSqlDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePostgreSqlDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdatePrestoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatePrestoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRadiantDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRadiantDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRdsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRdsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRedshiftAutoDiscoveredDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRedshiftAutoDiscoveredDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateRedshiftManualDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateRedshiftManualDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3AnalyticsDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3AnalyticsDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3DataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3DataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateS3TablesDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateS3TablesDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSalesforceDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSalesforceDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSapHanaDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSapHanaDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateServiceNowDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateServiceNowDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSnowflakeDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSnowflakeDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSparkDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSparkDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSqlServerDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSqlServerDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateSquareDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateSquareDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateStarburstDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateStarburstDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTeradataDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTeradataDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTimestreamDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTimestreamDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTrinoDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTrinoDataSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UpdateTwitterDataSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTwitterDataSource",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseBedrockManagedKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBedrockManagedKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseBoxKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBoxKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseBrowserExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useBrowserExtension",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseConfluenceKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useConfluenceKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseExcelAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useExcelAddInExtension",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseGoogleDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useGoogleDriveKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseIdcKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useIdcKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseOneDriveKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOneDriveKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseOutlookAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useOutlookAddInExtension",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UsePowerpointAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usePowerpointAddInExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseQBusinessKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useQBusinessKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseS3KnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useS3KnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseSharePointKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useSharePointKnowledgeBase",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseWebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useWebCrawlerKnowledgeBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) UseWordAddInExtension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"useWordAddInExtension",
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

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) WebCrawlerKnowledgeBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webCrawlerKnowledgeBase",
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

