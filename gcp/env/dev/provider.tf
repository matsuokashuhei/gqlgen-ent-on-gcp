provider "google" {
  project = var.project.id
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}

provider "google-beta" {
  project = var.project.id
  region  = "asia-northeast1"
}
