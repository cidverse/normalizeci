package gitlabci

func (n Normalizer) Denormalize(env map[string]string) map[string]string {
	data := make(map[string]string)

	data["GITLAB_CI"] = "true"

	return data
}
