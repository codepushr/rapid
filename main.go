package rapid

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"

	"github.com/labstack/echo"
)

// Rapid struct
type Rapid struct {
	Address    string
	Port       uint
	HideBanner bool
	httpServer *echo.Echo
}

// Constants
const (
	version = "1.0.0"
	website = "https://github.com/codepushr/rapid"
	banner  = `                 _     _ 
 _ __ __ _ _ __ (_) __| |
| '__/ _` + "`" + ` | '_ \| |/ _` + "`" + ` |
| | | (_| | |_) | | (_| |
|_|  \__,_| .__/|_|\__,_|
          |_| v%s
`
)

// New func
func New() (r *Rapid) {
	return &Rapid{
		Address:    "127.0.0.1",
		Port:       6700,
		httpServer: echo.New(),
	}
}

// Start func
func (r *Rapid) Start() error {
	r.httpServer.HideBanner = true
	r.httpServer.HidePort = true

	// Parse config file

	// Everything else ...
	// Handlers
	// Routes
	// Bindings
	// ...
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "model/user_model.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Scope.Objects)
	ast.Inspect(file, func(x ast.Node) bool {
		s, ok := x.(*ast.StructType)
		if !ok {
			return true
		}
		for _, field := range s.Fields.List {
			fmt.Printf("Field: %s\n", field.Names[0].Name)
			fmt.Printf("Tag:   %s\n", field.Tag.Value)
		}
		return false
	})

	if !r.HideBanner {
		fmt.Printf(banner, version+"\n")
	}

	return r.startServer()
}

func (r *Rapid) startServer() error {
	return r.httpServer.Start(fmt.Sprintf("%s:%s", r.Address, strconv.Itoa(int(r.Port))))
}
