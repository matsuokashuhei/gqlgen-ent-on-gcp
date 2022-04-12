resource "google_artifact_registry_repository" "landin" {
  provider = google-beta

  location      = var.project.region
  repository_id = "docker-repository"
  format        = "DOCKER"
}
