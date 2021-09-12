resource "google_vpc_access_connector" "connector" {
  name          = var.vpc_name
  project       = "jjkoh95"
  ip_cidr_range = "10.8.0.0/28"
  network       = "default"
  min_instances = 2
  max_instances = 3
  machine_type  = "f1-micro"
  region        = "asia-southeast1"
  provider      = google-beta
}

