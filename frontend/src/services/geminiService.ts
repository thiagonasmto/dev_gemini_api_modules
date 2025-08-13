import axios from "axios"

const api = axios.create({
    baseURL: "http://localhost:8080/gemini-service",
    headers: {
    "Content-Type": "application/json",
  },
});

export interface GeminiRequest {
  role: "user" | "model" | "system";
  content: string;
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
export async function sendMessage(data: GeminiRequest): Promise<string> {
  try {
    const response = await api.post("/", data);

    return response.data.response.candidates[0].content.parts[0].text;
  } catch (error) {
    console.error("Erro ao chamar Gemini Service:", error);
    throw error;
  }
}