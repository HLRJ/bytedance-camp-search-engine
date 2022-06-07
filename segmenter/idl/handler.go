package main

import (
	"context"
	"segmenter/idl/kitex_gen/segmenter_demo"
)

// SegmentServiceImpl implements the last service interface defined in the IDL.
type SegmentServiceImpl struct{}

// NGramSplit implements the SegmentServiceImpl interface.
func (s *SegmentServiceImpl) NGramSplit(ctx context.Context, req *segmenter_demo.SegmenterRequest) (resp *segmenter_demo.SegmenterResponse, err error) {
	// TODO: Your code here...
	return
}
