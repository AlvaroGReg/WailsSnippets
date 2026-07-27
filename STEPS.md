# Snippet Vault: plan de desarrollo

Aplicación de escritorio para guardar, buscar y copiar snippets de código.

## Objetivo del MVP

- Crear, editar y eliminar snippets.
- Guardar título, lenguaje, código y etiquetas.
- Buscar por título, lenguaje, etiquetas o contenido.
- Copiar código al portapapeles.
- Persistir los datos localmente.
- Mantener una interfaz sencilla y agradable.

## Día 1: Wails y la interfaz inicial

### 1. Ejecutar el proyecto

```bash
wails dev
```

Comprueba que la aplicación se abre y localiza estos archivos:

- `app.go`
- `main.go`
- `frontend/src/App.tsx`
- `frontend/src/App.css`

### 2. Crear la primera pantalla

Modifica `App.tsx` para mostrar el nombre de la aplicación:

```tsx
<h1>Snippet Vault</h1>
```

Comprueba que los cambios aparecen con recarga automática.

### 3. Diseñar el modelo de datos

Define en Go una estructura similar a esta:

```go
type Snippet struct {
    ID        int      `json:"id"`
    Title     string   `json:"title"`
    Language  string   `json:"language"`
    Code      string   `json:"code"`
    Tags      []string `json:"tags"`
    CreatedAt string   `json:"createdAt"`
}
```

No añadas todavía usuarios, sincronización ni cuentas.

## Día 2: backend en Go

### 4. Crear el servicio de snippets

Crea `snippet.go` y prepara métodos para:

```go
func (a *App) GetSnippets() []Snippet
func (a *App) CreateSnippet(snippet Snippet) Snippet
func (a *App) UpdateSnippet(snippet Snippet) Snippet
func (a *App) DeleteSnippet(id int) error
```

Wails generará los bindings necesarios para llamar a estos métodos desde TypeScript.

### 5. Implementar persistencia JSON

Guarda los datos en un archivo JSON dentro de la carpeta de configuración del usuario.
Usa `os.UserConfigDir()` para obtener una ubicación apropiada.

Separa la lógica en funciones como:

```go
func loadSnippets() ([]Snippet, error)
func saveSnippets(snippets []Snippet) error
```

Aprende a controlar errores de lectura, escritura y serialización.

### 6. Verificar el ciclo completo

Comprueba manualmente que puedes:

1. Crear un snippet.
2. Guardarlo.
3. Cerrar la aplicación.
4. Abrirla de nuevo.
5. Recuperarlo.

## Día 3: interfaz React

### 7. Dividir la interfaz en componentes

Una estructura posible:

```text
frontend/src/
├── components/
│   ├── Sidebar.tsx
│   ├── SnippetList.tsx
│   ├── SnippetEditor.tsx
│   └── SearchBar.tsx
├── App.tsx
└── App.css
```

### 8. Crear el flujo principal

Incluye:

- Barra lateral con lenguajes o etiquetas.
- Lista de snippets.
- Panel de edición.
- Botón “Nuevo snippet”.
- Botones “Guardar”, “Eliminar” y “Copiar”.

Usa datos falsos al principio para construir la interfaz antes de conectarla con Go.

### 9. Conectar React con Go

Importa los bindings generados por Wails:

```ts
import { GetSnippets, CreateSnippet } from "../wailsjs/go/main/App";
```

Carga los datos al iniciar la aplicación y conecta los botones con los métodos del backend.

## Día 4: búsqueda y acabado

### 10. Añadir búsqueda

Filtra por título, lenguaje, etiquetas y contenido del código.
La búsqueda debe ignorar mayúsculas y minúsculas.

### 11. Añadir copia al portapapeles

```tsx
await navigator.clipboard.writeText(snippet.code);
```

Muestra temporalmente un mensaje como “Copiado”.

### 12. Mejorar los estados de la aplicación

Incluye:

- Estado vacío.
- Indicador de guardado.
- Mensajes de error.
- Confirmación antes de eliminar.
- Estado de carga.

### 13. Pulir el diseño

Añade tema oscuro, espaciado consistente, estados hover y una presentación clara del código.
El resaltado de sintaxis puede dejarse para el final.

## Orden mínimo recomendado

Implementa las funcionalidades en este orden:

1. Lista de snippets de prueba.
2. Crear un snippet.
3. Editar un snippet.
4. Eliminar un snippet.
5. Guardar y cargar desde JSON.
6. Buscar snippets.
7. Copiar código.
8. Pulir estilos.
9. Generar una versión distribuible.

Si el tiempo se acaba después del punto 5, ya tendrás un MVP válido.

## Checklist de finalización

- [ ] La aplicación se inicia con `wails dev`.
- [ ] Se pueden crear snippets.
- [ ] Se pueden editar y eliminar.
- [ ] Los datos sobreviven al reinicio.
- [ ] La búsqueda funciona.
- [ ] El código se puede copiar.
- [ ] Los errores se muestran al usuario.
- [ ] El README contiene capturas y pasos de instalación.
- [ ] Se ha generado al menos un binario.

## README y portfolio

Incluye en el README:

- Captura de pantalla.
- Descripción y funcionalidades.
- Tecnologías utilizadas.
- Cómo ejecutar y compilar el proyecto.
- Decisiones técnicas.
- Mejoras futuras.

Descripción sugerida:

> Aplicación de escritorio multiplataforma desarrollada con Wails, Go, React y TypeScript para gestionar snippets de código localmente con búsqueda, etiquetas y persistencia offline.

## Funcionalidades opcionales

Solo añádelas si el MVP ya está terminado:

- Resaltado de sintaxis.
- Atajo `Ctrl/Cmd + K` para buscar.
- Importar y exportar snippets.
- Base de datos SQLite.
- Fijar snippets favoritos.
- Exportar un snippet como archivo.
