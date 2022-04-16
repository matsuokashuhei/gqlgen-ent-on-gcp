resource "google_dns_managed_zone" "landin" {
  name       = "landin"
  dns_name   = var.google_dns.dns_name
  visibility = "public"
}

resource "google_dns_record_set" "txt" {
  name         = resource.google_dns_managed_zone.landin.dns_name
  type         = "TXT"
  ttl          = 3600
  managed_zone = resource.google_dns_managed_zone.landin.name
  rrdatas      = [var.google_dns.txt_rrdata]
}

resource "google_dns_record_set" "frontend" {
  name         = resource.google_dns_managed_zone.landin.dns_name
  type         = "A"
  ttl          = 300
  managed_zone = resource.google_dns_managed_zone.landin.name
  rrdatas      = [google_compute_global_address.frontend.address]
}
