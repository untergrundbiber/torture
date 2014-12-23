package main

import (
	"fmt"
	"path"
	"log"
	"github.com/jlaffaye/ftp"
)

type FTP struct {
	url string
	running bool
	conn *ftp.ServerConn
}

func (elem *FTP) crawlDirectory(dir string) {
	list, err := elem.conn.List(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range list {
		ff := path.Join(dir, file.Name)

		// go deeper!
		if file.Type == ftp.EntryTypeFolder {
			elem.crawlDirectory(ff)
		}
		// into teh elastics
		if file.Type == ftp.EntryTypeFile {
			fmt.Println(ff)
		}
	}
}

func (elem *FTP) crawlFtpDirectories() {
	pwd, _ := elem.conn.CurrentDir()
	elem.crawlDirectory(pwd)
}