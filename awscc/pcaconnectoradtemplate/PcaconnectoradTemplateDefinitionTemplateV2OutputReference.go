// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcaconnectoradtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/pcaconnectoradtemplate/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type PcaconnectoradTemplateDefinitionTemplateV2OutputReference interface {
	cdktn.ComplexObject
	CertificateValidity() PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference
	CertificateValidityInput() interface{}
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
	EnrollmentFlags() PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference
	EnrollmentFlagsInput() interface{}
	Extensions() PcaconnectoradTemplateDefinitionTemplateV2ExtensionsOutputReference
	ExtensionsInput() interface{}
	// Experimental.
	Fqn() *string
	GeneralFlags() PcaconnectoradTemplateDefinitionTemplateV2GeneralFlagsOutputReference
	GeneralFlagsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PrivateKeyAttributes() PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyAttributesOutputReference
	PrivateKeyAttributesInput() interface{}
	PrivateKeyFlags() PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyFlagsOutputReference
	PrivateKeyFlagsInput() interface{}
	SubjectNameFlags() PcaconnectoradTemplateDefinitionTemplateV2SubjectNameFlagsOutputReference
	SubjectNameFlagsInput() interface{}
	SupersededTemplates() *[]*string
	SetSupersededTemplates(val *[]*string)
	SupersededTemplatesInput() *[]*string
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
	PutCertificateValidity(value *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity)
	PutEnrollmentFlags(value *PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlags)
	PutExtensions(value *PcaconnectoradTemplateDefinitionTemplateV2Extensions)
	PutGeneralFlags(value *PcaconnectoradTemplateDefinitionTemplateV2GeneralFlags)
	PutPrivateKeyAttributes(value *PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyAttributes)
	PutPrivateKeyFlags(value *PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyFlags)
	PutSubjectNameFlags(value *PcaconnectoradTemplateDefinitionTemplateV2SubjectNameFlags)
	ResetCertificateValidity()
	ResetEnrollmentFlags()
	ResetExtensions()
	ResetGeneralFlags()
	ResetPrivateKeyAttributes()
	ResetPrivateKeyFlags()
	ResetSubjectNameFlags()
	ResetSupersededTemplates()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PcaconnectoradTemplateDefinitionTemplateV2OutputReference
type jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) CertificateValidity() PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2CertificateValidityOutputReference
	_jsii_.Get(
		j,
		"certificateValidity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) CertificateValidityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateValidityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) EnrollmentFlags() PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlagsOutputReference
	_jsii_.Get(
		j,
		"enrollmentFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) EnrollmentFlagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrollmentFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) Extensions() PcaconnectoradTemplateDefinitionTemplateV2ExtensionsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2ExtensionsOutputReference
	_jsii_.Get(
		j,
		"extensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GeneralFlags() PcaconnectoradTemplateDefinitionTemplateV2GeneralFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2GeneralFlagsOutputReference
	_jsii_.Get(
		j,
		"generalFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GeneralFlagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generalFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PrivateKeyAttributes() PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyAttributesOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyAttributesOutputReference
	_jsii_.Get(
		j,
		"privateKeyAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PrivateKeyAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateKeyAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PrivateKeyFlags() PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyFlagsOutputReference
	_jsii_.Get(
		j,
		"privateKeyFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PrivateKeyFlagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateKeyFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) SubjectNameFlags() PcaconnectoradTemplateDefinitionTemplateV2SubjectNameFlagsOutputReference {
	var returns PcaconnectoradTemplateDefinitionTemplateV2SubjectNameFlagsOutputReference
	_jsii_.Get(
		j,
		"subjectNameFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) SubjectNameFlagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subjectNameFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) SupersededTemplates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supersededTemplates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) SupersededTemplatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supersededTemplatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPcaconnectoradTemplateDefinitionTemplateV2OutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) PcaconnectoradTemplateDefinitionTemplateV2OutputReference {
	_init_.Initialize()

	if err := validateNewPcaconnectoradTemplateDefinitionTemplateV2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPcaconnectoradTemplateDefinitionTemplateV2OutputReference_Override(p PcaconnectoradTemplateDefinitionTemplateV2OutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.pcaconnectoradTemplate.PcaconnectoradTemplateDefinitionTemplateV2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference)SetSupersededTemplates(val *[]*string) {
	if err := j.validateSetSupersededTemplatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supersededTemplates",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PutCertificateValidity(value *PcaconnectoradTemplateDefinitionTemplateV2CertificateValidity) {
	if err := p.validatePutCertificateValidityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCertificateValidity",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PutEnrollmentFlags(value *PcaconnectoradTemplateDefinitionTemplateV2EnrollmentFlags) {
	if err := p.validatePutEnrollmentFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putEnrollmentFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PutExtensions(value *PcaconnectoradTemplateDefinitionTemplateV2Extensions) {
	if err := p.validatePutExtensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putExtensions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PutGeneralFlags(value *PcaconnectoradTemplateDefinitionTemplateV2GeneralFlags) {
	if err := p.validatePutGeneralFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putGeneralFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PutPrivateKeyAttributes(value *PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyAttributes) {
	if err := p.validatePutPrivateKeyAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPrivateKeyAttributes",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PutPrivateKeyFlags(value *PcaconnectoradTemplateDefinitionTemplateV2PrivateKeyFlags) {
	if err := p.validatePutPrivateKeyFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPrivateKeyFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) PutSubjectNameFlags(value *PcaconnectoradTemplateDefinitionTemplateV2SubjectNameFlags) {
	if err := p.validatePutSubjectNameFlagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSubjectNameFlags",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetCertificateValidity() {
	_jsii_.InvokeVoid(
		p,
		"resetCertificateValidity",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetEnrollmentFlags() {
	_jsii_.InvokeVoid(
		p,
		"resetEnrollmentFlags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetExtensions() {
	_jsii_.InvokeVoid(
		p,
		"resetExtensions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetGeneralFlags() {
	_jsii_.InvokeVoid(
		p,
		"resetGeneralFlags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetPrivateKeyAttributes() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKeyAttributes",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetPrivateKeyFlags() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKeyFlags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetSubjectNameFlags() {
	_jsii_.InvokeVoid(
		p,
		"resetSubjectNameFlags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ResetSupersededTemplates() {
	_jsii_.InvokeVoid(
		p,
		"resetSupersededTemplates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (p *jsiiProxy_PcaconnectoradTemplateDefinitionTemplateV2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

