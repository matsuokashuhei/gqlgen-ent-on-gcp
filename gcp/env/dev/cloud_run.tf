resource "google_cloud_run_service" "slack_notifier" {
  name     = "slack-notifier"
  location = var.project.region
  template {
    spec {
      containers {
        image = "us-east1-docker.pkg.dev/gcb-release/cloud-build-notifiers/slack:slack-1.14.0"
        env {
          name  = "CONFIG_PATH"
          value = "gs://${google_storage_bucket.slack_notifier.name}/${google_storage_bucket_object.slack_notifier["manifest.yaml"].name}"
        }
      }
    }
  }
}
