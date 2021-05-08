package calc

import (
	"context"
	"embed"
	"io"
	"log"
	"mime"
	"path/filepath"

	gen "goa.design/examples/embed_files/gen/asset"
)

// asset service example implementation.
// The example methods log the requests and return zero values.
type assetsrvc struct {
	logger *log.Logger
}

// NewAsset returns the asset service implementation.
func NewAsset(logger *log.Logger) gen.Service {
	return &assetsrvc{logger}
}

//go:embed public
var public embed.FS

// Public implements public.
func (s *assetsrvc) Public(ctx context.Context, p string) (res *gen.PublicResult, resp io.ReadCloser, err error) {
	f, err := public.Open(filepath.Join("public", p))
	if err != nil {
		return nil, nil, gen.MakeInvalidFilePath(err)
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, nil, gen.MakeInternalError(err)
	}
	var ct *string
	if v := mime.TypeByExtension(filepath.Ext(p)); v != ""{
		ct = &v
	}
	return &gen.PublicResult{
		Length: fi.Size(),
		ContentType: ct,
	}, f, nil
}
