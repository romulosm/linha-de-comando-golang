package app

import (
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando para ser executada
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips de servicores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca Ips de servidores na internet",
			Action: buscaIps,
			Flags:  flags,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome de servidores na internet",
			Action: buscaServidores,
			Flags:  flags,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		log.Println(ip)
	}
}

func buscaServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		log.Println(servidor.Host)
	}
}
