package rancher2

import "fmt"

// Config name

func fieldNameForGenericConfig(p []interface{}) string {
	if len(p) == 0 || p[0] == nil {
		return ""
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["driver"].(string); ok && len(v) > 0 {
		return fmt.Sprintf("%sConfig", v)
	}

	return ""
}

// Expanders

func configForGenericConfig(p []interface{}) interface{} {
	if len(p) == 0 || p[0] == nil {
		return nil
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["config_json"].(string); ok && len(v) > 0 {
		v2, _ := jsonToMapInterface(v)
		return v2
	}

	return nil
}

func expandGenericConfig(p []interface{}) *genericConfig {
	obj := &genericConfig{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["driver"].(string); ok && len(v) > 0 {
		obj.Driver = v
	}

	if v, ok := in["config_json"].(string); ok && len(v) > 0 {
		obj.Config, _ = jsonToMapInterface(v)
	}

	return obj
}
