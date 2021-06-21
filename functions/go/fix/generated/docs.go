

// Code generated by "mdtogo"; DO NOT EDIT.
package generated

var FixShort = `Fix resources and make them compatible with kpt 1.0.`
var FixLong = `
` + "`" + `fix` + "`" + ` helps you migrate the resources from ` + "`" + `v1alpha1` + "`" + ` format to ` + "`" + `v1` + "`" + ` format.
This is an automated step to migrate kpt packages which are compatible with kpt v0.X.Y
versions of kpt, and make them compatible with kpt 1.0

Here are the automated changes performed by ` + "`" + `fix` + "`" + ` function on ` + "`" + `v1alpha1` + "`" + ` kpt package:

1. The ` + "`" + `packageMetaData` + "`" + ` section will be transformed to ` + "`" + `info` + "`" + ` section.
2. ` + "`" + `upstream` + "`" + ` section(if present), in the ` + "`" + `v1alpha1` + "`" + ` Kptfile is converted to ` + "`" + `upstream` + "`" + `
   and ` + "`" + `upstreamLock` + "`" + ` sections in ` + "`" + `v1` + "`" + ` version of Kptfile.
3. ` + "`" + `dependencies` + "`" + ` section will be removed from the Kptfile.
4. Setters no longer follow the OpenAPI format. The setters and substitutions will be converted
   to simple setter patterns. ` + "`" + `apply-setters` + "`" + ` function is declared in the ` + "`" + `pipeline` + "`" + ` section.
   Setters are configured using [ConfigMap] option.
5. Function annotation from function configs will be removed and corresponding
   function definitions will be [declared in pipeline] section of Kptfile. Reference
   to function config is added via [configPath] option.

Limitations of ` + "`" + `fix` + "`" + ` function:

1. All the functions are treated as mutators by the ` + "`" + `fix` + "`" + ` function while migrating and are added to
   the mutators section in the pipeline. Users must manually go through the functions
   and move the validator functions to the [validators] section in the pipeline section
   of ` + "`" + `v1` + "`" + ` Kptfile.
2. [Openapi validations] and [required setters] feature offered by v0.X.Y setters is
   no longer offered in v1.0 version of kpt. ` + "`" + `fix` + "`" + ` function will remove them.
   Users must write their own validation functions to achieve the functionality.
   Tip: Adding a [starlark function] would be an easier alternative to achieve the
   equivalent validation functionality.
3. If you have used [Starlark runtime] in v0.X, please checkout the new and improved
   [starlark function] and declare it in the pipeline as ` + "`" + `fix` + "`" + ` funtion will remove them.
4. [Auto-setters] feature is deprecated in v1.0 version of kpt. Since the setters are
   migrated to a new and simple declarative version, package consumers can easily
   declare all the setter values and render them all at once.
5. The ` + "`" + `fix` + "`" + ` function does not alter resources in live cluster.
   If you are using the [inventory object] to manage live cluster, please
   refer to [live migrate] docs to perform live migration separately.
`
var FixExamples = `
Let's start with a simple input resource which is compatible with kpt v0.X.Y

  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: my-nginx
  spec:
    replicas: 3 # {"$kpt-set":"replicas"}

Here is the corresponding v1alpha1 Kptfile in the package

  apiVersion: kpt.dev/v1alpha1
  kind: Kptfile
  metadata:
    name: nginx
  openAPI:
    definitions:
      io.k8s.cli.setters.replicas:
        x-k8s-cli:
          setter:
            name: replicas
            value: "3"

Invoke ` + "`" + `fix` + "`" + ` function on the package:

  $ kpt fn eval --image gcr.io/kpt-fn/fix:unstable --include-meta-resources

Here is the transformed resource

  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: my-nginx
  spec:
    replicas: 3 # kpt-set: ${replicas}

Here is the transformed ` + "`" + `v1` + "`" + ` Kptfile:

  apiVersion: kpt.dev/v1
  kind: Kptfile
  metadata:
    name: nginx
  pipeline:
    mutators:
      - image: gcr.io/kpt-fn/apply-setters:v0.1
        configMap:
          replicas: "3"
`
