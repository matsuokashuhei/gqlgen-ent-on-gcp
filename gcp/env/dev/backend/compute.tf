resource "google_compute_managed_ssl_certificate" "backend" {
  provider = google-beta
  name     = "backend"
  managed {
    domains = [join(".", ["api", data.google_dns_managed_zone.landin.dns_name])]
  }
}

resource "google_compute_region_network_endpoint_group" "backend" {
  provider              = google-beta
  name                  = "backend"
  network_endpoint_type = "SERVERLESS"
  region                = var.project.region
  cloud_run {
    service = google_cloud_run_service.backend.name
  }
}

resource "google_compute_backend_service" "backend" {
  name        = "backend"
  protocol    = "HTTP"
  port_name   = "http"
  timeout_sec = 30
  backend {
    group = google_compute_region_network_endpoint_group.backend.id
  }
}

resource "google_compute_url_map" "backend" {
  name            = "backend"
  default_service = google_compute_backend_service.backend.id
}

resource "google_compute_target_https_proxy" "backend" {
  name    = "backend"
  url_map = google_compute_url_map.backend.id
  ssl_certificates = [
    google_compute_managed_ssl_certificate.backend.id
  ]
}

# resource "google_compute_global_address" "backend" {
#   name = "backend"
# }

# resource "google_compute_global_forwarding_rule" "backend" {
#   name       = "backend"
#   target     = google_compute_target_https_proxy.backend.id
#   port_range = "443"
#   ip_address = google_compute_global_address.backend.address
# }
