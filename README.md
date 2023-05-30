# UNICyber-BackEnd
Sistema para el prestamo de equipos presentes en las Salas de UNICA

### Operaciones que debehacer

**Operaciones de Cuentas** ✅
- Coneccion para SignUp de Usuarios UNICA
- Inicio de Sesion SignIn usando UNIAccounts
  
**Operaciones de UI**
- Creacion de Aulas para un Usuario
- Eliminar Aula
- Cambiar datos de Aula (nombre)
- Actualizar posicion de una Aula
  
- Añadir Nueva Compu/Impresora/Scanner a una Aula
- Elimanar Compu de Una Aula
- Actualizar posicionde Compu
- Cambiar datos de Compu nombre/tipo

**Operaciones de Alumnos**

- Crear un Alumno
- Elimiar un Alumno
- Actualizar Un Alumno
- Obtener Datos de Un alumno

**Operaciones de Prestamo**

- Prestar una PC/Impresora/Scanner
- Sancionar, eliminar sancion
- Liberar Prestamo

**Operaciones de Estado de PCs**
- Cambiar de Estado
- Poner un mensaje 

**Operaciones de Estadisticas**
- Pos Sala/Usuario
- Por aula
- Por año
- Por mes
- Por Alumno
- Por Compu

**Operaciones de Mantenimiento**
- Apagar una PC
- Obtener datos de la PC
- Ejecutar programas/encuestas

```bash
source env/local/bin/activate
cd src
flask run

cd src
go run app.go

```

Dowload Modules
```bash
go mod download
```

Build reduce size
```bash
go build -ldflags="-s -w" app.go
```