// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package glueconnectiontype

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validatePutValueParameters(value *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterValue) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (g *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetDefaultValueParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameter:
		val := val.(*GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameter)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameter:
		val_ := val.(GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameter)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameter; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetKeyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetPropertyLocationParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_GlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewGlueConnectionTypeRestConfigurationGlobalSourceConfigurationPaginationConfigurationOffsetConfigurationOffsetParameterOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

