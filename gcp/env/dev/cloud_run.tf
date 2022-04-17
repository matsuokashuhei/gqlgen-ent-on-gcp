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

resource "google_cloud_run_service_iam_policy" "slack_notifier" {
  location    = google_cloud_run_service.slack_notifier.location
  project     = google_cloud_run_service.slack_notifier.project
  service     = google_cloud_run_service.slack_notifier.name
  policy_data = data.google_iam_policy.slack_notifier.policy_data
}

resource "google_cloud_run_service" "backend" {
  name     = "backend"
  location = var.project.region
  template {
    spec {
      containers {
        image = "${var.project.region}-docker.pkg.dev/${var.project.id}/${google_artifact_registry_repository.landin.repository_id}/backend"
      }
    }
  }
  lifecycle {
    ignore_changes = [
      template[0].spec[0].containers[0].image
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "backend" {
  location    = google_cloud_run_service.backend.location
  project     = google_cloud_run_service.backend.project
  service     = google_cloud_run_service.backend.name
  policy_data = data.google_iam_policy.backend.policy_data
}
