package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	Copy "github.com/otiai10/copy"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

func has_argument(argument string) int {
	for _, arg := range os.Args[1:] {
		if arg == argument {
			return 1
		}
	}
	return 0
}

func fread(filename string) string {
	contents, _ := ioutil.ReadFile(filename)
	return string(contents)
}

func fwrite(filename string, text string) {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(text)
	f.Close()
}

func truncate(text string, words int) string {
	m1 := regexp.MustCompile(`(?s)<.*?>`)
	replaced := m1.ReplaceAllString(text, " ")
	tokens := strings.Fields(replaced)
	max_len := int(math.Min(float64(len(tokens)), float64(words)))
	return strings.Join(tokens[:max_len], " ")
}

func read_headers(text string) (map[string]string, int) {
	headers := make(map[string]string)
	m1 := regexp.MustCompile(`\s*<!--\s*(.+?)\s*:\s*(.+?)\s*-->\s*|.+`)
	matches := m1.FindAllStringSubmatch(text, -1)
	for _, v := range matches {
		if len(strings.TrimSpace(v[1])) == 0 || len(strings.TrimSpace(v[2])) == 0 {
			continue
		}
		headers[v[1]] = v[2]
	}
	header_tag := "+++"
	header_len := len(header_tag)
	header_start := strings.Index(text, header_tag)
	header_end := 0
	if header_start != -1 {
		header_end = strings.Index(text[header_start+1:], header_tag)
		if header_end != -1 {
			header_text := text[header_start:header_end]
			for _, line := range strings.Split(header_text, "\n") {
				equals := strings.Index(line, "=")
				if equals == -1 {
					continue
				}
				tokens := strings.Split(line, "=")
				key := strings.TrimSpace(tokens[0])
				value := strings.ReplaceAll(strings.TrimSpace(tokens[1]), "\"", "")
				headers[key] = value
			}
		}
	}
	return headers, header_start + header_end + header_len + 1
}

func include_drafts() bool {
	return has_argument("--drafts") == 1
}

func debug_mode() bool {
	return has_argument("--debug") == 1
}

func fileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func endsWith(str string, candidates []string) bool {
	for _, candidate := range candidates {
		if strings.HasSuffix(str, candidate) {
			return true
		}
	}
	return false
}

func markdownToHtml(markdown string) string {
	var buf bytes.Buffer
	mark := goldmark.New(
		goldmark.WithRendererOptions(
			html.WithXHTML(),
			html.WithUnsafe(),
		),
	)
	if err := mark.Convert([]byte(markdown), &buf); err != nil {
		panic(err)
	}
	return string(buf.Bytes())
}

func read_content(filename string) map[string]string {
	text := fread(filename)
	content := make(map[string]string)
	content["date"] = "1970-01-01"
	content["slug"] = fileNameWithoutExtension(filepath.Base(filename))

	headers, end := read_headers(text)
	for key, val := range headers {
		content[key] = val
	}

	text = text[end:]

	if endsWith(filename, []string{".md", ".mkd", ".mkdn", ".mdown", ".markdown"}) {
		text = markdownToHtml(text)
	}

	content["content"] = text
	return content
}

func render(template string, params map[string]string) string {
	m1 := regexp.MustCompile(`{{\s*([^}\s]+)\s*}}`)
	matches := m1.FindAllStringSubmatch(template, -1)
	for _, v := range matches {
		param_key := v[1]
		if value, ok := params[param_key]; ok {
			template = strings.ReplaceAll(template, v[0], value)
		}
	}
	return template
}

func set_page_state(params map[string]string, build_drafts bool) bool {
	params["post_state"] = ""
	if val, ok := params["draft"]; ok && val == "true" {
		if !build_drafts {
			return false
		}
		params["post_state"] = "draft"
	}
	return true
}

func make_pages(src string, dst string, layout string, public_dir string, params map[string]string) []map[string]string {
	items := make([]map[string]string, 0)

	drafts := include_drafts()
	src_paths, _ := filepath.Glob(src)
	var wg sync.WaitGroup
	for _, src_path := range src_paths {
		content := read_content(src_path)
		for k, v := range params {
			if _, ok := content[k]; !ok {
				content[k] = v
			}
		}
		if set_page_state(content, drafts) == false {
			continue
		}
		rendered_content := render(content["content"], content)
		content["content"] = rendered_content
		items = append(items, content)

		dst_path := render(dst, content)
		output := render(layout, content)

		wg.Add(1)
		go func(wg *sync.WaitGroup, src_path string, dst_path string, content map[string]string, public_dir string, output string) {
			defer wg.Done()
			fmt.Println(fmt.Sprintf("Rendering %s => %s ...", src_path, dst_path))
			if val, ok := content["alias"]; ok {
				fwrite(public_dir+"/"+val+"/index.html", output)
			}
			fwrite(dst_path, output)
		}(&wg, src_path, dst_path, copy_params(content), public_dir, output)
	}
	wg.Wait()
	sort.Slice(items,
		func(i, j int) bool {
			return items[i]["date"] > items[j]["date"]
		})

	return items
}

func make_list_output(posts []map[string]string, list_layout string, item_layout string, params map[string]string) string {
	items := make([]string, 0)
	drafts := include_drafts()
	for _, post := range posts {
		for k, v := range params {
			if _, ok := post[k]; !ok {
				post[k] = v
			}
		}
		if set_page_state(post, drafts) == false {
			continue
		}
		post["summary"] = truncate(post["content"], 25)
		item := render(item_layout, post)
		items = append(items, item)
	}

	params["content"] = strings.Join(items, "")
	output := render(list_layout, params)
	return output
}

func make_list(posts []map[string]string, dst string, list_layout string, item_layout string, params map[string]string) string {
	output_c := make(chan string)
	go func() {
		dst_path := render(dst, params)
		output := make_list_output(posts, list_layout, item_layout, copy_params(params))
		fmt.Println(fmt.Sprintf("Rendering list => %s ...", dst_path))
		fwrite(dst_path, output)
		output_c <- output
	}()
	output := <-output_c
	return output
}

func copy_params(originalMap map[string]string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range originalMap {
		newMap[key] = value
	}
	return newMap
}

func main() {
	public_dir := "public"
	if _, err := os.Stat(public_dir); os.IsNotExist(err) {
	} else {
		os.RemoveAll(public_dir)
	}
	Copy.Copy("static", public_dir)

	debug := 0
	if debug_mode() {
		debug = 1
	}
	params := map[string]string{
		"base_path":         "",
		"subtitle":          "Jarmo Nikkanen",
		"author":            "jarmonik",
		"site_url":          "http://jarmonik.org",
		"current_year":      strconv.Itoa(time.Now().Year()),
		"home_recent_posts": "",
		"debug":             strconv.Itoa(debug),
	}

	page_layout := fread("layout/page.html")
	post_layout := fread("layout/post.html")
	list_layout := fread("layout/list.html")
	list_layout_vanilla := fread("layout/list.html")
	item_layout := fread("layout/item.html")

	post_layout = render(page_layout, map[string]string{"content": post_layout})
	list_layout = render(page_layout, map[string]string{"content": list_layout})

	list_params := copy_params(params)
	list_params["blog"] = "notes"
	blog_posts := make_pages("content/notes/*.md",
		public_dir+"/notes/{{ slug }}/index.html",
		post_layout, public_dir, list_params)

	list_params = copy_params(params)
	list_params["blog"] = "post"
	news_posts := make_pages("content/post/*.md",
		public_dir+"/post/{{ slug }}/index.html",
		post_layout, public_dir, list_params)

	list_params = copy_params(params)
	list_params["blog"] = "notes"
	list_params["title"] = "Notes"
	make_list(blog_posts, public_dir+"/notes/index.html",
		list_layout, item_layout, list_params)

	list_params = copy_params(params)
	list_params["blog"] = "post"
	list_params["title"] = "Posts"
	make_list(news_posts, public_dir+"/post/index.html",
		list_layout, item_layout, list_params)

	list_params = copy_params(params)
	list_params["blog"] = "post"
	list_params["title"] = ""

	list_params = copy_params(params)
	list_params["title"] = ""
	recent_posts := make_list_output(news_posts[0:int(math.Min(5, float64(len(news_posts))))],
		list_layout_vanilla, item_layout, list_params)

	make_pages("content/*.md", public_dir+"/{{ slug }}/index.html",
		page_layout, public_dir, copy_params(params))

	params["home_recent_posts"] = recent_posts
	make_pages("content/_index.md", public_dir+"/index.html",
		page_layout, public_dir, copy_params(params))
}
