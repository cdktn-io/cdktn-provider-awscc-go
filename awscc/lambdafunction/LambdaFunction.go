// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdafunction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lambdafunction/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function awscc_lambda_function}.
type LambdaFunction interface {
	cdktn.TerraformResource
	Architectures() *[]*string
	SetArchitectures(val *[]*string)
	ArchitecturesInput() *[]*string
	Arn() *string
	CapacityProviderConfig() LambdaFunctionCapacityProviderConfigOutputReference
	CapacityProviderConfigInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	Code() LambdaFunctionCodeOutputReference
	CodeInput() interface{}
	CodeSigningConfigArn() *string
	SetCodeSigningConfigArn(val *string)
	CodeSigningConfigArnInput() *string
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
	DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference
	DeadLetterConfigInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DurableConfig() LambdaFunctionDurableConfigOutputReference
	DurableConfigInput() interface{}
	Environment() LambdaFunctionEnvironmentOutputReference
	EnvironmentInput() interface{}
	EphemeralStorage() LambdaFunctionEphemeralStorageOutputReference
	EphemeralStorageInput() interface{}
	FileSystemConfigs() LambdaFunctionFileSystemConfigsList
	FileSystemConfigsInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	FunctionScalingConfig() LambdaFunctionFunctionScalingConfigOutputReference
	FunctionScalingConfigInput() interface{}
	Handler() *string
	SetHandler(val *string)
	HandlerInput() *string
	Id() *string
	ImageConfig() LambdaFunctionImageConfigOutputReference
	ImageConfigInput() interface{}
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	Layers() *[]*string
	SetLayers(val *[]*string)
	LayersInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	LoggingConfig() LambdaFunctionLoggingConfigOutputReference
	LoggingConfigInput() interface{}
	MemorySize() *float64
	SetMemorySize(val *float64)
	MemorySizeInput() *float64
	// The tree node.
	Node() constructs.Node
	PackageType() *string
	SetPackageType(val *string)
	PackageTypeInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublishToLatestPublished() interface{}
	SetPublishToLatestPublished(val interface{})
	PublishToLatestPublishedInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	RecursiveLoop() *string
	SetRecursiveLoop(val *string)
	RecursiveLoopInput() *string
	ReservedConcurrentExecutions() *float64
	SetReservedConcurrentExecutions(val *float64)
	ReservedConcurrentExecutionsInput() *float64
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Runtime() *string
	SetRuntime(val *string)
	RuntimeInput() *string
	RuntimeManagementConfig() LambdaFunctionRuntimeManagementConfigOutputReference
	RuntimeManagementConfigInput() interface{}
	SnapStart() LambdaFunctionSnapStartOutputReference
	SnapStartInput() interface{}
	SnapStartResponse() LambdaFunctionSnapStartResponseOutputReference
	Tags() LambdaFunctionTagsList
	TagsInput() interface{}
	TenancyConfig() LambdaFunctionTenancyConfigOutputReference
	TenancyConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
	TracingConfig() LambdaFunctionTracingConfigOutputReference
	TracingConfigInput() interface{}
	VpcConfig() LambdaFunctionVpcConfigOutputReference
	VpcConfigInput() interface{}
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
	PutCapacityProviderConfig(value *LambdaFunctionCapacityProviderConfig)
	PutCode(value *LambdaFunctionCode)
	PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig)
	PutDurableConfig(value *LambdaFunctionDurableConfig)
	PutEnvironment(value *LambdaFunctionEnvironment)
	PutEphemeralStorage(value *LambdaFunctionEphemeralStorage)
	PutFileSystemConfigs(value interface{})
	PutFunctionScalingConfig(value *LambdaFunctionFunctionScalingConfig)
	PutImageConfig(value *LambdaFunctionImageConfig)
	PutLoggingConfig(value *LambdaFunctionLoggingConfig)
	PutRuntimeManagementConfig(value *LambdaFunctionRuntimeManagementConfig)
	PutSnapStart(value *LambdaFunctionSnapStart)
	PutTags(value interface{})
	PutTenancyConfig(value *LambdaFunctionTenancyConfig)
	PutTracingConfig(value *LambdaFunctionTracingConfig)
	PutVpcConfig(value *LambdaFunctionVpcConfig)
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
	ResetArchitectures()
	ResetCapacityProviderConfig()
	ResetCodeSigningConfigArn()
	ResetDeadLetterConfig()
	ResetDescription()
	ResetDurableConfig()
	ResetEnvironment()
	ResetEphemeralStorage()
	ResetFileSystemConfigs()
	ResetFunctionName()
	ResetFunctionScalingConfig()
	ResetHandler()
	ResetImageConfig()
	ResetKmsKeyArn()
	ResetLayers()
	ResetLoggingConfig()
	ResetMemorySize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackageType()
	ResetPublishToLatestPublished()
	ResetRecursiveLoop()
	ResetReservedConcurrentExecutions()
	ResetRuntime()
	ResetRuntimeManagementConfig()
	ResetSnapStart()
	ResetTags()
	ResetTenancyConfig()
	ResetTimeout()
	ResetTracingConfig()
	ResetVpcConfig()
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

// The jsii proxy struct for LambdaFunction
type jsiiProxy_LambdaFunction struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_LambdaFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CapacityProviderConfig() LambdaFunctionCapacityProviderConfigOutputReference {
	var returns LambdaFunctionCapacityProviderConfigOutputReference
	_jsii_.Get(
		j,
		"capacityProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CapacityProviderConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Code() LambdaFunctionCodeOutputReference {
	var returns LambdaFunctionCodeOutputReference
	_jsii_.Get(
		j,
		"code",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"codeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) CodeSigningConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfig() LambdaFunctionDeadLetterConfigOutputReference {
	var returns LambdaFunctionDeadLetterConfigOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DeadLetterConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DurableConfig() LambdaFunctionDurableConfigOutputReference {
	var returns LambdaFunctionDurableConfigOutputReference
	_jsii_.Get(
		j,
		"durableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) DurableConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"durableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Environment() LambdaFunctionEnvironmentOutputReference {
	var returns LambdaFunctionEnvironmentOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EphemeralStorage() LambdaFunctionEphemeralStorageOutputReference {
	var returns LambdaFunctionEphemeralStorageOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) EphemeralStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfigs() LambdaFunctionFileSystemConfigsList {
	var returns LambdaFunctionFileSystemConfigsList
	_jsii_.Get(
		j,
		"fileSystemConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FileSystemConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileSystemConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionScalingConfig() LambdaFunctionFunctionScalingConfigOutputReference {
	var returns LambdaFunctionFunctionScalingConfigOutputReference
	_jsii_.Get(
		j,
		"functionScalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) FunctionScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionScalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfig() LambdaFunctionImageConfigOutputReference {
	var returns LambdaFunctionImageConfigOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ImageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LayersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LoggingConfig() LambdaFunctionLoggingConfigOutputReference {
	var returns LambdaFunctionLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) LoggingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) MemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PackageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PublishToLatestPublished() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishToLatestPublished",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) PublishToLatestPublishedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishToLatestPublishedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RecursiveLoop() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recursiveLoop",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RecursiveLoopInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recursiveLoopInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) ReservedConcurrentExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeManagementConfig() LambdaFunctionRuntimeManagementConfigOutputReference {
	var returns LambdaFunctionRuntimeManagementConfigOutputReference
	_jsii_.Get(
		j,
		"runtimeManagementConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) RuntimeManagementConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeManagementConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStart() LambdaFunctionSnapStartOutputReference {
	var returns LambdaFunctionSnapStartOutputReference
	_jsii_.Get(
		j,
		"snapStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) SnapStartResponse() LambdaFunctionSnapStartResponseOutputReference {
	var returns LambdaFunctionSnapStartResponseOutputReference
	_jsii_.Get(
		j,
		"snapStartResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Tags() LambdaFunctionTagsList {
	var returns LambdaFunctionTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TenancyConfig() LambdaFunctionTenancyConfigOutputReference {
	var returns LambdaFunctionTenancyConfigOutputReference
	_jsii_.Get(
		j,
		"tenancyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TenancyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tenancyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfig() LambdaFunctionTracingConfigOutputReference {
	var returns LambdaFunctionTracingConfigOutputReference
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) TracingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfig() LambdaFunctionVpcConfigOutputReference {
	var returns LambdaFunctionVpcConfigOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunction) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function awscc_lambda_function} Resource.
func NewLambdaFunction(scope constructs.Construct, id *string, config *LambdaFunctionConfig) LambdaFunction {
	_init_.Initialize()

	if err := validateNewLambdaFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunction{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaFunction.LambdaFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/lambda_function awscc_lambda_function} Resource.
func NewLambdaFunction_Override(l LambdaFunction, scope constructs.Construct, id *string, config *LambdaFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaFunction.LambdaFunction",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaFunction)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetCodeSigningConfigArn(val *string) {
	if err := j.validateSetCodeSigningConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetLayers(val *[]*string) {
	if err := j.validateSetLayersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetMemorySize(val *float64) {
	if err := j.validateSetMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetPackageType(val *string) {
	if err := j.validateSetPackageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetPublishToLatestPublished(val interface{}) {
	if err := j.validateSetPublishToLatestPublishedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishToLatestPublished",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetRecursiveLoop(val *string) {
	if err := j.validateSetRecursiveLoopParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recursiveLoop",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetReservedConcurrentExecutions(val *float64) {
	if err := j.validateSetReservedConcurrentExecutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_LambdaFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

// Generates CDKTN code for importing a LambdaFunction resource upon running "cdktn plan <stack-name>".
func LambdaFunction_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateLambdaFunction_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.lambdaFunction.LambdaFunction",
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
func LambdaFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.lambdaFunction.LambdaFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.lambdaFunction.LambdaFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-awscc.lambdaFunction.LambdaFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-awscc.lambdaFunction.LambdaFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaFunction) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LambdaFunction) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LambdaFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LambdaFunction) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := l.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LambdaFunction) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LambdaFunction) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LambdaFunction) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaFunction) PutCapacityProviderConfig(value *LambdaFunctionCapacityProviderConfig) {
	if err := l.validatePutCapacityProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCapacityProviderConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutCode(value *LambdaFunctionCode) {
	if err := l.validatePutCodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCode",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutDeadLetterConfig(value *LambdaFunctionDeadLetterConfig) {
	if err := l.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutDurableConfig(value *LambdaFunctionDurableConfig) {
	if err := l.validatePutDurableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDurableConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutEnvironment(value *LambdaFunctionEnvironment) {
	if err := l.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutEphemeralStorage(value *LambdaFunctionEphemeralStorage) {
	if err := l.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutFileSystemConfigs(value interface{}) {
	if err := l.validatePutFileSystemConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFileSystemConfigs",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutFunctionScalingConfig(value *LambdaFunctionFunctionScalingConfig) {
	if err := l.validatePutFunctionScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFunctionScalingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutImageConfig(value *LambdaFunctionImageConfig) {
	if err := l.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutLoggingConfig(value *LambdaFunctionLoggingConfig) {
	if err := l.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutRuntimeManagementConfig(value *LambdaFunctionRuntimeManagementConfig) {
	if err := l.validatePutRuntimeManagementConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRuntimeManagementConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutSnapStart(value *LambdaFunctionSnapStart) {
	if err := l.validatePutSnapStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSnapStart",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTags(value interface{}) {
	if err := l.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTags",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTenancyConfig(value *LambdaFunctionTenancyConfig) {
	if err := l.validatePutTenancyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTenancyConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutTracingConfig(value *LambdaFunctionTracingConfig) {
	if err := l.validatePutTracingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTracingConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) PutVpcConfig(value *LambdaFunctionVpcConfig) {
	if err := l.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LambdaFunction) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := l.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (l *jsiiProxy_LambdaFunction) ResetArchitectures() {
	_jsii_.InvokeVoid(
		l,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetCapacityProviderConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetCapacityProviderConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetCodeSigningConfigArn() {
	_jsii_.InvokeVoid(
		l,
		"resetCodeSigningConfigArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetDurableConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetDurableConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetEnvironment() {
	_jsii_.InvokeVoid(
		l,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		l,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFileSystemConfigs() {
	_jsii_.InvokeVoid(
		l,
		"resetFileSystemConfigs",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFunctionName() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionName",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetFunctionScalingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionScalingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetHandler() {
	_jsii_.InvokeVoid(
		l,
		"resetHandler",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetImageConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		l,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetLayers() {
	_jsii_.InvokeVoid(
		l,
		"resetLayers",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetMemorySize() {
	_jsii_.InvokeVoid(
		l,
		"resetMemorySize",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetPackageType() {
	_jsii_.InvokeVoid(
		l,
		"resetPackageType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetPublishToLatestPublished() {
	_jsii_.InvokeVoid(
		l,
		"resetPublishToLatestPublished",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetRecursiveLoop() {
	_jsii_.InvokeVoid(
		l,
		"resetRecursiveLoop",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetReservedConcurrentExecutions() {
	_jsii_.InvokeVoid(
		l,
		"resetReservedConcurrentExecutions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetRuntime() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntime",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetRuntimeManagementConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntimeManagementConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetSnapStart() {
	_jsii_.InvokeVoid(
		l,
		"resetSnapStart",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTenancyConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetTenancyConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeout",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		l,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunction) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		l,
		"with",
		args,
		&returns,
	)

	return returns
}

