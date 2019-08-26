package printers

func init() {
	disallowedPackagePrefixes = append(disallowedPackagePrefixes,
		"github.com/openshift/openshift-apiserver/pkg/apps/apis/",
		"github.com/openshift/openshift-apiserver/pkg/authorization/apis/",
		"github.com/openshift/openshift-apiserver/pkg/build/apis/",
		"github.com/openshift/openshift-apiserver/pkg/image/apis/",
		"github.com/openshift/openshift-apiserver/pkg/oauth/apis/",
		"github.com/openshift/openshift-apiserver/pkg/project/apis/",
		"github.com/openshift/openshift-apiserver/pkg/quota/apis/",
		"github.com/openshift/openshift-apiserver/pkg/route/apis/",
		"github.com/openshift/openshift-apiserver/pkg/security/apis/",
		"github.com/openshift/openshift-apiserver/pkg/template/apis/",
		"github.com/openshift/openshift-apiserver/pkg/user/apis/",
	)
}
