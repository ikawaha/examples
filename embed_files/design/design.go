package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("embed", func() {
	Title("Serving static files from embed.FS via HTTP")
	Server("embed", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("asset", func() {
	Method("public", func() {
		Payload(String, "filepath", func() {
			Description("Path to downloaded file.")
		})

		Result(func() {
			Attribute("length", Int64, func() {
				Description("Length is the downloaded content length in bytes.")
				Example(4 * 1024 * 1024)
			})
			Attribute("content_type", String, func() {
				Description("Content-Type is the content type of the downloaded file.")
				Example("application/json")
			})
			Required("length")
		})

		Error("invalid_file_path", ErrorResult, "Could not locate file for download")
		Error("internal_error", ErrorResult, "Fault while processing download.")

		HTTP(func() {
			GET("/public/{*filename}") // Encode payload in URL path
			SkipResponseBodyEncodeDecode()
			Response(func() {
				Header("length:Content-Length")
				Header("content_type:Content-Type")
			})
			Response("invalid_file_path", StatusNotFound)
			Response("internal_error", StatusInternalServerError)
		})
	})
})
