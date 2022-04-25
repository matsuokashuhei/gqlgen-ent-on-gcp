resource "google_storage_bucket" "cloudbuild" {
  name          = join("-", [var.project.id, "cloudbuild"])
  location      = "ASIA"
  force_destroy = true
}

resource "google_storage_bucket" "slack_notifier" {
  name          = join("-", [var.project.id, "slack-notifier"])
  location      = "ASIA"
  force_destroy = true
}

resource "google_storage_bucket_object" "slack_notifier_manifest" {
  name = "manifest.yaml"
  # source   = "slack_notifier/${each.key}"
  content = templatefile("slack_notifier/manifest.tftpl", { project_number = data.google_project.landin.number })
  bucket  = google_storage_bucket.slack_notifier.name
}

data "google_storage_bucket" "frontend" {
  name = var.frontend.google_storage_bucket.frontend.name
}
