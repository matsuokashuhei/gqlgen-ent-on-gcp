data "google_dns_managed_zone" "landin" {
  name = var.dns.google_dns_managed_zone.landin.name
}

# resource "google_dns_record_set" "backend" {
#   name         = join(".", ["api", data.google_dns_managed_zone.landin.dns_name])
#   type         = "A"
#   ttl          = 300
#   managed_zone = data.google_dns_managed_zone.landin.name
#   rrdatas      = [google_compute_global_address.backend.address]
# }
