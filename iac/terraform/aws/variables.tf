variable "region" {
  description = "AWS Region"
  type        = string
  default     = "sa-east-1"
}

variable "aws_account_id" {
  description = "AWS Account ID number"
  type        = number
  default     = 477408701252
}

variable "deploy_user_arn" {
  description = "ARN for IAM user with permission on infrastructure actions"
  type        = string
  default     = "arn:aws:iam::477408701252:user/rodolfo.nascimento.azevedo"
}

variable "domain_name" {
  description = "Ajuda domain name"
  type        = string
  default     = "ajuda.academy"
}

variable "application_domain_name" {
  description = "Domain name for this application"
  type        = string
}

variable "default_vpc_id" {
  description = "Default vpc id"
  type        = string
  default     = "vpc-312e0c58"
}

variable "default_subnets_ids" {
  description = "Default subnets ids"
  type        = list(string)
  default     = ["subnet-332e0c5a", "subnet-302e0c59", "subnet-0cd4814a"]
}

variable "route53_zone_id" {
  description = "Zone Id for Ajuda on Route53"
  type        = string
  default     = "Z071451172ANOXED8ZXY"
}

variable "log_group_name" {
  description = "CloudWatch log group name"
  type        = string
}

variable "ecr_repository_name" {
  description = "ECR repository name"
  type        = string
}

variable "ecr_policy_name" {
  description = "ECR policy name"
  type        = string
}

variable "ecs_cluster_name" {
  description = "Cluster name on ECS"
  type        = string
}

variable "ecs_service_name" {
  description = "Service name on cluster ECS"
  type        = string
}

variable "ecs_task_role_arn" {
  description = "ARN for ECS task role that gives permission to tasks"
  type        = string
  default     = "arn:aws:iam::477408701252:role/ecsTaskExecutionRole"
}

variable "alb_name" {
  description = "Name of application load balancer"
  type        = string
}

variable "alb_security_group_name" {
  description = "Name of security group for ELB"
  type        = string
}

variable "target_group_name" {
  description = "Name of target group to associate with the load balancer"
  type        = string
}

variable "ecs_security_group_name" {
  description = "Name of security group for ECS service"
  type        = string
}

variable "ecs_service_container_name" {
  description = "Container name on cluster ECS inside a service"
  type        = string
}

variable "ecs_service_container_port" {
  description = "Container port on cluster ECS inside a service"
  type        = number
  default     = 8080
}

variable "ecs_service_container_cpu" {
  description = "Container CPU capacity"
  type        = number
}

variable "ecs_service_container_memory" {
  description = "Container RAM memory capacity"
  type        = number
}

variable "ecs_service_role_for_ecs_arn" {
  description = "ARN for role to enable Amazon ECS to manage your cluster"
  type        = string
  default     = "arn:aws:iam::477408701252:role/aws-service-role/ecs.amazonaws.com/AWSServiceRoleForECS"
}

variable "ecs_autoscaling_min_capacity" {
  description = "Min capacity for autoscaling"
  type        = number
}

variable "ecs_autoscaling_max_capacity" {
  description = "Max capacity for autoscaling"
  type        = number
}

variable "ecs_autoscaling_memory_target_value" {
  description = "Target memory percentage to start autoscaling"
  type        = number
}

variable "ecs_autoscaling_cpu_target_value" {
  description = "Target cpu percentage to start autoscaling"
  type        = number
}

variable "rds_db_identifier" {
  description = "RDS database identifier"
  type        = string
}

variable "rds_db_name" {
  description = "RDS database name"
  type        = string
}

variable "rds_db_parameter_group_name" {
  description = "RDS database default parameter group"
  type        = string
}

variable "rds_default_subnet_group_name" {
  description = "RDS database default subnet group name"
  type        = string
}

variable "rds_db_instance_class" {
  description = "RDS database instance class"
  type        = string
}

variable "rds_db_port" {
  description = "RDS database port"
  type        = number
}

variable "rds_db_user" {
  description = "RDS database user"
  type        = string
}

variable "rds_db_subnet_group_name" {
  description = "RDS database subnet group name"
  type        = string
}

variable "rds_db_sg_name" {
  description = "RDS database security group name"
  type        = string
}

variable "rds_db_secret_name" {
  description = "RDS database secret name"
  type        = string
}
