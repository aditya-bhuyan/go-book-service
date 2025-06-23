// swagger-ui/swagger-initializer.js
window.onload = function() {
  // Customize to point to your swagger.json
  window.ui = SwaggerUIBundle({
    url: 'swagger.json',
    dom_id: '#swagger-ui',
    presets: [SwaggerUIBundle.presets.apis],
    layout: "BaseLayout"
  });
};
