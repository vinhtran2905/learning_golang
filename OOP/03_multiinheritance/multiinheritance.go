package main

import "fmt"

type Camera struct {
}

func (_ Camera) takePicture() string {
	return "click"
}

type Phone struct {
}

func (_ Phone) call() string {
	return "ring ring"
}

type CameraPhone struct {
	Camera
	phone Phone
}

func main() {
	cp := CameraPhone{}
	fmt.Println(cp.takePicture())
	fmt.Println(cp.phone.call())
}
