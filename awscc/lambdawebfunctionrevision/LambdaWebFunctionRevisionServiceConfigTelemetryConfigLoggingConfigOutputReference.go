// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdawebfunctionrevision

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/lambdawebfunctionrevision/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference interface {
	cdktn.ComplexObject
	ApplicationLogLevel() *string
	SetApplicationLogLevel(val *string)
	ApplicationLogLevelInput() *string
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
	LogGroup() *string
	SetLogGroup(val *string)
	LogGroupInput() *string
	SystemLogLevel() *string
	SetSystemLogLevel(val *string)
	SystemLogLevelInput() *string
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
	ResetApplicationLogLevel()
	ResetLogGroup()
	ResetSystemLogLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference
type jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ApplicationLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ApplicationLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) LogGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) LogGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) SystemLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) SystemLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"systemLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewLambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaWebFunctionRevision.LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference_Override(l LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.lambdaWebFunctionRevision.LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetApplicationLogLevel(val *string) {
	if err := j.validateSetApplicationLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationLogLevel",
		val,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetLogGroup(val *string) {
	if err := j.validateSetLogGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logGroup",
		val,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetSystemLogLevel(val *string) {
	if err := j.validateSetSystemLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemLogLevel",
		val,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ResetApplicationLogLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetApplicationLogLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ResetLogGroup() {
	_jsii_.InvokeVoid(
		l,
		"resetLogGroup",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ResetSystemLogLevel() {
	_jsii_.InvokeVoid(
		l,
		"resetSystemLogLevel",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (l *jsiiProxy_LambdaWebFunctionRevisionServiceConfigTelemetryConfigLoggingConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

