data "google_service_account" "compute" {
  account_id = "${data.google_project.landin.number}-compute@developer.gserviceaccount.com"
}

data "google_iam_policy" "secretmanager" {
  binding {
    role = "roles/secretmanager.secretAccessor"
    members = [
      "serviceAccount:${data.google_service_account.compute.email}"
    ]
  }
}

resource "google_service_account" "slack_notifier" {
  account_id   = "slack-notifier"
  display_name = "Cloud Run Pub/Sub Invoker"
}
