// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoicercsagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/smsvoicercsagent/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/smsvoice_rcs_agent awscc_smsvoice_rcs_agent}.
type SmsvoiceRcsAgent interface {
	cdktn.TerraformResource
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
	CreatedTimestamp() *string
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
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OptOutListName() *string
	SetOptOutListName(val *string)
	OptOutListNameInput() *string
	PoolId() *string
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
	RcsAgentArn() *string
	RcsAgentId() *string
	SelfManagedOptOutsEnabled() interface{}
	SetSelfManagedOptOutsEnabled(val interface{})
	SelfManagedOptOutsEnabledInput() interface{}
	Status() *string
	Tags() SmsvoiceRcsAgentTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TestingAgent() SmsvoiceRcsAgentTestingAgentOutputReference
	TwoWayChannelArn() *string
	SetTwoWayChannelArn(val *string)
	TwoWayChannelArnInput() *string
	TwoWayChannelRole() *string
	SetTwoWayChannelRole(val *string)
	TwoWayChannelRoleInput() *string
	TwoWayEnabled() interface{}
	SetTwoWayEnabled(val interface{})
	TwoWayEnabledInput() interface{}
	TwoWayMediaS3BucketName() *string
	SetTwoWayMediaS3BucketName(val *string)
	TwoWayMediaS3BucketNameInput() *string
	TwoWayMediaS3KeyPrefix() *string
	SetTwoWayMediaS3KeyPrefix(val *string)
	TwoWayMediaS3KeyPrefixInput() *string
	TwoWayMediaS3Role() *string
	SetTwoWayMediaS3Role(val *string)
	TwoWayMediaS3RoleInput() *string
	TwoWayRcsEventsEnabled() *[]*string
	SetTwoWayRcsEventsEnabled(val *[]*string)
	TwoWayRcsEventsEnabledInput() *[]*string
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
	ResetDeletionProtectionEnabled()
	ResetOptOutListName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSelfManagedOptOutsEnabled()
	ResetTags()
	ResetTwoWayChannelArn()
	ResetTwoWayChannelRole()
	ResetTwoWayEnabled()
	ResetTwoWayMediaS3BucketName()
	ResetTwoWayMediaS3KeyPrefix()
	ResetTwoWayMediaS3Role()
	ResetTwoWayRcsEventsEnabled()
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

// The jsii proxy struct for SmsvoiceRcsAgent
type jsiiProxy_SmsvoiceRcsAgent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_SmsvoiceRcsAgent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) CreatedTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) OptOutListName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) OptOutListNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optOutListNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) PoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"poolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) RcsAgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rcsAgentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) RcsAgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rcsAgentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) SelfManagedOptOutsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) SelfManagedOptOutsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selfManagedOptOutsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) Tags() SmsvoiceRcsAgentTagsList {
	var returns SmsvoiceRcsAgentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TestingAgent() SmsvoiceRcsAgentTestingAgentOutputReference {
	var returns SmsvoiceRcsAgentTestingAgentOutputReference
	_jsii_.Get(
		j,
		"testingAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayChannelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayChannelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayChannelRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayChannelRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayChannelRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"twoWayEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayMediaS3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayMediaS3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayMediaS3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayMediaS3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayMediaS3Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3Role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayMediaS3RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twoWayMediaS3RoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayRcsEventsEnabled() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"twoWayRcsEventsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SmsvoiceRcsAgent) TwoWayRcsEventsEnabledInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"twoWayRcsEventsEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/smsvoice_rcs_agent awscc_smsvoice_rcs_agent} Resource.
func NewSmsvoiceRcsAgent(scope constructs.Construct, id *string, config *SmsvoiceRcsAgentConfig) SmsvoiceRcsAgent {
	_init_.Initialize()

	if err := validateNewSmsvoiceRcsAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SmsvoiceRcsAgent{}

	_jsii_.Create(
		"@cdktn/provider-awscc.smsvoiceRcsAgent.SmsvoiceRcsAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/smsvoice_rcs_agent awscc_smsvoice_rcs_agent} Resource.
func NewSmsvoiceRcsAgent_Override(s SmsvoiceRcsAgent, scope constructs.Construct, id *string, config *SmsvoiceRcsAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.smsvoiceRcsAgent.SmsvoiceRcsAgent",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetOptOutListName(val *string) {
	if err := j.validateSetOptOutListNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optOutListName",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetSelfManagedOptOutsEnabled(val interface{}) {
	if err := j.validateSetSelfManagedOptOutsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfManagedOptOutsEnabled",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetTwoWayChannelArn(val *string) {
	if err := j.validateSetTwoWayChannelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelArn",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetTwoWayChannelRole(val *string) {
	if err := j.validateSetTwoWayChannelRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayChannelRole",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetTwoWayEnabled(val interface{}) {
	if err := j.validateSetTwoWayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayEnabled",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetTwoWayMediaS3BucketName(val *string) {
	if err := j.validateSetTwoWayMediaS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayMediaS3BucketName",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetTwoWayMediaS3KeyPrefix(val *string) {
	if err := j.validateSetTwoWayMediaS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayMediaS3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetTwoWayMediaS3Role(val *string) {
	if err := j.validateSetTwoWayMediaS3RoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayMediaS3Role",
		val,
	)
}

func (j *jsiiProxy_SmsvoiceRcsAgent)SetTwoWayRcsEventsEnabled(val *[]*string) {
	if err := j.validateSetTwoWayRcsEventsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twoWayRcsEventsEnabled",
		val,
	)
}

// Generates CDKTN code for importing a SmsvoiceRcsAgent resource upon running "cdktn plan <stack-name>".
func SmsvoiceRcsAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateSmsvoiceRcsAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.smsvoiceRcsAgent.SmsvoiceRcsAgent",
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
func SmsvoiceRcsAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSmsvoiceRcsAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.smsvoiceRcsAgent.SmsvoiceRcsAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SmsvoiceRcsAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSmsvoiceRcsAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.smsvoiceRcsAgent.SmsvoiceRcsAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SmsvoiceRcsAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSmsvoiceRcsAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.smsvoiceRcsAgent.SmsvoiceRcsAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SmsvoiceRcsAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.smsvoiceRcsAgent.SmsvoiceRcsAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := s.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := s.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetOptOutListName() {
	_jsii_.InvokeVoid(
		s,
		"resetOptOutListName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetSelfManagedOptOutsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSelfManagedOptOutsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTwoWayChannelArn() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoWayChannelArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTwoWayChannelRole() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoWayChannelRole",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTwoWayEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoWayEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTwoWayMediaS3BucketName() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoWayMediaS3BucketName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTwoWayMediaS3KeyPrefix() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoWayMediaS3KeyPrefix",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTwoWayMediaS3Role() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoWayMediaS3Role",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ResetTwoWayRcsEventsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetTwoWayRcsEventsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SmsvoiceRcsAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SmsvoiceRcsAgent) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		s,
		"with",
		args,
		&returns,
	)

	return returns
}

