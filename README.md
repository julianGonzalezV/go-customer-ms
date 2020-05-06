# ms-cliente-go
microservice Poc built in golang

En su local cree la siguiente estructura de carpetas para montar el proyecto, enla ruta que ud desee:
xxworkspace
    bin
    src
    pkg

Luego cree la variable de entorno GOPATH apuntando a la carpeta xxworkspace
en el PATH tambien  adicione GOPATH/bin (pendiente por validar en la  documentación)
    https://golang.org/doc/gopath_code.html

# Herramientas 
En este ejemplo nos apalancamos del paquete gorilla mux https://github.com/gorilla/mux para crear APIS mas avanzadas comparado con LinkToSimpleAPI.
    Nota: Importante que note que no se usó un framework para este ejemplo, digamos que hasta ahora no es necesario, además dado que Go se promociona por ser muy rápido no sé que tan recomendable sea adicionar otra capa, que en muchos casos le agrega complejidad al aplicativo.

# Referencia para la realización del presente ejemplo:
https://blog.friendsofgo.tech/posts/como_crear_una_api_rest_en_golang/ 


# Comandos usados
-Para instalar el paquete
$ go get -u github.com/gorilla/mux