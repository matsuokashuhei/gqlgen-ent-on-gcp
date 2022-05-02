resource "google_vpc_access_connector" "backend" {
  provider      = google-beta
  name          = "backend"
  region        = var.region
  ip_cidr_range = "10.8.0.0/28"
  network       = data.google_compute_network.landin.name
}
