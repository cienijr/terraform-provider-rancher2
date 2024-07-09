package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Types

type genericConfig struct {
	Driver string                 `json:"driver,omitempty" yaml:"driver,omitempty"`
	Config map[string]interface{} `json:"config,omitempty" yaml:"config,omitempty"`
}

//Schemas

func genericConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"driver": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   false,
			Description: "Driver type",
		},
		"config_json": {
			Type:        schema.TypeString,
			Required:    true,
			Sensitive:   true,
			Description: "Driver config encoded as JSON",
		},
	}

	return s
}
