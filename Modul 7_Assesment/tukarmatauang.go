package main

import "fmt" 

func main() {
	var dinar, dirham, fals, qirat int 

	fmt.Print("Masukkan Uang dalam Satuan Qirat : ")
	fmt.Scan(&qirat)

	dinar = qirat / 600	
	dirham = qirat % 600 / 60
	fals = qirat % 600 % 60 / 6
	qirat = qirat % 600 % 60 % 6 
	
	fmt.Println("Uang dalam satuan Dinar, Dirham, Fals, dan Qirat adalah ", dinar, dirham, fals, qirat)
	}