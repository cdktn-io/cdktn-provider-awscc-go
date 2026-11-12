// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package bedrockagentcoregatewayrule

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validatePutTriggerOnParameters(value *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackTriggerOn) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (b *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetFirstByteTimeoutMillisecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallback:
		val := val.(*BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallback)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallback:
		val_ := val.(BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallback)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallback; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetModelParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetTargetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_BedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewBedrockagentcoreGatewayRuleActionsRouteToTargetStaticRouteFallbackTargetFallbackOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

