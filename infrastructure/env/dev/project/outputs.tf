output "project" {
  value = {
    google_project = {
      landin = {
        project_id = data.google_project.landin.project_id
      }
    }
    google_service_account = {
      "backend" = {
        account_id = google_service_account.landin["backend"].account_id
      }
      "cloudbuild" = {
        account_id = google_service_account.landin["cloudbuild"].account_id
      }
      "slack_notifier" = {
        account_id = google_service_account.landin["slack_notifier"].account_id
      }
    }
  }
}
