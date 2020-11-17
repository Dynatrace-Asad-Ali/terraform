# Configure the dynatrace provider
provider "dynatrace" {
    dt_env_url    = var.environment 
    dt_api_token  = var.token
}

resource "dynatrace_management_zone" "sockshop_prod" {

  name = "sockshop_prod"

  rule{
    type = "SERVICE"
    enabled = true
    propagation_types = [
      "SERVICE_TO_HOST_LIKE", 
      "SERVICE_TO_PROCESS_GROUP_LIKE"
    ]
    condition {
      key {
        attribute = "SERVICE_TAGS"
      }
      comparison_info {
        type = "TAG"
        operator = "EQUALS"
        value          = jsonencode(
            {
                context = "CONTEXTLESS"
                key     = "product"
                value   = "carts"
            }
        )
        negate = false
      }
    }

    condition {
      key {
        attribute = "SERVICE_TAGS"
      }
      comparison_info {
        type = "TAG"
        operator = "EQUALS"
        value = jsonencode(
            {
                context = "CONTEXTLESS"
                key     = "product"
                value   = "sockshop"
            }
        )
        negate = false
      }
    }

  }

  rule{
    type = "SERVICE"
    enabled = true
    propagation_types = [
      "SERVICE_TO_HOST_LIKE", 
      "SERVICE_TO_PROCESS_GROUP_LIKE"
    ]
    condition {
      key {
        attribute = "SERVICE_TAGS"
      }
      comparison_info {
        type = "TAG"
        operator = "EQUALS"
        value = jsonencode(
            {
                context = "CONTEXTLESS"
                key     = "env"
                value   = "prod"
            }
        )
        negate = false
      }
    }

    condition {
      key {
        attribute = "HOST_GROUP_NAME"
      }
      comparison_info {
        type = "STRING"
        operator = "BEGINS_WITH"
        value = jsonencode("simpleapp")
        negate = false
        case_sensitive = false
      }
    }

  }

}
