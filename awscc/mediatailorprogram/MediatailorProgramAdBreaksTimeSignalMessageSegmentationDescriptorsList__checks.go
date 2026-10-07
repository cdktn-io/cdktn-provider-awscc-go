// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package mediatailorprogram

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptors:
		val := val.(*[]*MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptors)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptors:
		val_ := val.([]*MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptors)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptors; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewMediatailorProgramAdBreaksTimeSignalMessageSegmentationDescriptorsListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

