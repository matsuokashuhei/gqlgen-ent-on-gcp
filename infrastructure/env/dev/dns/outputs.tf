output "dns" {
  value = {
    google_dns_managed_zone = {
      landin = {
        name = google_dns_managed_zone.landin.name
      }
    }
  }
}
