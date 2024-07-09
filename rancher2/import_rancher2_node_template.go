package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceRancher2NodeTemplateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	client, err := meta.(*Config).ManagementClient()
	if err != nil {
		return []*schema.ResourceData{}, err
	}
	nodeTemplate := &NodeTemplate{}
	err = fetchNodeTemplate(client, d.Id(), nodeTemplate)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	err = flattenNodeTemplate(d, nodeTemplate)
	if err != nil {
		return []*schema.ResourceData{}, err
	}

	return []*schema.ResourceData{d}, nil
}
