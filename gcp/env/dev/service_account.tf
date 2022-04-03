data "google_service_account" "compute" {
  account_id = "339833052610-compute@developer.gserviceaccount.com"
}

data "google_iam_policy" "secretmanager" {
  binding {
    role = "roles/secretmanager.secretAccessor"
    members = [
      "serviceAccount:${data.google_service_account.compute.email}"
    ]
  }
}

resource "google_secret_manager_secret_iam_policy" "compute" {
  for_each    = toset(["slack_webhook_url"])
  project     = google_secret_manager_secret.landin[each.key].project
  secret_id   = google_secret_manager_secret.landin[each.key].secret_id
  policy_data = data.google_iam_policy.secretmanager.policy_data
}
