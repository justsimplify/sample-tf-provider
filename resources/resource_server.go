package resources

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/justsimplify/sample-tf-provider/utils"
	"io/ioutil"
)

func ResourceServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	headers := m.(map[string]string)

	k := d.Get("key").(string)
	v := d.Get("value").(string)
	url := fmt.Sprintf("http://0.0.0.0:8080/add/%s/%s", k, v)

	_, err := requestCreate("GET", url, headers)
	if err != nil {
		return err
	}

	d.SetId(k)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	k := d.Id()
	headers := m.(map[string]string)

	url := fmt.Sprintf("http://0.0.0.0:8080/get/%s", k)
	r, err := requestCreate("GET", url, headers)
	if err == nil {
		_ = d.Set("value", r.Message)
		return nil
	}

	d.SetId("")
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	d.Partial(true)
	headers := m.(map[string]string)

	if d.HasChange("value") {
		k := d.Id()
		v := d.Get("value").(string)
		url := fmt.Sprintf("http://0.0.0.0:8080/add/%s/%s", k, v)
		_, err := requestCreate("GET", url, headers)
		if err != nil {
			return err
		}
		d.SetPartial("value")
	}
	d.Partial(false)
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	headers := m.(map[string]string)

	k := d.Id()
	url := fmt.Sprintf("http://0.0.0.0:8080/delete/%s", k)

	_, err := requestCreate("GET", url, headers)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

func requestCreate(method, url string, headers map[string]string) (utils.Response, error) {
	res, err := utils.MakeRequest(method, url, headers)
	if err != nil {
		return utils.Response{}, err
	}
	var response utils.Response
	b, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(b, &response)
	if response.Error != nil {
		return utils.Response{}, fmt.Errorf(response.Error.(string))
	}
	return response, nil
}
