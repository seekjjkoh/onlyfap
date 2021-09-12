resource "google_cloud_run_service" "pongrun" {
  name     = "pongrun"
  location = "asia-southeast1"
  template {
    spec {
      timeout_seconds = 600
      containers {
        image = var.image_url
        env {
          name  = "REDISHOST"
          value = google_redis_instance.rediscache.host
        }
        env {
          name  = "REDISPORT"
          value = google_redis_instance.rediscache.port
        }
        resources {
          limits = {
            memory = "128Mi"
            cpu    = "1"
          }
        }
      }
      container_concurrency = 1
    }
    metadata {
      annotations = {
        "run.googleapis.com/vpc-access-connector" = var.vpc_name
        "autoscaling.knative.dev/maxScale"        = "10"
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.pongrun.location
  project     = google_cloud_run_service.pongrun.project
  service     = google_cloud_run_service.pongrun.name
  policy_data = data.google_iam_policy.noauth.policy_data
}