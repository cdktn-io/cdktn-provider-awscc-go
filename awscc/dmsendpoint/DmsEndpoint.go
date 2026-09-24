// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dmsendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dmsendpoint/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dms_endpoint awscc_dms_endpoint}.
type DmsEndpoint interface {
	cdktn.TerraformResource
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	CertificateArn() *string
	SetCertificateArn(val *string)
	CertificateArnInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DocDbSettings() DmsEndpointDocDbSettingsOutputReference
	DocDbSettingsInput() interface{}
	DynamoDbSettings() DmsEndpointDynamoDbSettingsOutputReference
	DynamoDbSettingsInput() interface{}
	ElasticsearchSettings() DmsEndpointElasticsearchSettingsOutputReference
	ElasticsearchSettingsInput() interface{}
	EndpointArn() *string
	EndpointIdentifier() *string
	SetEndpointIdentifier(val *string)
	EndpointIdentifierInput() *string
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	EngineName() *string
	SetEngineName(val *string)
	EngineNameInput() *string
	ExternalId() *string
	ExtraConnectionAttributes() *string
	SetExtraConnectionAttributes(val *string)
	ExtraConnectionAttributesInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcpMySqlSettings() DmsEndpointGcpMySqlSettingsOutputReference
	GcpMySqlSettingsInput() interface{}
	IbmDb2Settings() DmsEndpointIbmDb2SettingsOutputReference
	IbmDb2SettingsInput() interface{}
	Id() *string
	KafkaSettings() DmsEndpointKafkaSettingsOutputReference
	KafkaSettingsInput() interface{}
	KinesisSettings() DmsEndpointKinesisSettingsOutputReference
	KinesisSettingsInput() interface{}
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MicrosoftSqlServerSettings() DmsEndpointMicrosoftSqlServerSettingsOutputReference
	MicrosoftSqlServerSettingsInput() interface{}
	MongoDbSettings() DmsEndpointMongoDbSettingsOutputReference
	MongoDbSettingsInput() interface{}
	MySqlSettings() DmsEndpointMySqlSettingsOutputReference
	MySqlSettingsInput() interface{}
	NeptuneSettings() DmsEndpointNeptuneSettingsOutputReference
	NeptuneSettingsInput() interface{}
	// The tree node.
	Node() constructs.Node
	OracleSettings() DmsEndpointOracleSettingsOutputReference
	OracleSettingsInput() interface{}
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PostgreSqlSettings() DmsEndpointPostgreSqlSettingsOutputReference
	PostgreSqlSettingsInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RedisSettings() DmsEndpointRedisSettingsOutputReference
	RedisSettingsInput() interface{}
	RedshiftSettings() DmsEndpointRedshiftSettingsOutputReference
	RedshiftSettingsInput() interface{}
	ResourceIdentifier() *string
	SetResourceIdentifier(val *string)
	ResourceIdentifierInput() *string
	S3Settings() DmsEndpointS3SettingsOutputReference
	S3SettingsInput() interface{}
	ServerName() *string
	SetServerName(val *string)
	ServerNameInput() *string
	SslMode() *string
	SetSslMode(val *string)
	SslModeInput() *string
	SybaseSettings() DmsEndpointSybaseSettingsOutputReference
	SybaseSettingsInput() interface{}
	Tags() DmsEndpointTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDocDbSettings(value *DmsEndpointDocDbSettings)
	PutDynamoDbSettings(value *DmsEndpointDynamoDbSettings)
	PutElasticsearchSettings(value *DmsEndpointElasticsearchSettings)
	PutGcpMySqlSettings(value *DmsEndpointGcpMySqlSettings)
	PutIbmDb2Settings(value *DmsEndpointIbmDb2Settings)
	PutKafkaSettings(value *DmsEndpointKafkaSettings)
	PutKinesisSettings(value *DmsEndpointKinesisSettings)
	PutMicrosoftSqlServerSettings(value *DmsEndpointMicrosoftSqlServerSettings)
	PutMongoDbSettings(value *DmsEndpointMongoDbSettings)
	PutMySqlSettings(value *DmsEndpointMySqlSettings)
	PutNeptuneSettings(value *DmsEndpointNeptuneSettings)
	PutOracleSettings(value *DmsEndpointOracleSettings)
	PutPostgreSqlSettings(value *DmsEndpointPostgreSqlSettings)
	PutRedisSettings(value *DmsEndpointRedisSettings)
	PutRedshiftSettings(value *DmsEndpointRedshiftSettings)
	PutS3Settings(value *DmsEndpointS3Settings)
	PutSybaseSettings(value *DmsEndpointSybaseSettings)
	PutTags(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetCertificateArn()
	ResetDatabaseName()
	ResetDocDbSettings()
	ResetDynamoDbSettings()
	ResetElasticsearchSettings()
	ResetEndpointIdentifier()
	ResetExtraConnectionAttributes()
	ResetGcpMySqlSettings()
	ResetIbmDb2Settings()
	ResetKafkaSettings()
	ResetKinesisSettings()
	ResetKmsKeyId()
	ResetMicrosoftSqlServerSettings()
	ResetMongoDbSettings()
	ResetMySqlSettings()
	ResetNeptuneSettings()
	ResetOracleSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetPort()
	ResetPostgreSqlSettings()
	ResetRedisSettings()
	ResetRedshiftSettings()
	ResetResourceIdentifier()
	ResetS3Settings()
	ResetServerName()
	ResetSslMode()
	ResetSybaseSettings()
	ResetTags()
	ResetUsername()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for DmsEndpoint
type jsiiProxy_DmsEndpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DmsEndpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) CertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DocDbSettings() DmsEndpointDocDbSettingsOutputReference {
	var returns DmsEndpointDocDbSettingsOutputReference
	_jsii_.Get(
		j,
		"docDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DocDbSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"docDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DynamoDbSettings() DmsEndpointDynamoDbSettingsOutputReference {
	var returns DmsEndpointDynamoDbSettingsOutputReference
	_jsii_.Get(
		j,
		"dynamoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) DynamoDbSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamoDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ElasticsearchSettings() DmsEndpointElasticsearchSettingsOutputReference {
	var returns DmsEndpointElasticsearchSettingsOutputReference
	_jsii_.Get(
		j,
		"elasticsearchSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ElasticsearchSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EngineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) EngineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ExternalId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ExtraConnectionAttributes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ExtraConnectionAttributesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extraConnectionAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) GcpMySqlSettings() DmsEndpointGcpMySqlSettingsOutputReference {
	var returns DmsEndpointGcpMySqlSettingsOutputReference
	_jsii_.Get(
		j,
		"gcpMySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) GcpMySqlSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gcpMySqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) IbmDb2Settings() DmsEndpointIbmDb2SettingsOutputReference {
	var returns DmsEndpointIbmDb2SettingsOutputReference
	_jsii_.Get(
		j,
		"ibmDb2Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) IbmDb2SettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ibmDb2SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KafkaSettings() DmsEndpointKafkaSettingsOutputReference {
	var returns DmsEndpointKafkaSettingsOutputReference
	_jsii_.Get(
		j,
		"kafkaSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KafkaSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KinesisSettings() DmsEndpointKinesisSettingsOutputReference {
	var returns DmsEndpointKinesisSettingsOutputReference
	_jsii_.Get(
		j,
		"kinesisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KinesisSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MicrosoftSqlServerSettings() DmsEndpointMicrosoftSqlServerSettingsOutputReference {
	var returns DmsEndpointMicrosoftSqlServerSettingsOutputReference
	_jsii_.Get(
		j,
		"microsoftSqlServerSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MicrosoftSqlServerSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"microsoftSqlServerSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MongoDbSettings() DmsEndpointMongoDbSettingsOutputReference {
	var returns DmsEndpointMongoDbSettingsOutputReference
	_jsii_.Get(
		j,
		"mongoDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MongoDbSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongoDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MySqlSettings() DmsEndpointMySqlSettingsOutputReference {
	var returns DmsEndpointMySqlSettingsOutputReference
	_jsii_.Get(
		j,
		"mySqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) MySqlSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mySqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) NeptuneSettings() DmsEndpointNeptuneSettingsOutputReference {
	var returns DmsEndpointNeptuneSettingsOutputReference
	_jsii_.Get(
		j,
		"neptuneSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) NeptuneSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neptuneSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) OracleSettings() DmsEndpointOracleSettingsOutputReference {
	var returns DmsEndpointOracleSettingsOutputReference
	_jsii_.Get(
		j,
		"oracleSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) OracleSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PostgreSqlSettings() DmsEndpointPostgreSqlSettingsOutputReference {
	var returns DmsEndpointPostgreSqlSettingsOutputReference
	_jsii_.Get(
		j,
		"postgreSqlSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) PostgreSqlSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgreSqlSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedisSettings() DmsEndpointRedisSettingsOutputReference {
	var returns DmsEndpointRedisSettingsOutputReference
	_jsii_.Get(
		j,
		"redisSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedisSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedshiftSettings() DmsEndpointRedshiftSettingsOutputReference {
	var returns DmsEndpointRedshiftSettingsOutputReference
	_jsii_.Get(
		j,
		"redshiftSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) RedshiftSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ResourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ResourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) S3Settings() DmsEndpointS3SettingsOutputReference {
	var returns DmsEndpointS3SettingsOutputReference
	_jsii_.Get(
		j,
		"s3Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) S3SettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SslMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SslModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SybaseSettings() DmsEndpointSybaseSettingsOutputReference {
	var returns DmsEndpointSybaseSettingsOutputReference
	_jsii_.Get(
		j,
		"sybaseSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) SybaseSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sybaseSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Tags() DmsEndpointTagsList {
	var returns DmsEndpointTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsEndpoint) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dms_endpoint awscc_dms_endpoint} Resource.
func NewDmsEndpoint(scope constructs.Construct, id *string, config *DmsEndpointConfig) DmsEndpoint {
	_init_.Initialize()

	if err := validateNewDmsEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsEndpoint{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dms_endpoint awscc_dms_endpoint} Resource.
func NewDmsEndpoint_Override(d DmsEndpoint, scope constructs.Construct, id *string, config *DmsEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpoint",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetCertificateArn(val *string) {
	if err := j.validateSetCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateArn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetEndpointIdentifier(val *string) {
	if err := j.validateSetEndpointIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetEngineName(val *string) {
	if err := j.validateSetEngineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetExtraConnectionAttributes(val *string) {
	if err := j.validateSetExtraConnectionAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraConnectionAttributes",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetResourceIdentifier(val *string) {
	if err := j.validateSetResourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetSslMode(val *string) {
	if err := j.validateSetSslModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslMode",
		val,
	)
}

func (j *jsiiProxy_DmsEndpoint)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

// Generates CDKTN code for importing a DmsEndpoint resource upon running "cdktn plan <stack-name>".
func DmsEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDmsEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpoint",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DmsEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dmsEndpoint.DmsEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsEndpoint) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DmsEndpoint) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DmsEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DmsEndpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := d.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DmsEndpoint) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutDocDbSettings(value *DmsEndpointDocDbSettings) {
	if err := d.validatePutDocDbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDocDbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutDynamoDbSettings(value *DmsEndpointDynamoDbSettings) {
	if err := d.validatePutDynamoDbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDynamoDbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutElasticsearchSettings(value *DmsEndpointElasticsearchSettings) {
	if err := d.validatePutElasticsearchSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putElasticsearchSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutGcpMySqlSettings(value *DmsEndpointGcpMySqlSettings) {
	if err := d.validatePutGcpMySqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGcpMySqlSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutIbmDb2Settings(value *DmsEndpointIbmDb2Settings) {
	if err := d.validatePutIbmDb2SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putIbmDb2Settings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutKafkaSettings(value *DmsEndpointKafkaSettings) {
	if err := d.validatePutKafkaSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKafkaSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutKinesisSettings(value *DmsEndpointKinesisSettings) {
	if err := d.validatePutKinesisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKinesisSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutMicrosoftSqlServerSettings(value *DmsEndpointMicrosoftSqlServerSettings) {
	if err := d.validatePutMicrosoftSqlServerSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMicrosoftSqlServerSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutMongoDbSettings(value *DmsEndpointMongoDbSettings) {
	if err := d.validatePutMongoDbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMongoDbSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutMySqlSettings(value *DmsEndpointMySqlSettings) {
	if err := d.validatePutMySqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMySqlSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutNeptuneSettings(value *DmsEndpointNeptuneSettings) {
	if err := d.validatePutNeptuneSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNeptuneSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutOracleSettings(value *DmsEndpointOracleSettings) {
	if err := d.validatePutOracleSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracleSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutPostgreSqlSettings(value *DmsEndpointPostgreSqlSettings) {
	if err := d.validatePutPostgreSqlSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgreSqlSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutRedisSettings(value *DmsEndpointRedisSettings) {
	if err := d.validatePutRedisSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedisSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutRedshiftSettings(value *DmsEndpointRedshiftSettings) {
	if err := d.validatePutRedshiftSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedshiftSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutS3Settings(value *DmsEndpointS3Settings) {
	if err := d.validatePutS3SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3Settings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutSybaseSettings(value *DmsEndpointSybaseSettings) {
	if err := d.validatePutSybaseSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSybaseSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsEndpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetCertificateArn() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificateArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		d,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetDocDbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetDocDbSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetDynamoDbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamoDbSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetElasticsearchSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetElasticsearchSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetEndpointIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetEndpointIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetExtraConnectionAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetExtraConnectionAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetGcpMySqlSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetGcpMySqlSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetIbmDb2Settings() {
	_jsii_.InvokeVoid(
		d,
		"resetIbmDb2Settings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKafkaSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetKafkaSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKinesisSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetKinesisSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetMicrosoftSqlServerSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMicrosoftSqlServerSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetMongoDbSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMongoDbSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetMySqlSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetMySqlSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetNeptuneSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetNeptuneSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetOracleSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetOracleSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetPostgreSqlSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgreSqlSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetRedisSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetRedisSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetRedshiftSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshiftSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetResourceIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetResourceIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetS3Settings() {
	_jsii_.InvokeVoid(
		d,
		"resetS3Settings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetServerName() {
	_jsii_.InvokeVoid(
		d,
		"resetServerName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetSslMode() {
	_jsii_.InvokeVoid(
		d,
		"resetSslMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetSybaseSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSybaseSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

