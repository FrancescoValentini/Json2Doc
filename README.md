# Json2Doc

**Json2Doc** is a small command-line tool written in Go that renders documents from JSON data using a template file.  

You provide the JSON data, pick a template, and Json2Doc takes care of generating the final output.

This tool can be extremely useful and time-saving in contexts where it is necessary to produce periodic reports for humans from serialized data in a machine-readable format.

## 🚀 Usage 
```bash
# Example Command
json2doc -i data.json -t template.tpl -o output.txt
```

```bash
Options:
  -i    Input JSON file
  -o    Output file
  -t    Template file
  -v    Show version information
  -d    Enable verbose logging
```

> [!WARNING]
> If the output file already exists, it will be overwritten!

## 🧩 Templates

Templates use **Golang's `text/template` syntax**
([https://pkg.go.dev/text/template](https://pkg.go.dev/text/template)).

In this project, templates are used to **process JSON data and transform it into the desired output format** (Markdown, text files, reports, etc.).

Each JSON structure typically requires a **custom-made template**, designed specifically for the target document and use case.


---

### Template Example
**JSON data:** `data.json`
```json
[
  {
    "id": 1,
    "first_name": "Harmon",
    "last_name": "McGoldrick",
    "balance": "€5400,11"
  },{
    "id": 2,
    "first_name": "Julienne",
    "last_name": "Disbrey",
    "balance": "€1147,11"
  },{
    "id": 3,
    "first_name": "Julissa",
    "last_name": "Snoddy",
    "balance": "€6295,29"
  }
]
```

**Template:** `template.tpl`
```text
# Example Bank Report 
{{range . }}
## Account: N. {{ .id }} - {{ .first_name }} {{ .last_name }}
**First Name:** {{ .first_name }}  
**Last Name:** {{ .last_name }}  
**Balance:** {{ .balance }}  
---
{{- end }}
```

---

### Expected Output

**Command:**
```bash
json2doc -i data.json -t template.tpl -o BankReport.md
```

**Generated file:** `BankReport.md`
```md
# Example Bank Report 

## Account: N. 1 - Harmon McGoldrick
**First Name:** Harmon  
**Last Name:** McGoldrick  
**Balance:** €5400,11  
---
## Account: N. 2 - Julienne Disbrey
**First Name:** Julienne  
**Last Name:** Disbrey  
**Balance:** €1147,11  
---
## Account: N. 3 - Julissa Snoddy
**First Name:** Julissa  
**Last Name:** Snoddy  
**Balance:** €6295,29  
---
```


## ⚠️ Disclaimer ⚠️
This is an **amateur / personal project**, built primarily for personal use.
While it works for my needs, it may contain bugs, missing features. Use it at your own risk!