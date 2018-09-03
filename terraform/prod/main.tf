locals {
  instance_name = "grpc-server"
}

data "template_file" "template_file" {
  template = "${file("container-manifest.yaml")}"
  vars {
    instance_name = "${local.instance_name}"
    project =  "${var.project}"
    image_name =  "${var.image_name}"
  }
}

resource "google_compute_address" "api_address" {
  name = "api-address"
}


resource "google_compute_instance" "api" {
  name = "${local.instance_name}"
  machine_type = "f1-micro"
  zone = "${var.zone}"

  network_interface {
    network = "default"
    access_config {
      nat_ip = "${google_compute_address.api_address.address}"
    }
  }

  labels {
    container-vm = "cos-stable-68-10718-86-0"
  }
  
  metadata {
    gce-container-declaration = "${data.template_file.template_file.rendered}"
  }

  
  boot_disk {
    initialize_params {
      image = "https://www.googleapis.com/compute/v1/projects/cos-cloud/global/images/cos-stable-68-10718-86-0"
    }
  }
  
  service_account {
    scopes = [
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write" ,
      "https://www.googleapis.com/auth/service.management.readonly" ,
      "https://www.googleapis.com/auth/servicecontrol" ,
      "https://www.googleapis.com/auth/trace.append" ,
      "https://www.googleapis.com/auth/monitoring.write" ,
    ]
  }
}
