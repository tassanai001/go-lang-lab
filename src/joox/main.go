package main

import (
	"fmt"
	"log"
	"strconv"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UrlAllArtists struct {
	Url    string
	IsUsed bool
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	c := session.DB("server").C("url-all-artists-page")
	result := UrlAllArtists{}
	err = c.Find(bson.M{"isused": false}).One(&result)
	if err != nil {
		for i := 1; i < 283; i++ {
			itemsPerPage := 50
			currentPage := i
			from := (((currentPage - 1) * itemsPerPage) + 1) - 1
			to := (currentPage * itemsPerPage) - 1

			if from < to {
				url := "https://api-jooxtt.sanook.com/web-fcgi-bin/web_all_singer_list?country=th&lang=th&sin=" + strconv.Itoa(from) + "&ein=" + strconv.Itoa(to) + "&is_all=1"
				saveToDB(url)
			}

		}
		fmt.Println("Done...!")
		callGetSongListWithSongId()
	} else {
		callGetSongListWithSongId()
	}
}

func callGetSongListWithSongId() {
	session, err := mgo.Dial("localhost:27017")
	con := session.DB("server").C("url-all-artists-page")
	count, err := con.Find(bson.M{"isused": false}).Count()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Count:", count)
		for i := 1; i < count+1; i++ {
			result := UrlAllArtists{}
			err := con.Find(bson.M{"isused": false}).One(&result)
			if err != nil {
				fmt.Println(err)
			} else {
				// fmt.Println(result.Url)
				colQuerier := result
				change := bson.M{"$set": bson.M{"url": result.Url, "isused": true}}
				err = con.Update(colQuerier, change)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("update done...!")
				}
			}
		}
	}
	// "https: //api-jooxtt.sanook.com/web-fcgi-bin/web_album_singer?country=th&lang=th&cmd=2&sin=0&ein=200000&singerid=4082"
}

func saveToDB(url string) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("server").C("url-all-artists-page")
	err = c.Insert(&UrlAllArtists{url, false})

	if err != nil {
		log.Fatal(err)
	}
}
