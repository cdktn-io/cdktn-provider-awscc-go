// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build !no_runtime_type_checking

package mediaconnectsourceendpoint

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase:
		val := val.(*MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase:
		val_ := val.(MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateSetSecretArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewMediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphraseOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

