# About

The goal of this collection of libraries is to create a production-ready toolkit for building REST APIs quickly.
It accomplishes this by making opinionated architectural decisions aimed at developer productivity, development best practices,
operational effectiveness, and efficiency (in that order).

One thing that makes this particularly unique is explicit endpoint objects that can be easily serialized to documentation 
formats like (RAML, Swagger, etc). This desire for self-documenting code also encourages developers to explicitly document requirements and assumptions.

Our tools of choice are:

* http://www.gorillatoolkit.org/pkg/mux - For routing
* https://github.com/carbocation/interpose - For middleware
* https://github.com/mcuadros/go-defaults - For sane API defaults
* https://github.com/asaskevich/govalidator - For object validations
* https://github.com/go-kit/kit - For interfacing with distributed systems
* https://github.com/spf13/viper - For configuration
* https://github.com/Sirupsen/logrus - For logging

# Dev tools:

* https://github.com/derekparker/delve

## Godep

Install at least version v62 of Godep (github.com/tools/godep).
If you're using a version of Go prior to 1.6 make sure to set:

    GO15VENDOREXPERIMENT=1



To update vendored dependencies:

    godep save ./...



## Github pre-push hook

Enable pre-push hook. This hook will make sure the go code is
formatted, that the basic `go vet` checks pass, and that the code
builds and the tests pass.  To bypass the `go vet` check you can
set the `IGNORE_VET_WARNINGS` environment variable and vet errors
will not block the push.
For example `IGNORE_VET_WARNINGS=true git push`.
To enable the pre-push hook link the bash script:
```shell
$ ln -s ../../scripts/pre-push.bash .git/hooks/pre-push
```

# TODOS:

* Documentation rendering
* Tracing
