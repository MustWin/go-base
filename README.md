

The goal of this collection of libraries is to create a production-ready toolkit for building REST APIs quickly.
It accomplishes this by making opinionated architectural decisions aimed at developer productivity, development best practices,
operational effectiveness, and efficiency (in that order).

One thing that makes this particularly unique is explicit endpoint objects that can be easily serialized to documentation 
formats like (RAML, Swagger, etc). This desire for self-documenting code also encourages developers to explicitly document requirements and assumptions.

Our tools of choice are:

* https://github.com/carbocation/interpose - For middleware
* http://www.gorillatoolkit.org/pkg/mux - For routing
* https://github.com/mcuadros/go-defaults - For sane API defaults
* https://github.com/asaskevich/govalidator - For object validations
* https://github.com/go-kit/kit - For interfacing with distributed systems
* https://github.com/spf13/viper - For configuration
https://github.com/Sirupsen/logrus - For logging

Dev tools:
* https://github.com/derekparker/delve


TODOS: 

* Documentation rendering
* Tracing
