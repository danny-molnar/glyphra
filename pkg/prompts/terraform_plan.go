package prompts

func TerraformPlanPrompt(input string) string {
    return "Summarize this Terraform plan:\n" + input
}
