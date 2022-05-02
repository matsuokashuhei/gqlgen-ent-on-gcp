include {
  path = find_in_parent_folders()
}

dependency "frontend" {
  config_path = "../frontend"
}

dependency "backend" {
  config_path = "../backend"
}

dependency "db" {
  config_path = "../db"
}

dependency "network" {
  config_path = "../network"
}

dependency "project" {
  config_path = "../project"
}

inputs = {
  frontend = dependency.frontend.outputs.frontend
  backend  = dependency.backend.outputs.backend
  network  = dependency.network.outputs.network
  db       = dependency.db.outputs.db
  project  = dependency.project.outputs.project
}
