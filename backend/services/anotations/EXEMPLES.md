# 📌 Exemplos de Uso da API Gemini em Go

Este repositório contém exemplos práticos de como interagir com a API **Google Gemini** utilizando **Go**, explorando diferentes modos de requisição e recursos como **streaming**, **respostas JSON** e **entendimento de documentos**.

---

## 🚀 Funcionalidades Demonstradas

* 🔹 **Requisição simples** (sem streaming)
* 🔹 **Resposta em JSON**
* 🔹 **Streaming de respostas**
* 🔹 **Entendimento de documentos locais (PDF)**
* 🔹 **Entendimento de documentos online (PDF)**

---

## 📂 Exemplos de Código

### 1️⃣ Requisição simples (No Streaming Response)

```go
response, err := utilexamples.SimpleRequest(client, ctx, model, contents, config)
if err != nil {
    fmt.Println("erro na comunicação com o gemini: ", err)
}

// Summary da resposta
if debugResponse {
    utilexamples.SummaryResponse(response)
} else {
    fmt.Print(response.Text())
}
```

---

### 2️⃣ Resposta em JSON

```go
utilexamples.JSONSimpleRequest(response)
```

---

### 3️⃣ Streaming de Resposta

```go
utilexamples.StreamingRequest(client, ctx, model, prompt, config)
```

---

### 4️⃣ Entendimento de Documentos Locais

```go
prompt := "Sobre o que fala esses documentos? Me explique em poucas palavras."
pdfPaths := []string{
    "./docs/pdf_examples/webinar-ibef-sobena-apresentacao.pdf",
}

contents := geminiresources.UnderstandingLocalDocuments(pdfPaths, ctx, client, prompt)
response, err := utilexamples.SimpleRequest(client, ctx, model, contents, config)
if err != nil {
    fmt.Println("erro na comunicação com o gemini: ", err)
}

if debugResponse {
    utilexamples.SummaryResponse(response)
} else {
    fmt.Print(response.Text())
}
```

---

### 5️⃣ Entendimento de Documentos Online

```go
doc_url := "https://discovery.ucl.ac.uk/id/eprint/10089234/1/343019_3_art_0_py4t4l_convrt.pdf"

response, _ := geminiresources.UnderstandingInLineDocuments(ctx, client, model, config, prompt, doc_url)
if debugResponse {
    utilexamples.SummaryResponse(response)
} else {
    fmt.Print(response.Text())
}
```

---

## ⚙️ Requisitos

* Go 1.20+
* Chave de API Google Gemini
* Dependências listadas no `go.mod`