resource "google_compute_backend_bucket" "frontentd" {
  name        = "frontend"
  bucket_name = google_storage_bucket.frontend.name
}

resource "google_compute_url_map" "frontend" {
  name            = "frontend"
  default_service = google_compute_backend_bucket.frontentd.self_link
}

resource "google_compute_managed_ssl_certificate" "frontend" {
  name = "frontend"
  managed {
    domains = [data.google_dns_managed_zone.landin.dns_name]
  }
}

resource "google_compute_target_https_proxy" "frontend" {
  name             = "frontend"
  url_map          = google_compute_url_map.frontend.self_link
  ssl_certificates = [google_compute_managed_ssl_certificate.frontend.self_link]
}

resource "google_compute_global_address" "frontend" {
  name         = "frontend"
  ip_version   = "IPV4"
  address_type = "EXTERNAL"
}

resource "google_compute_global_forwarding_rule" "frontend" {
  name       = "frontend"
  target     = google_compute_target_https_proxy.frontend.self_link
  ip_address = google_compute_global_address.frontend.address
  port_range = "443"
}
