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

resource "google_service_account_iam_binding" "cloudbuild" {
  service_account_id = data.google_service_account.compute.id
  role               = "roles/iam.serviceAccountUser"
  members = [
    "serviceAccount:${data.google_project.landin.number}@cloudbuild.gserviceaccount.com"
  ]
}

resource "google_service_account" "slack_notifier" {
  account_id   = "slack-notifier"
  display_name = "Cloud Run Pub/Sub Invoker"
}

resource "google_service_account" "backend" {
  account_id   = "cloud-run"
  display_name = "Cloud Run"
}

