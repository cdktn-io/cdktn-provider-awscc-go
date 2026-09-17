// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package ec2ec2fleet

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateSetAvailabilityZoneParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateSetAvailabilityZoneIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZones:
		val := val.(*Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZones)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZones:
		val_ := val.(Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZones)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZones; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEc2Ec2FleetLaunchTemplateConfigsOverridesBlockDeviceMappingsEbsMultiAvailabilityZoneConfigurationStandbyAvailabilityZonesOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

