package restapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/server-middleware/docui"
	"gopkg.in/yaml.v3"

	"github.com/vitalyshatskikh/go-lib/http/restapi/assets"
)

const (
	specURL           = "/openapi.json"
	swaggerAssetsPath = "/swagger-ui-dist"
	swaggerAssetsURL  = swaggerAssetsPath + "/swagger-ui-bundle.js"
	swaggerPresetURL  = swaggerAssetsPath + "/swagger-ui-standalone-preset.js"
	swaggerStylesURL  = swaggerAssetsPath + "/swagger-ui.css"
	swaggerFavicon32  = swaggerAssetsPath + "/favicon-32x32.png"
	swaggerFavicon16  = swaggerAssetsPath + "/favicon-16x16.png"
)

// OpenAPIHandler returns an http.Handler that serves the OpenAPI spec at
// {basePath}/openapi.json and renders Swagger UI at {basePath}.
//
// NOTE: the basePath must be the same as handler mount point (/docs if empty)
func OpenAPIHandler(basePath string, jsonSpec []byte, next http.Handler) http.Handler {
	router := chi.NewRouter()

	if basePath == "" {
		basePath = docsPath
	}
	basePath = strings.TrimSuffix(basePath, "/")

	if next == nil {
		next = http.NotFoundHandler()
	}

	swaggerHandler := docui.SwaggerUI(
		next,
		docui.WithSpecURL(basePath+specURL),
		docui.WithUIAssetsURL(basePath+swaggerAssetsURL),
		docui.WithSwaggerUIOptions(docui.SwaggerUIOptions{
			SwaggerPresetURL: basePath + swaggerPresetURL,
			SwaggerStylesURL: basePath + swaggerStylesURL,
			Favicon32:        basePath + swaggerFavicon32,
			Favicon16:        basePath + swaggerFavicon16,
		}),
	)

	router.Handle(swaggerAssetsPath+"/*", http.StripPrefix(basePath+"/", http.FileServerFS(assets.SwaggerUIDist)))
	router.Handle("/", swaggerHandler)
	router.Get(specURL, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(jsonSpec)
	})

	return router
}

func parseSpec(spec io.Reader) ([]byte, error) {
	raw, err := io.ReadAll(spec)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec: %w", err)
	}

	if json.Valid(raw) {
		return raw, nil
	}

	var parsed any
	err = yaml.Unmarshal(raw, &parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to parse spec: %w", err)
	}

	jsonBytes, err := json.Marshal(parsed)
	if err != nil {
		return nil, fmt.Errorf("failed to convert spec to JSON: %w", err)
	}

	return jsonBytes, nil
}
