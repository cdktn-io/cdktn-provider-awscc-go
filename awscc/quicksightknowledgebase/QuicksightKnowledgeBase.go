// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightknowledgebase

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/quicksightknowledgebase/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_knowledge_base awscc_quicksight_knowledge_base}.
type QuicksightKnowledgeBase interface {
	cdktn.TerraformResource
	AccessControlConfiguration() QuicksightKnowledgeBaseAccessControlConfigurationOutputReference
	AccessControlConfigurationInput() interface{}
	AwsAccountId() *string
	SetAwsAccountId(val *string)
	AwsAccountIdInput() *string
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
	CreatedAt() *string
	DataSourceArn() *string
	SetDataSourceArn(val *string)
	DataSourceArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DocumentCount() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IsEmailNotificationOptedForIngestionFailures() interface{}
	SetIsEmailNotificationOptedForIngestionFailures(val interface{})
	IsEmailNotificationOptedForIngestionFailuresInput() interface{}
	KnowledgeBaseArn() *string
	KnowledgeBaseConfiguration() QuicksightKnowledgeBaseKnowledgeBaseConfigurationOutputReference
	KnowledgeBaseConfigurationInput() interface{}
	KnowledgeBaseId() *string
	SetKnowledgeBaseId(val *string)
	KnowledgeBaseIdInput() *string
	KnowledgeBaseSizeBytes() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MediaExtractionConfiguration() QuicksightKnowledgeBaseMediaExtractionConfigurationOutputReference
	MediaExtractionConfigurationInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Permissions() QuicksightKnowledgeBasePermissionsList
	PermissionsInput() interface{}
	PrimaryOwnerArn() *string
	SetPrimaryOwnerArn(val *string)
	PrimaryOwnerArnInput() *string
	PrimaryOwnerUsername() *string
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
	Status() *string
	Tags() QuicksightKnowledgeBaseTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	UpdatedAt() *string
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
	PutAccessControlConfiguration(value *QuicksightKnowledgeBaseAccessControlConfiguration)
	PutKnowledgeBaseConfiguration(value *QuicksightKnowledgeBaseKnowledgeBaseConfiguration)
	PutMediaExtractionConfiguration(value *QuicksightKnowledgeBaseMediaExtractionConfiguration)
	PutPermissions(value interface{})
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
	ResetAccessControlConfiguration()
	ResetDescription()
	ResetIsEmailNotificationOptedForIngestionFailures()
	ResetMediaExtractionConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissions()
	ResetPrimaryOwnerArn()
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

// The jsii proxy struct for QuicksightKnowledgeBase
type jsiiProxy_QuicksightKnowledgeBase struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_QuicksightKnowledgeBase) AccessControlConfiguration() QuicksightKnowledgeBaseAccessControlConfigurationOutputReference {
	var returns QuicksightKnowledgeBaseAccessControlConfigurationOutputReference
	_jsii_.Get(
		j,
		"accessControlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) AccessControlConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessControlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) AwsAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) DataSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) DataSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) DocumentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"documentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) IsEmailNotificationOptedForIngestionFailures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEmailNotificationOptedForIngestionFailures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) IsEmailNotificationOptedForIngestionFailuresInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isEmailNotificationOptedForIngestionFailuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) KnowledgeBaseArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) KnowledgeBaseConfiguration() QuicksightKnowledgeBaseKnowledgeBaseConfigurationOutputReference {
	var returns QuicksightKnowledgeBaseKnowledgeBaseConfigurationOutputReference
	_jsii_.Get(
		j,
		"knowledgeBaseConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) KnowledgeBaseConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"knowledgeBaseConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) KnowledgeBaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) KnowledgeBaseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"knowledgeBaseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) KnowledgeBaseSizeBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"knowledgeBaseSizeBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) MediaExtractionConfiguration() QuicksightKnowledgeBaseMediaExtractionConfigurationOutputReference {
	var returns QuicksightKnowledgeBaseMediaExtractionConfigurationOutputReference
	_jsii_.Get(
		j,
		"mediaExtractionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) MediaExtractionConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediaExtractionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Permissions() QuicksightKnowledgeBasePermissionsList {
	var returns QuicksightKnowledgeBasePermissionsList
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) PermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) PrimaryOwnerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOwnerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) PrimaryOwnerArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOwnerArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) PrimaryOwnerUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryOwnerUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Tags() QuicksightKnowledgeBaseTagsList {
	var returns QuicksightKnowledgeBaseTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightKnowledgeBase) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_knowledge_base awscc_quicksight_knowledge_base} Resource.
func NewQuicksightKnowledgeBase(scope constructs.Construct, id *string, config *QuicksightKnowledgeBaseConfig) QuicksightKnowledgeBase {
	_init_.Initialize()

	if err := validateNewQuicksightKnowledgeBaseParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightKnowledgeBase{}

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBase",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_knowledge_base awscc_quicksight_knowledge_base} Resource.
func NewQuicksightKnowledgeBase_Override(q QuicksightKnowledgeBase, scope constructs.Construct, id *string, config *QuicksightKnowledgeBaseConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBase",
		[]interface{}{scope, id, config},
		q,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetAwsAccountId(val *string) {
	if err := j.validateSetAwsAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"awsAccountId",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetDataSourceArn(val *string) {
	if err := j.validateSetDataSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataSourceArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetIsEmailNotificationOptedForIngestionFailures(val interface{}) {
	if err := j.validateSetIsEmailNotificationOptedForIngestionFailuresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isEmailNotificationOptedForIngestionFailures",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetKnowledgeBaseId(val *string) {
	if err := j.validateSetKnowledgeBaseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"knowledgeBaseId",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetPrimaryOwnerArn(val *string) {
	if err := j.validateSetPrimaryOwnerArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryOwnerArn",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_QuicksightKnowledgeBase)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a QuicksightKnowledgeBase resource upon running "cdktn plan <stack-name>".
func QuicksightKnowledgeBase_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateQuicksightKnowledgeBase_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBase",
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
func QuicksightKnowledgeBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightKnowledgeBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightKnowledgeBase_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightKnowledgeBase_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBase",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func QuicksightKnowledgeBase_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateQuicksightKnowledgeBase_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBase",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func QuicksightKnowledgeBase_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.quicksightKnowledgeBase.QuicksightKnowledgeBase",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) AddMoveTarget(moveTarget *string) {
	if err := q.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) AddOverride(path *string, value interface{}) {
	if err := q.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := q.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (q *jsiiProxy_QuicksightKnowledgeBase) MoveFromId(id *string) {
	if err := q.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveFromId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) MoveTo(moveTarget *string, index interface{}) {
	if err := q.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) MoveToId(id *string) {
	if err := q.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"moveToId",
		[]interface{}{id},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) OverrideLogicalId(newLogicalId *string) {
	if err := q.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) PutAccessControlConfiguration(value *QuicksightKnowledgeBaseAccessControlConfiguration) {
	if err := q.validatePutAccessControlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putAccessControlConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) PutKnowledgeBaseConfiguration(value *QuicksightKnowledgeBaseKnowledgeBaseConfiguration) {
	if err := q.validatePutKnowledgeBaseConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putKnowledgeBaseConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) PutMediaExtractionConfiguration(value *QuicksightKnowledgeBaseMediaExtractionConfiguration) {
	if err := q.validatePutMediaExtractionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putMediaExtractionConfiguration",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) PutPermissions(value interface{}) {
	if err := q.validatePutPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putPermissions",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) PutTags(value interface{}) {
	if err := q.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putTags",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := q.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetAccessControlConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetAccessControlConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetDescription() {
	_jsii_.InvokeVoid(
		q,
		"resetDescription",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetIsEmailNotificationOptedForIngestionFailures() {
	_jsii_.InvokeVoid(
		q,
		"resetIsEmailNotificationOptedForIngestionFailures",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetMediaExtractionConfiguration() {
	_jsii_.InvokeVoid(
		q,
		"resetMediaExtractionConfiguration",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		q,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetPermissions() {
	_jsii_.InvokeVoid(
		q,
		"resetPermissions",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetPrimaryOwnerArn() {
	_jsii_.InvokeVoid(
		q,
		"resetPrimaryOwnerArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ResetTags() {
	_jsii_.InvokeVoid(
		q,
		"resetTags",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightKnowledgeBase) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		q,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightKnowledgeBase) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

