package gohtml

import "testing"

const (
	unformattedHTML = `<!DOCTYPE html><html><head><title>This is a title.</title></head><body><p>Line1<br>Line2</p><br/><script>if (1) { alert('test'); }</script></body></html> <!-- aaa -->`
	formattedHTML   = `<!DOCTYPE html>
<html>
  <head>
    <title>
      This is a title.
    </title>
  </head>
  <body>
    <p>
      Line1
      <br>
      Line2
    </p>
    <br/>
    <script>
if (1) {
  alert('test');
}
</script>
  </body>
</html>
<!-- aaa -->`
)

func TestFormat(t *testing.T) {
	actual := Format(unformattedHTML)
	if actual != formattedHTML {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", formattedHTML, actual)
	}
}

func TestFormatBytes(t *testing.T) {
	actual := string(FormatBytes([]byte(unformattedHTML)))
	if actual != formattedHTML {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", formattedHTML, actual)
	}
}

func _TestFormatWithLineNo(t *testing.T) {
	actual := FormatWithLineNo(unformattedHTML)
	expected := ` 1  <!DOCTYPE html>
 2  <html>
 3    <head>
 4      <title>
 5        This is a title.
 6      </title>
 7    </head>
 8    <body>
 9      <p>
10        Line1
11        <br>
12        Line2
13      </p>
14      <br/>
15      <script>
16        alert('test');
17      </script>
18    </body>
19  </html>
20  <!-- aaa -->`
	if actual != expected {
		t.Errorf("Invalid result. [expected: %s][actual: %s]", expected, actual)
	}
}
