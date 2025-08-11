## 📥 Resposta da Requisição à API Gemini

A API Gemini pode retornar os dados em diferentes formatos, permitindo desde inspeção em **nível binário** até **uso programático** com structs Go.

---

### 1️⃣ Slice de bytes (ASCII) — Representação binária do JSON

```go
usageMetadata, err := json.MarshalIndent(response, "", "  ")
if err != nil {
	log.Fatal(err)
}

fmt.Println(usageMetadata)
```

**Exemplo de saída:**

```
[123 10 ... 32 32 ... 125]
```

---

### 2️⃣ Conteúdo da resposta do modelo Gemini (JSON formatado)

```go
fmt.Println(string(usageMetadata))
```

**Exemplo de saída:**

```json
{
  "sdkHttpResponse": {
    "headers": {
      "Alt-Svc": ["h3=\":443\"; ma=2592000,h3-29=\":443\"; ma=2592000"],
      "Content-Type": ["application/json; charset=UTF-8"],
      "Date": ["Tue, 05 Aug 2025 19:07:57 GMT"],
      "Server": ["scaffolding on HTTPServer2"],
      "Server-Timing": ["gfet4t7; dur=5470"],
      "Vary": ["Origin", "X-Origin", "Referer"],
      "X-Content-Type-Options": ["nosniff"],
      "X-Frame-Options": ["SAMEORIGIN"],
      "X-Xss-Protection": ["0"]
    }
  },
  "candidates": [
    {
      "content": {
        "parts": [
          {
            "text": "AI learns patterns from data to make predictions or solve problems."
          }
        ],
        "role": "model"
      },
      "finishReason": "STOP"
    }
  ],
  "modelVersion": "gemini-2.5-flash",
  "usageMetadata": {
    "candidatesTokenCount": 12,
    "promptTokenCount": 9,
    "promptTokensDetails": [
      {
        "modality": "TEXT",
        "tokenCount": 9
      }
    ],
    "thoughtsTokenCount": 787,
    "totalTokenCount": 808
  }
}
```

---

### 3️⃣ Struct Go Nativa (`UsageMetadata`) — Uso programático

```go
fmt.Println(response.UsageMetadata)
```

**Exemplo de saída:**

```
&{[] 0 11 [] 9 [0xc0003f9b18] 747 0 [] 767 }
```

---

## 💰 Preços e Tokens

O custo é calculado com base na soma de:

* **Tokens de saída** (`candidatesTokenCount`)
* **Tokens de pensamento** (`thoughtsTokenCount`)

📌 Você pode acessar esse número no campo `thoughtsTokenCount` do objeto `usageMetadata`.

---

## 📁 API File

* Permite armazenar **até 50 MB** de arquivos PDF
* Arquivos ficam disponíveis por **48 horas**
* Acesso somente via chave de API (download não suportado pela API)
* Disponível **sem custo adicional** em todas as regiões com suporte ao Gemini

---

## 📄 Entendendo Documentos — Detalhes Técnicos

* **Limite**: até **1.000 páginas** por documento
* **Tokens por página**: \~258
* **Resolução**:

  * Máx.: `3072 x 3072` (mantendo proporção)
  * Mín.: ajustado para `768 x 768`
* **Tipos aceitos**: PDF recomendado

  * Outros formatos (TXT, Markdown, HTML, XML) são tratados como **texto puro**
  * Perde-se formatação, gráficos, tabelas e elementos visuais

---

### 🛠 Boas Práticas

* 📐 Gire as páginas para a orientação correta antes do upload
* 🔍 Evite imagens desfocadas
* 📝 Se for enviar apenas **uma página**, coloque o prompt **após ela**

### Para Estudar

* Arquitetura Agent2Agent - A2A (Google)
