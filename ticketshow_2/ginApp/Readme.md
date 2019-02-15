# Explicación de las carpetas.

- Servidor va en la instancia 2, inicia el servidor de grpc. Este se conecta a la base de datos de manera remota, que está en la instancia 1.
- grpcClient es el cliente de grpc. Que se conecta con el servidor de la instancia 2.
- front contiene un archivo .go que corre la static file ubicada en ../dist 
- nginx Contiene las configuraciones para el Nginx.
- proto Contiene las confirugaciones del archivo .proto para poder usar el protobuffer
- views Es la parte donde se puede configurar todo lo del fron. En este se está usando el framework de React.js

