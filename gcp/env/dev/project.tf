data "google_project" "landin" {
  project_id = var.project.id
}

locals {
  services = [
    "artifactregistry.googleapis.com",
    "cloudbuild.googleapis.com",
    "compute.googleapis.com",
    "dns.googleapis.com",
    "pubsub.googleapis.com",
    "run.googleapis.com",
    "secretmanager.googleapis.com",
  ]
}

resource "google_project_service" "landin" {
  for_each = toset(local.services)
  project  = data.google_project.landin.id
  service  = each.key
  # disable_dependent_services = true
  # disable_on_destroy         = false
}
