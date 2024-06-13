package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/han-y-s/kitex_demo/kitex_gen/student_demo"
	"sync"
)

// StudentServiceImpl implements the last service interface defined in the IDL.
type StudentServiceImpl struct {
	studentCache map[int32]*student_demo.Student
	mu           sync.Mutex
}

// Register implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Register(ctx context.Context, student *student_demo.Student) (resp *student_demo.RegisterResp, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "start register student: %v", student)
	defer s.mu.Unlock()
	s.mu.Lock()
	if s.studentCache == nil {
		s.studentCache = make(map[int32]*student_demo.Student)
	}
	s.studentCache[student.Id] = student
	resp = &student_demo.RegisterResp{
		Success: true,
		Message: "register successfully",
	}
	klog.CtxInfof(ctx, "end register student: %v", student)
	return
}

// Query implements the StudentServiceImpl interface.
func (s *StudentServiceImpl) Query(ctx context.Context, req *student_demo.QueryReq) (resp *student_demo.Student, err error) {
	// TODO: Your code here...
	klog.CtxInfof(ctx, "start query student: %v", req.Id)
	defer s.mu.Unlock()
	s.mu.Lock()
	resp = s.studentCache[req.Id]
	klog.CtxInfof(ctx, "end query student: %v", req.Id)
	return
}
