package main

import "chat-server-golang/network"


func main(){
	n := network.NewServer();
	n.StartServer();

}