package backup

import (
	"regexp"
	"strings"

	api "github.com/appscode/stash/apis/stash/v1alpha1"
	"gopkg.in/ini.v1"
)

var (
	invalidLabelCharRE = regexp.MustCompile(`[^a-zA-Z0-9_]`)
)

func sanitizeLabelName(name string) string {
	return invalidLabelCharRE.ReplaceAllString(name, "_")
}

func sanitizeLabelValue(name string) string {
	return strings.Replace(name, "/", "|", -1)
}

func (c *Controller) JobName(resource *api.Restic) string {
	return sanitizeLabelValue(resource.Namespace + "-" + resource.Name)
}

func (c *Controller) GroupingKeys(resource *api.Restic) map[string]string {
	labels := make(map[string]string)
	labels["app"] = sanitizeLabelValue(c.opt.Workload.Name)
	labels["kind"] = sanitizeLabelValue(c.opt.Workload.Kind)
	labels["namespace"] = resource.Namespace
	labels["stash_config"] = resource.Name
	if cfg, err := ini.LooseLoad(c.opt.PodLabelsPath); err == nil {
		for _, section := range cfg.Sections() {
			for k, v := range section.KeysHash() {
				if k != "pod-template-hash" {
					labels["pod_"+sanitizeLabelName(k)] = sanitizeLabelValue(v)
				}
			}
		}
	}
	return labels
}
