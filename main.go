package main

import (
	"net/http"
	"text/template"


	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type Dados struct {
	Imagem string
	Nome   string
	Texto  string
	Dado   int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	fs := http.FileServer(http.Dir("templates/css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	http.Handle("/imagens/", http.StripPrefix("/imagens/", http.FileServer(http.Dir("templates/imagens"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cpuPercent, _ := cpu.Percent(0, false)
	virtualMem, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")

	cPU := int(cpuPercent[0])
	dISK := int(diskUsage.UsedPercent)
	rAM := int(virtualMem.UsedPercent)

	resultados := []Dados{
		{"/imagens/CPU.png", "CPU", "da CPU", cPU},
		{"/imagens/disk.png", "Disco", "do disco", dISK},
		{"/imagens/virtual-memory.png", "RAM", "da RAM", rAM},
	}

	temp.ExecuteTemplate(w, "Index", resultados)
}
