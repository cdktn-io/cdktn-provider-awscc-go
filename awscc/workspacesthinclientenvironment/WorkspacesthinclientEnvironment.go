// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesthinclientenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/workspacesthinclientenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/workspacesthinclient_environment awscc_workspacesthinclient_environment}.
type WorkspacesthinclientEnvironment interface {
	cdktn.TerraformResource
	ActivationCode() *string
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
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredSoftwareSetId() *string
	SetDesiredSoftwareSetId(val *string)
	DesiredSoftwareSetIdInput() *string
	DesktopArn() *string
	SetDesktopArn(val *string)
	DesktopArnInput() *string
	DesktopEndpoint() *string
	SetDesktopEndpoint(val *string)
	DesktopEndpointInput() *string
	DesktopType() *string
	DeviceCreationTags() WorkspacesthinclientEnvironmentDeviceCreationTagsList
	DeviceCreationTagsInput() interface{}
	EnvironmentId() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaintenanceWindow() WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference
	MaintenanceWindowInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PendingSoftwareSetId() *string
	PendingSoftwareSetVersion() *string
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
	RegisteredDevicesCount() *float64
	SoftwareSetComplianceStatus() *string
	SoftwareSetUpdateMode() *string
	SetSoftwareSetUpdateMode(val *string)
	SoftwareSetUpdateModeInput() *string
	SoftwareSetUpdateSchedule() *string
	SetSoftwareSetUpdateSchedule(val *string)
	SoftwareSetUpdateScheduleInput() *string
	Tags() WorkspacesthinclientEnvironmentTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutDeviceCreationTags(value interface{})
	PutMaintenanceWindow(value *WorkspacesthinclientEnvironmentMaintenanceWindow)
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
	ResetDesiredSoftwareSetId()
	ResetDesktopEndpoint()
	ResetDeviceCreationTags()
	ResetKmsKeyArn()
	ResetMaintenanceWindow()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSoftwareSetUpdateMode()
	ResetSoftwareSetUpdateSchedule()
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

// The jsii proxy struct for WorkspacesthinclientEnvironment
type jsiiProxy_WorkspacesthinclientEnvironment struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) ActivationCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DesiredSoftwareSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredSoftwareSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DesiredSoftwareSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredSoftwareSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DesktopArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desktopArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DesktopArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desktopArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DesktopEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desktopEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DesktopEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desktopEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DesktopType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desktopType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DeviceCreationTags() WorkspacesthinclientEnvironmentDeviceCreationTagsList {
	var returns WorkspacesthinclientEnvironmentDeviceCreationTagsList
	_jsii_.Get(
		j,
		"deviceCreationTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) DeviceCreationTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deviceCreationTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) EnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) MaintenanceWindow() WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference {
	var returns WorkspacesthinclientEnvironmentMaintenanceWindowOutputReference
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) MaintenanceWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) PendingSoftwareSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pendingSoftwareSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) PendingSoftwareSetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pendingSoftwareSetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) RegisteredDevicesCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"registeredDevicesCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) SoftwareSetComplianceStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"softwareSetComplianceStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) SoftwareSetUpdateMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"softwareSetUpdateMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) SoftwareSetUpdateModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"softwareSetUpdateModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) SoftwareSetUpdateSchedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"softwareSetUpdateSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) SoftwareSetUpdateScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"softwareSetUpdateScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) Tags() WorkspacesthinclientEnvironmentTagsList {
	var returns WorkspacesthinclientEnvironmentTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/workspacesthinclient_environment awscc_workspacesthinclient_environment} Resource.
func NewWorkspacesthinclientEnvironment(scope constructs.Construct, id *string, config *WorkspacesthinclientEnvironmentConfig) WorkspacesthinclientEnvironment {
	_init_.Initialize()

	if err := validateNewWorkspacesthinclientEnvironmentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkspacesthinclientEnvironment{}

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironment",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/workspacesthinclient_environment awscc_workspacesthinclient_environment} Resource.
func NewWorkspacesthinclientEnvironment_Override(w WorkspacesthinclientEnvironment, scope constructs.Construct, id *string, config *WorkspacesthinclientEnvironmentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironment",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetDesiredSoftwareSetId(val *string) {
	if err := j.validateSetDesiredSoftwareSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredSoftwareSetId",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetDesktopArn(val *string) {
	if err := j.validateSetDesktopArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desktopArn",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetDesktopEndpoint(val *string) {
	if err := j.validateSetDesktopEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desktopEndpoint",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetSoftwareSetUpdateMode(val *string) {
	if err := j.validateSetSoftwareSetUpdateModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softwareSetUpdateMode",
		val,
	)
}

func (j *jsiiProxy_WorkspacesthinclientEnvironment)SetSoftwareSetUpdateSchedule(val *string) {
	if err := j.validateSetSoftwareSetUpdateScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"softwareSetUpdateSchedule",
		val,
	)
}

// Generates CDKTN code for importing a WorkspacesthinclientEnvironment resource upon running "cdktn plan <stack-name>".
func WorkspacesthinclientEnvironment_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateWorkspacesthinclientEnvironment_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironment",
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
func WorkspacesthinclientEnvironment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspacesthinclientEnvironment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspacesthinclientEnvironment_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspacesthinclientEnvironment_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironment",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WorkspacesthinclientEnvironment_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWorkspacesthinclientEnvironment_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironment",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WorkspacesthinclientEnvironment_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.workspacesthinclientEnvironment.WorkspacesthinclientEnvironment",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := w.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) PutDeviceCreationTags(value interface{}) {
	if err := w.validatePutDeviceCreationTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDeviceCreationTags",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) PutMaintenanceWindow(value *WorkspacesthinclientEnvironmentMaintenanceWindow) {
	if err := w.validatePutMaintenanceWindowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putMaintenanceWindow",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) PutTags(value interface{}) {
	if err := w.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTags",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := w.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetDesiredSoftwareSetId() {
	_jsii_.InvokeVoid(
		w,
		"resetDesiredSoftwareSetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetDesktopEndpoint() {
	_jsii_.InvokeVoid(
		w,
		"resetDesktopEndpoint",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetDeviceCreationTags() {
	_jsii_.InvokeVoid(
		w,
		"resetDeviceCreationTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		w,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		w,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetSoftwareSetUpdateMode() {
	_jsii_.InvokeVoid(
		w,
		"resetSoftwareSetUpdateMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetSoftwareSetUpdateSchedule() {
	_jsii_.InvokeVoid(
		w,
		"resetSoftwareSetUpdateSchedule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WorkspacesthinclientEnvironment) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		w,
		"with",
		args,
		&returns,
	)

	return returns
}

