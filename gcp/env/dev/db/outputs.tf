output "db" {
  value = {
    google_sql_database = {
      landin = {
        name = google_sql_database.landin.name
      }
    }
    google_sql_database_instance = {
      landin = {
        name = google_sql_database_instance.landin.name
      }
    }
    google_compute_network = {
      db = {
        name = google_compute_network.db.name
      }
    }
  }
}
