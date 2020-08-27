package sentry

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Bool returns a pointer to the bool value.
func Bool(v bool) *bool {
	return &v
}

// Int returns a pointer to the int value.
func Int(v int) *int {
	return &v
}

// checkClientGet returns a `found` bool.
// The following return values are meaningful:
// `true`, `nil` => a resource was successfully found
// `false`, `nil` => a resource was successfully not found
func checkClientGet(resp *http.Response, err error, d *schema.ResourceData) (bool, error) {
	if err != nil {
		if resp != nil && resp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return false, nil
		}

		return false, err
	}

	return true, nil
}
