resource "yandex_vpc_network" "this" {
  name = "test"
}

resource "yandex_vpc_subnet" "this" {
  network_id = yandex_vpc_network.this.id
  v4_cidr_blocks = ["10.0.1.0/24"]
  zone = "ru-central1-a"
}

resource "yandex_alb_load_balancer" "this" {
//  name        = "my-load-balancer"

//  network_id  = yandex_vpc_network.this.id

  zone_id   = "ru-central1-a"

  allocation_policy {
    location {
      zone_id   = "ru-central1-as"
      subnet_id = yandex_vpc_subnet.this.id
    }
  }
}

resource "yandex_cm_certificate" "this" {
  managed {
    challenge_type = "DNS_CNAME"
  }
}