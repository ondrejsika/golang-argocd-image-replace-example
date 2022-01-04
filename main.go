package main

import (
	"fmt"
	"regexp"
)

var yaml string = `apiVersion: argoproj.io/v1alpha1
kind: Application
metadata:
  name: foo-bar
  namespace: argocd
spec:
  destination:
    namespace: foo-bar
    server: https://kubernetes.default.svc
  project: default
  source:
    repoURL: https://gitlab.com/foo/bar.git
    targetRevision: HEAD
    path: helm/trainingcrm
    helm:
      releaseName: foo-bar
      values: |
        imageFrontend: registry.com/foo/bar/fe:5d5e8df0-master-5100
        imageBackend: registry.com/foo/bar/be:5d5e8df0-master-5100
        registry: registry.com
        registryUsername: foo
        registryPassword: bar
  syncPolicy:
    automated:
      prune: true
      selfHeal: true
    syncOptions:
    - CreateNamespace=true
`

func replaceImage(s, key, value string) string {
	r := regexp.MustCompile(key + `: +([\w\./:_-]+)`)
	return r.ReplaceAllString(s, key+": "+value)
}

func main() {
	yaml = replaceImage(yaml, "imageFrontend", "new-fe")
	yaml = replaceImage(yaml, "imageBackend", "new-be")
	fmt.Println(yaml)
}
