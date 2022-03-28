provider "google" {
  credentials = file("gserviceaccount.json")
  # project     = "landin-337813"
  project = var.project.id
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}
