# Static assets

## Swagger UI

Run to update:

```shell
(
    cd swagger-ui-dist \
    && rm swagger-ui* favicon-* \
    && wget https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js \
    && wget https://unpkg.com/swagger-ui-dist/swagger-ui-standalone-preset.js \
    && wget https://unpkg.com/swagger-ui-dist/swagger-ui.css \
    && wget https://unpkg.com/swagger-ui-dist/favicon-16x16.png \
    && wget https://unpkg.com/swagger-ui-dist/favicon-32x32.png \
    && echo Done
)
```