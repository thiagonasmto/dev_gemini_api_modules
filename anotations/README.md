### Resposta da requisição à API Gemini

> Slice de bytes (ASCII) — representação binária do JSON
```
usageMetadata, err := json.MarshalIndent(response, "", "  ")
if err != nil {
	log.Fatal(err)
}

fmt.Println(usageMetadata)

Output:
[123 10 ... 32 32 ... 125]
```

> Conteúdo da resposta do modelo Gemini (completions, headers, tokens, etc.)
```
fmt.Println(string(usageMetadata))

Output:
{
  "sdkHttpResponse": {
    "headers": {
      "Alt-Svc": [
        "h3=\":443\"; ma=2592000,h3-29=\":443\"; ma=2592000"
      ],
      "Content-Type": [
        "application/json; charset=UTF-8"
      ],
      "Date": [
        "Tue, 05 Aug 2025 19:07:57 GMT"
      ],
      "Server": [
        "scaffolding on HTTPServer2"
      ],
      "Server-Timing": [
        "gfet4t7; dur=5470"
      ],
      "Vary": [
        "Origin",
        "X-Origin",
        "Referer"
      ],
      "X-Content-Type-Options": [
        "nosniff"
      ],
      "X-Frame-Options": [
        "SAMEORIGIN"
      ],
      "X-Xss-Protection": [
        "0"
      ]
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

> Struct Go nativa (UsageMetadata) — usada para análise programática
```
fmt.Println(response.UsageMetadata)

Output:
&{[] 0 11 [] 9 [0xc0003f9b18] 747 0 [] 767 }
```

### Preços
> O preço da resposta é a soma dos tokens de saída e de pensamento. É possível conferir o número total de tokens de pensamento gerados no campo `thoughtsTokenCount`.