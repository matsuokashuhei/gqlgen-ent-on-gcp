output "project" {
  value = {
    google_project = {
      landin = {
        project_id = data.google_project.landin.project_id
      }
    }
  }
}
