resource "random_id" "db" {
  byte_length = 4
}

resource "google_sql_database_instance" "landin" {
  name                = "landin-${random_id.db.hex}"
  database_version    = "MYSQL_8_0"
  region              = var.project.region
  deletion_protection = false
  settings {
    tier = "db-f1-micro"
    ip_configuration {
      ipv4_enabled    = false
      private_network = google_compute_network.db.id
    }
  }
  depends_on = [google_service_networking_connection.db]
}

resource "google_sql_database" "landin" {
  name     = "landin"
  instance = resource.google_sql_database_instance.landin.name
}

resource "google_sql_user" "landin" {
  name     = data.google_secret_manager_secret_version.landin["sql_user_name"].secret_data
  instance = resource.google_sql_database_instance.landin.name
  host     = "%"
  password = data.google_secret_manager_secret_version.landin["sql_user_password"].secret_data
}
