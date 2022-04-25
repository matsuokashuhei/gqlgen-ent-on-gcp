resource "google_cloud_run_service" "backend" {
  name     = "backend"
  location = var.project.region
  template {
    spec {
      containers {
        image = "us-docker.pkg.dev/cloudrun/container/hello"
        # image = "${var.project.region}-docker.pkg.dev/${var.project.id}/${google_artifact_registry_repository.landin.repository_id}/backend"
      }
    }
  }
  # lifecycle {
  #   ignore_changes = [
  #     template[0].spec[0].containers[0].image
  #   ]
  # }
}

resource "google_cloud_run_service_iam_policy" "backend" {
  location    = google_cloud_run_service.backend.location
  project     = google_cloud_run_service.backend.project
  service     = google_cloud_run_service.backend.name
  policy_data = data.google_iam_policy.backend.policy_data
}

data "google_iam_policy" "backend" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers"
    ]
  }
}
