resource "google_dns_managed_zone" "landin" {
  name       = "landin"
  dns_name   = var.google_dns_managed_zone.dns_name
  visibility = "public"
}

# resource "google_dns_record_set" "landin" {
#   for_each     = var.google_dns_record_set
#   name         = resource.google_dns_managed_zone.landin.dns_name
#   type         = upper(each.key)
#   ttl          = each.value.ttl
#   managed_zone = resource.google_dns_managed_zone.landin.name
#   rrdatas      = each.value.rrdatas
# }

resource "google_dns_record_set" "frontend" {
  name         = resource.google_dns_managed_zone.landin.dns_name
  type         = "A"
  ttl          = 300
  managed_zone = resource.google_dns_managed_zone.landin.name
  rrdatas      = [google_compute_global_address.frontend.address]
}
