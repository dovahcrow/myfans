package main

import (
	"log"
	"poke/models"
	"runtime"
	"snatch"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
}
func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	fids, err := snatch.GetFid()
	if err != nil {
		log.Fatalln(err)
	}

	avfids, err := snatch.ChkFidAv(fids)
	if err != nil {
		log.Fatalln(err)
	}

	err = models.DeleteAllFids()
	if err != nil {
		log.Println(`delete all fids fail`, err)
	}
	err = models.InsertFids(avfids)
	if err != nil {
		log.Fatalln(err)
	}

	err = models.DeleteAllTids()
	if err != nil {
		log.Println(`delete all tids fail`, err)
	}

	wg := sync.WaitGroup{}
	for _, v := range fids {
		time.Sleep(3 * time.Second)
		wg.Add(1)
		go func(v string) {
			defer func() {
				wg.Done()
			}()

			tids, err := snatch.GetTid(v)
			if err != nil {
				log.Println(err)
				return
			}

			avtids, err := snatch.ChkTidAv(tids)
			if err != nil {
				log.Println(err)
				return
			}

			err = models.InsertTids(avtids)
			if err != nil {
				log.Println(err)
				return
			}
		}(v)

	}
	wg.Wait()
	log.Println(`采集完毕`)
}
