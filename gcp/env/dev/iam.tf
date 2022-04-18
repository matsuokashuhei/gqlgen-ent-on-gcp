data "google_iam_policy" "slack_notifier" {
  binding {
    role = "roles/run.invoker"
    members = [
      "serviceAccount:${google_service_account.slack_notifier.email}"
    ]
  }
}

data "google_iam_policy" "backend" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers"
    ]
  }
}
