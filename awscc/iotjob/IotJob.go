// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/iotjob/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job awscc_iot_job}.
type IotJob interface {
	cdktn.TerraformResource
	AbortConfig() IotJobAbortConfigOutputReference
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
	CreatedAt() *string
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
	DocumentParameters() *map[string]*string
	SetDocumentParameters(val *map[string]*string)
	DocumentParametersInput() *map[string]*string
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
	JobExecutionsRetryConfig() IotJobJobExecutionsRetryConfigOutputReference
	JobExecutionsRetryConfigInput() interface{}
	JobExecutionsRolloutConfig() IotJobJobExecutionsRolloutConfigOutputReference
	JobExecutionsRolloutConfigInput() interface{}
	JobId() *string
	SetJobId(val *string)
	JobIdInput() *string
	JobTemplateArn() *string
	SetJobTemplateArn(val *string)
	JobTemplateArnInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PresignedUrlConfig() IotJobPresignedUrlConfigOutputReference
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
	SchedulingConfig() IotJobSchedulingConfigOutputReference
	SchedulingConfigInput() interface{}
	Tags() IotJobTagsList
	TagsInput() interface{}
	Targets() *[]*string
	SetTargets(val *[]*string)
	TargetSelection() *string
	SetTargetSelection(val *string)
	TargetSelectionInput() *string
	TargetsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeoutConfig() IotJobTimeoutConfigOutputReference
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
	PutAbortConfig(value *IotJobAbortConfig)
	PutJobExecutionsRetryConfig(value *IotJobJobExecutionsRetryConfig)
	PutJobExecutionsRolloutConfig(value *IotJobJobExecutionsRolloutConfig)
	PutPresignedUrlConfig(value *IotJobPresignedUrlConfig)
	PutSchedulingConfig(value *IotJobSchedulingConfig)
	PutTags(value interface{})
	PutTimeoutConfig(value *IotJobTimeoutConfig)
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
	ResetDescription()
	ResetDestinationPackageVersions()
	ResetDocument()
	ResetDocumentParameters()
	ResetDocumentSource()
	ResetJobExecutionsRetryConfig()
	ResetJobExecutionsRolloutConfig()
	ResetJobTemplateArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPresignedUrlConfig()
	ResetSchedulingConfig()
	ResetTags()
	ResetTargetSelection()
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

// The jsii proxy struct for IotJob
type jsiiProxy_IotJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_IotJob) AbortConfig() IotJobAbortConfigOutputReference {
	var returns IotJobAbortConfigOutputReference
	_jsii_.Get(
		j,
		"abortConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) AbortConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DestinationPackageVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPackageVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DestinationPackageVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"destinationPackageVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Document() *string {
	var returns *string
	_jsii_.Get(
		j,
		"document",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DocumentParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"documentParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DocumentParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"documentParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DocumentSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) DocumentSourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobExecutionsRetryConfig() IotJobJobExecutionsRetryConfigOutputReference {
	var returns IotJobJobExecutionsRetryConfigOutputReference
	_jsii_.Get(
		j,
		"jobExecutionsRetryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobExecutionsRetryConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobExecutionsRetryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobExecutionsRolloutConfig() IotJobJobExecutionsRolloutConfigOutputReference {
	var returns IotJobJobExecutionsRolloutConfigOutputReference
	_jsii_.Get(
		j,
		"jobExecutionsRolloutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobExecutionsRolloutConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobExecutionsRolloutConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobTemplateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTemplateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) JobTemplateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTemplateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) PresignedUrlConfig() IotJobPresignedUrlConfigOutputReference {
	var returns IotJobPresignedUrlConfigOutputReference
	_jsii_.Get(
		j,
		"presignedUrlConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) PresignedUrlConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"presignedUrlConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) SchedulingConfig() IotJobSchedulingConfigOutputReference {
	var returns IotJobSchedulingConfigOutputReference
	_jsii_.Get(
		j,
		"schedulingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) SchedulingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schedulingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Tags() IotJobTagsList {
	var returns IotJobTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) Targets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TargetSelection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetSelection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TargetSelectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetSelectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TimeoutConfig() IotJobTimeoutConfigOutputReference {
	var returns IotJobTimeoutConfigOutputReference
	_jsii_.Get(
		j,
		"timeoutConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotJob) TimeoutConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job awscc_iot_job} Resource.
func NewIotJob(scope constructs.Construct, id *string, config *IotJobConfig) IotJob {
	_init_.Initialize()

	if err := validateNewIotJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotJob{}

	_jsii_.Create(
		"@cdktn/provider-awscc.iotJob.IotJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iot_job awscc_iot_job} Resource.
func NewIotJob_Override(i IotJob, scope constructs.Construct, id *string, config *IotJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.iotJob.IotJob",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetDestinationPackageVersions(val *[]*string) {
	if err := j.validateSetDestinationPackageVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPackageVersions",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetDocument(val *string) {
	if err := j.validateSetDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"document",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetDocumentParameters(val *map[string]*string) {
	if err := j.validateSetDocumentParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentParameters",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetDocumentSource(val *string) {
	if err := j.validateSetDocumentSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentSource",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetJobId(val *string) {
	if err := j.validateSetJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobId",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetJobTemplateArn(val *string) {
	if err := j.validateSetJobTemplateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTemplateArn",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetTargets(val *[]*string) {
	if err := j.validateSetTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targets",
		val,
	)
}

func (j *jsiiProxy_IotJob)SetTargetSelection(val *string) {
	if err := j.validateSetTargetSelectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSelection",
		val,
	)
}

// Generates CDKTN code for importing a IotJob resource upon running "cdktn plan <stack-name>".
func IotJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateIotJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJob.IotJob",
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
func IotJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJob.IotJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJob.IotJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.iotJob.IotJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.iotJob.IotJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotJob) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotJob) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IotJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IotJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IotJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IotJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IotJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IotJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IotJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IotJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (i *jsiiProxy_IotJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (i *jsiiProxy_IotJob) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotJob) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotJob) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IotJob) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotJob) PutAbortConfig(value *IotJobAbortConfig) {
	if err := i.validatePutAbortConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAbortConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJob) PutJobExecutionsRetryConfig(value *IotJobJobExecutionsRetryConfig) {
	if err := i.validatePutJobExecutionsRetryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putJobExecutionsRetryConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJob) PutJobExecutionsRolloutConfig(value *IotJobJobExecutionsRolloutConfig) {
	if err := i.validatePutJobExecutionsRolloutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putJobExecutionsRolloutConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJob) PutPresignedUrlConfig(value *IotJobPresignedUrlConfig) {
	if err := i.validatePutPresignedUrlConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPresignedUrlConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJob) PutSchedulingConfig(value *IotJobSchedulingConfig) {
	if err := i.validatePutSchedulingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putSchedulingConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJob) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJob) PutTimeoutConfig(value *IotJobTimeoutConfig) {
	if err := i.validatePutTimeoutConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeoutConfig",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := i.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (i *jsiiProxy_IotJob) ResetAbortConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetAbortConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetDestinationPackageVersions() {
	_jsii_.InvokeVoid(
		i,
		"resetDestinationPackageVersions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetDocument() {
	_jsii_.InvokeVoid(
		i,
		"resetDocument",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetDocumentParameters() {
	_jsii_.InvokeVoid(
		i,
		"resetDocumentParameters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetDocumentSource() {
	_jsii_.InvokeVoid(
		i,
		"resetDocumentSource",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetJobExecutionsRetryConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetJobExecutionsRetryConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetJobExecutionsRolloutConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetJobExecutionsRolloutConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetJobTemplateArn() {
	_jsii_.InvokeVoid(
		i,
		"resetJobTemplateArn",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetPresignedUrlConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetPresignedUrlConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetSchedulingConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetSchedulingConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetTargetSelection() {
	_jsii_.InvokeVoid(
		i,
		"resetTargetSelection",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) ResetTimeoutConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeoutConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

