// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dynamodbglobaltable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table awscc_dynamodb_global_table}.
type DynamodbGlobalTable interface {
	cdktn.TerraformResource
	Arn() *string
	AttributeDefinitions() DynamodbGlobalTableAttributeDefinitionsList
	AttributeDefinitionsInput() interface{}
	BillingMode() *string
	SetBillingMode(val *string)
	BillingModeInput() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalSecondaryIndexes() DynamodbGlobalTableGlobalSecondaryIndexesList
	GlobalSecondaryIndexesInput() interface{}
	GlobalTableSourceArn() *string
	SetGlobalTableSourceArn(val *string)
	GlobalTableSourceArnInput() *string
	GlobalTableWitnesses() DynamodbGlobalTableGlobalTableWitnessesList
	GlobalTableWitnessesInput() interface{}
	Id() *string
	KeySchema() DynamodbGlobalTableKeySchemaList
	KeySchemaInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocalSecondaryIndexes() DynamodbGlobalTableLocalSecondaryIndexesList
	LocalSecondaryIndexesInput() interface{}
	MultiRegionConsistency() *string
	SetMultiRegionConsistency(val *string)
	MultiRegionConsistencyInput() *string
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
	ReadOnDemandThroughputSettings() DynamodbGlobalTableReadOnDemandThroughputSettingsOutputReference
	ReadOnDemandThroughputSettingsInput() interface{}
	ReadProvisionedThroughputSettings() DynamodbGlobalTableReadProvisionedThroughputSettingsOutputReference
	ReadProvisionedThroughputSettingsInput() interface{}
	Replicas() DynamodbGlobalTableReplicasList
	ReplicasInput() interface{}
	SseSpecification() DynamodbGlobalTableSseSpecificationOutputReference
	SseSpecificationInput() interface{}
	StreamArn() *string
	StreamSpecification() DynamodbGlobalTableStreamSpecificationOutputReference
	StreamSpecificationInput() interface{}
	TableId() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeToLiveSpecification() DynamodbGlobalTableTimeToLiveSpecificationOutputReference
	TimeToLiveSpecificationInput() interface{}
	VectorIndexes() DynamodbGlobalTableVectorIndexesList
	VectorIndexesInput() interface{}
	WarmThroughput() DynamodbGlobalTableWarmThroughputOutputReference
	WarmThroughputInput() interface{}
	WriteOnDemandThroughputSettings() DynamodbGlobalTableWriteOnDemandThroughputSettingsOutputReference
	WriteOnDemandThroughputSettingsInput() interface{}
	WriteProvisionedThroughputSettings() DynamodbGlobalTableWriteProvisionedThroughputSettingsOutputReference
	WriteProvisionedThroughputSettingsInput() interface{}
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
	PutAttributeDefinitions(value interface{})
	PutGlobalSecondaryIndexes(value interface{})
	PutGlobalTableWitnesses(value interface{})
	PutKeySchema(value interface{})
	PutLocalSecondaryIndexes(value interface{})
	PutReadOnDemandThroughputSettings(value *DynamodbGlobalTableReadOnDemandThroughputSettings)
	PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableReadProvisionedThroughputSettings)
	PutReplicas(value interface{})
	PutSseSpecification(value *DynamodbGlobalTableSseSpecification)
	PutStreamSpecification(value *DynamodbGlobalTableStreamSpecification)
	PutTimeToLiveSpecification(value *DynamodbGlobalTableTimeToLiveSpecification)
	PutVectorIndexes(value interface{})
	PutWarmThroughput(value *DynamodbGlobalTableWarmThroughput)
	PutWriteOnDemandThroughputSettings(value *DynamodbGlobalTableWriteOnDemandThroughputSettings)
	PutWriteProvisionedThroughputSettings(value *DynamodbGlobalTableWriteProvisionedThroughputSettings)
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
	ResetAttributeDefinitions()
	ResetBillingMode()
	ResetGlobalSecondaryIndexes()
	ResetGlobalTableSourceArn()
	ResetGlobalTableWitnesses()
	ResetKeySchema()
	ResetLocalSecondaryIndexes()
	ResetMultiRegionConsistency()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadOnDemandThroughputSettings()
	ResetReadProvisionedThroughputSettings()
	ResetSseSpecification()
	ResetStreamSpecification()
	ResetTableName()
	ResetTimeToLiveSpecification()
	ResetVectorIndexes()
	ResetWarmThroughput()
	ResetWriteOnDemandThroughputSettings()
	ResetWriteProvisionedThroughputSettings()
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

// The jsii proxy struct for DynamodbGlobalTable
type jsiiProxy_DynamodbGlobalTable struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DynamodbGlobalTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) AttributeDefinitions() DynamodbGlobalTableAttributeDefinitionsList {
	var returns DynamodbGlobalTableAttributeDefinitionsList
	_jsii_.Get(
		j,
		"attributeDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) AttributeDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) BillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalSecondaryIndexes() DynamodbGlobalTableGlobalSecondaryIndexesList {
	var returns DynamodbGlobalTableGlobalSecondaryIndexesList
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalTableSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalTableSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalTableSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalTableSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalTableWitnesses() DynamodbGlobalTableGlobalTableWitnessesList {
	var returns DynamodbGlobalTableGlobalTableWitnessesList
	_jsii_.Get(
		j,
		"globalTableWitnesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) GlobalTableWitnessesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalTableWitnessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) KeySchema() DynamodbGlobalTableKeySchemaList {
	var returns DynamodbGlobalTableKeySchemaList
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) KeySchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) LocalSecondaryIndexes() DynamodbGlobalTableLocalSecondaryIndexesList {
	var returns DynamodbGlobalTableLocalSecondaryIndexesList
	_jsii_.Get(
		j,
		"localSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) LocalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) MultiRegionConsistency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiRegionConsistency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) MultiRegionConsistencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiRegionConsistencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ReadOnDemandThroughputSettings() DynamodbGlobalTableReadOnDemandThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableReadOnDemandThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"readOnDemandThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ReadOnDemandThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnDemandThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ReadProvisionedThroughputSettings() DynamodbGlobalTableReadProvisionedThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableReadProvisionedThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ReadProvisionedThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readProvisionedThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) Replicas() DynamodbGlobalTableReplicasList {
	var returns DynamodbGlobalTableReplicasList
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) ReplicasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) SseSpecification() DynamodbGlobalTableSseSpecificationOutputReference {
	var returns DynamodbGlobalTableSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) SseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) StreamSpecification() DynamodbGlobalTableStreamSpecificationOutputReference {
	var returns DynamodbGlobalTableStreamSpecificationOutputReference
	_jsii_.Get(
		j,
		"streamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) StreamSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TableId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TimeToLiveSpecification() DynamodbGlobalTableTimeToLiveSpecificationOutputReference {
	var returns DynamodbGlobalTableTimeToLiveSpecificationOutputReference
	_jsii_.Get(
		j,
		"timeToLiveSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) TimeToLiveSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToLiveSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) VectorIndexes() DynamodbGlobalTableVectorIndexesList {
	var returns DynamodbGlobalTableVectorIndexesList
	_jsii_.Get(
		j,
		"vectorIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) VectorIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vectorIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WarmThroughput() DynamodbGlobalTableWarmThroughputOutputReference {
	var returns DynamodbGlobalTableWarmThroughputOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WarmThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WriteOnDemandThroughputSettings() DynamodbGlobalTableWriteOnDemandThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableWriteOnDemandThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"writeOnDemandThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WriteOnDemandThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeOnDemandThroughputSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WriteProvisionedThroughputSettings() DynamodbGlobalTableWriteProvisionedThroughputSettingsOutputReference {
	var returns DynamodbGlobalTableWriteProvisionedThroughputSettingsOutputReference
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbGlobalTable) WriteProvisionedThroughputSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeProvisionedThroughputSettingsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table awscc_dynamodb_global_table} Resource.
func NewDynamodbGlobalTable(scope constructs.Construct, id *string, config *DynamodbGlobalTableConfig) DynamodbGlobalTable {
	_init_.Initialize()

	if err := validateNewDynamodbGlobalTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbGlobalTable{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table awscc_dynamodb_global_table} Resource.
func NewDynamodbGlobalTable_Override(d DynamodbGlobalTable, scope constructs.Construct, id *string, config *DynamodbGlobalTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetBillingMode(val *string) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetGlobalTableSourceArn(val *string) {
	if err := j.validateSetGlobalTableSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalTableSourceArn",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetMultiRegionConsistency(val *string) {
	if err := j.validateSetMultiRegionConsistencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiRegionConsistency",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DynamodbGlobalTable)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Generates CDKTN code for importing a DynamodbGlobalTable resource upon running "cdktn plan <stack-name>".
func DynamodbGlobalTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTable",
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
func DynamodbGlobalTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbGlobalTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbGlobalTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbGlobalTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamodbGlobalTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dynamodbGlobalTable.DynamodbGlobalTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbGlobalTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbGlobalTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbGlobalTable) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (d *jsiiProxy_DynamodbGlobalTable) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutAttributeDefinitions(value interface{}) {
	if err := d.validatePutAttributeDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttributeDefinitions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutGlobalSecondaryIndexes(value interface{}) {
	if err := d.validatePutGlobalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlobalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutGlobalTableWitnesses(value interface{}) {
	if err := d.validatePutGlobalTableWitnessesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlobalTableWitnesses",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutKeySchema(value interface{}) {
	if err := d.validatePutKeySchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKeySchema",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutLocalSecondaryIndexes(value interface{}) {
	if err := d.validatePutLocalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutReadOnDemandThroughputSettings(value *DynamodbGlobalTableReadOnDemandThroughputSettings) {
	if err := d.validatePutReadOnDemandThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadOnDemandThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutReadProvisionedThroughputSettings(value *DynamodbGlobalTableReadProvisionedThroughputSettings) {
	if err := d.validatePutReadProvisionedThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReadProvisionedThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutReplicas(value interface{}) {
	if err := d.validatePutReplicasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReplicas",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutSseSpecification(value *DynamodbGlobalTableSseSpecification) {
	if err := d.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutStreamSpecification(value *DynamodbGlobalTableStreamSpecification) {
	if err := d.validatePutStreamSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStreamSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutTimeToLiveSpecification(value *DynamodbGlobalTableTimeToLiveSpecification) {
	if err := d.validatePutTimeToLiveSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeToLiveSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutVectorIndexes(value interface{}) {
	if err := d.validatePutVectorIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVectorIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutWarmThroughput(value *DynamodbGlobalTableWarmThroughput) {
	if err := d.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutWriteOnDemandThroughputSettings(value *DynamodbGlobalTableWriteOnDemandThroughputSettings) {
	if err := d.validatePutWriteOnDemandThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWriteOnDemandThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) PutWriteProvisionedThroughputSettings(value *DynamodbGlobalTableWriteProvisionedThroughputSettings) {
	if err := d.validatePutWriteProvisionedThroughputSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWriteProvisionedThroughputSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetAttributeDefinitions() {
	_jsii_.InvokeVoid(
		d,
		"resetAttributeDefinitions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetGlobalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetGlobalTableSourceArn() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalTableSourceArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetGlobalTableWitnesses() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalTableWitnesses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetKeySchema() {
	_jsii_.InvokeVoid(
		d,
		"resetKeySchema",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetLocalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetMultiRegionConsistency() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiRegionConsistency",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetReadOnDemandThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReadOnDemandThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetReadProvisionedThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetReadProvisionedThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetStreamSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetTimeToLiveSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeToLiveSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetVectorIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetVectorIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetWriteOnDemandThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteOnDemandThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) ResetWriteProvisionedThroughputSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteProvisionedThroughputSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbGlobalTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbGlobalTable) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

