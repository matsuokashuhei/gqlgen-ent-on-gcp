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

inputs = {
  frontend = dependency.frontend.outputs.frontend
  backend  = dependency.backend.outputs.backend
  db       = dependency.db.outputs.db
}
