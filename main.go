package main

import (
	"fmt"
	"time"
)

var show bool

type Process struct {
	id        int
	is_finish bool
}

func (p *Process) start(c chan string) {
	go printChannel(c)
	i := 0
	for {
		if p.is_finish == true {
			return
		}
		if show == true {
			s := fmt.Sprint("id ", p.id, ": ", i)
			c <- s
		}
		i++
		time.Sleep(time.Millisecond * 500)
	}
}

func (p *Process) stop() {
	p.is_finish = true
}

func printChannel(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {
	op := 1
	var id_count = 1
	var process_slide []*Process

	show = false
	c := make(chan string)

	for op != 0 {
		fmt.Println(".::PROCESOS::.")
		fmt.Println("1) Agregar proceso")
		fmt.Println("2) Mostrar procesos")
		fmt.Println("3) Finalizar proceso")
		fmt.Println("0) Salir")
		fmt.Println("Opción: ")
		fmt.Scan(&op)

		switch op {

		case 0:
			fmt.Println("Adios")

		case 1:
			p := &Process{
				id:        id_count,
				is_finish: false,
			}
			go p.start(c)
			process_slide = append(process_slide, p)

			fmt.Printf("Proceso %d agregado\n", id_count)
			id_count++

		case 2:
			show = !show

		case 3:
			var process_id int
			fmt.Print("Capture ID del proceso a finalizar: ")
			fmt.Scan(&process_id)
			for i, v := range process_slide {
				if v.id == process_id {
					fmt.Printf("Proceso %d Finalizado\n", v.id)
					v.stop()
					process_slide = append(process_slide[:i], process_slide[i+1:]...)
					break
				}
			}

		default:
			fmt.Println("Opción no existe")

		}
	}
}
