# Configure the dynatrace provider
provider "dynatrace" {
    dt_env_url    = var.environment 
    dt_api_token  = var.token
}

resource "dynatrace_maintenance_window" "windows_patching" {
  name = "windows patches"
  description = "Weekly udpate of windows servers"
  type = "PLANNED" 
  suppression = "DETECT_PROBLEMS_DONT_ALERT"
  scope {
    match {
      type = "HOST"
      tags {
        context = "CONTEXTLESS"
        key = "OS"  
        value = "windows"
      }
    }
    match {
      type = "HOST"
      tags {
        context = "CONTEXTLESS"
        key = "OS"  
        value = "WIN32"
      }
    }
  }
  schedule {
    recurrence_type = "WEEKLY"
    recurrence {
      day_of_week = "THURSDAY"
      start_time = "19:21"
      duration_minutes = 60
    }
    start = "2020-10-20 15:38"
    end = "2020-10-25 15:38"
    zone_id = "America/Chicago"
  }
  
}