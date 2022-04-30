data "google_compute_network" "landin" {
  name = var.network.google_compute_network.landin.name
}

resource "google_compute_global_address" "db" {
  provider      = google-beta
  name          = "db"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = data.google_compute_network.landin.id
}

resource "google_service_networking_connection" "db" {
  provider                = google-beta
  network                 = data.google_compute_network.landin.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.db.name]
}
