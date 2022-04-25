resource "google_compute_network" "db" {
  provider = google-beta
  name     = "db"
}

resource "google_compute_global_address" "db" {
  provider      = google-beta
  name          = "db"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 16
  network       = google_compute_network.db.id
}

resource "google_service_networking_connection" "db" {
  provider                = google-beta
  network                 = google_compute_network.db.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.db.name]
}
