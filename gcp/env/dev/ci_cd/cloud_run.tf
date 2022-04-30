resource "google_cloud_run_service" "slack_notifier" {
  name     = "slack-notifier"
  location = var.region
  template {
    spec {
      containers {
        image = "us-east1-docker.pkg.dev/gcb-release/cloud-build-notifiers/slack:slack-1.14.0"
        env {
          name  = "CONFIG_PATH"
          value = "gs://${google_storage_bucket.slack_notifier.name}/${google_storage_bucket_object.slack_notifier_manifest.name}"
        }
      }
      service_account_name = data.google_service_account.landin["slack_notifier"].email
    }
  }
}

# resource "google_service_account" "slack_notifier" {
#   account_id   = "slack-notifier"
#   display_name = "slack-notifier"
# }

# resource "google_project_iam_member" "slack_notifier" {
#   role    = "roles/storage.objectAdmin"
#   member  = "serviceAccount:${google_service_account.slack_notifier.email}"
#   project = data.google_project.landin.id
# }

# data "google_iam_policy" "slack_notifier" {
#   binding {
#     role = "roles/secretmanager.secretAccessor"
#     members = [
#       "serviceAccount:${google_service_account.slack_notifier.email}",
#     ]
#   }
# }

# resource "google_secret_manager_secret_iam_policy" "slack_notifier" {
#   for_each    = toset(["slack_webhook_url"])
#   project     = data.google_secret_manager_secret.landin[each.key].project
#   secret_id   = data.google_secret_manager_secret.landin[each.key].secret_id
#   policy_data = data.google_iam_policy.slack_notifier.policy_data
# }
