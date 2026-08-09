// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdamicrovmimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lambdamicrovmimage/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaMicrovmImageHooksMicrovmHooksOutputReference interface {
	cdktn.ComplexObject
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
	Resume() *string
	SetResume(val *string)
	ResumeInput() *string
	ResumeTimeoutInSeconds() *float64
	SetResumeTimeoutInSeconds(val *float64)
	ResumeTimeoutInSecondsInput() *float64
	Run() *string
	SetRun(val *string)
	RunInput() *string
	RunTimeoutInSeconds() *float64
	SetRunTimeoutInSeconds(val *float64)
	RunTimeoutInSecondsInput() *float64
	Suspend() *string
	SetSuspend(val *string)
	SuspendInput() *string
	SuspendTimeoutInSeconds() *float64
	SetSuspendTimeoutInSeconds(val *float64)
	SuspendTimeoutInSecondsInput() *float64
	Terminate() *string
	SetTerminate(val *string)
	TerminateInput() *string
	TerminateTimeoutInSeconds() *float64
	SetTerminateTimeoutInSeconds(val *float64)
	TerminateTimeoutInSecondsInput() *float64
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
	ResetResume()
	ResetResumeTimeoutInSeconds()
	ResetRun()
	ResetRunTimeoutInSeconds()
	ResetSuspend()
	ResetSuspendTimeoutInSeconds()
	ResetTerminate()
	ResetTerminateTimeoutInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaMicrovmImageHooksMicrovmHooksOutputReference
type jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) Resume() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResumeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResumeTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resumeTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResumeTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"resumeTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) Run() *string {
	var returns *string
	_jsii_.Get(
		j,
		"run",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) RunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) RunTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) RunTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) Suspend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) SuspendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suspendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) SuspendTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) SuspendTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"suspendTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) Terminate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) TerminateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terminateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) TerminateTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminateTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) TerminateTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminateTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaMicrovmImageHooksMicrovmHooksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LambdaMicrovmImageHooksMicrovmHooksOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaMicrovmImageHooksMicrovmHooksOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaMicrovmImage.LambdaMicrovmImageHooksMicrovmHooksOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLambdaMicrovmImageHooksMicrovmHooksOutputReference_Override(l LambdaMicrovmImageHooksMicrovmHooksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaMicrovmImage.LambdaMicrovmImageHooksMicrovmHooksOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetResume(val *string) {
	if err := j.validateSetResumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resume",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetResumeTimeoutInSeconds(val *float64) {
	if err := j.validateSetResumeTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resumeTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetRun(val *string) {
	if err := j.validateSetRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"run",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetRunTimeoutInSeconds(val *float64) {
	if err := j.validateSetRunTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetSuspend(val *string) {
	if err := j.validateSetSuspendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspend",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetSuspendTimeoutInSeconds(val *float64) {
	if err := j.validateSetSuspendTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suspendTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetTerminate(val *string) {
	if err := j.validateSetTerminateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminate",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetTerminateTimeoutInSeconds(val *float64) {
	if err := j.validateSetTerminateTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetResume() {
	_jsii_.InvokeVoid(
		l,
		"resetResume",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetResumeTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetResumeTimeoutInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetRun() {
	_jsii_.InvokeVoid(
		l,
		"resetRun",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetRunTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetRunTimeoutInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetSuspend() {
	_jsii_.InvokeVoid(
		l,
		"resetSuspend",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetSuspendTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetSuspendTimeoutInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetTerminate() {
	_jsii_.InvokeVoid(
		l,
		"resetTerminate",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ResetTerminateTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetTerminateTimeoutInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := l.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaMicrovmImageHooksMicrovmHooksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

