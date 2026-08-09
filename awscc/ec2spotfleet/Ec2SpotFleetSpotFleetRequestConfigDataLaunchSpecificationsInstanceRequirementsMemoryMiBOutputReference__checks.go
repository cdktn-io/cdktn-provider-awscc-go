// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package ec2spotfleet

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiB:
		val := val.(*Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiB)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiB:
		val_ := val.(Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiB)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiB; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateSetMaxParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateSetMinParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsInstanceRequirementsMemoryMiBOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

