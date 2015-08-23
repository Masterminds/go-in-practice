package main

import (
	"bytes"
	"testing"
	"text/template"
)

func BenchmarkTemplates(b *testing.B) {
	// b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t, _ := template.New("test").Parse(tpl)
		t.Execute(&buf, data)
		buf.Reset()
	}
}
func BenchmarkCompiledTemplates(b *testing.B) {
	// b.Logf("b.N is %d\n", b.N)
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	for i := 0; i < b.N; i++ {
		t.Execute(&buf, data)
		buf.Reset()
	}
}

func BenchmarkParallelTemplates(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}

func BenchmarkParallelOops(b *testing.B) {
	tpl := "Hello {{.Name}}"
	t, _ := template.New("test").Parse(tpl)
	data := &map[string]string{
		"Name": "World",
	}
	var buf bytes.Buffer
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t.Execute(&buf, data)
			buf.Reset()
		}
	})
}
