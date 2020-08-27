package sentry

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSentryRuleImporter(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	addrID := d.Id()

	log.Printf("[DEBUG] Importing rule using ADDR ID %s", addrID)

	parts := strings.Split(addrID, "/")

	if len(parts) != 3 {
		d.SetId("")
		return nil, nil
	}

	d.Set("organization", parts[0])
	d.Set("project", parts[1])
	d.SetId(parts[2])

	return []*schema.ResourceData{d}, nil
}
