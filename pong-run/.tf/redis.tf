resource "google_redis_instance" "rediscache" {
  name           = "pongrunredis"
  memory_size_gb = 1
  tier = "BASIC"
  region = "asia-southeast1"
}
