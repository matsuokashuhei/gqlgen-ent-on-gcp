resource "google_dns_managed_zone" "landin" {
  name       = "landin"
  dns_name   = var.google_dns_managed_zone.dns_name
  visibility = "public"
}
