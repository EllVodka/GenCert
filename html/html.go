package html

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"training.go/GenCert/cert"
)

type HtmlSaver struct {
	OutputDir string
}

func New(outputDir string) (*HtmlSaver, error) {
	var h *HtmlSaver
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return h, err
	}
	h = &HtmlSaver{
		OutputDir: outputDir,
	}
	return h, nil
}

func (h *HtmlSaver) Save(cert cert.Cert) error {
	t, err := template.New("certificate").Parse(tpl)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("%v.html", cert.LabelCompletion)
	path := path.Join(h.OutputDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, cert)
	if err != nil {
		return err
	}

	fmt.Printf("Saved certificate to '%v'\n", path)

	return nil
}

var tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>{{.LabelTitle}}</title>
		<style>
			body {
				text-align: center;
				font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
			}
			h1 {
				font-size: 3em;
			}
            .gophers{
                height: 10em;
                margin-top: 2em;
            }
            .l{
                float: left;
                margin-left: 5em
            }
            .r{
                float:right;
                margin-right: 5em
            }
            .stamp{
                height: 15em;
                float:right;
                margin-right: 5em
            }
		</style>
	</head>
	<body>
        <img src="../img/gopher.png" class="gophers l">
        <img src="../img/gopher.png" class="gophers r">
		<h1>{{.LabelCompletion}}</h1>
       
		<h2><em>{{.LabelPresented}}</em></h2>
		<h1>{{.Name}}</h1>
		<h2>{{.LabelParticipation}}</h2>
		<p>
			<em>{{.LabelDate}}</em>
		</p>
        <img src="../img/stamp.png" class="stamp">
	</body>
</html>	
`
