// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package rtbfabriclink

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validatePutActionParameters(value *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeAction) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validatePutFilterConfigurationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfiguration:
		value := value.(*[]*RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfiguration)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfiguration:
		value_ := value.([]*RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfiguration)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeFilterConfiguration; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (r *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateSetFilterTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateSetHoldbackPercentageParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttribute:
		val := val.(*RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttribute)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttribute:
		val_ := val.(RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttribute)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttribute; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_RtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewRtbfabricLinkModuleConfigurationListModuleParametersOpenRtbAttributeOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

