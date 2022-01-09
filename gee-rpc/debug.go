package gee_rpc

import (
	"fmt"
	"html/template"
	"net/http"
)

const debugText = `<html>
	<body>
	<title>GeeRPC Services</title>
	{{range .}}
	<hr>
	Service {{.Name}}
	<hr>
		<table>
		<th align=center>Methods</th><th align=center>Calls</th>
		{{range $name, $mtype := .Methods}}
			<tr>
			<td align=left font=fixed>{{$name}}({{$mtype.ArgType}}, {{$mtype.ReplyType}}) error</td>
			<td align=center>{{$mtype.NumCalls}}</td>
			</tr>
		{{end}}
		</table>
	{{end}}
	</body>
	</html>`

var debug = template.Must(template.New("RPC debug").Parse(debugText))

type debugHTTP struct {
	*Server
}

type debugService struct {
	Name    string
	Methods map[string]*methodType
}

// ServeHTTP Runs at /debug/geerpc
func (server *debugHTTP) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var services []debugService
	server.serviceMap.Range(
		func(namei, svci interface{}) bool {
			svc := svci.(*service)
			services = append(services, debugService{
				Name:    namei.(string),
				Methods: svc.method,
			})
			return true
		},
	)
	if err := debug.Execute(w, services); err != nil {
		_, _ = fmt.Fprintln(w, "rpc: error executing template:", err.Error())
	}
}
