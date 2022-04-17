resource "google_compute_backend_bucket" "frontentd" {
  name        = "frontend"
  bucket_name = google_storage_bucket.frontend.name
}

# resource "google_compute_url_map" "frontend" {
#   name            = "frontend"
#   default_service = google_compute_backend_bucket.frontentd.self_link
# }

resource "google_compute_managed_ssl_certificate" "frontend" {
  name = "frontend"

  managed {
    domains = [var.google_dns_managed_zone.dns_name]
  }
}

# resource "google_compute_target_https_proxy" "frontend" {
#   name             = "frontend"
#   url_map          = google_compute_url_map.frontend.self_link
#   ssl_certificates = [google_compute_managed_ssl_certificate.frontend.self_link]
# }

resource "google_compute_global_address" "frontend" {
  name         = "frontend"
  ip_version   = "IPV4"
  address_type = "EXTERNAL"
}

# resource "google_compute_global_forwarding_rule" "frontend" {
#   name       = "frontend"
#   target     = google_compute_target_https_proxy.frontend.self_link
#   ip_address = google_compute_global_address.frontend.address
#   port_range = "443"
# }

# backend
resource "google_compute_global_address" "backend" {
  name = "backend"
}

resource "google_compute_managed_ssl_certificate" "backend" {
  provider = google-beta
  name     = "backend"
  managed {
    domains = [join(".", ["graph", resource.google_dns_managed_zone.landin.dns_name])]
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

# resource "google_compute_url_map" "backend" {
#   name            = "backend"
#   default_service = google_compute_backend_service.backend.id
# }

# resource "google_compute_target_https_proxy" "backend" {
#   name    = "backend"
#   url_map = google_compute_url_map.backend.id
#   ssl_certificates = [
#     google_compute_managed_ssl_certificate.backend.id
#   ]
# }

# resource "google_compute_global_forwarding_rule" "backend" {
#   name       = "backend"
#   target     = google_compute_target_https_proxy.backend.id
#   port_range = "443"
#   ip_address = google_compute_global_address.backend.address
# }
