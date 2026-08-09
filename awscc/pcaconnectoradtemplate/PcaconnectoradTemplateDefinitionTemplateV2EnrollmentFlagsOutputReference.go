// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference interface {
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
	EnableKeyReuseOnNtTokenKeysetStorageFull() interface{}
	SetEnableKeyReuseOnNtTokenKeysetStorageFull(val interface{})
	EnableKeyReuseOnNtTokenKeysetStorageFullInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeSymmetricAlgorithms() interface{}
	SetIncludeSymmetricAlgorithms(val interface{})
	IncludeSymmetricAlgorithmsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	NoSecurityExtension() interface{}
	SetNoSecurityExtension(val interface{})
	NoSecurityExtensionInput() interface{}
	RemoveInvalidCertificateFromPersonalStore() interface{}
	SetRemoveInvalidCertificateFromPersonalStore(val interface{})
	RemoveInvalidCertificateFromPersonalStoreInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	UserInteractionRequired() interface{}
	SetUserInteractionRequired(val interface{})
	UserInteractionRequiredInput() interface{}
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
	ResetEnableKeyReuseOnNtTokenKeysetStorageFull()
	ResetIncludeSymmetricAlgorithms()
	ResetNoSecurityExtension()
	ResetRemoveInvalidCertificateFromPersonalStore()
	ResetUserInteractionRequired()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) EnableKeyReuseOnNtTokenKeysetStorageFull() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKeyReuseOnNtTokenKeysetStorageFull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) EnableKeyReuseOnNtTokenKeysetStorageFullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableKeyReuseOnNtTokenKeysetStorageFullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) IncludeSymmetricAlgorithms() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSymmetricAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) IncludeSymmetricAlgorithmsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSymmetricAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) NoSecurityExtension() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSecurityExtension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) NoSecurityExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noSecurityExtensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) RemoveInvalidCertificateFromPersonalStore() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInvalidCertificateFromPersonalStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) RemoveInvalidCertificateFromPersonalStoreInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeInvalidCertificateFromPersonalStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) UserInteractionRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInteractionRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) UserInteractionRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userInteractionRequiredInput",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetEnableKeyReuseOnNtTokenKeysetStorageFull(val interface{}) {
	if err := j.validateSetEnableKeyReuseOnNtTokenKeysetStorageFullParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableKeyReuseOnNtTokenKeysetStorageFull",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetIncludeSymmetricAlgorithms(val interface{}) {
	if err := j.validateSetIncludeSymmetricAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSymmetricAlgorithms",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetNoSecurityExtension(val interface{}) {
	if err := j.validateSetNoSecurityExtensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSecurityExtension",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetRemoveInvalidCertificateFromPersonalStore(val interface{}) {
	if err := j.validateSetRemoveInvalidCertificateFromPersonalStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeInvalidCertificateFromPersonalStore",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference)SetUserInteractionRequired(val interface{}) {
	if err := j.validateSetUserInteractionRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userInteractionRequired",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ResetEnableKeyReuseOnNtTokenKeysetStorageFull() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableKeyReuseOnNtTokenKeysetStorageFull",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ResetIncludeSymmetricAlgorithms() {
	_jsii_.InvokeVoid(
		p,
		"resetIncludeSymmetricAlgorithms",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ResetNoSecurityExtension() {
	_jsii_.InvokeVoid(
		p,
		"resetNoSecurityExtension",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ResetRemoveInvalidCertificateFromPersonalStore() {
	_jsii_.InvokeVoid(
		p,
		"resetRemoveInvalidCertificateFromPersonalStore",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ResetUserInteractionRequired() {
	_jsii_.InvokeVoid(
		p,
		"resetUserInteractionRequired",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := p.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

