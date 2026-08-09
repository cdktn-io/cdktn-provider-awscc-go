// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/resiliencehubv2service/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference interface {
	cdktn.ComplexObject
	CfnStackArn() *string
	SetCfnStackArn(val *string)
	CfnStackArnInput() *string
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
	DesignFileS3Url() *string
	SetDesignFileS3Url(val *string)
	DesignFileS3UrlInput() *string
	Eks() Resiliencehubv2ServiceInputSourcesResourceConfigurationEksOutputReference
	EksInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ResourceTags() Resiliencehubv2ServiceInputSourcesResourceConfigurationResourceTagsList
	ResourceTagsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TfStateFileUrl() *string
	SetTfStateFileUrl(val *string)
	TfStateFileUrlInput() *string
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
	PutEks(value *Resiliencehubv2ServiceInputSourcesResourceConfigurationEks)
	PutResourceTags(value interface{})
	ResetCfnStackArn()
	ResetDesignFileS3Url()
	ResetEks()
	ResetResourceTags()
	ResetTfStateFileUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference
type jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) CfnStackArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) CfnStackArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cfnStackArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) DesignFileS3Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3Url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) DesignFileS3UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"designFileS3UrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) Eks() Resiliencehubv2ServiceInputSourcesResourceConfigurationEksOutputReference {
	var returns Resiliencehubv2ServiceInputSourcesResourceConfigurationEksOutputReference
	_jsii_.Get(
		j,
		"eks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) EksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ResourceTags() Resiliencehubv2ServiceInputSourcesResourceConfigurationResourceTagsList {
	var returns Resiliencehubv2ServiceInputSourcesResourceConfigurationResourceTagsList
	_jsii_.Get(
		j,
		"resourceTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ResourceTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) TfStateFileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) TfStateFileUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tfStateFileUrlInput",
		&returns,
	)
	return returns
}


func NewResiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewResiliencehubv2ServiceInputSourcesResourceConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.resiliencehubv2Service.Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewResiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference_Override(r Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.resiliencehubv2Service.Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetCfnStackArn(val *string) {
	if err := j.validateSetCfnStackArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cfnStackArn",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetDesignFileS3Url(val *string) {
	if err := j.validateSetDesignFileS3UrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"designFileS3Url",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference)SetTfStateFileUrl(val *string) {
	if err := j.validateSetTfStateFileUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tfStateFileUrl",
		val,
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) PutEks(value *Resiliencehubv2ServiceInputSourcesResourceConfigurationEks) {
	if err := r.validatePutEksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEks",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) PutResourceTags(value interface{}) {
	if err := r.validatePutResourceTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putResourceTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ResetCfnStackArn() {
	_jsii_.InvokeVoid(
		r,
		"resetCfnStackArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ResetDesignFileS3Url() {
	_jsii_.InvokeVoid(
		r,
		"resetDesignFileS3Url",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ResetEks() {
	_jsii_.InvokeVoid(
		r,
		"resetEks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ResetResourceTags() {
	_jsii_.InvokeVoid(
		r,
		"resetResourceTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ResetTfStateFileUrl() {
	_jsii_.InvokeVoid(
		r,
		"resetTfStateFileUrl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := r.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Resiliencehubv2ServiceInputSourcesResourceConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

