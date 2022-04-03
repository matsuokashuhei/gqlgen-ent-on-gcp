resource "google_storage_bucket" "slack_notifier" {
  name          = join("-", [var.project.id, "slack-notifier"])
  location      = "ASIA"
  force_destroy = true
}

resource "google_storage_bucket_object" "slack_notifier" {
  for_each = toset(["manifest.yaml"])

  name   = each.key
  source = "slack_notifier/${each.key}"
  bucket = google_storage_bucket.slack_notifier.name
}
