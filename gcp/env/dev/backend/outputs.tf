output "backend" {
  value = {
    google_artifact_registry_repository = {
      landin = {
        repository_id = google_artifact_registry_repository.landin.repository_id
      }
    }
  }
}
