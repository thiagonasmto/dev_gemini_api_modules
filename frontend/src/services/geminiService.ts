import axios from "axios"

const api = axios.create({
    baseURL: "http://localhost:8080/gemini-service",
    headers: {
    "Content-Type": "application/json",
  },
});

export interface GeminiRequestMessage {
  role: "user" | "model" | "system";
  content: string;
  file?: File;
}

export interface GeminiResponse {
  response: {
    candidates: Array<{
      content: {
        parts: Array<{
          text: string;
        }>;
        role: string;
      };
      finishReason: string;
    }>;
    modelVersion: string;
    responseId: string;
    usageMetadata: {
      candidatesTokenCount: number;
      promptTokenCount: number;
      promptTokensDetails: Array<{
        modality: string;
        tokenCount: number;
      }>;
      thoughtsTokenCount: number;
      totalTokenCount: number;
    };
  };
}

// Função para enviar mensagem e extrair apenas o texto da resposta
export async function sendMessage(data: GeminiRequestMessage): Promise<string> {
  try {
    let response;

    if (data.file) {
      // Se houver arquivo, envia como FormData
      const formData = new FormData();
      formData.append("file", data.file);
      formData.append("prompt", data.content); // content já é o prompt

      response = await api.post("/gemini-understanding-doc", formData, {
        headers: { "Content-Type": "multipart/form-data" },
      });
    } else {
      // Apenas texto
      response = await api.post("/", data);
    }

    // Retorna apenas o texto principal da resposta
    return response.data.response.candidates?.[0]?.content?.parts?.[0]?.text || "";
  } catch (error) {
    console.error("Erro ao chamar Gemini Service:", error);
    throw error;
  }
}