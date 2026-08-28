// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluesession

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/gluesession/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_session awscc_glue_session}.
type GlueSession interface {
	cdktn.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Command() GlueSessionCommandOutputReference
	CommandInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	Connections() GlueSessionConnectionsOutputReference
	ConnectionsInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreatedOn() *string
	DefaultArguments() *map[string]*string
	SetDefaultArguments(val *map[string]*string)
	DefaultArgumentsInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlueVersion() *string
	SetGlueVersion(val *string)
	GlueVersionInput() *string
	Id() *string
	IdleTimeout() *float64
	SetIdleTimeout(val *float64)
	IdleTimeoutInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaxCapacity() *float64
	SetMaxCapacity(val *float64)
	MaxCapacityInput() *float64
	// The tree node.
	Node() constructs.Node
	NumberOfWorkers() *float64
	SetNumberOfWorkers(val *float64)
	NumberOfWorkersInput() *float64
	Progress() *float64
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
	RequestOrigin() *string
	SetRequestOrigin(val *string)
	RequestOriginInput() *string
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SecurityConfiguration() *string
	SetSecurityConfiguration(val *string)
	SecurityConfigurationInput() *string
	SessionId() *string
	SetSessionId(val *string)
	SessionIdInput() *string
	Status() *string
	Tags() GlueSessionTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	WorkerType() *string
	SetWorkerType(val *string)
	WorkerTypeInput() *string
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
	PutCommand(value *GlueSessionCommand)
	PutConnections(value *GlueSessionConnections)
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
	ResetConnections()
	ResetDefaultArguments()
	ResetDescription()
	ResetGlueVersion()
	ResetIdleTimeout()
	ResetMaxCapacity()
	ResetNumberOfWorkers()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequestOrigin()
	ResetSecurityConfiguration()
	ResetTags()
	ResetTimeout()
	ResetWorkerType()
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

// The jsii proxy struct for GlueSession
type jsiiProxy_GlueSession struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_GlueSession) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Command() GlueSessionCommandOutputReference {
	var returns GlueSessionCommandOutputReference
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) CommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Connections() GlueSessionConnectionsOutputReference {
	var returns GlueSessionConnectionsOutputReference
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) ConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) DefaultArguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) DefaultArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) GlueVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) IdleTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) IdleTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) NumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Progress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"progress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) RequestOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) RequestOriginInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestOriginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) SessionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) SessionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Tags() GlueSessionTagsList {
	var returns GlueSessionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueSession) WorkerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_session awscc_glue_session} Resource.
func NewGlueSession(scope constructs.Construct, id *string, config *GlueSessionConfig) GlueSession {
	_init_.Initialize()

	if err := validateNewGlueSessionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueSession{}

	_jsii_.Create(
		"@cdktn/provider-awscc.glueSession.GlueSession",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/glue_session awscc_glue_session} Resource.
func NewGlueSession_Override(g GlueSession, scope constructs.Construct, id *string, config *GlueSessionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.glueSession.GlueSession",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueSession)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetDefaultArguments(val *map[string]*string) {
	if err := j.validateSetDefaultArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultArguments",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetGlueVersion(val *string) {
	if err := j.validateSetGlueVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetIdleTimeout(val *float64) {
	if err := j.validateSetIdleTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleTimeout",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetNumberOfWorkers(val *float64) {
	if err := j.validateSetNumberOfWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetRequestOrigin(val *string) {
	if err := j.validateSetRequestOriginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestOrigin",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetSessionId(val *string) {
	if err := j.validateSetSessionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionId",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_GlueSession)SetWorkerType(val *string) {
	if err := j.validateSetWorkerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Generates CDKTN code for importing a GlueSession resource upon running "cdktn plan <stack-name>".
func GlueSession_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateGlueSession_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.glueSession.GlueSession",
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
func GlueSession_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueSession_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.glueSession.GlueSession",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlueSession_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueSession_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.glueSession.GlueSession",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlueSession_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueSession_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.glueSession.GlueSession",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueSession_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.glueSession.GlueSession",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlueSession) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GlueSession) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GlueSession) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GlueSession) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := g.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GlueSession) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GlueSession) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GlueSession) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueSession) PutCommand(value *GlueSessionCommand) {
	if err := g.validatePutCommandParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCommand",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueSession) PutConnections(value *GlueSessionConnections) {
	if err := g.validatePutConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConnections",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueSession) PutTags(value interface{}) {
	if err := g.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTags",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueSession) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := g.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (g *jsiiProxy_GlueSession) ResetConnections() {
	_jsii_.InvokeVoid(
		g,
		"resetConnections",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetDefaultArguments() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultArguments",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetGlueVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetGlueVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetIdleTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetIdleTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetNumberOfWorkers() {
	_jsii_.InvokeVoid(
		g,
		"resetNumberOfWorkers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetRequestOrigin() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestOrigin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) ResetWorkerType() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkerType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueSession) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueSession) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		g,
		"with",
		args,
		&returns,
	)

	return returns
}

