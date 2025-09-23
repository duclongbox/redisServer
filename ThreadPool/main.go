package main

import (
	"io"
	"log"
	"net"
)

// element in queue
type Job struct {
	conn net.Conn
}

// Threads in the pool
type Worker struct {
	id       int
	jobQueue chan Job
}

type Pool struct {
	//queue
	jobQueue chan Job
	workers  []*Worker
}

func NewPool(n int) *Pool {
	return &Pool{
		jobQueue: make(chan Job),
		workers:  make([]*Worker, n),
	}
}

func (p *Pool) AddJob(conn net.Conn) {
	p.jobQueue <- Job{conn: conn}
}

func  NewWorker(id int, jobQueue chan Job) *Worker {
	return &Worker{
		id:       id,
		jobQueue: jobQueue,
	}
}

func (p *Pool) Start() {
	for i := 0; i < len(p.workers); i++ {
		worker := NewWorker(i, p.jobQueue)
		p.workers[i] = worker
		worker.Start()

	}
}

func (w *Worker) Start() {
	go func() {
		for job := range w.jobQueue {
			log.Printf("Worker %d is processing from %s", w.id, job.conn.RemoteAddr())
			handleConnection(job.conn)
		}
	}()
}

func readCommand(c net.Conn) (string, error){
	var buf []byte = make([]byte, 512)

	// Blocking I/O : method
	n, err := c.Read(buf)
	if err != nil{
		return "", err
	}
	return string(buf[:n]), nil
}

func respond (cmd string, c net.Conn) error{
	if _,err := c.Write([]byte(cmd)); err != nil{
		return err
	}
	return nil
}

func handleConnection(conn net.Conn) {
	// // create buffer store recieved data
	// defer conn.Close()
	// var buf []byte = make([]byte, 1000)
	// _, err := conn.Read(buf) // -> blocking call
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //process req : eg, sleep 1 sec
	// time.Sleep(time.Second * 5)
	// conn.Write([]byte("HTTP/1.1 200 OK \n\nhello, world"))
	log.Println("handle conn from ", conn.RemoteAddr())
	for {
		cmd, err :=  readCommand(conn)
		log.Println("command: ", cmd)
		if err !=  nil{
			conn.Close()
			log.Println("client disconnected: ", conn.RemoteAddr())
			if err == io.EOF{
				break
			}
		}
		if err = respond(cmd, conn); err != nil{
			log.Println("err write: ", err)
		}
	}
	


}

func main() {
	listener, err := net.Listen("tcp", ":3000")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening at port 3000")
	defer listener.Close()

	pool := NewPool(2)
	pool.Start()

	// conn  == socket == communication channel

	// listen to multi clients with solution 1 connect - 1 threads
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// key word go to create a new thread
		// go handleConnection(conn)

		pool.AddJob(conn)

	}

}
