locals {
  service_account = {
    "backend" = {
      account_id = "backend"
      roles = [
        "roles/secretmanager.secretAccessor",
      ]
    }
    "cloudbuild" = {
      account_id = "cloudbuild"
      roles = [
        "roles/artifactregistry.writer",
        "roles/cloudbuild.builds.editor",
        "roles/cloudsql.client",
        "roles/compute.networkAdmin",
        "roles/iam.serviceAccountUser",
        "roles/logging.logWriter",
        "roles/run.admin",
        "roles/storage.admin",
      ]
    }
    "slack_notifier" = {
      account_id = "slack-notifier"
      roles = [
        "roles/storage.objectAdmin",
      ]
    }
  }
  # secret = {
  #   "cloudbuild" = {
  #     secret_ids = [
  #       "slack_webhook_url"
  #     ]
  #   }
  # }
}
resource "google_service_account" "landin" {
  for_each     = local.service_account
  account_id   = each.value.account_id
  display_name = each.value.account_id
}

resource "google_project_iam_member" "backend" {
  for_each = toset(local.service_account["backend"].roles)
  project  = data.google_project.landin.project_id
  role     = each.key
  member   = "serviceAccount:${google_service_account.landin["backend"].email}"
}

resource "google_project_iam_member" "cloudbuild" {
  for_each = toset(local.service_account["cloudbuild"].roles)
  project  = data.google_project.landin.project_id
  role     = each.key
  member   = "serviceAccount:${google_service_account.landin["cloudbuild"].email}"
}

resource "google_project_iam_member" "slack_notifier" {
  for_each = toset(local.service_account["slack_notifier"].roles)
  project  = data.google_project.landin.project_id
  role     = each.key
  member   = "serviceAccount:${google_service_account.landin["slack_notifier"].email}"
}
