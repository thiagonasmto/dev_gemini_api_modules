import { useState } from "react";
import { sendMessage, type GeminiRequest } from "./services/geminiService";
import TextareaAutosize from "react-textarea-autosize";
import ReactMarkdown from "react-markdown";
import "./App.css";

function App() {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState<GeminiRequest[]>([]);

  const handleSend = async () => {
    if (!message.trim()) return;

    // Adiciona mensagem do usuário
    const userMessage: GeminiRequest = {
      role: "user",
      content: message,
    };
    setMessages((prev) => [...prev, userMessage]);

    try {
      const responseText = await sendMessage(userMessage);

      const assistantMessage: GeminiRequest = {
        role: "model", // role retornado pelo Gemini
        content: responseText,
      };

      // Adiciona resposta do assistente
      setMessages((prev) => [...prev, assistantMessage]);
    } catch {
      setMessages((prev) => [
        ...prev,
        { role: "model", content: "Erro ao processar requisição" },
      ]);
    }

    setMessage("");
  };

  return (
    <div className="container-chat">
      <div className="chat-messages">
        {messages.map((m, i) => (
          <div
            key={i}
            className={`message-bubble ${m.role === "user" ? "user" : "model"}`}
          >
            <ReactMarkdown>{m.content}</ReactMarkdown>
          </div>
        ))}
      </div>

      <div className="prompt-input">
        <TextareaAutosize
          minRows={1}
          value={message}
          onChange={(e) => setMessage(e.target.value)}
          placeholder="Digite sua mensagem"
          onKeyDown={(e) => {
            if (e.key === "Enter" && !e.shiftKey) {
              e.preventDefault();
              handleSend();
            }
          }}
        />
        <button onClick={handleSend}>Enviar</button>
      </div>
    </div>
  );
}

export default App;
