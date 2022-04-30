output "backend" {
  value = {
    google_artifact_registry_repository = {
      landin = {
        repository_id = google_artifact_registry_repository.landin.repository_id
      }
    }
    google_vpc_access_connector = {
      backend = {
        name = google_vpc_access_connector.backend.name
      }
    }
  }
}
