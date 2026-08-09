// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evsenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/evsenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type EvsEnvironmentVcfHostnamesOutputReference interface {
	cdktn.ComplexObject
	CloudBuilder() *string
	SetCloudBuilder(val *string)
	CloudBuilderInput() *string
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
	Nsx() *string
	SetNsx(val *string)
	NsxEdge1() *string
	SetNsxEdge1(val *string)
	NsxEdge1Input() *string
	NsxEdge2() *string
	SetNsxEdge2(val *string)
	NsxEdge2Input() *string
	NsxInput() *string
	NsxManager1() *string
	SetNsxManager1(val *string)
	NsxManager1Input() *string
	NsxManager2() *string
	SetNsxManager2(val *string)
	NsxManager2Input() *string
	NsxManager3() *string
	SetNsxManager3(val *string)
	NsxManager3Input() *string
	SddcManager() *string
	SetSddcManager(val *string)
	SddcManagerInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VCenter() *string
	SetVCenter(val *string)
	VCenterInput() *string
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
	ResetCloudBuilder()
	ResetNsx()
	ResetNsxEdge1()
	ResetNsxEdge2()
	ResetNsxManager1()
	ResetNsxManager2()
	ResetNsxManager3()
	ResetSddcManager()
	ResetVCenter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EvsEnvironmentVcfHostnamesOutputReference
type jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) CloudBuilder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuilder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) CloudBuilderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudBuilderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) Nsx() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxEdge1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxEdge1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxEdge1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxEdge1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxEdge2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxEdge2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxEdge2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxEdge2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxManager1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxManager1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxManager1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxManager1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxManager2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxManager2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxManager2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxManager2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxManager3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxManager3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) NsxManager3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nsxManager3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) SddcManager() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sddcManager",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) SddcManagerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sddcManagerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) VCenter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vCenter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) VCenterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vCenterInput",
		&returns,
	)
	return returns
}


func NewEvsEnvironmentVcfHostnamesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) EvsEnvironmentVcfHostnamesOutputReference {
	_init_.Initialize()

	if err := validateNewEvsEnvironmentVcfHostnamesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironmentVcfHostnamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEvsEnvironmentVcfHostnamesOutputReference_Override(e EvsEnvironmentVcfHostnamesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.evsEnvironment.EvsEnvironmentVcfHostnamesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetCloudBuilder(val *string) {
	if err := j.validateSetCloudBuilderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudBuilder",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetNsx(val *string) {
	if err := j.validateSetNsxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsx",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetNsxEdge1(val *string) {
	if err := j.validateSetNsxEdge1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxEdge1",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetNsxEdge2(val *string) {
	if err := j.validateSetNsxEdge2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxEdge2",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetNsxManager1(val *string) {
	if err := j.validateSetNsxManager1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxManager1",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetNsxManager2(val *string) {
	if err := j.validateSetNsxManager2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxManager2",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetNsxManager3(val *string) {
	if err := j.validateSetNsxManager3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nsxManager3",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetSddcManager(val *string) {
	if err := j.validateSetSddcManagerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sddcManager",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference)SetVCenter(val *string) {
	if err := j.validateSetVCenterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vCenter",
		val,
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetCloudBuilder() {
	_jsii_.InvokeVoid(
		e,
		"resetCloudBuilder",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetNsx() {
	_jsii_.InvokeVoid(
		e,
		"resetNsx",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetNsxEdge1() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxEdge1",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetNsxEdge2() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxEdge2",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetNsxManager1() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxManager1",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetNsxManager2() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxManager2",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetNsxManager3() {
	_jsii_.InvokeVoid(
		e,
		"resetNsxManager3",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetSddcManager() {
	_jsii_.InvokeVoid(
		e,
		"resetSddcManager",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ResetVCenter() {
	_jsii_.InvokeVoid(
		e,
		"resetVCenter",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (e *jsiiProxy_EvsEnvironmentVcfHostnamesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

