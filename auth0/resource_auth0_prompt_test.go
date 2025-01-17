package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPrompt(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccPromptCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt.prompt", "universal_login_experience", "classic"),
					resource.TestCheckResourceAttr("auth0_prompt.prompt", "identifier_first", "false"),
				),
			},
			{
				Config: testAccPromptUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt.prompt", "universal_login_experience", "new"),
					resource.TestCheckResourceAttr("auth0_prompt.prompt", "identifier_first", "true"),
				),
			},
		},
	})
}

const testAccPromptCreate = `
resource "auth0_prompt" "prompt" {
  universal_login_experience = "classic"
  identifier_first = false
}
`

const testAccPromptUpdate = `
resource "auth0_prompt" "prompt" {
  universal_login_experience = "new"
  identifier_first = true
}
`
