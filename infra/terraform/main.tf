resource "aws_instance" "promo_proxy" {
  ami                    = "ami-05d34d340fb1d89e5"
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.promo_proxy.id]
  key_name               = "promo-proxy"
  tags                   = {
    Name  = "Promo Proxy"
    Owner = "Max Klymenko"
  }
}

resource "aws_security_group" "promo_proxy" {
  name                   = "promo_proxy"
  description            = "promo proxy sg"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    ipv6_cidr_blocks = ["::/0"]
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = ""
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [var.ssh_cidr_blocks]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name  = "Promo Proxy"
    Owner = "Max Klymenko"
  }
}
