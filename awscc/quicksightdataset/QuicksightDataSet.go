// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightdataset/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set awscc_quicksight_data_set}.
type QuicksightDataSet interface {
	cdktn.TerraformResource
	Arn() *string
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	ColumnGroups() QuicksightDataSetColumnGroupsList
	ColumnGroupsInput() interface{}
	ColumnLevelPermissionRules() QuicksightDataSetColumnLevelPermissionRulesList
	ColumnLevelPermissionRulesInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ConsumedSpiceCapacityInBytes() *float64
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedTime() *string
	DataSetId() *string
	SetDataSetId(val *string)
	DataSetIdInput() *string
	DataSetUsageConfiguration() QuicksightDataSetDataSetUsageConfigurationOutputReference
	DataSetUsageConfigurationInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FieldFolders() QuicksightDataSetFieldFoldersMap
	FieldFoldersInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	ImportMode() *string
	SetImportMode(val *string)
	ImportModeInput() *string
	IngestionWaitPolicy() QuicksightDataSetIngestionWaitPolicyOutputReference
	IngestionWaitPolicyInput() interface{}
	LastUpdatedTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LogicalTableMap() QuicksightDataSetLogicalTableMapMap
	LogicalTableMapInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputColumns() QuicksightDataSetOutputColumnsList
	Permissions() QuicksightDataSetPermissionsList
	PermissionsInput() interface{}
	PhysicalTableMap() QuicksightDataSetPhysicalTableMapMap
	PhysicalTableMapInput() interface{}
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
	RowLevelPermissionDataSet() QuicksightDataSetRowLevelPermissionDataSetOutputReference
	RowLevelPermissionDataSetInput() interface{}
	Tags() QuicksightDataSetTagsList
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
	PutColumnGroups(value interface{})
	PutColumnLevelPermissionRules(value interface{})
	PutDataSetUsageConfiguration(value *QuicksightDataSetDataSetUsageConfiguration)
	PutFieldFolders(value interface{})
	PutIngestionWaitPolicy(value *QuicksightDataSetIngestionWaitPolicy)
	PutLogicalTableMap(value interface{})
	PutPermissions(value interface{})
	PutPhysicalTableMap(value interface{})
	PutRowLevelPermissionDataSet(value *QuicksightDataSetRowLevelPermissionDataSet)
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
	ResetAwsAccountId()
	ResetColumnGroups()
	ResetColumnLevelPermissionRules()
	ResetDataSetId()
	ResetDataSetUsageConfiguration()
	ResetFieldFolders()
	ResetImportMode()
	ResetIngestionWaitPolicy()
	ResetLogicalTableMap()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissions()
	ResetPhysicalTableMap()
	ResetRowLevelPermissionDataSet()
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

// The jsii proxy struct for QuicksightDataSet
type jsiiProxy_QuicksightDataSet struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_QuicksightDataSet) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnGroups() QuicksightDataSetColumnGroupsList {
	var returns QuicksightDataSetColumnGroupsList
	_jsii_.Get(
		j,
		"columnGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnLevelPermissionRules() QuicksightDataSetColumnLevelPermissionRulesList {
	var returns QuicksightDataSetColumnLevelPermissionRulesList
	_jsii_.Get(
		j,
		"columnLevelPermissionRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ColumnLevelPermissionRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"columnLevelPermissionRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ConsumedSpiceCapacityInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"consumedSpiceCapacityInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) CreatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetUsageConfiguration() QuicksightDataSetDataSetUsageConfigurationOutputReference {
	var returns QuicksightDataSetDataSetUsageConfigurationOutputReference
	_jsii_.Get(
		j,
		"dataSetUsageConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DataSetUsageConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSetUsageConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) FieldFolders() QuicksightDataSetFieldFoldersMap {
	var returns QuicksightDataSetFieldFoldersMap
	_jsii_.Get(
		j,
		"fieldFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) FieldFoldersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fieldFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ImportMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) ImportModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) IngestionWaitPolicy() QuicksightDataSetIngestionWaitPolicyOutputReference {
	var returns QuicksightDataSetIngestionWaitPolicyOutputReference
	_jsii_.Get(
		j,
		"ingestionWaitPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) IngestionWaitPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ingestionWaitPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) LogicalTableMap() QuicksightDataSetLogicalTableMapMap {
	var returns QuicksightDataSetLogicalTableMapMap
	_jsii_.Get(
		j,
		"logicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) LogicalTableMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logicalTableMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) OutputColumns() QuicksightDataSetOutputColumnsList {
	var returns QuicksightDataSetOutputColumnsList
	_jsii_.Get(
		j,
		"outputColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Permissions() QuicksightDataSetPermissionsList {
	var returns QuicksightDataSetPermissionsList
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) PermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) PhysicalTableMap() QuicksightDataSetPhysicalTableMapMap {
	var returns QuicksightDataSetPhysicalTableMapMap
	_jsii_.Get(
		j,
		"physicalTableMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) PhysicalTableMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"physicalTableMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) RowLevelPermissionDataSet() QuicksightDataSetRowLevelPermissionDataSetOutputReference {
	var returns QuicksightDataSetRowLevelPermissionDataSetOutputReference
	_jsii_.Get(
		j,
		"rowLevelPermissionDataSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) RowLevelPermissionDataSetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowLevelPermissionDataSetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) Tags() QuicksightDataSetTagsList {
	var returns QuicksightDataSetTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set awscc_quicksight_data_set} Resource.
func NewQuicksightDataSet(scope constructs.Construct, id *string, config *QuicksightDataSetConfig) QuicksightDataSet {
	_init_.Initialize()

	if err := validateNewQuicksightDataSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSet{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set awscc_quicksight_data_set} Resource.
func NewQuicksightDataSet_Override(q QuicksightDataSet, scope constructs.Construct, id *string, config *QuicksightDataSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSet",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetDataSetId(val *string) {
	if err := j.validateSetDataSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSetId",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetImportMode(val *string) {
	if err := j.validateSetImportModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importMode",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a QuicksightDataSet resource upon running "cdktn plan <stack-name>".
func QuicksightDataSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateQuicksightDataSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSet",
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
func QuicksightDataSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightDataSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightDataSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightDataSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightDataSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightDataSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightDataSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.quicksightDataSet.QuicksightDataSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QuicksightDataSet) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QuicksightDataSet) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightDataSet) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightDataSet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightDataSet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightDataSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightDataSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightDataSet) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightDataSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightDataSet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QuicksightDataSet) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightDataSet) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := q.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightDataSet) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QuicksightDataSet) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightDataSet) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutColumnGroups(value interface{}) {
	if err := q.validatePutColumnGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putColumnGroups",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutColumnLevelPermissionRules(value interface{}) {
	if err := q.validatePutColumnLevelPermissionRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putColumnLevelPermissionRules",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutDataSetUsageConfiguration(value *QuicksightDataSetDataSetUsageConfiguration) {
	if err := q.validatePutDataSetUsageConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putDataSetUsageConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutFieldFolders(value interface{}) {
	if err := q.validatePutFieldFoldersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putFieldFolders",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutIngestionWaitPolicy(value *QuicksightDataSetIngestionWaitPolicy) {
	if err := q.validatePutIngestionWaitPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIngestionWaitPolicy",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutLogicalTableMap(value interface{}) {
	if err := q.validatePutLogicalTableMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putLogicalTableMap",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutPermissions(value interface{}) {
	if err := q.validatePutPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPermissions",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutPhysicalTableMap(value interface{}) {
	if err := q.validatePutPhysicalTableMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPhysicalTableMap",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutRowLevelPermissionDataSet(value *QuicksightDataSetRowLevelPermissionDataSet) {
	if err := q.validatePutRowLevelPermissionDataSetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putRowLevelPermissionDataSet",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) PutTags(value interface{}) {
	if err := q.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTags",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSet) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := q.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetAwsAccountId() {
	_jsii_.InvokeVoid(
		q,
		"resetAwsAccountId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetColumnGroups() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnGroups",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetColumnLevelPermissionRules() {
	_jsii_.InvokeVoid(
		q,
		"resetColumnLevelPermissionRules",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetDataSetId() {
	_jsii_.InvokeVoid(
		q,
		"resetDataSetId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetDataSetUsageConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetDataSetUsageConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetFieldFolders() {
	_jsii_.InvokeVoid(
		q,
		"resetFieldFolders",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetImportMode() {
	_jsii_.InvokeVoid(
		q,
		"resetImportMode",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetIngestionWaitPolicy() {
	_jsii_.InvokeVoid(
		q,
		"resetIngestionWaitPolicy",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetLogicalTableMap() {
	_jsii_.InvokeVoid(
		q,
		"resetLogicalTableMap",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetName() {
	_jsii_.InvokeVoid(
		q,
		"resetName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetPermissions() {
	_jsii_.InvokeVoid(
		q,
		"resetPermissions",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetPhysicalTableMap() {
	_jsii_.InvokeVoid(
		q,
		"resetPhysicalTableMap",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetRowLevelPermissionDataSet() {
	_jsii_.InvokeVoid(
		q,
		"resetRowLevelPermissionDataSet",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) ResetTags() {
	_jsii_.InvokeVoid(
		q,
		"resetTags",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSet) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		q,
		"with",
		args,
		&returns,
	)

	return returns
}

