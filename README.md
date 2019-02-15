# Non-ticketshow

Proyecto de Sistemas distribuídos, Segundo término 2018-2019

## Integrantes
- Xavier Idrovo Vallejo

## Instalación

Instalando Go
```sh
$ sudo apt-get install golang-go
```

##### Dependencias de Go
Realizar un go get -u [url-dependencia]
Las dependencias a continuación:
```sh
$ "github.com/gin-gonic/contrib/static"
$ "github.com/gin-gonic/gin"
$ "github.com/go-sql-driver/mysql"
$ "google.golang.org/grpc"
$ "google.golang.org/grpc/reflection"
$ "github.com/gin-gonic/contrib/static"
```
##### Nginx
Instalación de nginx
```sh
$ sudo apt-get install nginx-server
```
Reemplazar el archivo default ubicado en 
```sh
$ etc/nginx/sites-available/
```
Por el archivo default de la carpeta de Nginx.

##### Mysql
Como base de datos, se está usando mysql
```sh
$ sudo apt-get install mysql-server
```
