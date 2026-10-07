// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataawsccmediatailorprogram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-awscc-go/awscc/jsii"

	"github.com/cdktn-io/cdktn-provider-awscc-go/awscc/dataawsccmediatailorprogram/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference interface {
	cdktn.ComplexObject
	AdBreakMetadata() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksAdBreakMetadataList
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
	InternalValue() *DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaks
	SetInternalValue(val *DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaks)
	MessageType() *string
	OffsetMillis() *float64
	Slate() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksSlateOutputReference
	SpliceInsertMessage() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksSpliceInsertMessageOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	TimeSignalMessage() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksTimeSignalMessageOutputReference
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

// The jsii proxy struct for DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference
type jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) AdBreakMetadata() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksAdBreakMetadataList {
	var returns DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksAdBreakMetadataList
	_jsii_.Get(
		j,
		"adBreakMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) InternalValue() *DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaks {
	var returns *DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaks
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) MessageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) OffsetMillis() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"offsetMillis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) Slate() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksSlateOutputReference {
	var returns DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksSlateOutputReference
	_jsii_.Get(
		j,
		"slate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) SpliceInsertMessage() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksSpliceInsertMessageOutputReference {
	var returns DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksSpliceInsertMessageOutputReference
	_jsii_.Get(
		j,
		"spliceInsertMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) TimeSignalMessage() DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksTimeSignalMessageOutputReference {
	var returns DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksTimeSignalMessageOutputReference
	_jsii_.Get(
		j,
		"timeSignalMessage",
		&returns,
	)
	return returns
}


func NewDataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediatailorProgram.DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference_Override(d DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-awscc.dataAwsccMediatailorProgram.DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference)SetInternalValue(val *DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaks) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccMediatailorProgramAudienceMediaAlternateMediaAdBreaksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

