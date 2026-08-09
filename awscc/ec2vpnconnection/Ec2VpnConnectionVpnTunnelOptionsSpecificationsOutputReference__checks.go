// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package ec2vpnconnection

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutIkeVersionsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersions:
		value := value.(*[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersions)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersions:
		value_ := value.([]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersions)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsIkeVersions; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutLogOptionsParameters(value *Ec2VpnConnectionVpnTunnelOptionsSpecificationsLogOptions) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutPhase1DhGroupNumbersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbers:
		value := value.(*[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbers)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbers:
		value_ := value.([]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbers)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1DhGroupNumbers; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutPhase1EncryptionAlgorithmsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithms:
		value := value.(*[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithms)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithms:
		value_ := value.([]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithms)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1EncryptionAlgorithms; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutPhase1IntegrityAlgorithmsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithms:
		value := value.(*[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithms)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithms:
		value_ := value.([]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithms)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase1IntegrityAlgorithms; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutPhase2DhGroupNumbersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbers:
		value := value.(*[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbers)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbers:
		value_ := value.([]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbers)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbers; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutPhase2EncryptionAlgorithmsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithms:
		value := value.(*[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithms)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithms:
		value_ := value.([]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithms)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2EncryptionAlgorithms; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validatePutPhase2IntegrityAlgorithmsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithms:
		value := value.(*[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithms)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithms:
		value_ := value.([]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithms)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2IntegrityAlgorithms; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (e *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetDpdTimeoutActionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetDpdTimeoutSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetEnableTunnelLifecycleControlParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *Ec2VpnConnectionVpnTunnelOptionsSpecifications:
		val := val.(*Ec2VpnConnectionVpnTunnelOptionsSpecifications)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case Ec2VpnConnectionVpnTunnelOptionsSpecifications:
		val_ := val.(Ec2VpnConnectionVpnTunnelOptionsSpecifications)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *Ec2VpnConnectionVpnTunnelOptionsSpecifications; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetPhase1LifetimeSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetPhase2LifetimeSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetPreSharedKeyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetRekeyFuzzPercentageParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetRekeyMarginTimeSecondsParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetReplayWindowSizeParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetStartupActionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetTunnelInsideCidrParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Ec2VpnConnectionVpnTunnelOptionsSpecificationsOutputReference) validateSetTunnelInsideIpv6CidrParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewEc2VpnConnectionVpnTunnelOptionsSpecificationsOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

