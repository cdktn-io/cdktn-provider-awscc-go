// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package verifiedpermissionspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/verifiedpermissionspolicy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference interface {
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
	PolicyTemplateId() *string
	SetPolicyTemplateId(val *string)
	PolicyTemplateIdInput() *string
	Principal() VerifiedpermissionsPolicyDefinitionTemplateLinkedPrincipalOutputReference
	PrincipalInput() interface{}
	Resource() VerifiedpermissionsPolicyDefinitionTemplateLinkedResourceOutputReference
	ResourceInput() interface{}
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
	PutPrincipal(value *VerifiedpermissionsPolicyDefinitionTemplateLinkedPrincipal)
	PutResource(value *VerifiedpermissionsPolicyDefinitionTemplateLinkedResource)
	ResetPolicyTemplateId()
	ResetPrincipal()
	ResetResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference
type jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) PolicyTemplateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTemplateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) PolicyTemplateIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTemplateIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) Principal() VerifiedpermissionsPolicyDefinitionTemplateLinkedPrincipalOutputReference {
	var returns VerifiedpermissionsPolicyDefinitionTemplateLinkedPrincipalOutputReference
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) PrincipalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) Resource() VerifiedpermissionsPolicyDefinitionTemplateLinkedResourceOutputReference {
	var returns VerifiedpermissionsPolicyDefinitionTemplateLinkedResourceOutputReference
	_jsii_.Get(
		j,
		"resource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference {
	_init_.Initialize()

	if err := validateNewVerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.verifiedpermissionsPolicy.VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference_Override(v VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.verifiedpermissionsPolicy.VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference)SetPolicyTemplateId(val *string) {
	if err := j.validateSetPolicyTemplateIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyTemplateId",
		val,
	)
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) PutPrincipal(value *VerifiedpermissionsPolicyDefinitionTemplateLinkedPrincipal) {
	if err := v.validatePutPrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPrincipal",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) PutResource(value *VerifiedpermissionsPolicyDefinitionTemplateLinkedResource) {
	if err := v.validatePutResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putResource",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ResetPolicyTemplateId() {
	_jsii_.InvokeVoid(
		v,
		"resetPolicyTemplateId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ResetPrincipal() {
	_jsii_.InvokeVoid(
		v,
		"resetPrincipal",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ResetResource() {
	_jsii_.InvokeVoid(
		v,
		"resetResource",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VerifiedpermissionsPolicyDefinitionTemplateLinkedOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

