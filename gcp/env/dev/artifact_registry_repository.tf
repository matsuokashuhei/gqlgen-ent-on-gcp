resource "google_artifact_registry_repository" "landin" {
  for_each = toset(["backend"])

  provider = google-beta

  location      = var.project.region
  repository_id = each.key
  format        = "DOCKER"
}
