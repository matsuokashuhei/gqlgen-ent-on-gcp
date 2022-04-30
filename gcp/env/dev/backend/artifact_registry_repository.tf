resource "google_artifact_registry_repository" "landin" {
  provider      = google-beta
  location      = var.region
  repository_id = "landin"
  format        = "DOCKER"
}
