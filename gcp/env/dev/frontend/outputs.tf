output "frontend" {
  value = {
    google_storage_bucket = {
      frontend = {
        name = google_storage_bucket.frontend.name
      }
    }
  }
}
