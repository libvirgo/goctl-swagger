package action

import (
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/goctl-swagger/generate"
)

func Generator(ctx *cli.Context) error {
	fileName := ctx.String("filename")

	if len(fileName) == 0 {
		fileName = "rest.swagger.json"
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	basepath := ctx.String("basepath")
	host := ctx.String("host")
	https := ctx.Bool("https")
	if err = generate.Do(fileName, host, basepath, https, p); err != nil {
		return err
	}
	print("generate " + fileName + " success")
	return nil
}
