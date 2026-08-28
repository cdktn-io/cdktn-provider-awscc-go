// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecsdaemontaskdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ecsdaemontaskdefinition/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference interface {
	cdktn.ComplexObject
	Command() *[]*string
	SetCommand(val *[]*string)
	CommandInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Interval() *float64
	SetInterval(val *float64)
	IntervalInput() *float64
	Retries() *float64
	SetRetries(val *float64)
	RetriesInput() *float64
	StartPeriod() *float64
	SetStartPeriod(val *float64)
	StartPeriodInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Timeout() *float64
	SetTimeout(val *float64)
	TimeoutInput() *float64
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
	ResetCommand()
	ResetInterval()
	ResetRetries()
	ResetStartPeriod()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference
type jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) Command() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) CommandInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) Interval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"interval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) IntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) StartPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) StartPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"startPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewEcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference {
	_init_.Initialize()

	if err := validateNewEcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference_Override(e EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.ecsDaemonTaskDefinition.EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetCommand(val *[]*string) {
	if err := j.validateSetCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"command",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetInterval(val *float64) {
	if err := j.validateSetIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"interval",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetStartPeriod(val *float64) {
	if err := j.validateSetStartPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startPeriod",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ResetCommand() {
	_jsii_.InvokeVoid(
		e,
		"resetCommand",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ResetInterval() {
	_jsii_.InvokeVoid(
		e,
		"resetInterval",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ResetRetries() {
	_jsii_.InvokeVoid(
		e,
		"resetRetries",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ResetStartPeriod() {
	_jsii_.InvokeVoid(
		e,
		"resetStartPeriod",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		e,
		"resetTimeout",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EcsDaemonTaskDefinitionContainerDefinitionsHealthCheckOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

