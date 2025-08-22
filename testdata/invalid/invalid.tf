resource "aws_db_instance" "foobar" {
  identifier     = "mydb-instance"
  engine         = "mysql"
  engine_version = "8.0"
  instance_class = "db.t3.micro"

  db_name  = "myapp"
  username = "admin"
  password = "Password123!"

  allocated_storage = 20
  storage_type      = "gp2"

  multi_az = false # No redundancy - single point of failure

  skip_final_snapshot = true
}
