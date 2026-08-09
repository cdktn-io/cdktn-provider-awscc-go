// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package finspaceenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/finspaceenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FinspaceEnvironmentFederationParametersOutputReference interface {
	cdktn.ComplexObject
	ApplicationCallBackUrl() *string
	SetApplicationCallBackUrl(val *string)
	ApplicationCallBackUrlInput() *string
	AttributeMap() FinspaceEnvironmentFederationParametersAttributeMapList
	AttributeMapInput() interface{}
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
	FederationProviderName() *string
	SetFederationProviderName(val *string)
	FederationProviderNameInput() *string
	FederationUrn() *string
	SetFederationUrn(val *string)
	FederationUrnInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SamlMetadataDocument() *string
	SetSamlMetadataDocument(val *string)
	SamlMetadataDocumentInput() *string
	SamlMetadataUrl() *string
	SetSamlMetadataUrl(val *string)
	SamlMetadataUrlInput() *string
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
	PutAttributeMap(value interface{})
	ResetApplicationCallBackUrl()
	ResetAttributeMap()
	ResetFederationProviderName()
	ResetFederationUrn()
	ResetSamlMetadataDocument()
	ResetSamlMetadataUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FinspaceEnvironmentFederationParametersOutputReference
type jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ApplicationCallBackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationCallBackUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ApplicationCallBackUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationCallBackUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) AttributeMap() FinspaceEnvironmentFederationParametersAttributeMapList {
	var returns FinspaceEnvironmentFederationParametersAttributeMapList
	_jsii_.Get(
		j,
		"attributeMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) AttributeMapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeMapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) FederationProviderName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationProviderName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) FederationProviderNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationProviderNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) FederationUrn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationUrn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) FederationUrnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"federationUrnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) SamlMetadataDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlMetadataDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) SamlMetadataDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlMetadataDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) SamlMetadataUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlMetadataUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) SamlMetadataUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samlMetadataUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFinspaceEnvironmentFederationParametersOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) FinspaceEnvironmentFederationParametersOutputReference {
	_init_.Initialize()

	if err := validateNewFinspaceEnvironmentFederationParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.finspaceEnvironment.FinspaceEnvironmentFederationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFinspaceEnvironmentFederationParametersOutputReference_Override(f FinspaceEnvironmentFederationParametersOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.finspaceEnvironment.FinspaceEnvironmentFederationParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetApplicationCallBackUrl(val *string) {
	if err := j.validateSetApplicationCallBackUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationCallBackUrl",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetFederationProviderName(val *string) {
	if err := j.validateSetFederationProviderNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"federationProviderName",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetFederationUrn(val *string) {
	if err := j.validateSetFederationUrnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"federationUrn",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetSamlMetadataDocument(val *string) {
	if err := j.validateSetSamlMetadataDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlMetadataDocument",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetSamlMetadataUrl(val *string) {
	if err := j.validateSetSamlMetadataUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samlMetadataUrl",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) PutAttributeMap(value interface{}) {
	if err := f.validatePutAttributeMapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAttributeMap",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ResetApplicationCallBackUrl() {
	_jsii_.InvokeVoid(
		f,
		"resetApplicationCallBackUrl",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ResetAttributeMap() {
	_jsii_.InvokeVoid(
		f,
		"resetAttributeMap",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ResetFederationProviderName() {
	_jsii_.InvokeVoid(
		f,
		"resetFederationProviderName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ResetFederationUrn() {
	_jsii_.InvokeVoid(
		f,
		"resetFederationUrn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ResetSamlMetadataDocument() {
	_jsii_.InvokeVoid(
		f,
		"resetSamlMetadataDocument",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ResetSamlMetadataUrl() {
	_jsii_.InvokeVoid(
		f,
		"resetSamlMetadataUrl",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := f.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FinspaceEnvironmentFederationParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

