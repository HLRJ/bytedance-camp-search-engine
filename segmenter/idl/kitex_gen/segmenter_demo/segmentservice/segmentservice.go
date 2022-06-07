// Code generated by Kitex v0.3.1. DO NOT EDIT.

package segmentservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"segmenter/idl/kitex_gen/segmenter_demo"
)

func serviceInfo() *kitex.ServiceInfo {
	return segmentServiceServiceInfo
}

var segmentServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SegmentService"
	handlerType := (*segmenter_demo.SegmentService)(nil)
	methods := map[string]kitex.MethodInfo{
		"NGramSplit": kitex.NewMethodInfo(nGramSplitHandler, newNGramSplitArgs, newNGramSplitResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "segmenter",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func nGramSplitHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(segmenter_demo.SegmenterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(segmenter_demo.SegmentService).NGramSplit(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *NGramSplitArgs:
		success, err := handler.(segmenter_demo.SegmentService).NGramSplit(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*NGramSplitResult)
		realResult.Success = success
	}
	return nil
}
func newNGramSplitArgs() interface{} {
	return &NGramSplitArgs{}
}

func newNGramSplitResult() interface{} {
	return &NGramSplitResult{}
}

type NGramSplitArgs struct {
	Req *segmenter_demo.SegmenterRequest
}

func (p *NGramSplitArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in NGramSplitArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *NGramSplitArgs) Unmarshal(in []byte) error {
	msg := new(segmenter_demo.SegmenterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var NGramSplitArgs_Req_DEFAULT *segmenter_demo.SegmenterRequest

func (p *NGramSplitArgs) GetReq() *segmenter_demo.SegmenterRequest {
	if !p.IsSetReq() {
		return NGramSplitArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *NGramSplitArgs) IsSetReq() bool {
	return p.Req != nil
}

type NGramSplitResult struct {
	Success *segmenter_demo.SegmenterResponse
}

var NGramSplitResult_Success_DEFAULT *segmenter_demo.SegmenterResponse

func (p *NGramSplitResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in NGramSplitResult")
	}
	return proto.Marshal(p.Success)
}

func (p *NGramSplitResult) Unmarshal(in []byte) error {
	msg := new(segmenter_demo.SegmenterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *NGramSplitResult) GetSuccess() *segmenter_demo.SegmenterResponse {
	if !p.IsSetSuccess() {
		return NGramSplitResult_Success_DEFAULT
	}
	return p.Success
}

func (p *NGramSplitResult) SetSuccess(x interface{}) {
	p.Success = x.(*segmenter_demo.SegmenterResponse)
}

func (p *NGramSplitResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) NGramSplit(ctx context.Context, Req *segmenter_demo.SegmenterRequest) (r *segmenter_demo.SegmenterResponse, err error) {
	var _args NGramSplitArgs
	_args.Req = Req
	var _result NGramSplitResult
	if err = p.c.Call(ctx, "NGramSplit", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}