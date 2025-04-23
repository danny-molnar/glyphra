#!/bin/bash
cat <<EOF > plan.out
resource "aws_s3_bucket" "demo" {
  bucket = "glyphra-test"
  acl    = "private"
}
EOF

echo "✅ Generated dummy Terraform plan at plan.out"
