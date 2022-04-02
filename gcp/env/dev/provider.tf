provider "google" {
  project = var.project.id
  region  = var.project.region
  zone    = var.project.zone
}

provider "google-beta" {
  project = var.project.id
  region  = var.project.region
}
