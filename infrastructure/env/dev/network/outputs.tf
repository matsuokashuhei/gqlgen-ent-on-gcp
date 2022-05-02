output "network" {
  value = {
    google_compute_network = {
      landin = {
        name = google_compute_network.landin.name
      }
    }
  }
}
