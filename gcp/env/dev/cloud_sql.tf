# resource "google_sql_database_instance" "landin" {
#   name                = "landin"
#   database_version    = "MYSQL_8_0"
#   region              = var.project.region
#   deletion_protection = false
#   settings {
#     tier = "db-f1-micro"
#   }
# }

# resource "google_sql_database" "landin" {
#   name     = "landin"
#   instance = resource.google_sql_database_instance.landin.name
# }

# resource "google_sql_user" "landin" {
#   for_each = var.google_sql_user
#   name     = each.value.name
#   instance = resource.google_sql_database_instance.landin.name
#   host     = "%"
#   password = each.value.password
# }
