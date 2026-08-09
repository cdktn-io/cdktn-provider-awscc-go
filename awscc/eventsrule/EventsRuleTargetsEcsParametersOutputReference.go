// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventsrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/eventsrule/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EventsRuleTargetsEcsParametersOutputReference interface {
	cdktn.ComplexObject
	CapacityProviderStrategy() EventsRuleTargetsEcsParametersCapacityProviderStrategyList
	CapacityProviderStrategyInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableEcsManagedTags() interface{}
	SetEnableEcsManagedTags(val interface{})
	EnableEcsManagedTagsInput() interface{}
	EnableExecuteCommand() interface{}
	SetEnableExecuteCommand(val interface{})
	EnableExecuteCommandInput() interface{}
	// Experimental.
	Fqn() *string
	Group() *string
	SetGroup(val *string)
	GroupInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LaunchType() *string
	SetLaunchType(val *string)
	LaunchTypeInput() *string
	NetworkConfiguration() EventsRuleTargetsEcsParametersNetworkConfigurationOutputReference
	NetworkConfigurationInput() interface{}
	PlacementConstraints() EventsRuleTargetsEcsParametersPlacementConstraintsList
	PlacementConstraintsInput() interface{}
	PlacementStrategies() EventsRuleTargetsEcsParametersPlacementStrategiesList
	PlacementStrategiesInput() interface{}
	PlatformVersion() *string
	SetPlatformVersion(val *string)
	PlatformVersionInput() *string
	PropagateTags() *string
	SetPropagateTags(val *string)
	PropagateTagsInput() *string
	ReferenceId() *string
	SetReferenceId(val *string)
	ReferenceIdInput() *string
	TagList() EventsRuleTargetsEcsParametersTagListStructList
	TagListInput() interface{}
	TaskCount() *float64
	SetTaskCount(val *float64)
	TaskCountInput() *float64
	TaskDefinitionArn() *string
	SetTaskDefinitionArn(val *string)
	TaskDefinitionArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutCapacityProviderStrategy(value interface{})
	PutNetworkConfiguration(value *EventsRuleTargetsEcsParametersNetworkConfiguration)
	PutPlacementConstraints(value interface{})
	PutPlacementStrategies(value interface{})
	PutTagList(value interface{})
	ResetCapacityProviderStrategy()
	ResetEnableEcsManagedTags()
	ResetEnableExecuteCommand()
	ResetGroup()
	ResetLaunchType()
	ResetNetworkConfiguration()
	ResetPlacementConstraints()
	ResetPlacementStrategies()
	ResetPlatformVersion()
	ResetPropagateTags()
	ResetReferenceId()
	ResetTagList()
	ResetTaskCount()
	ResetTaskDefinitionArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventsRuleTargetsEcsParametersOutputReference
type jsiiProxy_EventsRuleTargetsEcsParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) CapacityProviderStrategy() EventsRuleTargetsEcsParametersCapacityProviderStrategyList {
	var returns EventsRuleTargetsEcsParametersCapacityProviderStrategyList
	_jsii_.Get(
		j,
		"capacityProviderStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) CapacityProviderStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"capacityProviderStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) EnableEcsManagedTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) EnableEcsManagedTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableEcsManagedTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) EnableExecuteCommand() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) EnableExecuteCommandInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableExecuteCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) Group() *string {
	var returns *string
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) LaunchType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) LaunchTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"launchTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) NetworkConfiguration() EventsRuleTargetsEcsParametersNetworkConfigurationOutputReference {
	var returns EventsRuleTargetsEcsParametersNetworkConfigurationOutputReference
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PlacementConstraints() EventsRuleTargetsEcsParametersPlacementConstraintsList {
	var returns EventsRuleTargetsEcsParametersPlacementConstraintsList
	_jsii_.Get(
		j,
		"placementConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PlacementConstraintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PlacementStrategies() EventsRuleTargetsEcsParametersPlacementStrategiesList {
	var returns EventsRuleTargetsEcsParametersPlacementStrategiesList
	_jsii_.Get(
		j,
		"placementStrategies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PlacementStrategiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"placementStrategiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PlatformVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PlatformVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ReferenceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ReferenceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TagList() EventsRuleTargetsEcsParametersTagListStructList {
	var returns EventsRuleTargetsEcsParametersTagListStructList
	_jsii_.Get(
		j,
		"tagList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TagListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"taskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TaskDefinitionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TaskDefinitionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskDefinitionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventsRuleTargetsEcsParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EventsRuleTargetsEcsParametersOutputReference {
	_init_.Initialize()

	if err := validateNewEventsRuleTargetsEcsParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventsRuleTargetsEcsParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsRule.EventsRuleTargetsEcsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventsRuleTargetsEcsParametersOutputReference_Override(e EventsRuleTargetsEcsParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eventsRule.EventsRuleTargetsEcsParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetEnableEcsManagedTags(val interface{}) {
	if err := j.validateSetEnableEcsManagedTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableEcsManagedTags",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetEnableExecuteCommand(val interface{}) {
	if err := j.validateSetEnableExecuteCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableExecuteCommand",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetGroup(val *string) {
	if err := j.validateSetGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"group",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetLaunchType(val *string) {
	if err := j.validateSetLaunchTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"launchType",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetPlatformVersion(val *string) {
	if err := j.validateSetPlatformVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformVersion",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetReferenceId(val *string) {
	if err := j.validateSetReferenceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceId",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetTaskCount(val *float64) {
	if err := j.validateSetTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskCount",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetTaskDefinitionArn(val *string) {
	if err := j.validateSetTaskDefinitionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taskDefinitionArn",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PutCapacityProviderStrategy(value interface{}) {
	if err := e.validatePutCapacityProviderStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putCapacityProviderStrategy",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PutNetworkConfiguration(value *EventsRuleTargetsEcsParametersNetworkConfiguration) {
	if err := e.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PutPlacementConstraints(value interface{}) {
	if err := e.validatePutPlacementConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacementConstraints",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PutPlacementStrategies(value interface{}) {
	if err := e.validatePutPlacementStrategiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putPlacementStrategies",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) PutTagList(value interface{}) {
	if err := e.validatePutTagListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTagList",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetCapacityProviderStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityProviderStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetEnableEcsManagedTags() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableEcsManagedTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetEnableExecuteCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetEnableExecuteCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetGroup() {
	_jsii_.InvokeVoid(
		e,
		"resetGroup",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetLaunchType() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetPlacementConstraints() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementConstraints",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetPlacementStrategies() {
	_jsii_.InvokeVoid(
		e,
		"resetPlacementStrategies",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetPlatformVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetPlatformVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		e,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetReferenceId() {
	_jsii_.InvokeVoid(
		e,
		"resetReferenceId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetTagList() {
	_jsii_.InvokeVoid(
		e,
		"resetTagList",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetTaskCount() {
	_jsii_.InvokeVoid(
		e,
		"resetTaskCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ResetTaskDefinitionArn() {
	_jsii_.InvokeVoid(
		e,
		"resetTaskDefinitionArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventsRuleTargetsEcsParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

