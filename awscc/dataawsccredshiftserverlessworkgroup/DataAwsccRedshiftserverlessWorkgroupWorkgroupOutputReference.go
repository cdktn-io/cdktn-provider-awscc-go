// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccredshiftserverlessworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccredshiftserverlessworkgroup/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference interface {
	cdktn.ComplexObject
	BaseCapacity() *float64
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
	ConfigParameters() DataAwsccRedshiftserverlessWorkgroupWorkgroupConfigParametersList
	CreationDate() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Endpoint() DataAwsccRedshiftserverlessWorkgroupWorkgroupEndpointOutputReference
	EnhancedVpcRouting() cdktn.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataAwsccRedshiftserverlessWorkgroupWorkgroup
	SetInternalValue(val *DataAwsccRedshiftserverlessWorkgroupWorkgroup)
	MaxCapacity() *float64
	NamespaceName() *string
	PricePerformanceTarget() DataAwsccRedshiftserverlessWorkgroupWorkgroupPricePerformanceTargetOutputReference
	PubliclyAccessible() cdktn.IResolvable
	SecurityGroupIds() *[]*string
	Status() *string
	SubnetIds() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TrackName() *string
	WorkgroupArn() *string
	WorkgroupId() *string
	WorkgroupName() *string
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

// The jsii proxy struct for DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference
type jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) BaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) ConfigParameters() DataAwsccRedshiftserverlessWorkgroupWorkgroupConfigParametersList {
	var returns DataAwsccRedshiftserverlessWorkgroupWorkgroupConfigParametersList
	_jsii_.Get(
		j,
		"configParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) Endpoint() DataAwsccRedshiftserverlessWorkgroupWorkgroupEndpointOutputReference {
	var returns DataAwsccRedshiftserverlessWorkgroupWorkgroupEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) EnhancedVpcRouting() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) InternalValue() *DataAwsccRedshiftserverlessWorkgroupWorkgroup {
	var returns *DataAwsccRedshiftserverlessWorkgroupWorkgroup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) PricePerformanceTarget() DataAwsccRedshiftserverlessWorkgroupWorkgroupPricePerformanceTargetOutputReference {
	var returns DataAwsccRedshiftserverlessWorkgroupWorkgroupPricePerformanceTargetOutputReference
	_jsii_.Get(
		j,
		"pricePerformanceTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) PubliclyAccessible() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) TrackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) WorkgroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) WorkgroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) WorkgroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupName",
		&returns,
	)
	return returns
}


func NewDataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRedshiftserverlessWorkgroup.DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference_Override(d DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccRedshiftserverlessWorkgroup.DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference)SetInternalValue(val *DataAwsccRedshiftserverlessWorkgroupWorkgroup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccRedshiftserverlessWorkgroupWorkgroupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

