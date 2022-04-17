variable "project" {
  type = map(any)
}

variable "github" {
  type = map(any)
}

variable "secret_manager" {
  type = map(any)
}

variable "google_dns_managed_zone" {
  type = map(any)
}

variable "google_dns_record_set" {
  type = map(any)
}

variable "google_sql_user" {
  type = map(any)
}
