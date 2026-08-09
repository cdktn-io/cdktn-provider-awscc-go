// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscapability

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/ekscapability/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EksCapabilityConfigurationArgoCdOutputReference interface {
	cdktn.ComplexObject
	AwsIdc() EksCapabilityConfigurationArgoCdAwsIdcOutputReference
	AwsIdcInput() interface{}
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
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	NetworkAccess() EksCapabilityConfigurationArgoCdNetworkAccessOutputReference
	NetworkAccessInput() interface{}
	RbacRoleMappings() EksCapabilityConfigurationArgoCdRbacRoleMappingsList
	RbacRoleMappingsInput() interface{}
	ServerUrl() *string
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
	PutAwsIdc(value *EksCapabilityConfigurationArgoCdAwsIdc)
	PutNetworkAccess(value *EksCapabilityConfigurationArgoCdNetworkAccess)
	PutRbacRoleMappings(value interface{})
	ResetAwsIdc()
	ResetNamespace()
	ResetNetworkAccess()
	ResetRbacRoleMappings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EksCapabilityConfigurationArgoCdOutputReference
type jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) AwsIdc() EksCapabilityConfigurationArgoCdAwsIdcOutputReference {
	var returns EksCapabilityConfigurationArgoCdAwsIdcOutputReference
	_jsii_.Get(
		j,
		"awsIdc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) AwsIdcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"awsIdcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) NetworkAccess() EksCapabilityConfigurationArgoCdNetworkAccessOutputReference {
	var returns EksCapabilityConfigurationArgoCdNetworkAccessOutputReference
	_jsii_.Get(
		j,
		"networkAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) NetworkAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) RbacRoleMappings() EksCapabilityConfigurationArgoCdRbacRoleMappingsList {
	var returns EksCapabilityConfigurationArgoCdRbacRoleMappingsList
	_jsii_.Get(
		j,
		"rbacRoleMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) RbacRoleMappingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rbacRoleMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ServerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEksCapabilityConfigurationArgoCdOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EksCapabilityConfigurationArgoCdOutputReference {
	_init_.Initialize()

	if err := validateNewEksCapabilityConfigurationArgoCdOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.eksCapability.EksCapabilityConfigurationArgoCdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEksCapabilityConfigurationArgoCdOutputReference_Override(e EksCapabilityConfigurationArgoCdOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.eksCapability.EksCapabilityConfigurationArgoCdOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) PutAwsIdc(value *EksCapabilityConfigurationArgoCdAwsIdc) {
	if err := e.validatePutAwsIdcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAwsIdc",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) PutNetworkAccess(value *EksCapabilityConfigurationArgoCdNetworkAccess) {
	if err := e.validatePutNetworkAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNetworkAccess",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) PutRbacRoleMappings(value interface{}) {
	if err := e.validatePutRbacRoleMappingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRbacRoleMappings",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ResetAwsIdc() {
	_jsii_.InvokeVoid(
		e,
		"resetAwsIdc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		e,
		"resetNamespace",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ResetNetworkAccess() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkAccess",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ResetRbacRoleMappings() {
	_jsii_.InvokeVoid(
		e,
		"resetRbacRoleMappings",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EksCapabilityConfigurationArgoCdOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

