// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisfirehose_delivery_stream awscc_kinesisfirehose_delivery_stream}.
type KinesisfirehoseDeliveryStream interface {
	cdktn.TerraformResource
	AmazonOpenSearchServerlessDestinationConfiguration() KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference
	AmazonOpenSearchServerlessDestinationConfigurationInput() interface{}
	AmazonopensearchserviceDestinationConfiguration() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference
	AmazonopensearchserviceDestinationConfigurationInput() interface{}
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	DatabaseSourceConfiguration() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference
	DatabaseSourceConfigurationInput() interface{}
	DeliveryStreamEncryptionConfigurationInput() KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference
	DeliveryStreamEncryptionConfigurationInputInput() interface{}
	DeliveryStreamName() *string
	SetDeliveryStreamName(val *string)
	DeliveryStreamNameInput() *string
	DeliveryStreamType() *string
	SetDeliveryStreamType(val *string)
	DeliveryStreamTypeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DirectPutSourceConfiguration() KinesisfirehoseDeliveryStreamDirectPutSourceConfigurationOutputReference
	DirectPutSourceConfigurationInput() interface{}
	ElasticsearchDestinationConfiguration() KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference
	ElasticsearchDestinationConfigurationInput() interface{}
	ExtendedS3DestinationConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference
	ExtendedS3DestinationConfigurationInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpEndpointDestinationConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference
	HttpEndpointDestinationConfigurationInput() interface{}
	IcebergDestinationConfiguration() KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationOutputReference
	IcebergDestinationConfigurationInput() interface{}
	Id() *string
	KinesisStreamSourceConfiguration() KinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference
	KinesisStreamSourceConfigurationInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MskSourceConfiguration() KinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference
	MskSourceConfigurationInput() interface{}
	// The tree node.
	Node() constructs.Node
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
	RedshiftDestinationConfiguration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference
	RedshiftDestinationConfigurationInput() interface{}
	S3DestinationConfiguration() KinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference
	S3DestinationConfigurationInput() interface{}
	SnowflakeDestinationConfiguration() KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference
	SnowflakeDestinationConfigurationInput() interface{}
	SplunkDestinationConfiguration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference
	SplunkDestinationConfigurationInput() interface{}
	Tags() KinesisfirehoseDeliveryStreamTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAmazonOpenSearchServerlessDestinationConfiguration(value *KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration)
	PutAmazonopensearchserviceDestinationConfiguration(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfiguration)
	PutDatabaseSourceConfiguration(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfiguration)
	PutDeliveryStreamEncryptionConfigurationInput(value *KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInput)
	PutDirectPutSourceConfiguration(value *KinesisfirehoseDeliveryStreamDirectPutSourceConfiguration)
	PutElasticsearchDestinationConfiguration(value *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfiguration)
	PutExtendedS3DestinationConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfiguration)
	PutHttpEndpointDestinationConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfiguration)
	PutIcebergDestinationConfiguration(value *KinesisfirehoseDeliveryStreamIcebergDestinationConfiguration)
	PutKinesisStreamSourceConfiguration(value *KinesisfirehoseDeliveryStreamKinesisStreamSourceConfiguration)
	PutMskSourceConfiguration(value *KinesisfirehoseDeliveryStreamMskSourceConfiguration)
	PutRedshiftDestinationConfiguration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfiguration)
	PutS3DestinationConfiguration(value *KinesisfirehoseDeliveryStreamS3DestinationConfiguration)
	PutSnowflakeDestinationConfiguration(value *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfiguration)
	PutSplunkDestinationConfiguration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfiguration)
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
	ResetAmazonOpenSearchServerlessDestinationConfiguration()
	ResetAmazonopensearchserviceDestinationConfiguration()
	ResetDatabaseSourceConfiguration()
	ResetDeliveryStreamEncryptionConfigurationInput()
	ResetDeliveryStreamName()
	ResetDeliveryStreamType()
	ResetDirectPutSourceConfiguration()
	ResetElasticsearchDestinationConfiguration()
	ResetExtendedS3DestinationConfiguration()
	ResetHttpEndpointDestinationConfiguration()
	ResetIcebergDestinationConfiguration()
	ResetKinesisStreamSourceConfiguration()
	ResetMskSourceConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRedshiftDestinationConfiguration()
	ResetS3DestinationConfiguration()
	ResetSnowflakeDestinationConfiguration()
	ResetSplunkDestinationConfiguration()
	ResetTags()
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

// The jsii proxy struct for KinesisfirehoseDeliveryStream
type jsiiProxy_KinesisfirehoseDeliveryStream struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) AmazonOpenSearchServerlessDestinationConfiguration() KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"amazonOpenSearchServerlessDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) AmazonOpenSearchServerlessDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonOpenSearchServerlessDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) AmazonopensearchserviceDestinationConfiguration() KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"amazonopensearchserviceDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) AmazonopensearchserviceDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"amazonopensearchserviceDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DatabaseSourceConfiguration() KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamDatabaseSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"databaseSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DatabaseSourceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DeliveryStreamEncryptionConfigurationInput() KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference {
	var returns KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInputOutputReference
	_jsii_.Get(
		j,
		"deliveryStreamEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DeliveryStreamEncryptionConfigurationInputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deliveryStreamEncryptionConfigurationInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DeliveryStreamName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DeliveryStreamNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DeliveryStreamType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DeliveryStreamTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryStreamTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DirectPutSourceConfiguration() KinesisfirehoseDeliveryStreamDirectPutSourceConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamDirectPutSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"directPutSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) DirectPutSourceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"directPutSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) ElasticsearchDestinationConfiguration() KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamElasticsearchDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"elasticsearchDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) ElasticsearchDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) ExtendedS3DestinationConfiguration() KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamExtendedS3DestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"extendedS3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) ExtendedS3DestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extendedS3DestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) HttpEndpointDestinationConfiguration() KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"httpEndpointDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) HttpEndpointDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpEndpointDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) IcebergDestinationConfiguration() KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamIcebergDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"icebergDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) IcebergDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"icebergDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) KinesisStreamSourceConfiguration() KinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamKinesisStreamSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) KinesisStreamSourceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) MskSourceConfiguration() KinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamMskSourceConfigurationOutputReference
	_jsii_.Get(
		j,
		"mskSourceConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) MskSourceConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mskSourceConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) RedshiftDestinationConfiguration() KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamRedshiftDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"redshiftDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) RedshiftDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) S3DestinationConfiguration() KinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamS3DestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"s3DestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) S3DestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3DestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) SnowflakeDestinationConfiguration() KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamSnowflakeDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"snowflakeDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) SnowflakeDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snowflakeDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) SplunkDestinationConfiguration() KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference {
	var returns KinesisfirehoseDeliveryStreamSplunkDestinationConfigurationOutputReference
	_jsii_.Get(
		j,
		"splunkDestinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) SplunkDestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"splunkDestinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) Tags() KinesisfirehoseDeliveryStreamTagsList {
	var returns KinesisfirehoseDeliveryStreamTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisfirehose_delivery_stream awscc_kinesisfirehose_delivery_stream} Resource.
func NewKinesisfirehoseDeliveryStream(scope constructs.Construct, id *string, config *KinesisfirehoseDeliveryStreamConfig) KinesisfirehoseDeliveryStream {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStream{}

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisfirehose_delivery_stream awscc_kinesisfirehose_delivery_stream} Resource.
func NewKinesisfirehoseDeliveryStream_Override(k KinesisfirehoseDeliveryStream, scope constructs.Construct, id *string, config *KinesisfirehoseDeliveryStreamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetDeliveryStreamName(val *string) {
	if err := j.validateSetDeliveryStreamNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryStreamName",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetDeliveryStreamType(val *string) {
	if err := j.validateSetDeliveryStreamTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryStreamType",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStream)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a KinesisfirehoseDeliveryStream resource upon running "cdktn plan <stack-name>".
func KinesisfirehoseDeliveryStream_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateKinesisfirehoseDeliveryStream_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
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
func KinesisfirehoseDeliveryStream_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisfirehoseDeliveryStream_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KinesisfirehoseDeliveryStream_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisfirehoseDeliveryStream_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KinesisfirehoseDeliveryStream_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKinesisfirehoseDeliveryStream_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KinesisfirehoseDeliveryStream_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStream",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := k.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutAmazonOpenSearchServerlessDestinationConfiguration(value *KinesisfirehoseDeliveryStreamAmazonOpenSearchServerlessDestinationConfiguration) {
	if err := k.validatePutAmazonOpenSearchServerlessDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAmazonOpenSearchServerlessDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutAmazonopensearchserviceDestinationConfiguration(value *KinesisfirehoseDeliveryStreamAmazonopensearchserviceDestinationConfiguration) {
	if err := k.validatePutAmazonopensearchserviceDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putAmazonopensearchserviceDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutDatabaseSourceConfiguration(value *KinesisfirehoseDeliveryStreamDatabaseSourceConfiguration) {
	if err := k.validatePutDatabaseSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDatabaseSourceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutDeliveryStreamEncryptionConfigurationInput(value *KinesisfirehoseDeliveryStreamDeliveryStreamEncryptionConfigurationInput) {
	if err := k.validatePutDeliveryStreamEncryptionConfigurationInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDeliveryStreamEncryptionConfigurationInput",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutDirectPutSourceConfiguration(value *KinesisfirehoseDeliveryStreamDirectPutSourceConfiguration) {
	if err := k.validatePutDirectPutSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putDirectPutSourceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutElasticsearchDestinationConfiguration(value *KinesisfirehoseDeliveryStreamElasticsearchDestinationConfiguration) {
	if err := k.validatePutElasticsearchDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putElasticsearchDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutExtendedS3DestinationConfiguration(value *KinesisfirehoseDeliveryStreamExtendedS3DestinationConfiguration) {
	if err := k.validatePutExtendedS3DestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putExtendedS3DestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutHttpEndpointDestinationConfiguration(value *KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfiguration) {
	if err := k.validatePutHttpEndpointDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putHttpEndpointDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutIcebergDestinationConfiguration(value *KinesisfirehoseDeliveryStreamIcebergDestinationConfiguration) {
	if err := k.validatePutIcebergDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putIcebergDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutKinesisStreamSourceConfiguration(value *KinesisfirehoseDeliveryStreamKinesisStreamSourceConfiguration) {
	if err := k.validatePutKinesisStreamSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putKinesisStreamSourceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutMskSourceConfiguration(value *KinesisfirehoseDeliveryStreamMskSourceConfiguration) {
	if err := k.validatePutMskSourceConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putMskSourceConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutRedshiftDestinationConfiguration(value *KinesisfirehoseDeliveryStreamRedshiftDestinationConfiguration) {
	if err := k.validatePutRedshiftDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putRedshiftDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutS3DestinationConfiguration(value *KinesisfirehoseDeliveryStreamS3DestinationConfiguration) {
	if err := k.validatePutS3DestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putS3DestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutSnowflakeDestinationConfiguration(value *KinesisfirehoseDeliveryStreamSnowflakeDestinationConfiguration) {
	if err := k.validatePutSnowflakeDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSnowflakeDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutSplunkDestinationConfiguration(value *KinesisfirehoseDeliveryStreamSplunkDestinationConfiguration) {
	if err := k.validatePutSplunkDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSplunkDestinationConfiguration",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) PutTags(value interface{}) {
	if err := k.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTags",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := k.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetAmazonOpenSearchServerlessDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetAmazonOpenSearchServerlessDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetAmazonopensearchserviceDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetAmazonopensearchserviceDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetDatabaseSourceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDatabaseSourceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetDeliveryStreamEncryptionConfigurationInput() {
	_jsii_.InvokeVoid(
		k,
		"resetDeliveryStreamEncryptionConfigurationInput",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetDeliveryStreamName() {
	_jsii_.InvokeVoid(
		k,
		"resetDeliveryStreamName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetDeliveryStreamType() {
	_jsii_.InvokeVoid(
		k,
		"resetDeliveryStreamType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetDirectPutSourceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetDirectPutSourceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetElasticsearchDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetElasticsearchDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetExtendedS3DestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetExtendedS3DestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetHttpEndpointDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpEndpointDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetIcebergDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetIcebergDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetKinesisStreamSourceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetKinesisStreamSourceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetMskSourceConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetMskSourceConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetRedshiftDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetRedshiftDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetS3DestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetS3DestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetSnowflakeDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSnowflakeDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetSplunkDestinationConfiguration() {
	_jsii_.InvokeVoid(
		k,
		"resetSplunkDestinationConfiguration",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStream) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		k,
		"with",
		args,
		&returns,
	)

	return returns
}

