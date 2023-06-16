# Usa la imagen base de Go
FROM golang:1.20.5-alpine3.17

# Establece el directorio de trabajo en el que se copiará el código
WORKDIR /app
RUN apk update && \
    apk add --no-cache gcc musl-dev

# Copia el archivo go.mod y go.sum al directorio de trabajo
COPY go.mod go.mod
COPY go.sum go.sum

# Copia el resto de los archivos al directorio de trabajo
COPY src/ ./src/

# Descarga e instala las dependencias del módulo

RUN go mod download
COPY src/config/.env config/.env
# Compila el código en un binario ejecutable
ENV CGO_ENABLED=1
RUN go build -o src/main ./src

# Expone el puerto en el que se ejecuta la aplicación
EXPOSE 3000

# Establece el directorio de trabajo en el que se ejecutará la aplicación
#WORKDIR /app/src
# Ejecuta la aplicación cuando se inicie el contenedor
CMD ["./src/main"]
