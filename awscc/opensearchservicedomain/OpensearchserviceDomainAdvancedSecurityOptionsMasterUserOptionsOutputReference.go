// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package opensearchservicedomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/opensearchservicedomain/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference interface {
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
	MasterUserArn() *string
	SetMasterUserArn(val *string)
	MasterUserArnInput() *string
	MasterUserName() *string
	SetMasterUserName(val *string)
	MasterUserNameInput() *string
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MasterUserPasswordInput() *string
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
	ResetMasterUserArn()
	ResetMasterUserName()
	ResetMasterUserPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference
type jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.opensearchserviceDomain.OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference_Override(o OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.opensearchserviceDomain.OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetMasterUserArn(val *string) {
	if err := j.validateSetMasterUserArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserArn",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetMasterUserName(val *string) {
	if err := j.validateSetMasterUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserName",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserArn() {
	_jsii_.InvokeVoid(
		o,
		"resetMasterUserArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserName() {
	_jsii_.InvokeVoid(
		o,
		"resetMasterUserName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := o.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchserviceDomainAdvancedSecurityOptionsMasterUserOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

