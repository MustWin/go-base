

The goal of this collection of libraries it to create a production-ready toolkit for building REST APIs quickly.
It accomplishes this by making opinionated architectural decisions aimed at developer productivity, developer best practices,
operational effectiveness, and efficiency (in that order).

Ont thing that makes this particularly unique is explicit endpoint objects that can be easily serialized to documentation 
formats like (RAML, Swagger, etc). This desire for self-documenting code also encourages developers to explicitly document requirements and assumptions.

Our tools of choice are:

* https://github.com/codegangsta/negroni
* http://www.gorillatoolkit.org/pkg/mux
* https://github.com/mcuadros/go-defaults
* https://github.com/asaskevich/govalidator
* https://github.com/go-kit/kit
* https://github.com/spf13/viper
* https://github.com/spf13/cobra

Dev tools:
* https://github.com/derekparker/delve


TODOS: 

* Logging
* Testing
* Documentation rendering
