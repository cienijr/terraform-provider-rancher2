package rancher2

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
	"github.com/stretchr/testify/assert"
)

var (
	testGenericNodeTemplateNodeTaintsConf      []managementClient.Taint
	testGenericNodeTemplateNodeTaintsInterface interface{}
	testNodeTemplateGenericConf                genericConfig
	testGenericNodeTemplateInterface           map[string]interface{}
	testGenericNodeTemplateConf                *NodeTemplate
	testExpandedGenericNodeTemplateConf        *NodeTemplate
	testNodeTemplateSquashGenericConfInterface map[string]interface{}
	testNodeTemplateExpandGenericConfInterface map[string]interface{}
)

func init() {
	testGenericNodeTemplateNodeTaintsConf = []managementClient.Taint{
		{
			Key:       "key",
			Value:     "value",
			Effect:    "recipient",
			TimeAdded: "time_added",
		},
	}
	testGenericNodeTemplateNodeTaintsInterface = []interface{}{
		map[string]interface{}{
			"key":        "key",
			"value":      "value",
			"effect":     "recipient",
			"time_added": "time_added",
		},
	}

	configMap := map[string]interface{}{
		"type":  "4cpu-8gb",
		"image": "ubuntu-22-04",
		"tags":  "tag-a,tag-b",
		"network": map[string]interface{}{
			"public_ipv4": true,
			"public_ipv6": false,
		},
	}
	configJson, _ := mapInterfaceToJSON(configMap)

	testNodeTemplateGenericConf = genericConfig{
		Driver: "mycustomdriver",
		Config: configMap,
	}
	testGenericNodeTemplateInterface = map[string]interface{}{
		"driver":      "mycustomdriver",
		"config_json": configJson,
	}
	testGenericNodeTemplateAnnotationsConf := map[string]string{
		"key": "value",
	}
	testGenericNodeTemplateAnnotationsInterface := map[string]interface{}{
		"key": "value",
	}

	useInternalIP := false
	testGenericNodeTemplateConf = &NodeTemplate{
		NodeTemplate: managementClient.NodeTemplate{
			Driver:               "mycustomdriver",
			UseInternalIPAddress: &useInternalIP,
			Annotations:          testGenericNodeTemplateAnnotationsConf,
			NodeTaints:           testGenericNodeTemplateNodeTaintsConf,
			EngineInstallURL:     "http://fake.url",
			Name:                 "test-node-template",
		},
		GenericConfig:   &testNodeTemplateGenericConf,
		rawDriverConfig: configMap,
	}
	testExpandedGenericNodeTemplateConf = &NodeTemplate{
		NodeTemplate: managementClient.NodeTemplate{
			Driver:               "mycustomdriver",
			UseInternalIPAddress: &useInternalIP,
			Annotations:          testGenericNodeTemplateAnnotationsConf,
			NodeTaints:           testGenericNodeTemplateNodeTaintsConf,
			EngineInstallURL:     "http://fake.url",
			Name:                 "test-node-template",
		},
		GenericConfig: &testNodeTemplateGenericConf,
	}

	testNodeTemplateSquashGenericConfInterface = map[string]interface{}{
		"annotations":             testGenericNodeTemplateAnnotationsInterface,
		"driver":                  "mycustomdriver",
		"use_internal_ip_address": useInternalIP,
		"engine_install_url":      "http://fake.url",
		"name":                    "test-node-template",
	}

	testNodeTemplateExpandGenericConfInterface = map[string]interface{}{
		"annotations":             testGenericNodeTemplateAnnotationsInterface,
		"node_taints":             testGenericNodeTemplateNodeTaintsInterface,
		"driver":                  "mycustomdriver",
		"use_internal_ip_address": useInternalIP,
		"engine_install_url":      "http://fake.url",
		"name":                    "test-node-template",
		"generic_config":          []interface{}{testGenericNodeTemplateInterface},
	}
}

func TestFlattenGenericNodeTemplate(t *testing.T) {
	cases := []struct {
		Input          *NodeTemplate
		ExpectedOutput map[string]interface{}
	}{
		{
			testGenericNodeTemplateConf,
			testNodeTemplateSquashGenericConfInterface,
		},
	}

	for _, tc := range cases {
		output := schema.TestResourceDataRaw(t, nodeTemplateFields(), map[string]interface{}{})
		err := flattenNodeTemplate(output, tc.Input)
		if err != nil {
			assert.FailNow(t, "[ERROR] on flattener: %#v", err)
		}
		filteredOutput := map[string]interface{}{}
		for k := range tc.ExpectedOutput {
			filteredOutput[k] = output.Get(k)
		}
		if !reflect.DeepEqual(filteredOutput, tc.ExpectedOutput) {
			assert.FailNow(t, "Unexpected output from flattener", "Expected: %#v\nGiven: %#v", tc.ExpectedOutput, filteredOutput)
		}
	}
}

func TestExpandGenericNodeTemplate(t *testing.T) {
	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput *NodeTemplate
	}{
		{
			Input:          testNodeTemplateExpandGenericConfInterface,
			ExpectedOutput: testExpandedGenericNodeTemplateConf,
		},
	}

	for _, tc := range cases {
		inputData := schema.TestResourceDataRaw(t, nodeTemplateFields(), tc.Input)
		output := expandNodeTemplate(inputData)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			assert.FailNow(t, "Unexpected output from flattener", "Expected: %#v\nGiven: %#v", tc.ExpectedOutput, output)
		}
	}
}
