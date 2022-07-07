package main

import (
	"bytes"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var (
	mdProcessor = goldmark.New(
		goldmark.WithExtensions(extension.Linkify),
		goldmark.WithRendererOptions(
			html.WithUnsafe()))

	basicTemplate *template.Template
)

func init() {
	t, err := template.ParseFiles("templates/basic.html")
	if err != nil {
		log.Fatalf("failed to parse template basic.html: %v", err)
	}

	basicTemplate = t
}

type ServerConfig struct {
	RootPath   string
	ListenAddr string
	ListenPort string
}

func main() {

	conf := ServerConfig{}

	flag.StringVar(&conf.ListenAddr, "listen", "0.0.0.0", "ip address to listen for new connections")
	flag.StringVar(&conf.ListenPort, "port", "8888", "port to listen for new connections")
	flag.StringVar(&conf.RootPath, "root", "./docs", "path to directory containing files to serve")
	flag.Parse()

	log.Printf("starting server at http://%s:%s", conf.ListenAddr, conf.ListenPort)

	root := NewMarkdownServer(http.Dir(conf.RootPath))

	log.Fatal(http.ListenAndServe(conf.ListenAddr+":"+conf.ListenPort, root))
}

type MarkdownServer struct {
	root       http.FileSystem
	fileServer http.Handler
}

func NewMarkdownServer(dir http.Dir) http.Handler {

	s := MarkdownServer{
		root:       dir,
		fileServer: http.FileServer(dir),
	}

	return &s
}

func (mds *MarkdownServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlPath := path.Clean(ensureLeadingSlash(r.URL.Path))

	f, canRead := mds.tryFileAndIndexMD(urlPath)
	if !canRead {
		mds.fileServer.ServeHTTP(w, r)
		return
	}
	defer f.Close()

	source, err := io.ReadAll(f)
	if err != nil {
		log.Printf("failed to read already opened markdown %s: %v", urlPath, err)
		mds.fileServer.ServeHTTP(w, r)
		return
	}

	var buf bytes.Buffer
	if err := mdProcessor.Convert(source, &buf); err != nil {
		log.Printf("failed to convert markdown: %v", err)
		mds.fileServer.ServeHTTP(w, r)
		return
	}

	data := map[string]any{
		"Title":   "sample document",
		"Content": template.HTML(buf.String()),
	}

	err = basicTemplate.Execute(w, data)
	if err != nil {
		log.Printf("failed to execute basicTemplat: %v", err)
	}
	// w.Write(buf.Bytes())
}

func (mds *MarkdownServer) tryFileAndIndexMD(urlPath string) (http.File, bool) {

	f, canRead := mds.openWithMDSuffix(urlPath)
	if canRead {
		return f, canRead
	}

	return mds.openWithMDSuffix(urlPath + "/index")
}

func (mds *MarkdownServer) openWithMDSuffix(urlPath string) (http.File, bool) {

	mdName := urlPath + ".md"

	f, err := mds.root.Open(mdName)
	if err != nil {
		return nil, false
	}

	d, err := f.Stat()
	if err != nil {
		return nil, false
	}

	if d.IsDir() {
		return nil, false
	}

	return f, true
}

func ensureLeadingSlash(path string) string {
	if strings.HasPrefix(path, "/") {
		return path
	}

	return "/" + path
}
