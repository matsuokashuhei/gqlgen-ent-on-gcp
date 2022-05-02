data "google_project" "landin" {
  project_id = var.project.id
}

locals {
  services = [
    "artifactregistry.googleapis.com",
    "cloudbuild.googleapis.com",
    "compute.googleapis.com",
    "dns.googleapis.com",
    "iam.googleapis.com",
    "networkmanagement.googleapis.com",
    "pubsub.googleapis.com",
    "run.googleapis.com",
    "secretmanager.googleapis.com",
    "servicenetworking.googleapis.com",
    "sqladmin.googleapis.com",
    "vpcaccess.googleapis.com",
  ]
}

resource "google_project_service" "landin" {
  for_each = toset(local.services)
  project  = data.google_project.landin.id
  service  = each.key
}
