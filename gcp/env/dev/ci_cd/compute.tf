data "google_compute_network" "db" {
  name = var.db.google_compute_network.db.name
}
