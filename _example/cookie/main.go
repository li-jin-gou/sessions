// The MIT License (MIT)
//
// Copyright (c) 2016 Bo-Yi Wu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// This file may have been modified by CloudWeGo authors. All CloudWeGo
// Modifications are Copyright 2022 CloudWeGo Authors.

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/cookie"
)

func main() {
	h := server.New(server.WithHostPorts(":8000"))
	store := cookie.NewStore([]byte("secret"))
	h.Use(sessions.New("mysession", store))
	h.GET("/incr", func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v != nil {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, utils.H{"count": count})
	})
	h.Spin()
}
