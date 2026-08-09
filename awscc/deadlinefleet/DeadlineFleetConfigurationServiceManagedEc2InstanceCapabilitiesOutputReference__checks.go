// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package deadlinefleet

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validatePutAcceleratorCapabilitiesParameters(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesAcceleratorCapabilities) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validatePutCustomAmountsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmounts:
		value := value.(*[]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmounts)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmounts:
		value_ := value.([]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmounts)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAmounts; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validatePutCustomAttributesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributes:
		value := value.(*[]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributes)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributes:
		value_ := value.([]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributes)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesCustomAttributes; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validatePutMemoryMiBParameters(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesMemoryMiB) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validatePutRootEbsVolumeParameters(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesRootEbsVolume) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validatePutVCpuCountParameters(value *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesVCpuCount) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetAllowedInstanceTypesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetCpuArchitectureTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetExcludedInstanceTypesParameters(val *[]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities:
		val := val.(*DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities:
		val_ := val.(DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilities; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetOsFamilyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDeadlineFleetConfigurationServiceManagedEc2InstanceCapabilitiesOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

