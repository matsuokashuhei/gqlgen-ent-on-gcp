resource "google_secret_manager_secret" "landin" {
  for_each  = var.secret_manager
  secret_id = each.key
  labels = {
    label = each.value.label
  }
  replication {
    automatic = true
  }
}

resource "google_secret_manager_secret_version" "landin" {
  for_each    = var.secret_manager
  secret      = google_secret_manager_secret.landin[each.key].id
  secret_data = each.value.value
}

resource "google_secret_manager_secret_iam_policy" "compute" {
  for_each    = toset(["slack_webhook_url"])
  project     = google_secret_manager_secret.landin[each.key].project
  secret_id   = google_secret_manager_secret.landin[each.key].secret_id
  policy_data = data.google_iam_policy.secretmanager.policy_data
}
