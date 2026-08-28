// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotjobtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_job_template awscc_iot_job_template}.
type IotJobTemplate interface {
	cdktn.TerraformResource
	AbortConfig() IotJobTemplateAbortConfigOutputReference
	AbortConfigInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DestinationPackageVersions() *[]*string
	SetDestinationPackageVersions(val *[]*string)
	DestinationPackageVersionsInput() *[]*string
	Document() *string
	SetDocument(val *string)
	DocumentInput() *string
	DocumentSource() *string
	SetDocumentSource(val *string)
	DocumentSourceInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	JobArn() *string
	SetJobArn(val *string)
	JobArnInput() *string
	JobExecutionsRetryConfig() IotJobTemplateJobExecutionsRetryConfigOutputReference
	JobExecutionsRetryConfigInput() interface{}
	JobExecutionsRolloutConfig() IotJobTemplateJobExecutionsRolloutConfigOutputReference
	JobExecutionsRolloutConfigInput() interface{}
	JobTemplateId() *string
	SetJobTemplateId(val *string)
	JobTemplateIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	MaintenanceWindows() IotJobTemplateMaintenanceWindowsList
	MaintenanceWindowsInput() interface{}
	// The tree node.
	Node() constructs.Node
	PresignedUrlConfig() IotJobTemplatePresignedUrlConfigOutputReference
	PresignedUrlConfigInput() interface{}
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
	Tags() IotJobTemplateTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutConfig() IotJobTemplateTimeoutConfigOutputReference
	TimeoutConfigInput() interface{}
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
	PutAbortConfig(value *IotJobTemplateAbortConfig)
	PutJobExecutionsRetryConfig(value *IotJobTemplateJobExecutionsRetryConfig)
	PutJobExecutionsRolloutConfig(value *IotJobTemplateJobExecutionsRolloutConfig)
	PutMaintenanceWindows(value interface{})
	PutPresignedUrlConfig(value *IotJobTemplatePresignedUrlConfig)
	PutTags(value interface{})
	PutTimeoutConfig(value *IotJobTemplateTimeoutConfig)
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
	ResetAbortConfig()
	ResetDestinationPackageVersions()
	ResetDocument()
	ResetDocumentSource()
	ResetJobArn()
	ResetJobExecutionsRetryConfig()
	ResetJobExecutionsRolloutConfig()
	ResetMaintenanceWindows()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPresignedUrlConfig()
	ResetTags()
	ResetTimeoutConfig()
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

// The jsii proxy struct for IotJobTemplate
type jsiiProxy_IotJobTemplate struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IotJobTemplate) AbortConfig() IotJobTemplateAbortConfigOutputReference {
	var returns IotJobTemplateAbortConfigOutputReference
	_jsii_.Get(
		j,
		"abortConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) AbortConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) DestinationPackageVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPackageVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) DestinationPackageVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPackageVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Document() *string {
	var returns *string
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) DocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) DocumentSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) DocumentSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobExecutionsRetryConfig() IotJobTemplateJobExecutionsRetryConfigOutputReference {
	var returns IotJobTemplateJobExecutionsRetryConfigOutputReference
	_jsii_.Get(
		j,
		"jobExecutionsRetryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobExecutionsRetryConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobExecutionsRetryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobExecutionsRolloutConfig() IotJobTemplateJobExecutionsRolloutConfigOutputReference {
	var returns IotJobTemplateJobExecutionsRolloutConfigOutputReference
	_jsii_.Get(
		j,
		"jobExecutionsRolloutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobExecutionsRolloutConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobExecutionsRolloutConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) JobTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) MaintenanceWindows() IotJobTemplateMaintenanceWindowsList {
	var returns IotJobTemplateMaintenanceWindowsList
	_jsii_.Get(
		j,
		"maintenanceWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) MaintenanceWindowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) PresignedUrlConfig() IotJobTemplatePresignedUrlConfigOutputReference {
	var returns IotJobTemplatePresignedUrlConfigOutputReference
	_jsii_.Get(
		j,
		"presignedUrlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) PresignedUrlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presignedUrlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) Tags() IotJobTemplateTagsList {
	var returns IotJobTemplateTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) TimeoutConfig() IotJobTemplateTimeoutConfigOutputReference {
	var returns IotJobTemplateTimeoutConfigOutputReference
	_jsii_.Get(
		j,
		"timeoutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJobTemplate) TimeoutConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_job_template awscc_iot_job_template} Resource.
func NewIotJobTemplate(scope constructs.Construct, id *string, config *IotJobTemplateConfig) IotJobTemplate {
	_init_.Initialize()

	if err := validateNewIotJobTemplateParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotJobTemplate{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplate",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_job_template awscc_iot_job_template} Resource.
func NewIotJobTemplate_Override(i IotJobTemplate, scope constructs.Construct, id *string, config *IotJobTemplateConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplate",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetDestinationPackageVersions(val *[]*string) {
	if err := j.validateSetDestinationPackageVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPackageVersions",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetDocument(val *string) {
	if err := j.validateSetDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"document",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetDocumentSource(val *string) {
	if err := j.validateSetDocumentSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentSource",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetJobArn(val *string) {
	if err := j.validateSetJobArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobArn",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetJobTemplateId(val *string) {
	if err := j.validateSetJobTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTemplateId",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotJobTemplate)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTN code for importing a IotJobTemplate resource upon running "cdktn plan <stack-name>".
func IotJobTemplate_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIotJobTemplate_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplate",
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
func IotJobTemplate_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotJobTemplate_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplate",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotJobTemplate_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotJobTemplate_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplate",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotJobTemplate_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotJobTemplate_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplate",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotJobTemplate_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.iotJobTemplate.IotJobTemplate",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotJobTemplate) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotJobTemplate) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotJobTemplate) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotJobTemplate) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := i.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotJobTemplate) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotJobTemplate) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotJobTemplate) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotJobTemplate) PutAbortConfig(value *IotJobTemplateAbortConfig) {
	if err := i.validatePutAbortConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAbortConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJobTemplate) PutJobExecutionsRetryConfig(value *IotJobTemplateJobExecutionsRetryConfig) {
	if err := i.validatePutJobExecutionsRetryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putJobExecutionsRetryConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJobTemplate) PutJobExecutionsRolloutConfig(value *IotJobTemplateJobExecutionsRolloutConfig) {
	if err := i.validatePutJobExecutionsRolloutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putJobExecutionsRolloutConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJobTemplate) PutMaintenanceWindows(value interface{}) {
	if err := i.validatePutMaintenanceWindowsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putMaintenanceWindows",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJobTemplate) PutPresignedUrlConfig(value *IotJobTemplatePresignedUrlConfig) {
	if err := i.validatePutPresignedUrlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPresignedUrlConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJobTemplate) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJobTemplate) PutTimeoutConfig(value *IotJobTemplateTimeoutConfig) {
	if err := i.validatePutTimeoutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeoutConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJobTemplate) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := i.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetAbortConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAbortConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetDestinationPackageVersions() {
	_jsii_.InvokeVoid(
		i,
		"resetDestinationPackageVersions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetDocument() {
	_jsii_.InvokeVoid(
		i,
		"resetDocument",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetDocumentSource() {
	_jsii_.InvokeVoid(
		i,
		"resetDocumentSource",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetJobArn() {
	_jsii_.InvokeVoid(
		i,
		"resetJobArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetJobExecutionsRetryConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetJobExecutionsRetryConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetJobExecutionsRolloutConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetJobExecutionsRolloutConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetMaintenanceWindows() {
	_jsii_.InvokeVoid(
		i,
		"resetMaintenanceWindows",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetPresignedUrlConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetPresignedUrlConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) ResetTimeoutConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeoutConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJobTemplate) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJobTemplate) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

