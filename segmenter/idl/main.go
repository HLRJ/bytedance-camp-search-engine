package main

import (
	"log"
	segmenter_demo "segmenter/idl/kitex_gen/segmenter_demo/segmentservice"
)

func main() {
	svr := segmenter_demo.NewServer(new(SegmentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
