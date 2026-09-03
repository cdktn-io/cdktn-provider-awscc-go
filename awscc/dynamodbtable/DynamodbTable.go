// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dynamodbtable/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_table awscc_dynamodb_table}.
type DynamodbTable interface {
	cdktn.TerraformResource
	Arn() *string
	AttributeDefinitions() DynamodbTableAttributeDefinitionsList
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
	ContributorInsightsSpecification() DynamodbTableContributorInsightsSpecificationOutputReference
	ContributorInsightsSpecificationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
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
	GlobalSecondaryIndexes() DynamodbTableGlobalSecondaryIndexesList
	GlobalSecondaryIndexesInput() interface{}
	Id() *string
	ImportSourceSpecification() DynamodbTableImportSourceSpecificationOutputReference
	ImportSourceSpecificationInput() interface{}
	KeySchema() *string
	SetKeySchema(val *string)
	KeySchemaInput() *string
	KinesisStreamSpecification() DynamodbTableKinesisStreamSpecificationOutputReference
	KinesisStreamSpecificationInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LocalSecondaryIndexes() DynamodbTableLocalSecondaryIndexesList
	LocalSecondaryIndexesInput() interface{}
	// The tree node.
	Node() constructs.Node
	OnDemandThroughput() DynamodbTableOnDemandThroughputOutputReference
	OnDemandThroughputInput() interface{}
	PointInTimeRecoverySpecification() DynamodbTablePointInTimeRecoverySpecificationOutputReference
	PointInTimeRecoverySpecificationInput() interface{}
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	ProvisionedThroughput() DynamodbTableProvisionedThroughputOutputReference
	ProvisionedThroughputInput() interface{}
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourcePolicy() DynamodbTableResourcePolicyOutputReference
	ResourcePolicyInput() interface{}
	SseSpecification() DynamodbTableSseSpecificationOutputReference
	SseSpecificationInput() interface{}
	StreamArn() *string
	StreamSpecification() DynamodbTableStreamSpecificationOutputReference
	StreamSpecificationInput() interface{}
	TableClass() *string
	SetTableClass(val *string)
	TableClassInput() *string
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	Tags() DynamodbTableTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeToLiveSpecification() DynamodbTableTimeToLiveSpecificationOutputReference
	TimeToLiveSpecificationInput() interface{}
	WarmThroughput() DynamodbTableWarmThroughputOutputReference
	WarmThroughputInput() interface{}
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
	PutContributorInsightsSpecification(value *DynamodbTableContributorInsightsSpecification)
	PutGlobalSecondaryIndexes(value interface{})
	PutImportSourceSpecification(value *DynamodbTableImportSourceSpecification)
	PutKinesisStreamSpecification(value *DynamodbTableKinesisStreamSpecification)
	PutLocalSecondaryIndexes(value interface{})
	PutOnDemandThroughput(value *DynamodbTableOnDemandThroughput)
	PutPointInTimeRecoverySpecification(value *DynamodbTablePointInTimeRecoverySpecification)
	PutProvisionedThroughput(value *DynamodbTableProvisionedThroughput)
	PutResourcePolicy(value *DynamodbTableResourcePolicy)
	PutSseSpecification(value *DynamodbTableSseSpecification)
	PutStreamSpecification(value *DynamodbTableStreamSpecification)
	PutTags(value interface{})
	PutTimeToLiveSpecification(value *DynamodbTableTimeToLiveSpecification)
	PutWarmThroughput(value *DynamodbTableWarmThroughput)
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
	ResetContributorInsightsSpecification()
	ResetDeletionProtectionEnabled()
	ResetGlobalSecondaryIndexes()
	ResetImportSourceSpecification()
	ResetKinesisStreamSpecification()
	ResetLocalSecondaryIndexes()
	ResetOnDemandThroughput()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRecoverySpecification()
	ResetProvisionedThroughput()
	ResetResourcePolicy()
	ResetSseSpecification()
	ResetStreamSpecification()
	ResetTableClass()
	ResetTableName()
	ResetTags()
	ResetTimeToLiveSpecification()
	ResetWarmThroughput()
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

// The jsii proxy struct for DynamodbTable
type jsiiProxy_DynamodbTable struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_DynamodbTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) AttributeDefinitions() DynamodbTableAttributeDefinitionsList {
	var returns DynamodbTableAttributeDefinitionsList
	_jsii_.Get(
		j,
		"attributeDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) AttributeDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) BillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ContributorInsightsSpecification() DynamodbTableContributorInsightsSpecificationOutputReference {
	var returns DynamodbTableContributorInsightsSpecificationOutputReference
	_jsii_.Get(
		j,
		"contributorInsightsSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ContributorInsightsSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contributorInsightsSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) GlobalSecondaryIndexes() DynamodbTableGlobalSecondaryIndexesList {
	var returns DynamodbTableGlobalSecondaryIndexesList
	_jsii_.Get(
		j,
		"globalSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) GlobalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ImportSourceSpecification() DynamodbTableImportSourceSpecificationOutputReference {
	var returns DynamodbTableImportSourceSpecificationOutputReference
	_jsii_.Get(
		j,
		"importSourceSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ImportSourceSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"importSourceSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KeySchema() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KeySchemaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keySchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KinesisStreamSpecification() DynamodbTableKinesisStreamSpecificationOutputReference {
	var returns DynamodbTableKinesisStreamSpecificationOutputReference
	_jsii_.Get(
		j,
		"kinesisStreamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) KinesisStreamSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisStreamSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) LocalSecondaryIndexes() DynamodbTableLocalSecondaryIndexesList {
	var returns DynamodbTableLocalSecondaryIndexesList
	_jsii_.Get(
		j,
		"localSecondaryIndexes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) LocalSecondaryIndexesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) OnDemandThroughput() DynamodbTableOnDemandThroughputOutputReference {
	var returns DynamodbTableOnDemandThroughputOutputReference
	_jsii_.Get(
		j,
		"onDemandThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) OnDemandThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) PointInTimeRecoverySpecification() DynamodbTablePointInTimeRecoverySpecificationOutputReference {
	var returns DynamodbTablePointInTimeRecoverySpecificationOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecoverySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) PointInTimeRecoverySpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pointInTimeRecoverySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ProvisionedThroughput() DynamodbTableProvisionedThroughputOutputReference {
	var returns DynamodbTableProvisionedThroughputOutputReference
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ProvisionedThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ResourcePolicy() DynamodbTableResourcePolicyOutputReference {
	var returns DynamodbTableResourcePolicyOutputReference
	_jsii_.Get(
		j,
		"resourcePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ResourcePolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourcePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) SseSpecification() DynamodbTableSseSpecificationOutputReference {
	var returns DynamodbTableSseSpecificationOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) SseSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamSpecification() DynamodbTableStreamSpecificationOutputReference {
	var returns DynamodbTableStreamSpecificationOutputReference
	_jsii_.Get(
		j,
		"streamSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Tags() DynamodbTableTagsList {
	var returns DynamodbTableTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TimeToLiveSpecification() DynamodbTableTimeToLiveSpecificationOutputReference {
	var returns DynamodbTableTimeToLiveSpecificationOutputReference
	_jsii_.Get(
		j,
		"timeToLiveSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TimeToLiveSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToLiveSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) WarmThroughput() DynamodbTableWarmThroughputOutputReference {
	var returns DynamodbTableWarmThroughputOutputReference
	_jsii_.Get(
		j,
		"warmThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) WarmThroughputInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"warmThroughputInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_table awscc_dynamodb_table} Resource.
func NewDynamodbTable(scope constructs.Construct, id *string, config *DynamodbTableConfig) DynamodbTable {
	_init_.Initialize()

	if err := validateNewDynamodbTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTable{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_table awscc_dynamodb_table} Resource.
func NewDynamodbTable_Override(d DynamodbTable, scope constructs.Construct, id *string, config *DynamodbTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTable",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamodbTable)SetBillingMode(val *string) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetKeySchema(val *string) {
	if err := j.validateSetKeySchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySchema",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetTableClass(val *string) {
	if err := j.validateSetTableClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableClass",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

// Generates CDKTN code for importing a DynamodbTable resource upon running "cdktn plan <stack-name>".
func DynamodbTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDynamodbTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTable",
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
func DynamodbTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamodbTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.dynamodbTable.DynamodbTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamodbTable) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamodbTable) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamodbTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTable) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamodbTable) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DynamodbTable) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (d *jsiiProxy_DynamodbTable) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbTable) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamodbTable) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbTable) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamodbTable) PutAttributeDefinitions(value interface{}) {
	if err := d.validatePutAttributeDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttributeDefinitions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutContributorInsightsSpecification(value *DynamodbTableContributorInsightsSpecification) {
	if err := d.validatePutContributorInsightsSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putContributorInsightsSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutGlobalSecondaryIndexes(value interface{}) {
	if err := d.validatePutGlobalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlobalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutImportSourceSpecification(value *DynamodbTableImportSourceSpecification) {
	if err := d.validatePutImportSourceSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImportSourceSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutKinesisStreamSpecification(value *DynamodbTableKinesisStreamSpecification) {
	if err := d.validatePutKinesisStreamSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putKinesisStreamSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutLocalSecondaryIndexes(value interface{}) {
	if err := d.validatePutLocalSecondaryIndexesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocalSecondaryIndexes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutOnDemandThroughput(value *DynamodbTableOnDemandThroughput) {
	if err := d.validatePutOnDemandThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOnDemandThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutPointInTimeRecoverySpecification(value *DynamodbTablePointInTimeRecoverySpecification) {
	if err := d.validatePutPointInTimeRecoverySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPointInTimeRecoverySpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutProvisionedThroughput(value *DynamodbTableProvisionedThroughput) {
	if err := d.validatePutProvisionedThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProvisionedThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutResourcePolicy(value *DynamodbTableResourcePolicy) {
	if err := d.validatePutResourcePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putResourcePolicy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutSseSpecification(value *DynamodbTableSseSpecification) {
	if err := d.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutStreamSpecification(value *DynamodbTableStreamSpecification) {
	if err := d.validatePutStreamSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putStreamSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutTimeToLiveSpecification(value *DynamodbTableTimeToLiveSpecification) {
	if err := d.validatePutTimeToLiveSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeToLiveSpecification",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutWarmThroughput(value *DynamodbTableWarmThroughput) {
	if err := d.validatePutWarmThroughputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWarmThroughput",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DynamodbTable) ResetAttributeDefinitions() {
	_jsii_.InvokeVoid(
		d,
		"resetAttributeDefinitions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetContributorInsightsSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetContributorInsightsSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetGlobalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetImportSourceSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetImportSourceSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetKinesisStreamSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetKinesisStreamSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetLocalSecondaryIndexes() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSecondaryIndexes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetOnDemandThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetOnDemandThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetPointInTimeRecoverySpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetPointInTimeRecoverySpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetResourcePolicy() {
	_jsii_.InvokeVoid(
		d,
		"resetResourcePolicy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetStreamSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTableClass() {
	_jsii_.InvokeVoid(
		d,
		"resetTableClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTableName() {
	_jsii_.InvokeVoid(
		d,
		"resetTableName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTimeToLiveSpecification() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeToLiveSpecification",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetWarmThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetWarmThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

