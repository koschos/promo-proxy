variable "region" {
  type    = string
  default = "eu-central-1"

  validation {
    condition     = substr(var.region, 0, 3) == "eu-"
    error_message = "Must be an EUROPE AWS Region, like \"eu-\"."
  }
}

variable "ssh_cidr_blocks" {
  default = "1.1.1.1/32"
}