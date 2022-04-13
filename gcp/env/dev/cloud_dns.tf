resource "google_dns_managed_zone" "landin" {
  name       = "landin"
  dns_name   = var.google_dns.dns_name
  visibility = "public"
}

resource "google_dns_record_set" "txt" {
  managed_zone = resource.google_dns_managed_zone.landin.name
  name         = resource.google_dns_managed_zone.landin.dns_name
  type         = "TXT"
  ttl          = 3600
  rrdatas      = [var.google_dns.txt_rrdata]
}
