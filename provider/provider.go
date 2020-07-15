package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/justsimplify/sample-tf-provider/resources"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"redis_host": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("redis_host", ""),
			},
			"redis_port": {
				Type: schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("redis_port", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"redis-object": resources.ResourceServer(),
		},
		ConfigureFunc: configRedis,
	}
}

func configRedis(d *schema.ResourceData) (interface{}, error) {
	m := make(map[string]string)
	m["redis_host"] = d.Get("redis_host").(string)
	m["redis_port"] = d.Get("redis_port").(string)
	return m, nil
}
