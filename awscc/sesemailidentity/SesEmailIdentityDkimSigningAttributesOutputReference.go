// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesemailidentity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/sesemailidentity/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type SesEmailIdentityDkimSigningAttributesOutputReference interface {
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
	DomainSigningPrivateKey() *string
	SetDomainSigningPrivateKey(val *string)
	DomainSigningPrivateKeyInput() *string
	DomainSigningSelector() *string
	SetDomainSigningSelector(val *string)
	DomainSigningSelectorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NextSigningKeyLength() *string
	SetNextSigningKeyLength(val *string)
	NextSigningKeyLengthInput() *string
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
	ResetDomainSigningPrivateKey()
	ResetDomainSigningSelector()
	ResetNextSigningKeyLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SesEmailIdentityDkimSigningAttributesOutputReference
type jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) DomainSigningPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) DomainSigningPrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) DomainSigningSelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) DomainSigningSelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) NextSigningKeyLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) NextSigningKeyLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSesEmailIdentityDkimSigningAttributesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) SesEmailIdentityDkimSigningAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewSesEmailIdentityDkimSigningAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.sesEmailIdentity.SesEmailIdentityDkimSigningAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesEmailIdentityDkimSigningAttributesOutputReference_Override(s SesEmailIdentityDkimSigningAttributesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.sesEmailIdentity.SesEmailIdentityDkimSigningAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetDomainSigningPrivateKey(val *string) {
	if err := j.validateSetDomainSigningPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningPrivateKey",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetDomainSigningSelector(val *string) {
	if err := j.validateSetDomainSigningSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningSelector",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetNextSigningKeyLength(val *string) {
	if err := j.validateSetNextSigningKeyLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextSigningKeyLength",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) ResetDomainSigningPrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainSigningPrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) ResetDomainSigningSelector() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainSigningSelector",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) ResetNextSigningKeyLength() {
	_jsii_.InvokeVoid(
		s,
		"resetNextSigningKeyLength",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SesEmailIdentityDkimSigningAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

