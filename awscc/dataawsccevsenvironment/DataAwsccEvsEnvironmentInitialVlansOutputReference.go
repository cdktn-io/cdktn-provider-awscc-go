// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccevsenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccevsenvironment/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccEvsEnvironmentInitialVlansOutputReference interface {
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
	EdgeVTep() DataAwsccEvsEnvironmentInitialVlansEdgeVTepOutputReference
	ExpansionVlan1() DataAwsccEvsEnvironmentInitialVlansExpansionVlan1OutputReference
	ExpansionVlan2() DataAwsccEvsEnvironmentInitialVlansExpansionVlan2OutputReference
	// Experimental.
	Fqn() *string
	Hcx() DataAwsccEvsEnvironmentInitialVlansHcxOutputReference
	HcxNetworkAclId() *string
	InternalValue() *DataAwsccEvsEnvironmentInitialVlans
	SetInternalValue(val *DataAwsccEvsEnvironmentInitialVlans)
	IsHcxPublic() cdktn.IResolvable
	NsxUpLink() DataAwsccEvsEnvironmentInitialVlansNsxUpLinkOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	VmkManagement() DataAwsccEvsEnvironmentInitialVlansVmkManagementOutputReference
	VmManagement() DataAwsccEvsEnvironmentInitialVlansVmManagementOutputReference
	VMotion() DataAwsccEvsEnvironmentInitialVlansVMotionOutputReference
	VSan() DataAwsccEvsEnvironmentInitialVlansVSanOutputReference
	VTep() DataAwsccEvsEnvironmentInitialVlansVTepOutputReference
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccEvsEnvironmentInitialVlansOutputReference
type jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) EdgeVTep() DataAwsccEvsEnvironmentInitialVlansEdgeVTepOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansEdgeVTepOutputReference
	_jsii_.Get(
		j,
		"edgeVTep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) ExpansionVlan1() DataAwsccEvsEnvironmentInitialVlansExpansionVlan1OutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansExpansionVlan1OutputReference
	_jsii_.Get(
		j,
		"expansionVlan1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) ExpansionVlan2() DataAwsccEvsEnvironmentInitialVlansExpansionVlan2OutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansExpansionVlan2OutputReference
	_jsii_.Get(
		j,
		"expansionVlan2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) Hcx() DataAwsccEvsEnvironmentInitialVlansHcxOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansHcxOutputReference
	_jsii_.Get(
		j,
		"hcx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) HcxNetworkAclId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hcxNetworkAclId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) InternalValue() *DataAwsccEvsEnvironmentInitialVlans {
	var returns *DataAwsccEvsEnvironmentInitialVlans
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) IsHcxPublic() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"isHcxPublic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) NsxUpLink() DataAwsccEvsEnvironmentInitialVlansNsxUpLinkOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansNsxUpLinkOutputReference
	_jsii_.Get(
		j,
		"nsxUpLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) VmkManagement() DataAwsccEvsEnvironmentInitialVlansVmkManagementOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansVmkManagementOutputReference
	_jsii_.Get(
		j,
		"vmkManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) VmManagement() DataAwsccEvsEnvironmentInitialVlansVmManagementOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansVmManagementOutputReference
	_jsii_.Get(
		j,
		"vmManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) VMotion() DataAwsccEvsEnvironmentInitialVlansVMotionOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansVMotionOutputReference
	_jsii_.Get(
		j,
		"vMotion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) VSan() DataAwsccEvsEnvironmentInitialVlansVSanOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansVSanOutputReference
	_jsii_.Get(
		j,
		"vSan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) VTep() DataAwsccEvsEnvironmentInitialVlansVTepOutputReference {
	var returns DataAwsccEvsEnvironmentInitialVlansVTepOutputReference
	_jsii_.Get(
		j,
		"vTep",
		&returns,
	)
	return returns
}


func NewDataAwsccEvsEnvironmentInitialVlansOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccEvsEnvironmentInitialVlansOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccEvsEnvironmentInitialVlansOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEvsEnvironment.DataAwsccEvsEnvironmentInitialVlansOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccEvsEnvironmentInitialVlansOutputReference_Override(d DataAwsccEvsEnvironmentInitialVlansOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccEvsEnvironment.DataAwsccEvsEnvironmentInitialVlansOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference)SetInternalValue(val *DataAwsccEvsEnvironmentInitialVlans) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEvsEnvironmentInitialVlansOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

