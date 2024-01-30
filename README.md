# DU Capstone Pipeline Infrastructure


## Terraform Commands
terraform init
terraform plan -var-file="$TEAM/develop.tfvars"
terraform apply -var-file="$TEAM/develop.tfvars"
terraform destroy -var-file="$TEAM/develop.tfvars"
